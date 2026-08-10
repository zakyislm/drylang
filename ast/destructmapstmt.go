package ast

type DestructMapStmt struct {
	Keys  []string
	Value Expr
	Line  int
	Col   int
}

func (n *DestructMapStmt) stmtNode() {}

func (n *DestructMapStmt) TokenLine() int { return n.Line }

func (n *DestructMapStmt) TokenCol() int { return n.Col }
