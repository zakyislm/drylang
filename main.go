package main

import (
	"drylang/compiler"
	"drylang/errfmt"
	"drylang/lexer"
	"drylang/parser"
	"drylang/vm"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		runREPL()
		os.Exit(0)
	}

	target := os.Args[1]
	var files []string

	if target == "all" {
		// Run all .y files in current directory
		entries, err := os.ReadDir(".")
		if err != nil {
			die(err)
		}
		for _, e := range entries {
			if !e.IsDir() && strings.HasSuffix(e.Name(), ".y") {
				files = append(files, e.Name())
			}
		}
		if len(files) == 0 {
			fmt.Println("0:0 no .y files")
			os.Exit(1)
		}
	} else if strings.Contains(target, ",") {
		// Multiple files: y file1.y,file2.y
		files = strings.Split(target, ",")
	} else {
		// Single file or folder or URL
		if strings.HasPrefix(target, "http://") || strings.HasPrefix(target, "https://") {
			files = []string{target}
		} else {
			info, err := os.Stat(target)
			if err != nil {
				die(err)
			}
			if info.IsDir() {
				entries, err := os.ReadDir(target)
				if err != nil {
					die(err)
				}
				for _, e := range entries {
					if !e.IsDir() && strings.HasSuffix(e.Name(), ".y") {
						files = append(files, filepath.Join(target, e.Name()))
					}
				}
			} else {
				files = []string{target}
			}
		}
	}

	ldr := newLoader()
	var finalProg parser.Program

	for _, file := range files {
		prog, err := ldr.loadAndParse(file, ".")
		if err != nil {
			die(err)
		}
		finalProg.Stmts = append(finalProg.Stmts, prog.Stmts...)
	}

	errfmt.Init("")

	if err := runAST(&finalProg); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

type loader struct {
	visited map[string]bool
}

func newLoader() *loader {
	return &loader{visited: make(map[string]bool)}
}

func (l *loader) loadAndParse(target string, baseDir string) (*parser.Program, error) {
	var fullPath string
	isURL := false

	target = strings.TrimSpace(target)

	if strings.HasPrefix(target, "http://") || strings.HasPrefix(target, "https://") {
		isURL = true
		fullPath = target
	} else if strings.HasPrefix(baseDir, "http://") || strings.HasPrefix(baseDir, "https://") {
		isURL = true
		baseURL, _ := url.Parse(baseDir)
		targetURL, _ := url.Parse(target)
		fullPath = baseURL.ResolveReference(targetURL).String()
	} else {
		if filepath.IsAbs(target) {
			fullPath = filepath.Clean(target)
		} else {
			fullPath = filepath.Clean(filepath.Join(baseDir, target))
		}
	}

	if l.visited[fullPath] {
		return &parser.Program{}, nil // Skip already loaded to prevent cycles
	}
	l.visited[fullPath] = true

	var src []byte
	var err error

	if isURL {
		resp, err := http.Get(fullPath)
		if err != nil {
			return nil, fmt.Errorf("failed to fetch %s: %v", fullPath, err)
		}
		defer resp.Body.Close()
		src, err = io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
	} else {
		src, err = os.ReadFile(fullPath)
		if err != nil {
			return nil, fmt.Errorf("failed to read %s: %v", fullPath, err)
		}
	}

	lex := lexer.New(string(src))
	tokens, err := lex.Tokenize()
	if err != nil {
		return nil, err
	}

	p := parser.New(tokens)
	prog, err := p.Parse()
	if err != nil {
		return nil, err
	}

	var mergedStmts []parser.Stmt
	newBaseDir := baseDir
	if isURL {
		u, _ := url.Parse(fullPath)
		u.Path = filepath.Dir(u.Path)
		if !strings.HasSuffix(u.Path, "/") {
			u.Path += "/"
		}
		newBaseDir = u.String()
	} else {
		newBaseDir = filepath.Dir(fullPath)
	}

	for _, stmt := range prog.Stmts {
		if useStmt, ok := stmt.(*parser.UseStmt); ok {
			subProg, err := l.loadAndParse(useStmt.Path, newBaseDir)
			if err != nil {
				return nil, err
			}
			mergedStmts = append(mergedStmts, subProg.Stmts...)
		} else {
			mergedStmts = append(mergedStmts, stmt)
		}
	}

	prog.Stmts = mergedStmts
	return prog, nil
}

func runAST(prog *parser.Program) error {
	comp := compiler.New()
	chunk, fns, err := comp.Compile(prog)
	if err != nil {
		return err
	}

	machine := vm.New(chunk, fns)
	if err := machine.Run(); err != nil {
		return err
	}

	return nil
}

func run(source string) error {
	errfmt.Init(source)

	lex := lexer.New(source)
	tokens, err := lex.Tokenize()
	if err != nil {
		return err
	}

	p := parser.New(tokens)
	prog, err := p.Parse()
	if err != nil {
		return err
	}

	comp := compiler.New()
	chunk, fns, err := comp.Compile(prog)
	if err != nil {
		return err
	}

	machine := vm.New(chunk, fns)
	if err := machine.Run(); err != nil {
		return err
	}

	return nil
}

func die(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}
