package ast

// Stmt nodes (perform actions)
type Stmt interface {
	Node
	stmtNode()
}
