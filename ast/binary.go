package ast

import "drylang/lexer"

type BinaryExpr struct {
	Left  Expr
	Op    lexer.TokenType
	Right Expr
	Line  int
	Col   int
}

func (n *BinaryExpr) exprNode() {}

func (n *BinaryExpr) TokenLine() int { return n.Line }

func (n *BinaryExpr) TokenCol() int { return n.Col }
