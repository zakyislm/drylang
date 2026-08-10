package typehandler

import (
	"drylang/ast"
	"drylang/core"
)

func ParseUnknown(p core.ParserCore) (ast.Expr, error) {
	tok := p.Advance()
	return &ast.UnknownLit{Line: tok.Line, Col: tok.Col}, nil
}
