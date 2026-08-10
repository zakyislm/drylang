package loophandler
import ("drylang/ast"; "drylang/core")
func CompileDone(c core.CompilerCore, s *ast.DoneStmt) error {
	jmp := c.Emit(core.OpJump, 0, s.Line, s.Col)
	c.AddBreak(jmp)
	return nil
}