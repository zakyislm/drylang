package exprhandler
import ("drylang/ast"; "drylang/core")
func CompileExprStmt(c core.CompilerCore, s *ast.ExprStmt) error {
	if err := c.CompileExpr(s.Expression); err != nil { return err }
	c.Emit(core.OpPop, 0, s.Line, s.Col)
	return nil
}