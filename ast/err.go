package ast

type ThrowStmt struct {
	Value Expr
	Line  int
	Col   int
}

func (n *ThrowStmt) stmtNode() {}

func (n *ThrowStmt) TokenLine() int { return n.Line }

func (n *ThrowStmt) TokenCol() int { return n.Col }
