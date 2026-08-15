package compiler

import (
	"drylang/ast"
	"drylang/core"
)

func (c *Compiler) compileFnDecl(s *ast.FnDeclStmt) error {
	fnCompiler := New()
	fnCompiler.globals = c.globals
	fnCompiler.depth = 1
	for _, p := range s.Params {
		fnCompiler.addLocal(p)
	}
	for _, stmt := range s.Body {
		if err := fnCompiler.compileStmt(stmt); err != nil {
			return err
		}
	}
	fnCompiler.emit(core.OpUnknown, 0, s.Line, s.Col) // default return unknown
	fnCompiler.emit(core.OpReturn, 0, s.Line, s.Col)
	fnCompiler.chunk.LocalsCount = fnCompiler.maxLocals

	fn := &core.CompiledFn{
		Chunk:      fnCompiler.chunk,
		Name:       s.Name,
		ParamCount: len(s.Params),
		IsAsync:    s.IsAsync,
		LocalNames: fnCompiler.slotNames,
	}

	c.fns = append(c.fns, fn)
	fi := c.addConst(fn)
	c.emit(core.OpClosure, fi, s.Line, s.Col)

	if c.depth > 0 {
		slot := c.addLocal(s.Name)
		c.emit(core.OpSetLocal, slot, s.Line, s.Col)
	} else {
		c.globals[s.Name] = true
		ci := c.addConst(s.Name)
		c.emit(core.OpSetGlobal, ci, s.Line, s.Col)
	}

	return nil
}
