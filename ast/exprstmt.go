package ast

type ExprStmt struct {
	Expression Expr
	Line       int
	Col        int
}

func (n *ExprStmt) stmtNode() {}

func (n *ExprStmt) TokenLine() int { return n.Line }

func (n *ExprStmt) TokenCol() int { return n.Col }
