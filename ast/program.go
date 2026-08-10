package ast

// Program is the root node.
type Program struct {
	Stmts []Stmt
}

func (p *Program) TokenLine() int { return 0 }

func (p *Program) TokenCol() int { return 0 }
