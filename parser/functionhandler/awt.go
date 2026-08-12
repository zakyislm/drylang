package functionhandler

import (
	"drylang/ast"
	"drylang/core"
)

func ParseAwt(p core.ParserCore) (ast.Stmt, error) {
	tok := p.Advance() // Consume 'awt'
	return &ast.AwaitStmt{Line: tok.Line, Col: tok.Col}, nil
}
