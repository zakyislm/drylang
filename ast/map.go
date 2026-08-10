package ast

type MapLit struct {
	Keys   []Expr
	Values []Expr
	Line   int
	Col    int
}

func (n *MapLit) exprNode() {}

func (n *MapLit) TokenLine() int { return n.Line }

func (n *MapLit) TokenCol() int { return n.Col }
