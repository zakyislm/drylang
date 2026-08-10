package varhandler

import (
	"drylang/ast"
	"drylang/core"
	"drylang/lexer"
)

func ParseConstDecl(p core.ParserCore) (ast.Stmt, error) {
	tok := p.Advance() // consume cns
	name, err := p.Expect(lexer.TOKEN_IDENT)
	if err != nil {
		return nil, err
	}

	// Optional =
	if p.Current().Type == lexer.TOKEN_ASSIGN {
		p.Advance()
	}

	val, err := p.ParseExpression(core.PREC_LOWEST)
	if err != nil {
		return nil, err
	}

	return &ast.ConstDeclStmt{Name: name.Literal, Value: val, Line: tok.Line, Col: tok.Col}, nil
}

