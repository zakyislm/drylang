package compiler

import (
	"drylang/ast"
	"drylang/core"
)

func (c *Compiler) Emit(op core.Opcode, operand int, line int, col int) int {
	return c.emit(op, operand, line, col)
}

func (c *Compiler) Emit2(op core.Opcode, operand, operand2 int, line int, col int) int {
	return c.emit2(op, operand, operand2, line, col)
}

func (c *Compiler) AddConst(val interface{}) int {
	return c.addConst(val)
}

func (c *Compiler) AddLocal(name string) int {
	return c.addLocal(name)
}

func (c *Compiler) ResolveLocal(name string) int {
	return c.resolveLocal(name)
}

func (c *Compiler) SetGlobal(name string) {
	c.globals[name] = true
}

func (c *Compiler) IsGlobal(name string) bool {
	return c.globals[name]
}

func (c *Compiler) BeginScope() {
	c.beginScope()
}

func (c *Compiler) EndScope() {
	c.endScope()
}

func (c *Compiler) GetDepth() int {
	return c.depth
}

func (c *Compiler) CompileExpr(expr ast.Expr) error {
	return c.compileExpr(expr)
}

func (c *Compiler) CompileStmt(stmt ast.Stmt) error {
	return c.compileStmt(stmt)
}

func (c *Compiler) PushLoopCtx(start int) {
	c.loopCtx = append(c.loopCtx, loopContext{start: start})
}

func (c *Compiler) PopLoopCtx() {
	c.loopCtx = c.loopCtx[:len(c.loopCtx)-1]
}

func (c *Compiler) AddBreak(jump int) {
	c.loopCtx[len(c.loopCtx)-1].breaks = append(c.loopCtx[len(c.loopCtx)-1].breaks, jump)
}

func (c *Compiler) GetLoopStart() int {
	return c.loopCtx[len(c.loopCtx)-1].start
}

func (c *Compiler) AddFunc(fn *core.CompiledFn) {
	c.fns = append(c.fns, fn)
}

func (c *Compiler) Errorf(code string, line, col int, format string, args ...interface{}) error {
	return c.errorf(code, line, col, format, args...)
}
