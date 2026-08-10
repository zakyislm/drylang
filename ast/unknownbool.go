package ast

type UnknownBoolStmt struct {
	Name string
	Line int
	Col  int
}

func (n *UnknownBoolStmt) stmtNode() {}

func (n *UnknownBoolStmt) TokenLine() int { return n.Line }

func (n *UnknownBoolStmt) TokenCol() int { return n.Col }
