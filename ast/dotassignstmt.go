package ast

type DotAssignStmt struct {
	Object Expr
	Field  string
	Value  Expr
	Line   int
	Col    int
}

func (n *DotAssignStmt) stmtNode() {}

func (n *DotAssignStmt) TokenLine() int { return n.Line }

func (n *DotAssignStmt) TokenCol() int { return n.Col }
