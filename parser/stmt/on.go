package stmt

import (
	"drylang/ast"
	"drylang/core"
	"drylang/lexer"
)

func ParseOn(p core.ParserCore) (ast.Stmt, error) {
	tok := p.Advance() // consume on

	if _, err := p.Expect(lexer.TOKEN_LPAREN); err != nil {
		return nil, err
	}
	val, err := p.ParseExpression(core.PREC_LOWEST)
	if err != nil {
		return nil, err
	}
	if _, err := p.Expect(lexer.TOKEN_RPAREN); err != nil {
		return nil, err
	}

	if _, err := p.Expect(lexer.TOKEN_LBRACE); err != nil {
		return nil, err
	}
	p.SkipSemicolons()

	stmt := &ast.OnStmt{Value: val, Line: tok.Line, Col: tok.Col}

	for p.Current().Type != lexer.TOKEN_RBRACE && p.Current().Type != lexer.TOKEN_EOF {
		caseVal, err := p.ParseExpression(core.PREC_LOWEST)
		if err != nil {
			return nil, err
		}
		caseBody, err := p.ParseBlock()
		if err != nil {
			return nil, err
		}
		stmt.Cases = append(stmt.Cases, ast.OnCase{Value: caseVal, Body: caseBody})
		p.SkipSemicolons()
	}

	if _, err := p.Expect(lexer.TOKEN_RBRACE); err != nil {
		return nil, err
	}

	return stmt, nil
}

