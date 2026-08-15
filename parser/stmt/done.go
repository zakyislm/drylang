package stmt

import (
	"drylang/ast"
	"drylang/core"
)

func ParseDone(p core.ParserCore) (ast.Stmt, error) {
	tok := p.Advance()
	return &ast.DoneStmt{Line: tok.Line, Col: tok.Col}, nil
}
