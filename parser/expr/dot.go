package expr

import (
	"drylang/ast"
	"drylang/core"
	"drylang/lexer"
)

func ParseDotExpr(p core.ParserCore, object ast.Expr) (ast.Expr, error) {
	tok := p.Advance() // Consume '.' or '?.'
	optional := tok.Type == lexer.TOKEN_QMARK_DOT

	// Property/method name. Keywords are allowed as property names
	// (e.g. rt.on, pipe.close) — the literal text is the field.
	propertyTok := p.Advance()
	if propertyTok.Type == lexer.TOKEN_EOF {
		return nil, p.Errorf("E107", "needs property after dot")
	}
	switch propertyTok.Type {
	case lexer.TOKEN_LPAREN, lexer.TOKEN_RPAREN, lexer.TOKEN_LBRACE,
		lexer.TOKEN_RBRACE, lexer.TOKEN_LBRACKET, lexer.TOKEN_RBRACKET,
		lexer.TOKEN_COMMA, lexer.TOKEN_SEMICOLON, lexer.TOKEN_DOT:
		return nil, p.Errorf("E107", "needs property after dot, got %s", propertyTok.Literal)
	}

	return &ast.DotExpr{Object: object, Field: propertyTok.Literal, Optional: optional, Line: tok.Line, Col: tok.Col}, nil
}
