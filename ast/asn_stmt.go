package ast

type IndexAssignStmt struct {
	Object Expr
	Index  Expr
	Value  Expr
	Line   int
	Col    int
}

func (n *IndexAssignStmt) stmtNode() {}

func (n *IndexAssignStmt) TokenLine() int { return n.Line }

func (n *IndexAssignStmt) TokenCol() int { return n.Col }
