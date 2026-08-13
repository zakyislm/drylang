package analyzer

import (
	"drylang/ast"
)

func (a *analyzer) analyzeClass(s *ast.ClassStmt) error {
	a.declare(s.Name, s.Line, s.Col, false, false)

	// Extends references mark the parent classes as used.
	for _, parent := range s.Extends {
		a.markUsed(parent)
	}

	// Register private fields and methods of this class, including those
	// inherited from parents so pv access on a child instance is caught.
	priv := make(map[string]bool)
	for _, parent := range s.Extends {
		if pp, ok := a.privates[parent]; ok {
			for name := range pp {
				priv[name] = true
			}
		}
	}
	for _, f := range s.PrivateFields {
		priv[f] = true
	}
	for _, m := range s.Methods {
		if m.Visibility == "pv" {
			priv[m.Name] = true
		}
	}
	a.privates[s.Name] = priv

	// pv cl marks the class itself private.
	if s.Visibility == "pv" {
		a.privateClasses[s.Name] = true
	}

	prevClass := a.currentClass
	a.currentClass = s.Name

	for _, m := range s.Methods {
		// 'this' is an implicit parameter.
		params := append([]string{"this"}, m.Params...)
		if err := a.analyzeFnBody(params, m.Body, m.IsAsync); err != nil {
			return err
		}
	}

	a.currentClass = prevClass
	return nil
}

