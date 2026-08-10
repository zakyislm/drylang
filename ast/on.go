package ast

type OnStmt struct {
	Value Expr
	Cases []OnCase
	Line  int
	Col   int
}

func (n *OnStmt) stmtNode() {}

func (n *OnStmt) TokenLine() int { return n.Line }

func (n *OnStmt) TokenCol() int { return n.Col }
