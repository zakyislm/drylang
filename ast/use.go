package ast

type UseStmt struct {
	Path string
	Line int
	Col  int
}

func (n *UseStmt) stmtNode() {}

func (n *UseStmt) TokenLine() int { return n.Line }

func (n *UseStmt) TokenCol() int { return n.Col }
