package expr

import (
	"drylang/ast"
	"drylang/core"
	"drylang/lexer"
)

func ParseIndexExpr(p core.ParserCore, left ast.Expr) (ast.Expr, error) {
	tok := p.Advance() // Consume '['
	index, err := p.ParseExpression(core.PREC_LOWEST)
	if err != nil {
		return nil, err
	}

	if _, err := p.Expect(lexer.TOKEN_RBRACKET); err != nil {
		return nil, err
	}

	return &ast.IndexExpr{Object: left, Index: index, Line: tok.Line, Col: tok.Col}, nil
}
