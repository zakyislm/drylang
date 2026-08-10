package ast

type MethodDecl struct {
	Name       string
	Params     []string
	Body       []Stmt
	Visibility string // "pub", "pv", or ""
	IsAsync    bool
	Line       int
	Col        int
}

func (n *MethodDecl) stmtNode() {}

func (n *MethodDecl) TokenLine() int { return n.Line }

func (n *MethodDecl) TokenCol() int { return n.Col }
