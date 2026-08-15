package analyzer

import "drylang/ast"

func (a *analyzer) analyzeIndexAssign(s *ast.IndexAssignStmt) error {
	if err := a.analyzeExpr(s.Object); err != nil {
		return err
	}
	if err := a.analyzeExpr(s.Index); err != nil {
		return err
	}
	return a.analyzeExpr(s.Value)
}

func (a *analyzer) analyzeDotAssign(s *ast.DotAssignStmt) error {
	if err := a.analyzeExpr(s.Object); err != nil {
		return err
	}
	if err := a.checkPv(s.Object, s.Field, s.Line, s.Col); err != nil {
		return err
	}

	// Prevent method reassignment (Bug 4)
	if ident, ok := s.Object.(*ast.Ident); ok {
		var className string
		if ident.Name == "this" {
			className = a.currentClass
		} else if cls, ok := a.instances[ident.Name]; ok {
			className = cls
		}
		if className != "" && a.classMethods != nil && a.classMethods[className][s.Field] {
			return a.errorf(s.Line, s.Col, "cannot assign to method")
		}
	}

	return a.analyzeExpr(s.Value)
}
