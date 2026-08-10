package ast

type DotExpr struct {
	Object Expr
	Field  string
	Optional bool
	Line   int
	Col    int
}

func (n *DotExpr) exprNode() {}

func (n *DotExpr) TokenLine() int { return n.Line }

func (n *DotExpr) TokenCol() int { return n.Col }
