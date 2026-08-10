package functionhandler

import (
	"drylang/ast"
	"drylang/core"
	"drylang/lexer"
)

func ParseCallExpr(p core.ParserCore, function ast.Expr) (ast.Expr, error) {
	tok := p.Advance() // Consume '('
	var args []ast.Expr

	if p.Current().Type != lexer.TOKEN_RPAREN {
		expr, err := p.ParseExpression(core.PREC_LOWEST)
		if err != nil {
			return nil, err
		}
		args = append(args, expr)

		for p.Current().Type == lexer.TOKEN_COMMA {
			p.Advance() // Consume ','
			expr, err := p.ParseExpression(core.PREC_LOWEST)
			if err != nil {
				return nil, err
			}
			args = append(args, expr)
		}
	}

	if _, err := p.Expect(lexer.TOKEN_RPAREN); err != nil {
		return nil, err
	}

	return &ast.CallExpr{Callee: function, Args: args, Line: tok.Line, Col: tok.Col}, nil
}
