package ast

type AssignStmt struct {
	Name  string
	Value Expr
	Line  int
	Col   int
}

func (n *AssignStmt) stmtNode() {}

func (n *AssignStmt) TokenLine() int { return n.Line }

func (n *AssignStmt) TokenCol() int { return n.Col }
