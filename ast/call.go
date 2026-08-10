package ast

type CallExpr struct {
	Callee Expr
	Args   []Expr
	Line   int
	Col    int
}

func (n *CallExpr) exprNode() {}

func (n *CallExpr) TokenLine() int { return n.Line }

func (n *CallExpr) TokenCol() int { return n.Col }
