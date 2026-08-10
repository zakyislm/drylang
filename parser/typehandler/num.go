package typehandler

import (
	"drylang/ast"
	"drylang/core"
)

func ParseNum(p core.ParserCore) (ast.Expr, error) {
	tok := p.Advance()
	return &ast.NumberLit{Value: tok.Literal, Line: tok.Line, Col: tok.Col}, nil
}
