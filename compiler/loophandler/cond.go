package loophandler
import ("drylang/ast"; "drylang/core")
func CompileCon(c core.CompilerCore, s *ast.ConStmt) error {
	c.Emit(core.OpLoop, c.GetLoopStart(), s.Line, s.Col)
	return nil
}