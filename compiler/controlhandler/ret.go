package controlhandler
import ("drylang/ast"; "drylang/core")
func CompileReturn(c core.CompilerCore, s *ast.ReturnStmt) error {
	if s.Value != nil {
		if err := c.CompileExpr(s.Value); err != nil { return err }
	} else {
		c.Emit(core.OpUnknown, 0, s.Line, s.Col)
	}
	c.Emit(core.OpReturn, 0, s.Line, s.Col)
	return nil
}