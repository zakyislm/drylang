package ast

type FnDeclStmt struct {
	Name    string
	Params  []string
	Body    []Stmt
	IsAsync bool
	Line    int
	Col     int
}

func (n *FnDeclStmt) stmtNode() {}

func (n *FnDeclStmt) TokenLine() int { return n.Line }

func (n *FnDeclStmt) TokenCol() int { return n.Col }
