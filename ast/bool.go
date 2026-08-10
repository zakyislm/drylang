package ast

type BoolLit struct {
	Value bool
	Line  int
	Col   int
}

func (n *BoolLit) exprNode() {}

func (n *BoolLit) TokenLine() int { return n.Line }

func (n *BoolLit) TokenCol() int { return n.Col }
