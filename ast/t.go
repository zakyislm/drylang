package ast

type TryStmt struct {
	Body    []Stmt
	ErrName string // variable name in err(e)
	Catch   []Stmt
	Line    int
	Col     int
}

func (n *TryStmt) stmtNode() {}

func (n *TryStmt) TokenLine() int { return n.Line }

func (n *TryStmt) TokenCol() int { return n.Col }
