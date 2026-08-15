package compiler

import (
	"drylang/ast"
	"drylang/core"
)

func (c *Compiler) compileTry(s *ast.TryStmt) error {
	tryJump := c.emit(core.OpTry, 0, s.Line, s.Col)

	c.beginScope()
	for _, stmt := range s.Body {
		if err := c.compileStmt(stmt); err != nil {
			return err
		}
	}
	c.endScope()

	c.emit(core.OpEndTry, 0, s.Line, s.Col)
	endJump := c.emit(core.OpJump, 0, s.Line, s.Col)

	// Patch try jump to catch block
	c.chunk.Code[tryJump].Operand = len(c.chunk.Code)

	c.beginScope()
	errSlot := c.addLocal(s.ErrName) // error variable
	c.emit(core.OpSetLocal, errSlot, s.Line, s.Col) // Throw() pushed error to TOS
	for _, stmt := range s.Catch {
		if err := c.compileStmt(stmt); err != nil {
			return err
		}
	}
	c.endScope()

	c.chunk.Code[endJump].Operand = len(c.chunk.Code)

	return nil
}
