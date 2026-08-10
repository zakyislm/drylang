package ast

type ArrayLit struct {
	Items []Expr
	Line  int
	Col   int
}

func (n *ArrayLit) exprNode() {}

func (n *ArrayLit) TokenLine() int { return n.Line }

func (n *ArrayLit) TokenCol() int { return n.Col }
