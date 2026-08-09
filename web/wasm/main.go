//go:build js && wasm

package main

import (
	"syscall/js"

	"drylang/compiler"
	"drylang/lexer"
	"drylang/parser"
	"drylang/vm"
)

func rundryLang(this js.Value, args []js.Value) interface{} {
	if len(args) == 0 {
		return "Error: no code provided"
	}
	source := args[0].String()
	err := run(source)
	if err != nil {
		return err.Error()
	}
	return nil
}

func run(source string) error {
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
	return machine.Run()
}

func main() {
	js.Global().Set("rundryLang", js.FuncOf(rundryLang))
	
	// Keep the Wasm binary running so we can call rundryLang multiple times
	<-make(chan bool)
}
