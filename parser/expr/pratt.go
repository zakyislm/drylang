package expr

import (
	"drylang/ast"
	"drylang/core"
	"drylang/lexer"
)

func ParseExpression(p core.ParserCore, prec int) (ast.Expr, error) {
	left, err := p.ParsePrefix()
	if err != nil {
		return nil, err
	}

	for prec < p.CurrentPrecedence() {
		left, err = p.ParseInfix(left)
		if err != nil {
			return nil, err
		}
	}

	return left, nil
}

func ParsePrefix(p core.ParserCore) (ast.Expr, error) {
	switch p.Current().Type {
	case lexer.TOKEN_NUMBER:
		return ParseNum(p)
	case lexer.TOKEN_STRING:
		return ParseStr(p)
	case lexer.TOKEN_RAW_STRING:
		return ParseRawStr(p)
	case lexer.TOKEN_STRING_PART:
		return ParseStringInterp(p)
	case lexer.TOKEN_TRUE, lexer.TOKEN_FALSE:
		return ParseBool(p)
	case lexer.TOKEN_UNKNOWN:
		return ParseUnknown(p)
	case lexer.TOKEN_IDENT:
		return ParseIdent(p)
	case lexer.TOKEN_LPAREN:
		// Arrow function params: (a, b) -> body
		if isArrowParens(p) {
			return ParseArrowFnPrefix(p)
		}
		p.Advance() // Consume '('
		expr, err := p.ParseExpression(core.PREC_LOWEST)
		if err != nil {
			return nil, err
		}
		if _, err := p.Expect(lexer.TOKEN_RPAREN); err != nil {
			return nil, err
		}
		return expr, nil
	case lexer.TOKEN_LBRACKET:
		return ParseArrayLit(p)
	case lexer.TOKEN_LBRACE:
		return ParseMapLit(p)
	case lexer.TOKEN_MINUS, lexer.TOKEN_NOT:
		tok := p.Advance()
		operand, err := p.ParseExpression(core.PREC_UNARY)
		if err != nil {
			return nil, err
		}
		return &ast.UnaryExpr{Op: tok.Type, Operand: operand, Line: tok.Line, Col: tok.Col}, nil
	case lexer.TOKEN_ARROW:
		return nil, p.Errorf("E109", "unexpected arrow in prefix")
	default:
		return nil, p.Errorf("E109", "illegal token %s in prefix", p.Current().Literal)
	}
}

func ParseInfix(p core.ParserCore, left ast.Expr) (ast.Expr, error) {
	switch p.Current().Type {
	case lexer.TOKEN_PLUS, lexer.TOKEN_MINUS, lexer.TOKEN_STAR, lexer.TOKEN_SLASH,
		lexer.TOKEN_PERCENT, lexer.TOKEN_LT, lexer.TOKEN_GT, lexer.TOKEN_LT_EQ,
		lexer.TOKEN_GT_EQ, lexer.TOKEN_ASSIGN, lexer.TOKEN_NOT_EQ,
		lexer.TOKEN_AND, lexer.TOKEN_OR, lexer.TOKEN_QQ:
		tok := p.Advance()
		return ParseBinaryExpr(p, left, tok)
	case lexer.TOKEN_LPAREN:
		return ParseCallExpr(p, left)
	case lexer.TOKEN_LBRACKET:
		return ParseIndexExpr(p, left)
	case lexer.TOKEN_DOT, lexer.TOKEN_QMARK_DOT:
		return ParseDotExpr(p, left)
	case lexer.TOKEN_ARROW:
		return ParseArrowFnInfix(p, left)
	default:
		return left, nil
	}
}

// isArrowParens reports whether the '(' at the current position opens an
// arrow-function parameter list: (a, b) -> ... or () -> ...
func isArrowParens(p core.ParserCore) bool {
	depth := 0
	for i := 0; ; i++ {
		tok := p.PeekAt(i)
		switch tok.Type {
		case lexer.TOKEN_LPAREN:
			depth++
		case lexer.TOKEN_RPAREN:
			depth--
			if depth == 0 {
				return p.PeekAt(i+1).Type == lexer.TOKEN_ARROW
			}
		case lexer.TOKEN_EOF:
			return false
		}
	}
}

func ParseBinaryExpr(p core.ParserCore, left ast.Expr, op lexer.Token) (ast.Expr, error) {
	prec := getPrecedence(op.Type)
	right, err := p.ParseExpression(prec)
	if err != nil {
		return nil, err
	}
	return &ast.BinaryExpr{Left: left, Op: op.Type, Right: right, Line: op.Line, Col: op.Col}, nil
}

func getPrecedence(t lexer.TokenType) int {
	switch t {
	case lexer.TOKEN_QQ: return core.PREC_COALESCE
	case lexer.TOKEN_OR: return core.PREC_OR
	case lexer.TOKEN_AND: return core.PREC_AND
	case lexer.TOKEN_ASSIGN, lexer.TOKEN_NOT_EQ: return core.PREC_EQUALITY
	case lexer.TOKEN_LT, lexer.TOKEN_GT, lexer.TOKEN_LT_EQ, lexer.TOKEN_GT_EQ: return core.PREC_COMPARISON
	case lexer.TOKEN_PLUS, lexer.TOKEN_MINUS: return core.PREC_SUM
	case lexer.TOKEN_STAR, lexer.TOKEN_SLASH, lexer.TOKEN_PERCENT: return core.PREC_PRODUCT
	case lexer.TOKEN_LPAREN, lexer.TOKEN_LBRACKET, lexer.TOKEN_DOT, lexer.TOKEN_QMARK_DOT: return core.PREC_CALL
	}
	return core.PREC_LOWEST
}
