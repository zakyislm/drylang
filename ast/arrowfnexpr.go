package ast

type ArrowFnExpr struct {
	Params []string
	Body   []Stmt
	Line   int
	Col    int
}

func (n *ArrowFnExpr) exprNode() {}

func (n *ArrowFnExpr) TokenLine() int { return n.Line }

func (n *ArrowFnExpr) TokenCol() int { return n.Col }
