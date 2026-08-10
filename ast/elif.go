package ast

type ElIfClause struct {
	Condition Expr
	Body      []Stmt
}
