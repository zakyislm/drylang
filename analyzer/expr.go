package analyzer

import (
	"drylang/ast"
	"drylang/core"
)

func (a *analyzer) analyzeExpr(expr ast.Expr) error {
	switch e := expr.(type) {
	case *ast.Ident:
		if err := a.checkIdent(e); err != nil {
			return err
		}
	case *ast.BinaryExpr:
		if err := a.analyzeExpr(e.Left); err != nil {
			return err
		}
		return a.analyzeExpr(e.Right)
	case *ast.UnaryExpr:
		return a.analyzeExpr(e.Operand)
	case *ast.CallExpr:
		if err := a.analyzeExpr(e.Callee); err != nil {
			return err
		}
		// pv cl instantiation is allowed globally (same module context because ASTs are merged)
		for _, arg := range e.Args {
			if err := a.analyzeExpr(arg); err != nil {
				return err
			}
		}
	case *ast.IndexExpr:
		if err := a.analyzeExpr(e.Object); err != nil {
			return err
		}
		return a.analyzeExpr(e.Index)
	case *ast.DotExpr:
		if err := a.analyzeExpr(e.Object); err != nil {
			return err
		}
		return a.checkPv(e.Object, e.Field, e.Line, e.Col)
	case *ast.ArrayLit:
		for _, item := range e.Items {
			if err := a.analyzeExpr(item); err != nil {
				return err
			}
		}
	case *ast.MapLit:
		for i := range e.Keys {
			if err := a.analyzeExpr(e.Keys[i]); err != nil {
				return err
			}
			if err := a.analyzeExpr(e.Values[i]); err != nil {
				return err
			}
		}
	case *ast.ArrowFnExpr:
		return a.analyzeFnBody(e.Params, e.Body, false)
	case *ast.StringInterp:
		for _, part := range e.Parts {
			if err := a.analyzeExpr(part); err != nil {
				return err
			}
		}
	}
	return nil
}

// checkIdent validates an identifier reference: builtins, params, loop
// counters, 'this', and declared names are allowed; anything else is an
// undefined variable.
func (a *analyzer) checkIdent(e *ast.Ident) error {
	name := e.Name
	if _, isBuiltin := core.BuiltinNames[name]; isBuiltin {
		a.markUsed(name)
		return nil
	}
	if info, ok := a.declared[name]; ok {
		if !info.isParam {
			a.markUsed(name)
		}
		return nil
	}
	if name == "i" && a.inLoop {
		return nil
	}
	if name == "this" {
		return nil
	}
	return a.errorf(e.Line, e.Col, "undefined %s", name)
}

// checkPv enforces private field/method access. Access to a private member
// is only allowed from within the class that declares it.
func (a *analyzer) checkPv(object ast.Expr, field string, line, col int) error {
	ident, ok := object.(*ast.Ident)
	if !ok {
		return nil
	}
	// Builtin modules (math, jwt, ...) expose public methods only.
	if _, isBuiltin := core.BuiltinNames[ident.Name]; isBuiltin {
		return nil
	}

	// Determine the class of the object.
	var className string
	if ident.Name == "this" {
		className = a.currentClass
	} else if cls, ok := a.instances[ident.Name]; ok {
		className = cls
	} else {
		return nil // unknown object type; allow
	}
	if className == "" {
		return nil
	}

	if priv := a.privates[className]; priv[field] && className != a.currentClass {
		return a.errorf(line, col, "pv access %s.%s", ident.Name, field)
	}
	return nil
}
