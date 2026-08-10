package ast

type DoneStmt struct {
	Line int
	Col  int
}

func (n *DoneStmt) stmtNode() {}

func (n *DoneStmt) TokenLine() int { return n.Line }

func (n *DoneStmt) TokenCol() int { return n.Col }
