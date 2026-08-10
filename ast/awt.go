package ast

type AwaitExpr struct {
	Value Expr
	Line  int
	Col   int
}

func (n *AwaitExpr) exprNode() {}

func (n *AwaitExpr) TokenLine() int { return n.Line }

func (n *AwaitExpr) TokenCol() int { return n.Col }
