package main

import (
	"drylang/compiler"
	"drylang/lexer"
	"drylang/parser"
	"drylang/vm"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("y <file.y|folder|all>")
		os.Exit(1)
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
		// Single file or folder
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

	// Read and combine sources
	var combined strings.Builder
	for _, file := range files {
		src, err := os.ReadFile(strings.TrimSpace(file))
		if err != nil {
			die(fmt.Errorf("0:0 read %s", file))
		}
		combined.Write(src)
		combined.WriteString("\n")
	}

	// Run pipeline
	if err := run(combined.String()); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(source string) error {
	// Lex
	lex := lexer.New(source)
	tokens, err := lex.Tokenize()
	if err != nil {
		return err
	}

	// Parse
	p := parser.New(tokens)
	prog, err := p.Parse()
	if err != nil {
		return err
	}

	// Compile (skip analyzer for now, enable later)
	comp := compiler.New()
	chunk, fns, err := comp.Compile(prog)
	if err != nil {
		return err
	}

	// Run
	machine := vm.New(chunk, fns)
	return machine.Run()
}

func die(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}
