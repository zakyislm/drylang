package expr

import (
	"drylang/ast"
	"drylang/core"
)

func ParseRawStr(p core.ParserCore) (ast.Expr, error) {
	tok := p.Advance()
	return &ast.RawStringLit{Value: tok.Literal, Line: tok.Line, Col: tok.Col}, nil
}
