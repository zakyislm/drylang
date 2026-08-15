package cli

import (
	"drylang/analyzer"
	"drylang/ast"
	"drylang/compiler"
	"drylang/errfmt"
	"drylang/lexer"
	"drylang/parser"
	"drylang/vm"
)

func runAST(prog *ast.Program) error {
	if err := analyzer.Analyze(prog); err != nil {
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

	return runAST(prog)
}
