package exprhandler

import (
	"drylang/ast"
	"drylang/core"
	"drylang/lexer"
	"drylang/parser/classhandler"
	"drylang/parser/colshandler"
	"drylang/parser/functionhandler"
	"drylang/parser/typehandler"
	"drylang/parser/varhandler"
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
		return typehandler.ParseNum(p)
	case lexer.TOKEN_STRING:
		return typehandler.ParseStr(p)
	case lexer.TOKEN_RAW_STRING:
		return typehandler.ParseRawStr(p)
	case lexer.TOKEN_STRING_PART:
		return typehandler.ParseStringInterp(p)
	case lexer.TOKEN_TRUE, lexer.TOKEN_FALSE:
		return typehandler.ParseBool(p)
	case lexer.TOKEN_UNKNOWN:
		return typehandler.ParseUnknown(p)
	case lexer.TOKEN_IDENT:
		return varhandler.ParseIdent(p)
	case lexer.TOKEN_LPAREN:
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
		return colshandler.ParseArrayLit(p)
	case lexer.TOKEN_LBRACE:
		return colshandler.ParseMapLit(p)
	case lexer.TOKEN_MINUS, lexer.TOKEN_NOT:
		tok := p.Advance()
		operand, err := p.ParseExpression(core.PREC_UNARY)
		if err != nil {
			return nil, err
		}
		return &ast.UnaryExpr{Op: tok.Type, Operand: operand, Line: tok.Line, Col: tok.Col}, nil
	case lexer.TOKEN_ARROW:
		return functionhandler.ParseArrowFn(p)
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
		return functionhandler.ParseCallExpr(p, left)
	case lexer.TOKEN_LBRACKET:
		return colshandler.ParseIndexExpr(p, left)
	case lexer.TOKEN_DOT, lexer.TOKEN_QMARK_DOT:
		return classhandler.ParseDotExpr(p, left)
	default:
		return left, nil
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
