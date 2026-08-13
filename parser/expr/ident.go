package expr

import (
	"drylang/ast"
	"drylang/core"
)

func ParseIdent(p core.ParserCore) (ast.Expr, error) {
	tok := p.Advance()
	return &ast.Ident{Name: tok.Literal, Line: tok.Line, Col: tok.Col}, nil
}
