package expr

import (
	"drylang/ast"
	"drylang/core"
	"drylang/lexer"
)

// ParseArrowFnPrefix parses an arrow function starting at '('.
// Form: (a, b) -> body  |  () -> body
func ParseArrowFnPrefix(p core.ParserCore) (ast.Expr, error) {
	tok := p.Current()
	params, err := p.ParseParamList()
	if err != nil {
		return nil, err
	}

	if _, err := p.Expect(lexer.TOKEN_ARROW); err != nil {
		return nil, err
	}

	body, err := parseArrowBody(p)
	if err != nil {
		return nil, err
	}

	return &ast.ArrowFnExpr{Params: params, Body: body, Line: tok.Line, Col: tok.Col}, nil
}

// ParseArrowFnInfix parses an arrow function where the left side is a
// single identifier (the parameter) and the current token is '->'.
// Form: x -> body
func ParseArrowFnInfix(p core.ParserCore, left ast.Expr) (ast.Expr, error) {
	arrow := p.Advance() // consume '->'

	params, ok := paramNames(left)
	if !ok {
		return nil, p.Errorf("E109", "invalid arrow parameter: %s", arrow.Literal)
	}

	body, err := parseArrowBody(p)
	if err != nil {
		return nil, err
	}

	return &ast.ArrowFnExpr{Params: params, Body: body, Line: arrow.Line, Col: arrow.Col}, nil
}

// paramNames extracts parameter names from the arrow's left-hand side.
func paramNames(e ast.Expr) ([]string, bool) {
	switch v := e.(type) {
	case *ast.Ident:
		return []string{v.Name}, true
	case *ast.ArrayLit:
		var names []string
		for _, item := range v.Items {
			ident, ok := item.(*ast.Ident)
			if !ok {
				return nil, false
			}
			names = append(names, ident.Name)
		}
		return names, true
	}
	return nil, false
}

// parseArrowBody parses the body after '->': either a block or an expression.
func parseArrowBody(p core.ParserCore) ([]ast.Stmt, error) {
	if p.Current().Type == lexer.TOKEN_LBRACE {
		return p.ParseBlock()
	}

	expr, err := p.ParseExpression(core.PREC_LOWEST)
	if err != nil {
		return nil, err
	}
	return []ast.Stmt{&ast.ReturnStmt{Value: expr, Line: expr.TokenLine(), Col: expr.TokenCol()}}, nil
}
