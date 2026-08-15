package compiler

import (
	"drylang/ast"
	"drylang/core"
)

func (c *Compiler) compileIf(s *ast.IfStmt) error {
	if err := c.compileExpr(s.Condition); err != nil {
		return err
	}

	// Jump to elif/else if false
	falseJump := c.emit(core.OpJumpIfFalse, 0, s.Line, s.Col)

	c.beginScope()
	for _, stmt := range s.Body {
		if err := c.compileStmt(stmt); err != nil {
			return err
		}
	}
	c.endScope()

	// Jump over elif/else blocks
	var endJumps []int
	endJump := c.emit(core.OpJump, 0, s.Line, s.Col)
	endJumps = append(endJumps, endJump)

	// Patch false jump
	c.chunk.Code[falseJump].Operand = len(c.chunk.Code)

	// Compile elif chains
	for _, elif := range s.ElIfs {
		if err := c.compileExpr(elif.Condition); err != nil {
			return err
		}
		elifFalseJump := c.emit(core.OpJumpIfFalse, 0, s.Line, s.Col)

		c.beginScope()
		for _, stmt := range elif.Body {
			if err := c.compileStmt(stmt); err != nil {
				return err
			}
		}
		c.endScope()

		ej := c.emit(core.OpJump, 0, s.Line, s.Col)
		endJumps = append(endJumps, ej)
		c.chunk.Code[elifFalseJump].Operand = len(c.chunk.Code)
	}

	// Compile else
	if len(s.Else) > 0 {
		c.beginScope()
		for _, stmt := range s.Else {
			if err := c.compileStmt(stmt); err != nil {
				return err
			}
		}
		c.endScope()
	}

	// Patch all end jumps
	for _, ej := range endJumps {
		c.chunk.Code[ej].Operand = len(c.chunk.Code)
	}

	return nil
}

func (c *Compiler) compileOn(s *ast.OnStmt) error {
	if err := c.compileExpr(s.Value); err != nil {
		return err
	}

	// Save switch value to a temp local
	c.beginScope()
	switchSlot := c.addLocal("__switch__")
	c.emit(core.OpSetLocal, switchSlot, s.Line, s.Col)

	var endJumps []int

	for _, cas := range s.Cases {
		// Re-read switch value for comparison
		c.emit(core.OpGetLocal, switchSlot, s.Line, s.Col)
		if err := c.compileExpr(cas.Value); err != nil {
			return err
		}
		c.emit(core.OpEqual, 0, s.Line, s.Col)
		falseJump := c.emit(core.OpJumpIfFalse, 0, s.Line, s.Col)

		c.beginScope()
		for _, stmt := range cas.Body {
			if err := c.compileStmt(stmt); err != nil {
				return err
			}
		}
		c.endScope()

		ej := c.emit(core.OpJump, 0, s.Line, s.Col)
		endJumps = append(endJumps, ej)
		c.chunk.Code[falseJump].Operand = len(c.chunk.Code)
	}

	for _, ej := range endJumps {
		c.chunk.Code[ej].Operand = len(c.chunk.Code)
	}

	c.endScope() // remove __switch__ temp local

	return nil
}

func (c *Compiler) compileLoop(s *ast.LoopStmt) error {
	loopStart := len(c.chunk.Code)
	c.loopCtx = append(c.loopCtx, loopContext{start: loopStart})

	if s.Limit != nil {
		// Counted loop: emit counter init and check
		ci := c.addConst(float64(0))
		c.emit(core.OpConst, ci, s.Line, s.Col) // push 0 as counter
		iSlot := c.addLocal("i")
		c.emit(core.OpSetLocal, iSlot, s.Line, s.Col)

		loopStart = len(c.chunk.Code) // actual loop start
		c.loopCtx[len(c.loopCtx)-1].start = loopStart

		// Check: i < limit
		idx := c.resolveLocal("i")
		c.emit(core.OpGetLocal, idx, s.Line, s.Col)
		if err := c.compileExpr(s.Limit); err != nil {
			return err
		}
		c.emit(core.OpLess, 0, s.Line, s.Col)
		exitJump := c.emit(core.OpJumpIfFalse, 0, s.Line, s.Col)

		c.beginScope()
		for _, stmt := range s.Body {
			if err := c.compileStmt(stmt); err != nil {
				return err
			}
		}
		c.endScope()

		// Increment i — con jumps here
		incrementStart := len(c.chunk.Code)
		// Patch continues to jump here
		for _, cj := range c.loopCtx[len(c.loopCtx)-1].continues {
			c.chunk.Code[cj].Operand = incrementStart
		}
		c.loopCtx[len(c.loopCtx)-1].continues = nil
		c.emit(core.OpGetLocal, idx, s.Line, s.Col)
		c.emit(core.OpConst, c.addConst(float64(1)), s.Line, s.Col)
		c.emit(core.OpAdd, 0, s.Line, s.Col)
		c.emit(core.OpSetLocal, idx, s.Line, s.Col)

		c.emit(core.OpLoop, len(c.chunk.Code)+1-loopStart, s.Line, s.Col)
		c.chunk.Code[exitJump].Operand = len(c.chunk.Code)

		// Remove counter local
		c.locals = c.locals[:len(c.locals)-1]
	} else {
		// Infinite loop
		c.beginScope()
		for _, stmt := range s.Body {
			if err := c.compileStmt(stmt); err != nil {
				return err
			}
		}
		c.endScope()
		c.emit(core.OpLoop, len(c.chunk.Code)+1-loopStart, s.Line, s.Col)
	}

	// Patch breaks
	ctx := c.loopCtx[len(c.loopCtx)-1]
	for _, b := range ctx.breaks {
		c.chunk.Code[b].Operand = len(c.chunk.Code)
	}
	// Patch any remaining continues (infinite loop: jump to loop start)
	for _, cj := range ctx.continues {
		c.chunk.Code[cj].Operand = loopStart
	}
	c.loopCtx = c.loopCtx[:len(c.loopCtx)-1]

	return nil
}
