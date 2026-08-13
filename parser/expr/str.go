package expr

import (
	"drylang/ast"
	"drylang/core"
)

func ParseStr(p core.ParserCore) (ast.Expr, error) {
	tok := p.Advance()
	return &ast.StringLit{Value: tok.Literal, Line: tok.Line, Col: tok.Col}, nil
}
