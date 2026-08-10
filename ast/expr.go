package ast

// Expr nodes (produce values)
type Expr interface {
	Node
	exprNode()
}
