package errorhandler
import ("drylang/ast"; "drylang/core")
func CompileThrow(c core.CompilerCore, s *ast.ThrowStmt) error {
	if err := c.CompileExpr(s.Value); err != nil { return err }
	c.Emit(core.OpThrow, 0, s.Line, s.Col)
	return nil
}