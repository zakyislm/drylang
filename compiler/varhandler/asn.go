package varhandler
import ("drylang/ast"; "drylang/core")
func CompileAssign(c core.CompilerCore, s *ast.AssignStmt) error {
	if err := c.CompileExpr(s.Value); err != nil { return err }
	idx := c.ResolveLocal(s.Name)
	if idx >= 0 {
		c.Emit(core.OpSetLocal, idx, s.Line, s.Col)
	} else if c.IsGlobal(s.Name) {
		ci := c.AddConst(s.Name)
		c.Emit(core.OpSetGlobal, ci, s.Line, s.Col)
	} else if c.GetDepth() > 0 {
		slot := c.AddLocal(s.Name)
		c.Emit(core.OpSetLocal, slot, s.Line, s.Col)
	} else {
		c.SetGlobal(s.Name)
		ci := c.AddConst(s.Name)
		c.Emit(core.OpSetGlobal, ci, s.Line, s.Col)
	}
	return nil
}