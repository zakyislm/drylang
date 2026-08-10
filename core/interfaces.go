package core

import (
	"drylang/ast"
	"drylang/lexer"
)

// ParserCore defines the interface for the main parser.
// Sub-handlers will receive this interface to recursively parse blocks/expressions.
type ParserCore interface {
	ParseStatement() (ast.Stmt, error)
	ParseExpression(precedence int) (ast.Expr, error)
	ParsePrefix() (ast.Expr, error)
	ParseInfix(left ast.Expr) (ast.Expr, error)
	PeekPrecedence() int
	CurrentPrecedence() int
	ParseBlock() ([]ast.Stmt, error)
	ParseParamList() ([]string, error)
	SkipSemicolons()
	Advance() lexer.Token
	Expect(typ lexer.TokenType) (lexer.Token, error)
	Peek() lexer.Token
	Current() lexer.Token
	Errorf(code, format string, args ...interface{}) error
}

// CompilerCore defines the interface for the main compiler.
type CompilerCore interface {
	Emit(op Opcode, operand int, line int, col int) int
	Emit2(op Opcode, operand, operand2 int, line int, col int) int
	AddConst(val interface{}) int
	AddLocal(name string) int
	ResolveLocal(name string) int
	SetGlobal(name string)
	IsGlobal(name string) bool
	BeginScope()
	EndScope()
	GetDepth() int
	CompileExpr(expr ast.Expr) error
	CompileStmt(stmt ast.Stmt) error

	// Loop management
	PushLoopCtx(start int)
	PopLoopCtx()
	AddBreak(jump int)
	GetLoopStart() int

	// Function management
	AddFunc(fn *CompiledFn)

	// Error handling
	Errorf(code string, line, col int, format string, args ...interface{}) error
}

// VMCore defines the interface for the Virtual Machine.
type VMCore interface {
	Push(val Value)
	Pop() Value
	Peek() Value
	SetGlobal(name string, val Value)
	GetGlobal(name string) (Value, bool)
	SetLocal(slot int, val Value)
	GetLocal(slot int) Value
	CallFunction(fn *CompiledFn, argCount int) error
	Try(catchIP int) error
	EndTry() error
	Throw() error
	GetIP() int
	SetIP(ip int)
	GetChunk() *Chunk
	PushFrame(fn *CompiledFn, argCount int) error
	PopFrame()
	Errorf(format string, args ...interface{}) error
}
