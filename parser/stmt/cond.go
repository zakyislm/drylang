package stmt

import (
	"drylang/ast"
	"drylang/core"
)

func ParseCon(p core.ParserCore) (ast.Stmt, error) {
	tok := p.Advance()
	return &ast.ConStmt{Line: tok.Line, Col: tok.Col}, nil
}
