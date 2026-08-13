package stmt

import (
	"drylang/ast"
	"drylang/core"
	"drylang/lexer"
)

func ParseIf(p core.ParserCore) (ast.Stmt, error) {
	tok := p.Advance() // consume if
	cond, err := p.ParseExpression(core.PREC_LOWEST)
	if err != nil {
		return nil, err
	}

	body, err := p.ParseBlock()
	if err != nil {
		return nil, err
	}

	stmt := &ast.IfStmt{Condition: cond, Body: body, Line: tok.Line, Col: tok.Col}

	p.SkipSemicolons()

	// Parse elif chains
	for p.Current().Type == lexer.TOKEN_ELIF {
		p.Advance()
		elifCond, err := p.ParseExpression(core.PREC_LOWEST)
		if err != nil {
			return nil, err
		}
		elifBody, err := p.ParseBlock()
		if err != nil {
			return nil, err
		}
		stmt.ElIfs = append(stmt.ElIfs, ast.ElIfClause{Condition: elifCond, Body: elifBody})
		p.SkipSemicolons()
	}

	// Parse el (else)
	if p.Current().Type == lexer.TOKEN_EL {
		p.Advance()
		elseBody, err := p.ParseBlock()
		if err != nil {
			return nil, err
		}
		stmt.Else = elseBody
	}

	return stmt, nil
}

