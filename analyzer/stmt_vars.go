package analyzer

import (
	"drylang/ast"
	"drylang/core"
	"unicode"
)

func isFullUppercase(s string) bool {
	hasLetter := false
	for _, r := range s {
		if unicode.IsLetter(r) {
			hasLetter = true
			if !unicode.IsUpper(r) {
				return false
			}
		}
	}
	return hasLetter
}

func (a *analyzer) analyzeAssign(s *ast.AssignStmt) error {
	if _, isBuiltin := core.BuiltinNames[s.Name]; isBuiltin {
		return a.errorf(s.Line, s.Col, "cannot assign to builtin")
	}
	if info, exists := a.declared[s.Name]; exists {
		if info.isConst {
			return a.errorf(s.Line, s.Col, "cannot assign to cns")
		}
	} else {
		isConst := isFullUppercase(s.Name)
		a.declare(s.Name, s.Line, s.Col, isConst, false)
	}

	if _, isClass := a.privates[s.Name]; isClass {
		return a.errorf(s.Line, s.Col, "cannot assign to class")
	}

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
