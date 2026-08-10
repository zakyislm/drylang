package ast

type IfStmt struct {
	Condition Expr
	Body      []Stmt
	ElIfs     []ElIfClause
	Else      []Stmt // el block
	Line      int
	Col       int
}

func (n *IfStmt) stmtNode() {}

func (n *IfStmt) TokenLine() int { return n.Line }

func (n *IfStmt) TokenCol() int { return n.Col }
