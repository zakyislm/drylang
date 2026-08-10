package ast

type OnCase struct {
	Value Expr
	Body  []Stmt
}
