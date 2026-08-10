package ast

type ReturnStmt struct {
	Value Expr // nil if no value
	Line  int
	Col   int
}

func (n *ReturnStmt) stmtNode() {}

func (n *ReturnStmt) TokenLine() int { return n.Line }

func (n *ReturnStmt) TokenCol() int { return n.Col }
