package ast

type ConStmt struct {
	Line int
	Col  int
}

func (n *ConStmt) stmtNode() {}

func (n *ConStmt) TokenLine() int { return n.Line }

func (n *ConStmt) TokenCol() int { return n.Col }
