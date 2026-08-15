package ast

type ClassStmt struct {
	Name          string
	Extends       []string
	Fields        []string
	PrivateFields []string
	Methods       []*MethodDecl
	Visibility    string // "pub", "pv", or ""
	Line          int
	Col           int
}

func (n *ClassStmt) stmtNode() {}

func (n *ClassStmt) TokenLine() int { return n.Line }

func (n *ClassStmt) TokenCol() int { return n.Col }
