package analyzer

import "drylang/ast"

func (a *analyzer) analyzeTry(s *ast.TryStmt) error {
	for _, stmt := range s.Body {
		if err := a.analyzeStmt(stmt); err != nil {
			return err
		}
	}
	a.declare(s.ErrName, s.Line, s.Col, false, true)
	for _, stmt := range s.Catch {
		if err := a.analyzeStmt(stmt); err != nil {
			return err
		}
	}
	return nil
}
