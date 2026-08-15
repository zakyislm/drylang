package compiler

import (
	"drylang/ast"
	"drylang/core"
)

func (c *Compiler) compileReturn(s *ast.ReturnStmt) error {
	if s.Value != nil {
		if err := c.compileExpr(s.Value); err != nil {
			return err
		}
	} else {
		c.emit(core.OpUnknown, 0, s.Line, s.Col) // default return unknown
	}
	c.emit(core.OpReturn, 0, s.Line, s.Col)
	return nil
}

func (c *Compiler) compileMulCall(s *ast.MulCallStmt) error {
	if err := c.compileExpr(s.Call.Callee); err != nil {
		return err
	}
	for _, arg := range s.Call.Args {
		if err := c.compileExpr(arg); err != nil {
			return err
		}
	}
	if s.Workers == 1 {
		c.emit(core.OpCall, len(s.Call.Args), s.Line, s.Col)
		c.emit(core.OpPop, 0, s.Line, s.Col)
	} else {
		c.emit2(core.OpAsyncCall, len(s.Call.Args), s.Workers, s.Line, s.Col)
	}
	return nil
}

func (c *Compiler) compileDone(s *ast.DoneStmt) error {
	if len(c.loopCtx) == 0 {
		return c.errorf("E203", s.Line, s.Col, "stray done")
	}
	ctx := &c.loopCtx[len(c.loopCtx)-1]
	jmp := c.emit(core.OpJump, 0, s.Line, s.Col)
	ctx.breaks = append(ctx.breaks, jmp)
	return nil
}

func (c *Compiler) compileCon(s *ast.ConStmt) error {
	if len(c.loopCtx) == 0 {
		return c.errorf("E204", s.Line, s.Col, "stray con")
	}
	ctx := &c.loopCtx[len(c.loopCtx)-1]
	jmp := c.emit(core.OpJump, 0, s.Line, s.Col)
	ctx.continues = append(ctx.continues, jmp)
	return nil
}
