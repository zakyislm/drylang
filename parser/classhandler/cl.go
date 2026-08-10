package classhandler

import (
	"drylang/ast"
	"drylang/core"
	"drylang/lexer"
)

func ParseClass(p core.ParserCore) (ast.Stmt, error) {
	tok := p.Advance() // consume cl
	name, err := p.Expect(lexer.TOKEN_IDENT)
	if err != nil {
		return nil, err
	}

	var extends []string
	if p.Current().Type == lexer.TOKEN_LARROW {
		p.Advance()
		for {
			parentName, err := p.Expect(lexer.TOKEN_IDENT)
			if err != nil {
				return nil, err
			}
			extends = append(extends, parentName.Literal)
			if p.Current().Type == lexer.TOKEN_COMMA {
				p.Advance()
			} else {
				break
			}
		}
	}

	if _, err := p.Expect(lexer.TOKEN_LBRACE); err != nil {
		return nil, err
	}
	p.SkipSemicolons()

	var fields []string
	var methods []*ast.MethodDecl

	for p.Current().Type != lexer.TOKEN_RBRACE && p.Current().Type != lexer.TOKEN_EOF {
		vis := "" // default public
		if p.Current().Type == lexer.TOKEN_PV {
			vis = "pv"
			p.Advance()
		}

		if p.Current().Type == lexer.TOKEN_IDENT {
			fields = append(fields, p.Current().Literal) // In class, fields must be declared too
			p.Advance()
		} else if p.Current().Type == lexer.TOKEN_FN || p.Current().Type == lexer.TOKEN_ASN {
			isAsync := false
			if p.Current().Type == lexer.TOKEN_ASN {
				isAsync = true
				p.Advance()
			}
			fnTok := p.Current()
			if _, err := p.Expect(lexer.TOKEN_FN); err != nil {
				return nil, err
			}
			methodName, err := p.Expect(lexer.TOKEN_IDENT)
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

			methods = append(methods, &ast.MethodDecl{
				Name: methodName.Literal, Params: params, Body: body, Visibility: vis, IsAsync: isAsync, Line: fnTok.Line, Col: fnTok.Col,
			})
		} else {
			return nil, p.Errorf("E101", "expected class field or method")
		}
		p.SkipSemicolons()
	}
	if _, err := p.Expect(lexer.TOKEN_RBRACE); err != nil {
		return nil, err
	}

	return &ast.ClassStmt{Name: name.Literal, Extends: extends, Fields: fields, Methods: methods, Line: tok.Line, Col: tok.Col}, nil
}

