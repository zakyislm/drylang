package colshandler

import (
	"drylang/ast"
	"drylang/core"
	"drylang/lexer"
)

func ParseArrayLit(p core.ParserCore) (ast.Expr, error) {
	tok := p.Advance() // Consume '['
	var elements []ast.Expr

	if p.Current().Type != lexer.TOKEN_RBRACKET {
		expr, err := p.ParseExpression(core.PREC_LOWEST)
		if err != nil {
			return nil, err
		}
		elements = append(elements, expr)

		for p.Current().Type == lexer.TOKEN_COMMA {
			p.Advance() // Consume ','
			expr, err := p.ParseExpression(core.PREC_LOWEST)
			if err != nil {
				return nil, err
			}
			elements = append(elements, expr)
		}
	}

	if _, err := p.Expect(lexer.TOKEN_RBRACKET); err != nil {
		return nil, err
	}

	return &ast.ArrayLit{Items: elements, Line: tok.Line, Col: tok.Col}, nil
}
