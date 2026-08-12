package ast

// MulCallStmt represents a concurrent function call: mul [n] fn(args)
type MulCallStmt struct {
	Call    *CallExpr
	Workers int // default 2, or user-specified
	Line    int
	Col     int
}

func (s *MulCallStmt) stmtNode()      {}
func (s *MulCallStmt) TokenLine() int  { return s.Line }
func (s *MulCallStmt) TokenCol() int   { return s.Col }
