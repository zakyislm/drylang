package stmt

import (
	"drylang/ast"
	"drylang/core"
	"drylang/lexer"
)

func ParseLoop(p core.ParserCore) (ast.Stmt, error) {
	tok := p.Advance() // consume lp
	var limit ast.Expr

	// lp N { ... } or lp { ... }
	if p.Current().Type != lexer.TOKEN_LBRACE {
		var err error
		limit, err = p.ParseExpression(core.PREC_LOWEST)
		if err != nil {
			return nil, err
		}
	}

	body, err := p.ParseBlock()
	if err != nil {
		return nil, err
	}

	return &ast.LoopStmt{Limit: limit, Body: body, Line: tok.Line, Col: tok.Col}, nil
}

