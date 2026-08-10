package ast

// StringInterp represents a string with ${} interpolation.
// Parts alternate between string segments and expressions.
type StringInterp struct {
	Parts []Expr // StringLit and other exprs interleaved
	Line  int
	Col   int
}

func (n *StringInterp) exprNode() {}

func (n *StringInterp) TokenLine() int { return n.Line }

func (n *StringInterp) TokenCol() int { return n.Col }
