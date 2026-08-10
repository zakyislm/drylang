package ast

type Ident struct {
	Name string
	Line int
	Col  int
}

func (n *Ident) exprNode() {}

func (n *Ident) TokenLine() int { return n.Line }

func (n *Ident) TokenCol() int { return n.Col }
