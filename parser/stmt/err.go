package stmt

import (
	"drylang/ast"
	"drylang/core"
)

func ParseThrow(p core.ParserCore) (ast.Stmt, error) {
	tok := p.Advance() // consume err
	val, err := p.ParseExpression(core.PREC_LOWEST)
	if err != nil {
		return nil, err
	}
	return &ast.ThrowStmt{Value: val, Line: tok.Line, Col: tok.Col}, nil
}

