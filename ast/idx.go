package ast

type IndexExpr struct {
	Object Expr
	Index  Expr
	Line   int
	Col    int
}

func (n *IndexExpr) exprNode() {}

func (n *IndexExpr) TokenLine() int { return n.Line }

func (n *IndexExpr) TokenCol() int { return n.Col }
