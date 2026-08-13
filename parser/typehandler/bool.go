package typehandler

import (
	"drylang/ast"
	"drylang/core"
)

func ParseBool(p core.ParserCore) (ast.Expr, error) {
	tok := p.Advance()
	val := tok.Literal == "t"
	return &ast.BoolLit{Value: val, Line: tok.Line, Col: tok.Col}, nil
}
