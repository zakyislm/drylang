package classhandler
import ("drylang/ast"; "drylang/core")
func CompilePrivate(c core.CompilerCore, s *ast.PrivateStmt) error {
	return c.CompileStmt(s.Inner)
}