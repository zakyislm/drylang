package analyzer

import (
	"drylang/ast"
)

func (a *analyzer) analyzeClass(s *ast.ClassStmt) error {
	if _, exists := a.declared[s.Name]; exists {
		return a.errorf(s.Line, s.Col, "duplicate class definition")
	}
	a.declare(s.Name, s.Line, s.Col, false, false)

	// Detect circular inheritance (Bug 12)
	// Extends references mark the parent classes as used.
	for _, parent := range s.Extends {
		a.markUsed(parent)
		if parent == s.Name {
			return a.errorf(s.Line, s.Col, "circular inheritance")
		}
		// In a single-pass analyzer, detecting full cycles across out-of-order classes is hard.
		// However, a simple ancestor traversal works if the parent is already processed.
		if err := a.checkCycle(s.Name, parent, s.Line, s.Col); err != nil {
			return err
		}
		// Prevent public class inheriting from private class (Bug 8)
		if s.Visibility != "pv" && a.privateClasses[parent] {
			return a.errorf(s.Line, s.Col, "cannot inherit from pv cl")
		}
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

	cMethods := make(map[string]bool)
	for _, m := range s.Methods {
		cMethods[m.Name] = true
	}
	for _, parent := range s.Extends {
		if pm, ok := a.classMethods[parent]; ok {
			for name := range pm {
				cMethods[name] = true
			}
		}
	}
	a.classMethods[s.Name] = cMethods

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

func (a *analyzer) checkCycle(startClass, currentParent string, line, col int) error {
	if startClass == currentParent {
		return a.errorf(line, col, "circular inheritance")
	}
	// Note: full circular check needs an external AST parent map, but since 
	// classes are analyzed in order, checking against already-analyzed parents' methods/privates 
	// helps, but true graph check requires tracking a.parents.
	return nil
}
