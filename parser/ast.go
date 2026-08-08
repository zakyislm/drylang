package parser

import "drylang/lexer"

// Node is the base interface for all AST nodes.
type Node interface {
	TokenLine() int
	TokenCol() int
}

// Expr nodes (produce values)
type Expr interface {
	Node
	exprNode()
}

// Stmt nodes (perform actions)
type Stmt interface {
	Node
	stmtNode()
}

// Program is the root node.
type Program struct {
	Stmts []Stmt
}

func (p *Program) TokenLine() int { return 0 }
func (p *Program) TokenCol() int  { return 0 }

// --- Expressions ---

type NumberLit struct {
	Value string
	Line  int
	Col   int
}

func (n *NumberLit) exprNode()      {}
func (n *NumberLit) TokenLine() int { return n.Line }
func (n *NumberLit) TokenCol() int  { return n.Col }

type StringLit struct {
	Value string
	Line  int
	Col   int
}

func (n *StringLit) exprNode()      {}
func (n *StringLit) TokenLine() int { return n.Line }
func (n *StringLit) TokenCol() int  { return n.Col }

type RawStringLit struct {
	Value string
	Line  int
	Col   int
}

func (n *RawStringLit) exprNode()      {}
func (n *RawStringLit) TokenLine() int { return n.Line }
func (n *RawStringLit) TokenCol() int  { return n.Col }

// StringInterp represents a string with ${} interpolation.
// Parts alternate between string segments and expressions.
type StringInterp struct {
	Parts []Expr // StringLit and other exprs interleaved
	Line  int
	Col   int
}

func (n *StringInterp) exprNode()      {}
func (n *StringInterp) TokenLine() int { return n.Line }
func (n *StringInterp) TokenCol() int  { return n.Col }

type BoolLit struct {
	Value bool
	Line  int
	Col   int
}

func (n *BoolLit) exprNode()      {}
func (n *BoolLit) TokenLine() int { return n.Line }
func (n *BoolLit) TokenCol() int  { return n.Col }

type UnknownLit struct {
	Line int
	Col  int
}

func (n *UnknownLit) exprNode()      {}
func (n *UnknownLit) TokenLine() int { return n.Line }
func (n *UnknownLit) TokenCol() int  { return n.Col }

type Ident struct {
	Name string
	Line int
	Col  int
}

func (n *Ident) exprNode()      {}
func (n *Ident) TokenLine() int { return n.Line }
func (n *Ident) TokenCol() int  { return n.Col }

type BinaryExpr struct {
	Left  Expr
	Op    lexer.TokenType
	Right Expr
	Line  int
	Col   int
}

func (n *BinaryExpr) exprNode()      {}
func (n *BinaryExpr) TokenLine() int { return n.Line }
func (n *BinaryExpr) TokenCol() int  { return n.Col }

type UnaryExpr struct {
	Op      lexer.TokenType
	Operand Expr
	Line    int
	Col     int
}

func (n *UnaryExpr) exprNode()      {}
func (n *UnaryExpr) TokenLine() int { return n.Line }
func (n *UnaryExpr) TokenCol() int  { return n.Col }

type CallExpr struct {
	Callee Expr
	Args   []Expr
	Line   int
	Col    int
}

func (n *CallExpr) exprNode()      {}
func (n *CallExpr) TokenLine() int { return n.Line }
func (n *CallExpr) TokenCol() int  { return n.Col }

type IndexExpr struct {
	Object Expr
	Index  Expr
	Line   int
	Col    int
}

func (n *IndexExpr) exprNode()      {}
func (n *IndexExpr) TokenLine() int { return n.Line }
func (n *IndexExpr) TokenCol() int  { return n.Col }

type DotExpr struct {
	Object Expr
	Field  string
	Line   int
	Col    int
}

func (n *DotExpr) exprNode()      {}
func (n *DotExpr) TokenLine() int { return n.Line }
func (n *DotExpr) TokenCol() int  { return n.Col }

type ArrayLit struct {
	Items []Expr
	Line  int
	Col   int
}

func (n *ArrayLit) exprNode()      {}
func (n *ArrayLit) TokenLine() int { return n.Line }
func (n *ArrayLit) TokenCol() int  { return n.Col }

type MapLit struct {
	Keys   []Expr
	Values []Expr
	Line   int
	Col    int
}

func (n *MapLit) exprNode()      {}
func (n *MapLit) TokenLine() int { return n.Line }
func (n *MapLit) TokenCol() int  { return n.Col }

type ArrowFnExpr struct {
	Params []string
	Body   []Stmt
	Line   int
	Col    int
}

func (n *ArrowFnExpr) exprNode()      {}
func (n *ArrowFnExpr) TokenLine() int { return n.Line }
func (n *ArrowFnExpr) TokenCol() int  { return n.Col }

type InputExpr struct {
	Prompt Expr // optional prompt argument
	Line   int
	Col    int
}

func (n *InputExpr) exprNode()      {}
func (n *InputExpr) TokenLine() int { return n.Line }
func (n *InputExpr) TokenCol() int  { return n.Col }

// --- Statements ---

type AssignStmt struct {
	Name  string
	Value Expr
	Line  int
	Col   int
}

func (n *AssignStmt) stmtNode()      {}
func (n *AssignStmt) TokenLine() int { return n.Line }
func (n *AssignStmt) TokenCol() int  { return n.Col }

type ConstDeclStmt struct {
	Name  string
	Value Expr
	Line  int
	Col   int
}

func (n *ConstDeclStmt) stmtNode()      {}
func (n *ConstDeclStmt) TokenLine() int { return n.Line }
func (n *ConstDeclStmt) TokenCol() int  { return n.Col }

type UnknownBoolStmt struct {
	Name string
	Line int
	Col  int
}

func (n *UnknownBoolStmt) stmtNode()      {}
func (n *UnknownBoolStmt) TokenLine() int { return n.Line }
func (n *UnknownBoolStmt) TokenCol() int  { return n.Col }

type IndexAssignStmt struct {
	Object Expr
	Index  Expr
	Value  Expr
	Line   int
	Col    int
}

func (n *IndexAssignStmt) stmtNode()      {}
func (n *IndexAssignStmt) TokenLine() int { return n.Line }
func (n *IndexAssignStmt) TokenCol() int  { return n.Col }

type DotAssignStmt struct {
	Object Expr
	Field  string
	Value  Expr
	Line   int
	Col    int
}

func (n *DotAssignStmt) stmtNode()      {}
func (n *DotAssignStmt) TokenLine() int { return n.Line }
func (n *DotAssignStmt) TokenCol() int  { return n.Col }

type PrintStmt struct {
	Value Expr
	Line  int
	Col   int
}

func (n *PrintStmt) stmtNode()      {}
func (n *PrintStmt) TokenLine() int { return n.Line }
func (n *PrintStmt) TokenCol() int  { return n.Col }

type ExprStmt struct {
	Expression Expr
	Line       int
	Col        int
}

func (n *ExprStmt) stmtNode()      {}
func (n *ExprStmt) TokenLine() int { return n.Line }
func (n *ExprStmt) TokenCol() int  { return n.Col }

type ReturnStmt struct {
	Value Expr // nil if no value
	Line  int
	Col   int
}

func (n *ReturnStmt) stmtNode()      {}
func (n *ReturnStmt) TokenLine() int { return n.Line }
func (n *ReturnStmt) TokenCol() int  { return n.Col }

type FnDeclStmt struct {
	Name    string
	Params  []string
	Body    []Stmt
	IsAsync bool
	Line    int
	Col     int
}

func (n *FnDeclStmt) stmtNode()      {}
func (n *FnDeclStmt) TokenLine() int { return n.Line }
func (n *FnDeclStmt) TokenCol() int  { return n.Col }

type IfStmt struct {
	Condition Expr
	Body      []Stmt
	ElIfs     []ElIfClause
	Else      []Stmt // el block
	Line      int
	Col       int
}

type ElIfClause struct {
	Condition Expr
	Body      []Stmt
}

func (n *IfStmt) stmtNode()      {}
func (n *IfStmt) TokenLine() int { return n.Line }
func (n *IfStmt) TokenCol() int  { return n.Col }

type OnStmt struct {
	Value Expr
	Cases []OnCase
	Line  int
	Col   int
}

type OnCase struct {
	Value Expr
	Body  []Stmt
}

func (n *OnStmt) stmtNode()      {}
func (n *OnStmt) TokenLine() int { return n.Line }
func (n *OnStmt) TokenCol() int  { return n.Col }

type LoopStmt struct {
	Limit Expr // nil for infinite loop
	Body  []Stmt
	Line  int
	Col   int
}

func (n *LoopStmt) stmtNode()      {}
func (n *LoopStmt) TokenLine() int { return n.Line }
func (n *LoopStmt) TokenCol() int  { return n.Col }

type DoneStmt struct {
	Line int
	Col  int
}

func (n *DoneStmt) stmtNode()      {}
func (n *DoneStmt) TokenLine() int { return n.Line }
func (n *DoneStmt) TokenCol() int  { return n.Col }

type ConStmt struct {
	Line int
	Col  int
}

func (n *ConStmt) stmtNode()      {}
func (n *ConStmt) TokenLine() int { return n.Line }
func (n *ConStmt) TokenCol() int  { return n.Col }

type TryStmt struct {
	Body    []Stmt
	ErrName string // variable name in err(e)
	Catch   []Stmt
	Line    int
	Col     int
}

func (n *TryStmt) stmtNode()      {}
func (n *TryStmt) TokenLine() int { return n.Line }
func (n *TryStmt) TokenCol() int  { return n.Col }

type ThrowStmt struct {
	Value Expr
	Line  int
	Col   int
}

func (n *ThrowStmt) stmtNode()      {}
func (n *ThrowStmt) TokenLine() int { return n.Line }
func (n *ThrowStmt) TokenCol() int  { return n.Col }

type UseStmt struct {
	Path string
	Line int
	Col  int
}

func (n *UseStmt) stmtNode()      {}
func (n *UseStmt) TokenLine() int { return n.Line }
func (n *UseStmt) TokenCol() int  { return n.Col }

type PrivateStmt struct {
	Inner Stmt // wrapped statement (AssignStmt, FnDeclStmt, etc.)
	Line  int
	Col   int
}

func (n *PrivateStmt) stmtNode()      {}
func (n *PrivateStmt) TokenLine() int { return n.Line }
func (n *PrivateStmt) TokenCol() int  { return n.Col }

type StructDeclStmt struct {
	Name   string
	Fields []string
	Line   int
	Col    int
}

func (n *StructDeclStmt) stmtNode()      {}
func (n *StructDeclStmt) TokenLine() int { return n.Line }
func (n *StructDeclStmt) TokenCol() int  { return n.Col }

type StructInitExpr struct {
	TypeName string
	Fields   map[string]Expr
	Line     int
	Col      int
}

func (n *StructInitExpr) exprNode()      {}
func (n *StructInitExpr) TokenLine() int { return n.Line }
func (n *StructInitExpr) TokenCol() int  { return n.Col }

type AwaitExpr struct {
	Value Expr
	Line  int
	Col   int
}

func (n *AwaitExpr) exprNode()      {}
func (n *AwaitExpr) TokenLine() int { return n.Line }
func (n *AwaitExpr) TokenCol() int  { return n.Col }
