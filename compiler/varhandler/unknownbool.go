package varhandler
import ("drylang/ast"; "drylang/core")
func CompileUnknownBool(c core.CompilerCore, s *ast.UnknownBoolStmt) error {
	c.Emit(core.OpUnknown, 0, s.Line, s.Col)
	if c.GetDepth() > 0 {
		c.AddLocal(s.Name)
	} else {
		c.SetGlobal(s.Name)
		ci := c.AddConst(s.Name)
		c.Emit(core.OpSetGlobal, ci, s.Line, s.Col)
	}
	return nil
}