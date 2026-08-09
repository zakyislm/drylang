package main

import (
	"drylang/compiler"
	"drylang/errfmt"
	"drylang/lexer"
	"drylang/parser"
	"drylang/vm"
	"embed"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

//go:embed templates/*
var templateFiles embed.FS

const Version = "1.0.0"

func main() {
	if len(os.Args) < 2 {
		runREPL()
		os.Exit(0)
	}

	target := os.Args[1]

	if target == "--help" || target == "-h" {
		printHelp()
		os.Exit(0)
	}

	if target == "--version" || target == "-v" {
		fmt.Printf("dryLang version v%s\n", Version)
		os.Exit(0)
	}

	if target == "init" {
		handleInit()
		os.Exit(0)
	}

	var files []string

	if target == "all" {
		// Run all .y files in current directory
		entries, err := os.ReadDir(".")
		if err != nil {
			die(err)
		}
		for _, e := range entries {
			if !e.IsDir() && (strings.HasSuffix(e.Name(), ".y") || strings.HasSuffix(e.Name(), ".dry")) {
				files = append(files, e.Name())
			}
		}
		if len(files) == 0 {
			fmt.Println("0:0 no .y or .dry files")
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
					if !e.IsDir() && (strings.HasSuffix(e.Name(), ".y") || strings.HasSuffix(e.Name(), ".dry")) {
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

	if strings.HasPrefix(target, "github.com/") {
		// Translate github.com/user/repo to raw.githubusercontent.com
		parts := strings.Split(target, "/")
		if len(parts) >= 3 {
			user := parts[1]
			repo := parts[2]
			var path string
			if len(parts) > 3 {
				path = strings.Join(parts[3:], "/")
			} else {
				path = "idx.y" // default entrypoint
			}
			target = fmt.Sprintf("https://raw.githubusercontent.com/%s/%s/main/%s", user, repo, path)
		}
	}

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

func printHelp() {
	fmt.Println(`dryLang - Writeless, get more.

Usage:
  dry [file|folder|url|github_repo]
  dry init [dir] [template]

Commands:
  init        Scaffold a new project in the specified directory.
              Templates: api, web, crud, fetch, cli, automation, scraper, hello.
              (If no template is specified, creates an empty idx.y)

Options:
  -h, --help     Show this help message.
  -v, --version  Show version information.

Examples:
  dry main.y                   # Run a local file
  dry .                        # Run all files in current directory
  dry github.com/user/repo     # Run from a GitHub repo
  dry init . api               # Create a REST API project in current dir
`)
}

func handleInit() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: dry init <folder> [template]")
		os.Exit(1)
	}

	targetDir := os.Args[2]
	if targetDir != "." {
		if err := os.MkdirAll(targetDir, 0755); err != nil {
			die(fmt.Errorf("failed to create directory: %v", err))
		}
	}

	targetFile := filepath.Join(targetDir, "idx.y")
	if _, err := os.Stat(targetFile); err == nil {
		fmt.Printf("Error: %s already exists\n", targetFile)
		os.Exit(1)
	}

	var templateName string
	if len(os.Args) >= 4 {
		templateName = os.Args[3]
	}

	templateMap := map[string]string{
		"api":        "rest-api.y",
		"web":        "file-server.y",
		"crud":       "crud.y",
		"fetch":      "fetch-json.y",
		"cli":        "cli-tool.y",
		"automation": "automation.y",
		"scraper":    "scraper.y",
		"hello":      "hello.y",
	}

	var content []byte
	if templateName == "" {
		content = []byte("// idx.y - Entry point\npt \"Hello, dryLang!\"\n")
	} else {
		fileName, ok := templateMap[templateName]
		if !ok {
			fmt.Printf("Unknown template: %s\n", templateName)
			fmt.Println("Available templates: api, web, crud, fetch, cli, automation, scraper, hello")
			os.Exit(1)
		}
		
		var err error
		content, err = templateFiles.ReadFile("templates/" + fileName)
		if err != nil {
			die(fmt.Errorf("failed to load template: %v", err))
		}
	}

	if err := os.WriteFile(targetFile, content, 0644); err != nil {
		die(fmt.Errorf("failed to write file: %v", err))
	}

	fmt.Printf("Successfully created %s", targetFile)
	if templateName != "" {
		fmt.Printf(" using template '%s'", templateName)
	}
	fmt.Println("\nRun it with: dry " + targetDir)
}
