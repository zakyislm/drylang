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
	ParseBlock() (*ast.BlockStmt, error)
	Advance() lexer.Token
	Expect(typ lexer.TokenType) (lexer.Token, error)
	Peek() lexer.Token
	Current() lexer.Token
	Errorf(code, format string, args ...interface{}) error
}

// CompilerCore defines the interface for the main compiler.
type CompilerCore interface {
	Emit(op byte, operand int, line int, col int)
	CompileExpr(expr ast.Expr) error
	CompileStmt(stmt ast.Stmt) error
	// TODO: Add other necessary compiler methods
}

// VMCore defines the interface for the Virtual Machine.
type VMCore interface {
	Push(val interface{}) // Use interface{} or generic Value type for now
	Pop() interface{}
	// TODO: Add other VM methods
}
