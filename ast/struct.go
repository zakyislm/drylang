package ast

type StructDeclStmt struct {
	Name       string
	Fields     []string
	Visibility string // "pub", "pv", or ""
	Line       int
	Col        int
}

func (n *StructDeclStmt) stmtNode() {}

func (n *StructDeclStmt) TokenLine() int { return n.Line }

func (n *StructDeclStmt) TokenCol() int { return n.Col }
