package stmt

import (
	"drylang/ast"
	"drylang/core"
)

func ParsePrivate(p core.ParserCore) (ast.Stmt, error) {
	tok := p.Advance() // consume pv
	inner, err := p.ParseStatement()
	if err != nil {
		return nil, err
	}

	// If inner is a ClassStmt or ast.StructDeclStmt, inject "pv"
	if cls, ok := inner.(*ast.ClassStmt); ok {
		cls.Visibility = "pv"
	} else if strct, ok := inner.(*ast.StructDeclStmt); ok {
		strct.Visibility = "pv"
	}

	return &ast.PrivateStmt{Inner: inner, Line: tok.Line, Col: tok.Col}, nil
}

