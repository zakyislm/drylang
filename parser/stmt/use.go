package stmt

import (
	"drylang/ast"
	"drylang/core"
	"drylang/lexer"
)

func ParseUse(p core.ParserCore) (ast.Stmt, error) {
	tok := p.Advance()
	if p.GetBlockDepth() > 0 {
		return nil, p.Errorf("E109", "use must be at top level")
	}
	path, err := p.Expect(lexer.TOKEN_STRING)
	if err != nil {
		return nil, err
	}
	return &ast.UseStmt{Path: path.Literal, Line: tok.Line, Col: tok.Col}, nil
}

