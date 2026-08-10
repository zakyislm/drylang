package varhandler
import ("drylang/ast"; "drylang/core")
func CompileConstDecl(c core.CompilerCore, s *ast.ConstDeclStmt) error {
	if err := c.CompileExpr(s.Value); err != nil { return err }
	if c.GetDepth() > 0 {
		c.AddLocal(s.Name)
	} else {
		c.SetGlobal(s.Name)
		ci := c.AddConst(s.Name)
		c.Emit(core.OpSetGlobal, ci, s.Line, s.Col)
	}
	return nil
}