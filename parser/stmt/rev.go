package stmt

import (
	"drylang/ast"
	"drylang/core"
	"drylang/lexer"
)

func ParseReturn(p core.ParserCore) (ast.Stmt, error) {
	tok := p.Advance() // consume rev
	var val ast.Expr

	// rev can be without value (returns unknown)
	if p.Current().Type != lexer.TOKEN_SEMICOLON && p.Current().Type != lexer.TOKEN_RBRACE && p.Current().Type != lexer.TOKEN_EOF {
		var err error
		val, err = p.ParseExpression(core.PREC_LOWEST)
		if err != nil {
			return nil, err
		}
	}

	return &ast.ReturnStmt{Value: val, Line: tok.Line, Col: tok.Col}, nil
}

