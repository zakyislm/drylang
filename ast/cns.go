package ast

type ConstDeclStmt struct {
	Name  string
	Value Expr
	Line  int
	Col   int
}

func (n *ConstDeclStmt) stmtNode() {}

func (n *ConstDeclStmt) TokenLine() int { return n.Line }

func (n *ConstDeclStmt) TokenCol() int { return n.Col }
