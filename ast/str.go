package ast

type StringLit struct {
	Value string
	Line  int
	Col   int
}

func (n *StringLit) exprNode() {}

func (n *StringLit) TokenLine() int { return n.Line }

func (n *StringLit) TokenCol() int { return n.Col }
