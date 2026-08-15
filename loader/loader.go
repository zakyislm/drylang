package loader

import (
	"drylang/ast"
	"drylang/core"
	"drylang/errfmt"
	"drylang/lexer"
	"drylang/parser"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

type Loader struct {
	visited map[string]bool
}

func New() *Loader {
	return &Loader{visited: make(map[string]bool)}
}

func (l *Loader) LoadAndParse(target string, baseDir string) (*ast.Program, error) {
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
				path = "idx.dry" // default entrypoint
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
		return &ast.Program{}, nil // Skip already loaded to prevent cycles
	}
	l.visited[fullPath] = true

	var src []byte
	var err error

	if isURL {
		u, err := url.Parse(fullPath)
		if err != nil || !core.GetSandbox().AllowURL(u.Host) {
			return nil, fmt.Errorf("URL denied: %s", fullPath)
		}
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
		// Directory import: load every source file inside, sorted by name.
		if info, err := os.Stat(fullPath); err == nil && info.IsDir() {
			return l.loadDir(fullPath)
		}
		src, err = os.ReadFile(fullPath)
		if err != nil {
			return nil, fmt.Errorf("failed to read %s: %v", fullPath, err)
		}
	}

	errfmt.Init(string(src))
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

	var mergedStmts []ast.Stmt
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
		if useStmt, ok := stmt.(*ast.UseStmt); ok {
			subProg, err := l.LoadAndParse(useStmt.Path, newBaseDir)
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

// loadDir loads every source file directly inside a directory, in sorted
// order with .dry prioritized over .y over other files. Subdirectories are
// not recursed; nested modules must be imported explicitly.
func (l *Loader) loadDir(dir string) (*ast.Program, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, fmt.Errorf("failed to read dir %s: %v", dir, err)
	}

	var files []string
	for _, e := range entries {
		if !e.IsDir() {
			files = append(files, e.Name())
		}
	}
	sort.Slice(files, func(i, j int) bool {
		pi, pj := filePriority(files[i]), filePriority(files[j])
		if pi != pj {
			return pi < pj
		}
		return files[i] < files[j]
	})

	merged := &ast.Program{}
	for _, name := range files {
		prog, err := l.LoadAndParse(filepath.Join(dir, name), dir)
		if err != nil {
			return nil, err
		}
		merged.Stmts = append(merged.Stmts, prog.Stmts...)
	}
	return merged, nil
}

// filePriority ranks extensions: .dry first, .y second, anything else last.
func filePriority(name string) int {
	switch {
	case strings.HasSuffix(name, ".dry"):
		return 0
	case strings.HasSuffix(name, ".y"):
		return 1
	default:
		return 2
	}
}
