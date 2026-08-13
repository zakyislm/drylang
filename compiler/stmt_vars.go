package compiler

import (
	"drylang/ast"
	"drylang/core"
)

func (c *Compiler) compileAssign(s *ast.AssignStmt) error {
	if err := c.compileExpr(s.Value); err != nil {
		return err
	}
	idx := c.resolveLocal(s.Name)
	if idx >= 0 {
		c.emit(core.OpSetLocal, idx, s.Line, s.Col)
	} else if c.globals[s.Name] {
		ci := c.addConst(s.Name)
		c.emit(core.OpSetGlobal, ci, s.Line, s.Col)
	} else if c.depth > 0 {
		slot := c.addLocal(s.Name)
		c.emit(core.OpSetLocal, slot, s.Line, s.Col)
	} else {
		c.globals[s.Name] = true
		ci := c.addConst(s.Name)
		c.emit(core.OpSetGlobal, ci, s.Line, s.Col)
	}
	return nil
}

func (c *Compiler) compileConstDecl(s *ast.ConstDeclStmt) error {
	if err := c.compileExpr(s.Value); err != nil {
		return err
	}
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

func (c *Compiler) compileUnknownBool(s *ast.UnknownBoolStmt) error {
	c.emit(core.OpUnknown, 0, s.Line, s.Col)
	if c.depth > 0 {
		c.addLocal(s.Name)
	} else {
		c.globals[s.Name] = true
		ci := c.addConst(s.Name)
		c.emit(core.OpSetGlobal, ci, s.Line, s.Col)
	}
	return nil
}

func (c *Compiler) compileExprStmt(s *ast.ExprStmt) error {
	if err := c.compileExpr(s.Expression); err != nil {
		return err
	}
	c.emit(core.OpPop, 0, s.Line, s.Col)
	return nil
}

func (c *Compiler) compileDestructArray(s *ast.DestructArrayStmt) error {
	if err := c.compileExpr(s.Value); err != nil {
		return err
	}

	if c.depth > 0 {
		tmpSlot := c.addLocal("")
		c.emit(core.OpSetLocal, tmpSlot, s.Line, s.Col)

		for i, name := range s.Names {
			c.emit(core.OpGetLocal, tmpSlot, s.Line, s.Col)
			c.emit(core.OpConst, c.addConst(float64(i)), s.Line, s.Col)
			c.emit(core.OpIndex, 0, s.Line, s.Col)

			idx := c.resolveLocal(name)
			if idx >= 0 {
				c.emit(core.OpSetLocal, idx, s.Line, s.Col)
			} else {
				slot := c.addLocal(name)
				c.emit(core.OpSetLocal, slot, s.Line, s.Col)
			}
		}
	} else {
		ci := c.addConst("__destruct__")
		c.emit(core.OpSetGlobal, ci, s.Line, s.Col)

		for i, name := range s.Names {
			c.emit(core.OpGetGlobal, ci, s.Line, s.Col)
			c.emit(core.OpConst, c.addConst(float64(i)), s.Line, s.Col)
			c.emit(core.OpIndex, 0, s.Line, s.Col)

			c.globals[name] = true
			nameConst := c.addConst(name)
			c.emit(core.OpSetGlobal, nameConst, s.Line, s.Col)
		}
	}
	return nil
}

func (c *Compiler) compileDestructMap(s *ast.DestructMapStmt) error {
	if err := c.compileExpr(s.Value); err != nil {
		return err
	}

	if c.depth > 0 {
		tmpSlot := c.addLocal("")
		c.emit(core.OpSetLocal, tmpSlot, s.Line, s.Col)

		for _, key := range s.Keys {
			c.emit(core.OpGetLocal, tmpSlot, s.Line, s.Col)
			c.emit(core.OpConst, c.addConst(key), s.Line, s.Col)
			c.emit(core.OpIndex, 0, s.Line, s.Col)

			idx := c.resolveLocal(key)
			if idx >= 0 {
				c.emit(core.OpSetLocal, idx, s.Line, s.Col)
			} else {
				slot := c.addLocal(key)
				c.emit(core.OpSetLocal, slot, s.Line, s.Col)
			}
		}
	} else {
		ci := c.addConst("__destruct__")
		c.emit(core.OpSetGlobal, ci, s.Line, s.Col)

		for _, key := range s.Keys {
			c.emit(core.OpGetGlobal, ci, s.Line, s.Col)
			c.emit(core.OpConst, c.addConst(key), s.Line, s.Col)
			c.emit(core.OpIndex, 0, s.Line, s.Col)

			c.globals[key] = true
			nameConst := c.addConst(key)
			c.emit(core.OpSetGlobal, nameConst, s.Line, s.Col)
		}
	}
	return nil
}

func (c *Compiler) compileIndexAssign(s *ast.IndexAssignStmt) error {
	if err := c.compileExpr(s.Object); err != nil {
		return err
	}
	if err := c.compileExpr(s.Index); err != nil {
		return err
	}
	if err := c.compileExpr(s.Value); err != nil {
		return err
	}
	c.emit(core.OpSetIndex, 0, s.Line, s.Col)
	return nil
}

func (c *Compiler) compileDotAssign(s *ast.DotAssignStmt) error {
	if err := c.compileExpr(s.Object); err != nil {
		return err
	}
	fi := c.addConst(s.Field)
	if err := c.compileExpr(s.Value); err != nil {
		return err
	}
	c.emit(core.OpDotSet, fi, s.Line, s.Col)
	return nil
}
