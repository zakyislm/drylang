package ast

type PrivateStmt struct {
	Inner Stmt // wrapped statement (AssignStmt, FnDeclStmt, etc.)
	Line  int
	Col   int
}

func (n *PrivateStmt) stmtNode() {}

func (n *PrivateStmt) TokenLine() int { return n.Line }

func (n *PrivateStmt) TokenCol() int { return n.Col }
