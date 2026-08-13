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
	return a.analyzeExpr(s.Value)
}
