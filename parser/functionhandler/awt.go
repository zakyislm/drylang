package functionhandler

import (
	"drylang/ast"
	"drylang/core"
)

func ParseAwt(p core.ParserCore) (ast.Expr, error) {
	tok := p.Advance() // Consume 'awt'
	val, err := p.ParseExpression(core.PREC_UNARY)
	if err != nil {
		return nil, err
	}
	return &ast.AwaitExpr{Value: val, Line: tok.Line, Col: tok.Col}, nil
}
