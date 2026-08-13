package parser

import (
	"drylang/ast"
	"drylang/core"
	"drylang/errfmt"
	"drylang/lexer"
	"fmt"

	"drylang/parser/expr"
	"drylang/parser/stmt"
)



var precedences = map[lexer.TokenType]int{
	lexer.TOKEN_OR:       core.PREC_OR,
	lexer.TOKEN_AND:      core.PREC_AND,
	lexer.TOKEN_ASSIGN:   core.PREC_EQUALITY,
	lexer.TOKEN_NOT_EQ:   core.PREC_EQUALITY,
	lexer.TOKEN_LT:       core.PREC_COMPARISON,
	lexer.TOKEN_GT:       core.PREC_COMPARISON,
	lexer.TOKEN_LT_EQ:    core.PREC_COMPARISON,
	lexer.TOKEN_GT_EQ:    core.PREC_COMPARISON,
	lexer.TOKEN_PLUS:     core.PREC_SUM,
	lexer.TOKEN_MINUS:    core.PREC_SUM,
	lexer.TOKEN_STAR:     core.PREC_PRODUCT,
	lexer.TOKEN_SLASH:    core.PREC_PRODUCT,
	lexer.TOKEN_PERCENT:  core.PREC_PRODUCT,
	lexer.TOKEN_LPAREN:   core.PREC_CALL,
	lexer.TOKEN_LBRACKET: core.PREC_CALL,
	lexer.TOKEN_DOT:      core.PREC_CALL,
	lexer.TOKEN_QMARK_DOT: core.PREC_CALL,
	lexer.TOKEN_ARROW:     core.PREC_CALL,
	lexer.TOKEN_QQ:        core.PREC_COALESCE,
}

type Parser struct {
	tokens  []lexer.Token
	pos     int
	current lexer.Token
}

var _ core.ParserCore = (*Parser)(nil)

func New(tokens []lexer.Token) *Parser {
	p := &Parser{tokens: tokens, pos: 0}
	if len(tokens) > 0 {
		p.current = tokens[0]
	}
	return p
}

func (p *Parser) Peek() lexer.Token {
	if p.pos+1 < len(p.tokens) {
		return p.tokens[p.pos+1]
	}
	return lexer.Token{Type: lexer.TOKEN_EOF}
}

// PeekAt returns the token at the given offset from the current position.
func (p *Parser) PeekAt(offset int) lexer.Token {
	if p.pos+offset < len(p.tokens) {
		return p.tokens[p.pos+offset]
	}
	return lexer.Token{Type: lexer.TOKEN_EOF}
}

func (p *Parser) Advance() lexer.Token {
	tok := p.current
	p.pos++
	if p.pos < len(p.tokens) {
		p.current = p.tokens[p.pos]
	} else {
		p.current = lexer.Token{Type: lexer.TOKEN_EOF}
	}
	return tok
}

func (p *Parser) Current() lexer.Token {
	return p.current
}

func errorCodeForMissingToken(typ lexer.TokenType) string {
	switch typ {
	case lexer.TOKEN_RBRACE:
		return "E102"
	case lexer.TOKEN_RPAREN:
		return "E103"
	case lexer.TOKEN_RBRACKET:
		return "E104"
	case lexer.TOKEN_IDENT:
		return "E107"
	case lexer.TOKEN_STRING:
		return "E108"
	}
	return "E109"
}

func (p *Parser) Expect(typ lexer.TokenType) (lexer.Token, error) {
	if p.current.Type != typ {
		code := errorCodeForMissingToken(typ)
		return lexer.Token{}, p.Errorf(code, "needs %s", typ)
	}
	return p.Advance(), nil
}

func (p *Parser) Errorf(code, format string, args ...interface{}) error {
	msg := fmt.Sprintf(format, args...)
	return errfmt.Format(code, p.current.Line, p.current.Col, msg)
}

func (p *Parser) PeekPrecedence() int {
	if prec, ok := precedences[p.Peek().Type]; ok {
		return prec
	}
	return core.PREC_LOWEST
}

func (p *Parser) CurrentPrecedence() int {
	if prec, ok := precedences[p.current.Type]; ok {
		return prec
	}
	return core.PREC_LOWEST
}

func (p *Parser) Parse() (*ast.Program, error) {
	program := &ast.Program{}
	p.SkipSemicolons()

	for p.current.Type != lexer.TOKEN_EOF {
		stmt, err := p.ParseStatement()
		if err != nil {
			return nil, err
		}
		if stmt != nil {
			program.Stmts = append(program.Stmts, stmt)
		}
		p.SkipSemicolons()
	}

	return program, nil
}

func (p *Parser) ParseStatement() (ast.Stmt, error) {
	switch p.current.Type {
	case lexer.TOKEN_CNS:
		return stmt.ParseConstDecl(p)
	case lexer.TOKEN_FN:
		return stmt.ParseFnDecl(p, false)
	case lexer.TOKEN_ASN:
		// Async Function
		p.Advance()
		if p.Current().Type != lexer.TOKEN_FN {
			return nil, p.Errorf("E105", "needs fn")
		}
		return stmt.ParseFnDecl(p, true)
	case lexer.TOKEN_REV:
		return stmt.ParseReturn(p)
	case lexer.TOKEN_IF:
		return stmt.ParseIf(p)
	case lexer.TOKEN_ON:
		return stmt.ParseOn(p)
	case lexer.TOKEN_LP:
		return stmt.ParseLoop(p)
	case lexer.TOKEN_DONE:
		return stmt.ParseDone(p)
	case lexer.TOKEN_CON:
		return stmt.ParseCon(p)
	case lexer.TOKEN_MUL:
		return stmt.ParseMulCall(p)
	case lexer.TOKEN_UNI:
		return stmt.ParseUniCall(p)
	case lexer.TOKEN_AWT:
		return stmt.ParseAwt(p)
	case lexer.TOKEN_TRY:
		return stmt.ParseTry(p)
	case lexer.TOKEN_ERR:
		return stmt.ParseThrow(p)
	case lexer.TOKEN_USE:
		return stmt.ParseUse(p)
	case lexer.TOKEN_PV:
		return stmt.ParsePrivate(p)
	case lexer.TOKEN_CL:
		return stmt.ParseClass(p)
	case lexer.TOKEN_QUESTION:
		return stmt.ParseUnknownBool(p)
	case lexer.TOKEN_IDENT:
		if p.Peek().Type == lexer.TOKEN_LBRACE {
			return stmt.ParseStruct(p)
		}
		fallthrough
	default:
		return p.ParseExpressionOrAssign()
	}
}

func (p *Parser) SkipSemicolons() {
	for p.current.Type == lexer.TOKEN_SEMICOLON {
		p.Advance()
	}
}

func (p *Parser) ParseExpressionOrAssign() (ast.Stmt, error) {
	return stmt.ParseAssignOrExpr(p)
}

func (p *Parser) ParseExpression(prec int) (ast.Expr, error) {
	return expr.ParseExpression(p, prec)
}

func (p *Parser) ParsePrefix() (ast.Expr, error) {
	return expr.ParsePrefix(p)
}

func (p *Parser) ParseInfix(left ast.Expr) (ast.Expr, error) {
	return expr.ParseInfix(p, left)
}

func (p *Parser) ParseBlock() ([]ast.Stmt, error) {
	if _, err := p.Expect(lexer.TOKEN_LBRACE); err != nil {
		return nil, err
	}
	p.SkipSemicolons()

	var stmts []ast.Stmt
	for p.current.Type != lexer.TOKEN_RBRACE && p.current.Type != lexer.TOKEN_EOF {
		stmt, err := p.ParseStatement()
		if err != nil {
			return nil, err
		}
		if stmt != nil {
			stmts = append(stmts, stmt)
		}
		p.SkipSemicolons()
	}

	if _, err := p.Expect(lexer.TOKEN_RBRACE); err != nil {
		return nil, err
	}

	return stmts, nil
}

func (p *Parser) ParseParamList() ([]string, error) {
	if _, err := p.Expect(lexer.TOKEN_LPAREN); err != nil {
		return nil, err
	}

	var params []string
	if p.current.Type != lexer.TOKEN_RPAREN {
		name, err := p.Expect(lexer.TOKEN_IDENT)
		if err != nil {
			return nil, err
		}
		params = append(params, name.Literal)

		for p.current.Type == lexer.TOKEN_COMMA {
			p.Advance()
			name, err := p.Expect(lexer.TOKEN_IDENT)
			if err != nil {
				return nil, err
			}
			params = append(params, name.Literal)
		}
	}

	if _, err := p.Expect(lexer.TOKEN_RPAREN); err != nil {
		return nil, err
	}

	return params, nil
}



