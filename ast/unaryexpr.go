package ast

import "drylang/lexer"

type UnaryExpr struct {
	Op      lexer.TokenType
	Operand Expr
	Line    int
	Col     int
}

func (n *UnaryExpr) exprNode() {}

func (n *UnaryExpr) TokenLine() int { return n.Line }

func (n *UnaryExpr) TokenCol() int { return n.Col }
