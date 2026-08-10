package ast

type DestructArrayStmt struct {
	Names []string
	Value Expr
	Line  int
	Col   int
}

func (n *DestructArrayStmt) stmtNode() {}

func (n *DestructArrayStmt) TokenLine() int { return n.Line }

func (n *DestructArrayStmt) TokenCol() int { return n.Col }
