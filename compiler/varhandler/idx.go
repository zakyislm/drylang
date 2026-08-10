package varhandler
import ("drylang/ast"; "drylang/core")
func CompileIndexAssign(c core.CompilerCore, s *ast.IndexAssignStmt) error {
	if err := c.CompileExpr(s.Object); err != nil { return err }
	if err := c.CompileExpr(s.Index); err != nil { return err }
	if err := c.CompileExpr(s.Value); err != nil { return err }
	c.Emit(core.OpSetIndex, 0, s.Line, s.Col)
	return nil
}