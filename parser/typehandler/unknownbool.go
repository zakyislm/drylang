package typehandler

import (
	"drylang/ast"
	"drylang/core"
	"drylang/lexer"
)

func ParseUnknownBool(p core.ParserCore) (ast.Stmt, error) {
	tok := p.Advance() // consume '?'

	ident, err := p.Expect(lexer.TOKEN_IDENT)
	if err != nil {
		return nil, err
	}

	return &ast.UnknownBoolStmt{Name: ident.Literal, Line: tok.Line, Col: tok.Col}, nil
}
