package ast

type LoopStmt struct {
	Limit Expr // nil for infinite loop
	Body  []Stmt
	Line  int
	Col   int
}

func (n *LoopStmt) stmtNode() {}

func (n *LoopStmt) TokenLine() int { return n.Line }

func (n *LoopStmt) TokenCol() int { return n.Col }
