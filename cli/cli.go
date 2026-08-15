package cli

import (
	"drylang/ast"
	"drylang/loader"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/joho/godotenv"
)

func RunCLI(args []string) int {
	// Attempt to load .env file from the current directory, ignoring errors if it doesn't exist.
	_ = godotenv.Load()
	if len(args) < 1 {
		runREPL()
		return 0
	}

	target := args[0]

	if target == "--help" || target == "-h" {
		printHelp()
		return 0
	}

	if target == "--version" || target == "-v" {
		fmt.Printf("dryLang version v%s\n", strings.TrimPrefix(Version, "v"))
		return 0
	}

	if target == "init" {
		handleInit(args[1:])
		return 0
	}

	var files []string

	if target == "all" {
		// Run all .y files in current directory
		entries, err := os.ReadDir(".")
		if err != nil {
			die(err)
		}
		for _, e := range entries {
			if !e.IsDir() && (strings.HasSuffix(e.Name(), ".dry") || strings.HasSuffix(e.Name(), ".y")) {
				files = append(files, e.Name())
			}
		}
		if len(files) == 0 {
			fmt.Println("0:0 no .dry or .y files")
			return 1
		}
	} else if strings.Contains(target, ",") {
		// Multiple files: y file1.dry,file2.y
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
					if !e.IsDir() && (strings.HasSuffix(e.Name(), ".dry") || strings.HasSuffix(e.Name(), ".y")) {
						files = append(files, filepath.Join(target, e.Name()))
					}
				}
			} else {
				files = []string{target}
			}
		}
	}

	ldr := loader.New()
	var finalProg ast.Program

	for _, file := range files {
		prog, err := ldr.LoadAndParse(file, ".")
		if err != nil {
			die(err)
		}
		finalProg.Stmts = append(finalProg.Stmts, prog.Stmts...)
	}

	if err := runAST(&finalProg); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	return 0
}

func die(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}
