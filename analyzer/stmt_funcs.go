package analyzer

import "drylang/ast"

func (a *analyzer) analyzeReturn(s *ast.ReturnStmt) error {
	if !a.inFn {
		return a.errorf(s.Line, s.Col, "stray rev")
	}
	if s.Value != nil {
		return a.analyzeExpr(s.Value)
	}
	return nil
}

func (a *analyzer) analyzeFnDecl(s *ast.FnDeclStmt) error {
	a.declare(s.Name, s.Line, s.Col, false, false)
	return a.analyzeFnBody(s.Params, s.Body, s.IsAsync)
}

func (a *analyzer) analyzeFnBody(params []string, body []ast.Stmt, isAsync bool) error {
	prevFn := a.inFn
	prevAsync := a.inAsync
	a.inFn = true
	a.inAsync = isAsync

	seenParams := make(map[string]bool)
	for _, p := range params {
		if seenParams[p] {
			return a.errorf(0, 0, "duplicate parameter")
		}
		seenParams[p] = true
		a.declare(p, 0, 0, false, true)
	}

	for _, stmt := range body {
		if err := a.analyzeStmt(stmt); err != nil {
			return err
		}
	}

	a.inFn = prevFn
	a.inAsync = prevAsync
	return nil
}
