package ast

type NumberLit struct {
	Value string
	Line  int
	Col   int
}

func (n *NumberLit) exprNode() {}

func (n *NumberLit) TokenLine() int { return n.Line }

func (n *NumberLit) TokenCol() int { return n.Col }
