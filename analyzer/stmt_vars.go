package analyzer

import "drylang/ast"

func (a *analyzer) analyzeAssign(s *ast.AssignStmt) error {
	a.declare(s.Name, s.Line, s.Col, false, false)
	if err := a.analyzeExpr(s.Value); err != nil {
		return err
	}
	a.trackInstance(s.Name, s.Value)
	return nil
}

// trackInstance records s.Name as an instance of a class when the RHS is a
// class/struct constructor call, so pv access can be enforced later.
func (a *analyzer) trackInstance(name string, value ast.Expr) {
	call, ok := value.(*ast.CallExpr)
	if !ok {
		return
	}
	ident, ok := call.Callee.(*ast.Ident)
	if !ok {
		return
	}
	info, declared := a.declared[ident.Name]
	if declared && !info.isParam {
		a.instances[name] = ident.Name
	}
}

func (a *analyzer) analyzeConstDecl(s *ast.ConstDeclStmt) error {
	if info, exists := a.declared[s.Name]; exists && info.isConst {
		return a.errorf(s.Line, s.Col, "locked %s", s.Name)
	}
	a.declare(s.Name, s.Line, s.Col, true, false)
	return a.analyzeExpr(s.Value)
}
