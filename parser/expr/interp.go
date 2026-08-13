package expr

import (
	"drylang/ast"
	"drylang/core"
	"drylang/lexer"
)

// ParseStringInterp parses a string with ${expr} interpolation.
// Token stream produced by the lexer:
//   STRING_PART? ( INTERP_START <expr> INTERP_END )* STRING_PART?
func ParseStringInterp(p core.ParserCore) (ast.Expr, error) {
	tok := p.Advance() // consume leading STRING_PART

	parts := []ast.Expr{&ast.StringLit{Value: tok.Literal, Line: tok.Line, Col: tok.Col}}

	for p.Current().Type == lexer.TOKEN_INTERP_START {
		p.Advance() // consume INTERP_START
		e, err := p.ParseExpression(core.PREC_LOWEST)
		if err != nil {
			return nil, err
		}
		parts = append(parts, e)
		if _, err := p.Expect(lexer.TOKEN_INTERP_END); err != nil {
			return nil, err
		}
	}

	// Optional trailing string part (e.g. "a${x}b")
	if p.Current().Type == lexer.TOKEN_STRING_PART {
		ttok := p.Advance()
		parts = append(parts, &ast.StringLit{Value: ttok.Literal, Line: ttok.Line, Col: ttok.Col})
	}

	return &ast.StringInterp{Parts: parts, Line: tok.Line, Col: tok.Col}, nil
}
