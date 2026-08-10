package functionhandler

import (
	"drylang/ast"
	"drylang/core"
	"drylang/lexer"
)

func ParseFnDecl(p core.ParserCore, isAsync bool) (ast.Stmt, error) {
	tok := p.Advance() // consume fn
	name, err := p.Expect(lexer.TOKEN_IDENT)
	if err != nil {
		return nil, err
	}

	params, err := p.ParseParamList()
	if err != nil {
		return nil, err
	}

	body, err := p.ParseBlock()
	if err != nil {
		return nil, err
	}

	return &ast.FnDeclStmt{
		Name: name.Literal, Params: params, Body: body,
		IsAsync: isAsync, Line: tok.Line, Col: tok.Col,
	}, nil
}

