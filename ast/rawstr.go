package ast

type RawStringLit struct {
	Value string
	Line  int
	Col   int
}

func (n *RawStringLit) exprNode() {}

func (n *RawStringLit) TokenLine() int { return n.Line }

func (n *RawStringLit) TokenCol() int { return n.Col }
