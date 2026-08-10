package ast

type StructInitExpr struct {
	TypeName string
	Fields   map[string]Expr
	Line     int
	Col      int
}

func (n *StructInitExpr) exprNode() {}

func (n *StructInitExpr) TokenLine() int { return n.Line }

func (n *StructInitExpr) TokenCol() int { return n.Col }
