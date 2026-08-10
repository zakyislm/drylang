package varhandler

import (
	"drylang/ast"
	"drylang/core"
	"drylang/lexer"
)

func ParseAssignOrExpr(p core.ParserCore) (ast.Stmt, error) {
	expr, err := p.ParseExpression(core.PREC_LOWEST)
	if err != nil {
		return nil, err
	}

	if binExpr, ok := expr.(*ast.BinaryExpr); ok && binExpr.Op == lexer.TOKEN_ASSIGN {
		val := binExpr.Right
		
		if arrLit, ok := binExpr.Left.(*ast.ArrayLit); ok {
			var names []string
			for _, elem := range arrLit.Items {
				if ident, ok := elem.(*ast.Ident); ok {
					names = append(names, ident.Name)
				} else {
					return nil, p.Errorf("E109", "invalid array destructuring target")
				}
			}
			return &ast.DestructArrayStmt{
				Names: names,
				Value: val,
				Line:  arrLit.Line,
				Col:   arrLit.Col,
			}, nil
		}

		if mapLit, ok := binExpr.Left.(*ast.MapLit); ok {
			var keys []string
			// In MapLit, Keys contains Expr (usually Ident or String).
			// We only allow Ident for destructuring to bind to local variables.
			for _, key := range mapLit.Keys {
				if ident, ok := key.(*ast.Ident); ok {
					keys = append(keys, ident.Name)
				} else {
					return nil, p.Errorf("E109", "invalid map destructuring target")
				}
			}
			return &ast.DestructMapStmt{
				Keys:  keys,
				Value: val,
				Line:  mapLit.Line,
				Col:   mapLit.Col,
			}, nil
		}

		if dotExpr, ok := binExpr.Left.(*ast.DotExpr); ok {
			return &ast.DotAssignStmt{
				Object:   dotExpr.Object,
				Field:    dotExpr.Field,
				Value:    val,
				Line:     dotExpr.Line,
				Col:      dotExpr.Col,
			}, nil
		}

		if ident, ok := binExpr.Left.(*ast.Ident); ok {
			return &ast.AssignStmt{
				Name:  ident.Name,
				Value: val,
				Line:  ident.Line,
				Col:   ident.Col,
			}, nil
		}

		if idxExpr, ok := binExpr.Left.(*ast.IndexExpr); ok {
			return &ast.IndexAssignStmt{
				Object: idxExpr.Object,
				Index:  idxExpr.Index,
				Value:  val,
				Line:   idxExpr.Line,
				Col:    idxExpr.Col,
			}, nil
		}

		return nil, p.Errorf("E109", "invalid assignment target: %T = %T", binExpr.Left, binExpr.Right)
	}

	return &ast.ExprStmt{Expression: expr, Line: expr.TokenLine(), Col: expr.TokenCol()}, nil
}
