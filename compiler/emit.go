package compiler

import (
	"drylang/core"
	"drylang/errfmt"
	"fmt"
)

func (c *Compiler) emit(op core.Opcode, operand int, line, col int) int {
	return c.chunk.Emit(op, operand, line, col)
}

func (c *Compiler) emit2(op core.Opcode, operand, operand2 int, line, col int) int {
	return c.chunk.Emit2(op, operand, operand2, line, col)
}

func (c *Compiler) addConst(val interface{}) int {
	return c.chunk.AddConst(val)
}

func (c *Compiler) errorf(code string, line, col int, format string, args ...interface{}) error {
	return errfmt.Format(code, line, col, fmt.Sprintf(format, args...))
}
