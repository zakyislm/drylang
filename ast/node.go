package ast

// Node is the base interface for all AST nodes.
type Node interface {
	TokenLine() int
	TokenCol() int
}
