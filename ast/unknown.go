package ast

type UnknownLit struct {
	Line int
	Col  int
}

func (n *UnknownLit) exprNode() {}

func (n *UnknownLit) TokenLine() int { return n.Line }

func (n *UnknownLit) TokenCol() int { return n.Col }
