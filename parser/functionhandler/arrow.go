package functionhandler

import (
	"drylang/ast"
	"drylang/core"
	"drylang/lexer"
)

func ParseArrowFn(p core.ParserCore) (ast.Expr, error) {
	tok := p.Advance() // Consume '=>'

	var params []string
	if p.Current().Type == lexer.TOKEN_LPAREN {
		pList, err := p.ParseParamList()
		if err != nil {
			return nil, err
		}
		params = pList
	} else if p.Current().Type == lexer.TOKEN_IDENT {
		identTok, _ := p.Expect(lexer.TOKEN_IDENT)
		params = append(params, identTok.Literal)
	} else {
		return nil, p.Errorf("E107", "arrow function needs params")
	}

	var body []ast.Stmt
	if p.Current().Type == lexer.TOKEN_LBRACE {
		b, err := p.ParseBlock()
		if err != nil {
			return nil, err
		}
		body = b
	} else {
		// Single expression body
		expr, err := p.ParseExpression(core.PREC_LOWEST)
		if err != nil {
			return nil, err
		}
		// Convert expression to return statement or expr statement
		// For arrow functions, an expression body usually acts as a return
		body = []ast.Stmt{&ast.ReturnStmt{Value: expr, Line: expr.TokenLine(), Col: expr.TokenCol()}}
	}

	return &ast.ArrowFnExpr{Params: params, Body: body, Line: tok.Line, Col: tok.Col}, nil
}
