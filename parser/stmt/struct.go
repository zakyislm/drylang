package stmt

import (
	"drylang/ast"
	"drylang/core"
	"drylang/lexer"
)

func ParseStruct(p core.ParserCore) (ast.Stmt, error) {
	// The current token is IDENT (the struct name)
	nameTok := p.Current()
	name := nameTok.Literal

	p.Advance() // Consume IDENT

	// Next should be LBRACE
	if _, err := p.Expect(lexer.TOKEN_LBRACE); err != nil {
		return nil, err
	}

	var fields []string

	// Parse fields
	for p.Current().Type != lexer.TOKEN_RBRACE && p.Current().Type != lexer.TOKEN_EOF {
		// allow skipping commas or semicolons if they somehow get in?
		// Actually, let's just parse IDENT
		if p.Current().Type == lexer.TOKEN_COMMA || p.Current().Type == lexer.TOKEN_SEMICOLON {
			p.Advance()
			continue
		}

		// pv marker on a struct field is accepted (structs are data-only).
		if p.Current().Type == lexer.TOKEN_PV {
			p.Advance()
		}

		fieldTok, err := p.Expect(lexer.TOKEN_IDENT)
		if err != nil {
			return nil, err
		}
		fields = append(fields, fieldTok.Literal)
	}

	if _, err := p.Expect(lexer.TOKEN_RBRACE); err != nil {
		return nil, err
	}

	return &ast.StructDeclStmt{
		Name:   name,
		Fields: fields,
		Line:   nameTok.Line,
		Col:    nameTok.Col,
	}, nil
}
