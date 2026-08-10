package varhandler
import ("drylang/ast"; "drylang/core")
func CompileDotAssign(c core.CompilerCore, s *ast.DotAssignStmt) error {
	if err := c.CompileExpr(s.Object); err != nil { return err }
	fi := c.AddConst(s.Field)
	if err := c.CompileExpr(s.Value); err != nil { return err }
	c.Emit(core.OpDotSet, fi, s.Line, s.Col)
	return nil
}