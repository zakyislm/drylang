package classhandler

import (
	"drylang/ast"
	"drylang/core"
	"drylang/lexer"
)

func ParseDotExpr(p core.ParserCore, object ast.Expr) (ast.Expr, error) {
	tok := p.Advance() // Consume '.' or '?.'
	optional := tok.Type == lexer.TOKEN_QMARK_DOT
	
	propertyTok, err := p.Expect(lexer.TOKEN_IDENT)
	if err != nil {
		return nil, err
	}
	
	return &ast.DotExpr{Object: object, Field: propertyTok.Literal, Optional: optional, Line: tok.Line, Col: tok.Col}, nil
}
