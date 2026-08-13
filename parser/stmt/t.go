package stmt

import (
	"drylang/ast"
	"drylang/core"
	"drylang/lexer"
)

func ParseTry(p core.ParserCore) (ast.Stmt, error) {
	tok := p.Advance() // consume try
	body, err := p.ParseBlock()
	if err != nil {
		return nil, err
	}

	p.SkipSemicolons()

	if p.Current().Type != lexer.TOKEN_ERR {
		return nil, p.Errorf("E106", "needs err")
	}
	p.Advance() // consume err

	if _, err := p.Expect(lexer.TOKEN_LPAREN); err != nil {
		return nil, err
	}
	errName, err := p.Expect(lexer.TOKEN_IDENT)
	if err != nil {
		return nil, err
	}
	if _, err := p.Expect(lexer.TOKEN_RPAREN); err != nil {
		return nil, err
	}

	catchBody, err := p.ParseBlock()
	if err != nil {
		return nil, err
	}

	return &ast.TryStmt{Body: body, ErrName: errName.Literal, Catch: catchBody, Line: tok.Line, Col: tok.Col}, nil
}

