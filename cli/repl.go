package cli

import (
	"bufio"
	"drylang/analyzer"
	"drylang/compiler"
	"drylang/errfmt"
	"drylang/lexer"
	"drylang/parser"
	"drylang/vm"
	"fmt"
	"os"
	"strings"
)

// runREPL starts an interactive Read-Eval-Print Loop.
func runREPL() {
	fmt.Println("dryLang REPL v0.1.0")
	fmt.Println("Type your code and press Enter. (Ctrl+C to exit)")

	scanner := bufio.NewScanner(os.Stdin)

	// Create persistent VM instance
	machine := vm.New(nil, nil)

	var buffer strings.Builder
	prompt := ">>> "

	for {
		fmt.Print(prompt)
		if !scanner.Scan() {
			break // EOF
		}

		line := scanner.Text()

		// If line is empty and we aren't buffering, skip
		if strings.TrimSpace(line) == "" && buffer.Len() == 0 {
			continue
		}

		if buffer.Len() == 0 && (strings.TrimSpace(line) == "exit" || strings.TrimSpace(line) == "quit") {
			break
		}

		buffer.WriteString(line)
		buffer.WriteString("\n")

		source := buffer.String()

		// Setup error formatting for the current snippet
		errfmt.Init(source)

		// Create a fresh compiler for this snippet
		comp := compiler.New()

		// 1. Lex
		lex := lexer.New(source)
		tokens, err := lex.Tokenize()
		if err != nil {
			// If it's a "needs closing X" error and we're at EOF, we wait for more input
			if strings.Contains(err.Error(), "needs closing") {
				prompt = "... "
				continue
			}
			fmt.Println(err)
			buffer.Reset()
			prompt = ">>> "
			continue
		}

		// 2. Parse
		p := parser.New(tokens)
		prog, err := p.Parse()
		if err != nil {
			if strings.Contains(err.Error(), "needs }") || strings.Contains(err.Error(), "needs )") || strings.Contains(err.Error(), "needs ]") {
				prompt = "... "
				continue
			}
			fmt.Println(err)
			buffer.Reset()
			prompt = ">>> "
			continue
		}

		// 3. Analyze
		if err := analyzer.Analyze(prog); err != nil {
			fmt.Println(err)
			buffer.Reset()
			prompt = ">>> "
			continue
		}

		// 4. Compile
		chunk, fns, err := comp.Compile(prog)
		if err != nil {
			fmt.Println(err)
			buffer.Reset()
			prompt = ">>> "
			continue
		}

		// 4. Run (only if there are instructions)
		if len(chunk.Code) > 0 {
			machine.Update(chunk, fns)
			if err := machine.Run(); err != nil {
				fmt.Println(err)
			}
		}

		// Successfully compiled and ran, reset buffer
		buffer.Reset()
		prompt = ">>> "
	}
}
