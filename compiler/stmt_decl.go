package compiler

import (
	"drylang/ast"
	"drylang/core"
)

func (c *Compiler) compileStructDecl(s *ast.StructDeclStmt) error {
	// Store struct definition as a constant for runtime
	ci := c.addConst(core.StructDef{Name: s.Name, Fields: s.Fields, Visibility: s.Visibility})
	nameIdx := c.addConst(s.Name)
	c.emit(core.OpConst, ci, s.Line, s.Col)
	c.globals[s.Name] = true
	c.emit(core.OpSetGlobal, nameIdx, s.Line, s.Col)
	return nil
}

func (c *Compiler) compileClass(s *ast.ClassStmt) error {
	methods := make(map[string]core.ClassMethod)
	for _, m := range s.Methods {
		// Compile each method as a core.Chunk
		sub := New()
		sub.depth = 1 // Inside a method, we have local scope + `this`

		// Pre-declare `this` as local 0, which corresponds to the callee slot
		sub.addLocal("this")

		// Pre-declare parameters starting from local 1
		for _, param := range m.Params {
			sub.addLocal(param)
		}

		for _, bStmt := range m.Body {
			if err := sub.compileStmt(bStmt); err != nil {
				return err
			}
		}
		// Implicit return unknown if no return statement
		if len(m.Body) == 0 || !isReturn(m.Body[len(m.Body)-1]) {
			sub.emit(core.OpUnknown, 0, m.Line, m.Col)
			sub.emit(core.OpReturn, 0, m.Line, m.Col)
		}
		sub.chunk.LocalsCount = sub.maxLocals

		methods[m.Name] = core.ClassMethod{
			Chunk:      sub.chunk,
			Visibility: m.Visibility,
			IsAsync:    m.IsAsync,
		}
	}

	ci := c.addConst(core.ClassDef{Name: s.Name, ParentNames: s.Extends, Fields: s.Fields, Methods: methods, Visibility: s.Visibility})
	nameIdx := c.addConst(s.Name)
	c.emit(core.OpConst, ci, s.Line, s.Col)
	c.globals[s.Name] = true
	c.emit(core.OpSetGlobal, nameIdx, s.Line, s.Col)
	return nil
}

func isReturn(stmt ast.Stmt) bool {
	_, ok := stmt.(*ast.ReturnStmt)
	return ok
}
