package analyzer

import "drylang/ast"

func (a *analyzer) analyzeIf(s *ast.IfStmt) error {
	if err := a.analyzeExpr(s.Condition); err != nil {
		return err
	}
	for _, stmt := range s.Body {
		if err := a.analyzeStmt(stmt); err != nil {
			return err
		}
	}
	for _, elif := range s.ElIfs {
		if err := a.analyzeExpr(elif.Condition); err != nil {
			return err
		}
		for _, stmt := range elif.Body {
			if err := a.analyzeStmt(stmt); err != nil {
				return err
			}
		}
	}
	for _, stmt := range s.Else {
		if err := a.analyzeStmt(stmt); err != nil {
			return err
		}
	}
	return nil
}

func (a *analyzer) analyzeOn(s *ast.OnStmt) error {
	if err := a.analyzeExpr(s.Value); err != nil {
		return err
	}
	for _, c := range s.Cases {
		if err := a.analyzeExpr(c.Value); err != nil {
			return err
		}
		for _, stmt := range c.Body {
			if err := a.analyzeStmt(stmt); err != nil {
				return err
			}
		}
	}
	return nil
}

func (a *analyzer) analyzeLoop(s *ast.LoopStmt) error {
	if s.Limit != nil {
		if err := a.analyzeExpr(s.Limit); err != nil {
			return err
		}
	}
	prev := a.inLoop
	a.inLoop = true
	// Loop counter 'i' is implicitly declared
	a.declare("i", s.Line, s.Col, false, true)
	for _, stmt := range s.Body {
		if err := a.analyzeStmt(stmt); err != nil {
			return err
		}
	}
	a.inLoop = prev
	return nil
}
