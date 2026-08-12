package ast

type AwaitStmt struct {
	Line  int
	Col   int
}

func (n *AwaitStmt) stmtNode() {}

func (n *AwaitStmt) TokenLine() int { return n.Line }

func (n *AwaitStmt) TokenCol() int { return n.Col }
