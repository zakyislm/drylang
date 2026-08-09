package parser

import (
	"drylang/errfmt"
	"drylang/lexer"
	"fmt"
	"strings"
)

// Precedence levels for Pratt parser.
const (
	_ int = iota
	PREC_LOWEST
	PREC_OR         // |
	PREC_AND        // &
	PREC_EQUALITY   // = !=
	PREC_COMPARISON // < > <= >=
	PREC_SUM        // + -
	PREC_PRODUCT    // * / %
	PREC_UNARY      // ! -
	PREC_CALL       // fn() arr[0] obj.key
)

var precedences = map[lexer.TokenType]int{
	lexer.TOKEN_OR:      PREC_OR,
	lexer.TOKEN_AND:     PREC_AND,
	lexer.TOKEN_ASSIGN:  PREC_EQUALITY,
	lexer.TOKEN_NOT_EQ:  PREC_EQUALITY,
	lexer.TOKEN_LT:      PREC_COMPARISON,
	lexer.TOKEN_GT:      PREC_COMPARISON,
	lexer.TOKEN_LT_EQ:   PREC_COMPARISON,
	lexer.TOKEN_GT_EQ:   PREC_COMPARISON,
	lexer.TOKEN_PLUS:    PREC_SUM,
	lexer.TOKEN_MINUS:   PREC_SUM,
	lexer.TOKEN_STAR:    PREC_PRODUCT,
	lexer.TOKEN_SLASH:   PREC_PRODUCT,
	lexer.TOKEN_PERCENT: PREC_PRODUCT,
	lexer.TOKEN_LPAREN:  PREC_CALL,
	lexer.TOKEN_LBRACKET: PREC_CALL,
	lexer.TOKEN_DOT:     PREC_CALL,
}

// Parser parses dryLang tokens into an AST.
type Parser struct {
	tokens  []lexer.Token
	pos     int
	current lexer.Token
}

// New creates a parser from tokens.
func New(tokens []lexer.Token) *Parser {
	p := &Parser{tokens: tokens, pos: 0}
	if len(tokens) > 0 {
		p.current = tokens[0]
	}
	return p
}

func (p *Parser) peek() lexer.Token {
	if p.pos+1 < len(p.tokens) {
		return p.tokens[p.pos+1]
	}
	return lexer.Token{Type: lexer.TOKEN_EOF}
}

func (p *Parser) advance() lexer.Token {
	tok := p.current
	p.pos++
	if p.pos < len(p.tokens) {
		p.current = p.tokens[p.pos]
	} else {
		p.current = lexer.Token{Type: lexer.TOKEN_EOF}
	}
	return tok
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

func (p *Parser) expect(typ lexer.TokenType) (lexer.Token, error) {
	if p.current.Type != typ {
		code := errorCodeForMissingToken(typ)
		return lexer.Token{}, p.errorf(code, "needs %s", typ)
	}
	return p.advance(), nil
}

func (p *Parser) errorf(code, format string, args ...interface{}) error {
	return errfmt.Format(code, p.current.Line, p.current.Col, fmt.Sprintf(format, args...))
}

func (p *Parser) skipSemicolons() {
	for p.current.Type == lexer.TOKEN_SEMICOLON {
		p.advance()
	}
}

func (p *Parser) peekPrecedence() int {
	if prec, ok := precedences[p.current.Type]; ok {
		return prec
	}
	return PREC_LOWEST
}

// Parse parses the token stream into a Program AST.
func (p *Parser) Parse() (*Program, error) {
	program := &Program{}
	p.skipSemicolons()

	for p.current.Type != lexer.TOKEN_EOF {
		stmt, err := p.parseStatement()
		if err != nil {
			return nil, err
		}
		if stmt != nil {
			program.Stmts = append(program.Stmts, stmt)
		}
		p.skipSemicolons()
	}

	return program, nil
}

func (p *Parser) parseStatement() (Stmt, error) {
	switch p.current.Type {
	case lexer.TOKEN_CNS:
		return p.parseConstDecl()
	case lexer.TOKEN_FN:
		return p.parseFnDecl(false)
	case lexer.TOKEN_ASN:
		return p.parseAsyncFnDecl()
	case lexer.TOKEN_REV:
		return p.parseReturn()
	case lexer.TOKEN_IF:
		return p.parseIf()
	case lexer.TOKEN_ON:
		return p.parseOn()
	case lexer.TOKEN_LP:
		return p.parseLoop()
	case lexer.TOKEN_DONE:
		tok := p.advance()
		return &DoneStmt{Line: tok.Line, Col: tok.Col}, nil
	case lexer.TOKEN_CON:
		tok := p.advance()
		return &ConStmt{Line: tok.Line, Col: tok.Col}, nil
	case lexer.TOKEN_TRY:
		return p.parseTry()
	case lexer.TOKEN_ERR:
		return p.parseThrow()
	case lexer.TOKEN_USE:
		return p.parseUse()
	case lexer.TOKEN_PV:
		return p.parsePrivate()
	case lexer.TOKEN_PT:
		return p.parsePrint()
	case lexer.TOKEN_QUESTION:
		return p.parseUnknownBool()
	default:
		return p.parseExpressionOrAssign()
	}
}

func (p *Parser) parseConstDecl() (Stmt, error) {
	tok := p.advance() // consume cns
	name, err := p.expect(lexer.TOKEN_IDENT)
	if err != nil {
		return nil, err
	}

	// Optional =
	if p.current.Type == lexer.TOKEN_ASSIGN {
		p.advance()
	}

	val, err := p.parseExpression(PREC_LOWEST)
	if err != nil {
		return nil, err
	}

	return &ConstDeclStmt{Name: name.Literal, Value: val, Line: tok.Line, Col: tok.Col}, nil
}

func (p *Parser) parseFnDecl(isAsync bool) (Stmt, error) {
	tok := p.advance() // consume fn
	name, err := p.expect(lexer.TOKEN_IDENT)
	if err != nil {
		return nil, err
	}

	params, err := p.parseParamList()
	if err != nil {
		return nil, err
	}

	body, err := p.parseBlock()
	if err != nil {
		return nil, err
	}

	return &FnDeclStmt{
		Name: name.Literal, Params: params, Body: body,
		IsAsync: isAsync, Line: tok.Line, Col: tok.Col,
	}, nil
}

func (p *Parser) parseAsyncFnDecl() (Stmt, error) {
	p.advance() // consume asn
	if p.current.Type != lexer.TOKEN_FN {
		return nil, p.errorf("E105", "needs fn")
	}
	return p.parseFnDecl(true)
}

func (p *Parser) parseReturn() (Stmt, error) {
	tok := p.advance() // consume rev
	var val Expr

	// rev can be without value (returns unknown)
	if p.current.Type != lexer.TOKEN_SEMICOLON && p.current.Type != lexer.TOKEN_RBRACE && p.current.Type != lexer.TOKEN_EOF {
		var err error
		val, err = p.parseExpression(PREC_LOWEST)
		if err != nil {
			return nil, err
		}
	}

	return &ReturnStmt{Value: val, Line: tok.Line, Col: tok.Col}, nil
}

func (p *Parser) parseIf() (Stmt, error) {
	tok := p.advance() // consume if
	cond, err := p.parseExpression(PREC_LOWEST)
	if err != nil {
		return nil, err
	}

	body, err := p.parseBlock()
	if err != nil {
		return nil, err
	}

	stmt := &IfStmt{Condition: cond, Body: body, Line: tok.Line, Col: tok.Col}

	p.skipSemicolons()

	// Parse elif chains
	for p.current.Type == lexer.TOKEN_ELIF {
		p.advance()
		elifCond, err := p.parseExpression(PREC_LOWEST)
		if err != nil {
			return nil, err
		}
		elifBody, err := p.parseBlock()
		if err != nil {
			return nil, err
		}
		stmt.ElIfs = append(stmt.ElIfs, ElIfClause{Condition: elifCond, Body: elifBody})
		p.skipSemicolons()
	}

	// Parse el (else)
	if p.current.Type == lexer.TOKEN_EL {
		p.advance()
		elseBody, err := p.parseBlock()
		if err != nil {
			return nil, err
		}
		stmt.Else = elseBody
	}

	return stmt, nil
}

func (p *Parser) parseOn() (Stmt, error) {
	tok := p.advance() // consume on

	if _, err := p.expect(lexer.TOKEN_LPAREN); err != nil {
		return nil, err
	}
	val, err := p.parseExpression(PREC_LOWEST)
	if err != nil {
		return nil, err
	}
	if _, err := p.expect(lexer.TOKEN_RPAREN); err != nil {
		return nil, err
	}

	if _, err := p.expect(lexer.TOKEN_LBRACE); err != nil {
		return nil, err
	}
	p.skipSemicolons()

	stmt := &OnStmt{Value: val, Line: tok.Line, Col: tok.Col}

	for p.current.Type != lexer.TOKEN_RBRACE && p.current.Type != lexer.TOKEN_EOF {
		caseVal, err := p.parseExpression(PREC_LOWEST)
		if err != nil {
			return nil, err
		}
		caseBody, err := p.parseBlock()
		if err != nil {
			return nil, err
		}
		stmt.Cases = append(stmt.Cases, OnCase{Value: caseVal, Body: caseBody})
		p.skipSemicolons()
	}

	if _, err := p.expect(lexer.TOKEN_RBRACE); err != nil {
		return nil, err
	}

	return stmt, nil
}

func (p *Parser) parseLoop() (Stmt, error) {
	tok := p.advance() // consume lp
	var limit Expr

	// lp N { ... } or lp { ... }
	if p.current.Type != lexer.TOKEN_LBRACE {
		var err error
		limit, err = p.parseExpression(PREC_LOWEST)
		if err != nil {
			return nil, err
		}
	}

	body, err := p.parseBlock()
	if err != nil {
		return nil, err
	}

	return &LoopStmt{Limit: limit, Body: body, Line: tok.Line, Col: tok.Col}, nil
}

func (p *Parser) parseTry() (Stmt, error) {
	tok := p.advance() // consume try
	body, err := p.parseBlock()
	if err != nil {
		return nil, err
	}

	p.skipSemicolons()

	if p.current.Type != lexer.TOKEN_ERR {
		return nil, p.errorf("E106", "needs err")
	}
	p.advance() // consume err

	if _, err := p.expect(lexer.TOKEN_LPAREN); err != nil {
		return nil, err
	}
	errName, err := p.expect(lexer.TOKEN_IDENT)
	if err != nil {
		return nil, err
	}
	if _, err := p.expect(lexer.TOKEN_RPAREN); err != nil {
		return nil, err
	}

	catchBody, err := p.parseBlock()
	if err != nil {
		return nil, err
	}

	return &TryStmt{Body: body, ErrName: errName.Literal, Catch: catchBody, Line: tok.Line, Col: tok.Col}, nil
}

func (p *Parser) parseThrow() (Stmt, error) {
	tok := p.advance() // consume err
	val, err := p.parseExpression(PREC_LOWEST)
	if err != nil {
		return nil, err
	}
	return &ThrowStmt{Value: val, Line: tok.Line, Col: tok.Col}, nil
}

func (p *Parser) parseUse() (Stmt, error) {
	tok := p.advance() // consume use
	path, err := p.expect(lexer.TOKEN_STRING)
	if err != nil {
		return nil, err
	}
	return &UseStmt{Path: path.Literal, Line: tok.Line, Col: tok.Col}, nil
}

func (p *Parser) parsePrivate() (Stmt, error) {
	tok := p.advance() // consume pv
	inner, err := p.parseStatement()
	if err != nil {
		return nil, err
	}
	return &PrivateStmt{Inner: inner, Line: tok.Line, Col: tok.Col}, nil
}

func (p *Parser) parsePrint() (Stmt, error) {
	tok := p.advance() // consume pt
	val, err := p.parseExpression(PREC_LOWEST)
	if err != nil {
		return nil, err
	}
	return &PrintStmt{Value: val, Line: tok.Line, Col: tok.Col}, nil
}

func (p *Parser) parseUnknownBool() (Stmt, error) {
	tok := p.advance() // consume ?
	name, err := p.expect(lexer.TOKEN_IDENT)
	if err != nil {
		return nil, err
	}
	return &UnknownBoolStmt{Name: name.Literal, Line: tok.Line, Col: tok.Col}, nil
}

// parseExpressionOrAssign handles:
// - identifier = expr (assignment with =)
// - identifier expr (assignment without =, space-separated)
// - identifier { fields } (struct declaration — fields are bare idents)
// - varname structname { field val ... } (struct instance)
// - expression statement (function call, etc.)
func (p *Parser) parseExpressionOrAssign() (Stmt, error) {
	// Check for struct declaration: ident { bareIdent bareIdent ... }
	if p.current.Type == lexer.TOKEN_IDENT && p.peek().Type == lexer.TOKEN_LBRACE {
		if p.isStructDecl() {
			return p.parseStructDecl()
		}
	}
	// Check for ident expr (space-separated assignment) before parsing expression
	// This avoids parsing it as two separate expressions or failing
	if p.current.Type == lexer.TOKEN_IDENT {
		if p.peek().Type != lexer.TOKEN_SEMICOLON && p.peek().Type != lexer.TOKEN_EOF &&
			p.peek().Type != lexer.TOKEN_RBRACE && p.peek().Type != lexer.TOKEN_ASSIGN &&
			p.peekPrecedence() == PREC_LOWEST {
			
			// We might be looking at `ident expr`
			// Wait, let's just let it fall through to the expression parser,
			// and if it's an Ident, we can check for space-separated assignment afterwards.
		}
	}

	expr, err := p.parseExpression(PREC_LOWEST)
	if err != nil {
		return nil, err
	}

	// 1. Check if the parsed expression is a BinaryExpr with '='
	// Since '=' has PREC_EQUALITY, it parses as an infix operator.
	// But in a statement context, it's an assignment!
	if binExpr, ok := expr.(*BinaryExpr); ok && binExpr.Op == lexer.TOKEN_ASSIGN {
		switch target := binExpr.Left.(type) {
		case *Ident:
			if isAllUpper(target.Name) {
				return &ConstDeclStmt{Name: target.Name, Value: binExpr.Right, Line: target.Line, Col: target.Col}, nil
			}
			return &AssignStmt{Name: target.Name, Value: binExpr.Right, Line: target.Line, Col: target.Col}, nil
		case *IndexExpr:
			return &IndexAssignStmt{Object: target.Object, Index: target.Index, Value: binExpr.Right, Line: target.Line, Col: target.Col}, nil
		case *DotExpr:
			return &DotAssignStmt{Object: target.Object, Field: target.Field, Value: binExpr.Right, Line: target.Line, Col: target.Col}, nil
		}
	}

	// 2. Check for space-separated assignment (ident expr)
	if ident, isIdent := expr.(*Ident); isIdent {
		// Only if next token starts an expression (not end of statement)
		if p.current.Type != lexer.TOKEN_SEMICOLON && p.current.Type != lexer.TOKEN_EOF &&
			p.current.Type != lexer.TOKEN_RBRACE && p.isExprStart() {
			val, err := p.parseExpression(PREC_LOWEST)
			if err != nil {
				return nil, err
			}
			if isAllUpper(ident.Name) {
				return &ConstDeclStmt{Name: ident.Name, Value: val, Line: ident.Line, Col: ident.Col}, nil
			}
			return &AssignStmt{Name: ident.Name, Value: val, Line: ident.Line, Col: ident.Col}, nil
		}
	}

	return &ExprStmt{Expression: expr, Line: expr.TokenLine(), Col: expr.TokenCol()}, nil
}

func (p *Parser) isExprStart() bool {
	switch p.current.Type {
	case lexer.TOKEN_NUMBER, lexer.TOKEN_STRING, lexer.TOKEN_RAW_STRING,
		lexer.TOKEN_STRING_PART, lexer.TOKEN_TRUE, lexer.TOKEN_FALSE,
		lexer.TOKEN_UNKNOWN, lexer.TOKEN_IDENT, lexer.TOKEN_LPAREN,
		lexer.TOKEN_LBRACKET, lexer.TOKEN_MINUS, lexer.TOKEN_NOT,
		lexer.TOKEN_IN:
		return true
	}
	return false
}

// isStructDecl peeks ahead to check if ident { ... } is a struct declaration.
// Struct decl: all content inside {} is bare identifiers (fields), no values.
func (p *Parser) isStructDecl() bool {
	saved := p.pos
	defer func() { p.pos = saved; p.current = p.tokens[p.pos] }()

	p.advance() // skip ident
	p.advance() // skip {

	for p.current.Type != lexer.TOKEN_RBRACE && p.current.Type != lexer.TOKEN_EOF {
		if p.current.Type == lexer.TOKEN_SEMICOLON {
			p.advance()
			continue
		}
		if p.current.Type != lexer.TOKEN_IDENT {
			return false // has non-ident content → not struct decl
		}
		p.advance()
		// If next token is also ident or semicolon or rbrace → still struct decl
		// If next token is a value (string, number, etc.) → struct init, not decl
		if p.current.Type != lexer.TOKEN_IDENT && p.current.Type != lexer.TOKEN_SEMICOLON &&
			p.current.Type != lexer.TOKEN_RBRACE {
			return false
		}
	}
	return true
}

func (p *Parser) parseStructDecl() (Stmt, error) {
	name := p.advance() // struct name
	p.advance()         // consume {
	p.skipSemicolons()

	var fields []string
	for p.current.Type != lexer.TOKEN_RBRACE && p.current.Type != lexer.TOKEN_EOF {
		field, err := p.expect(lexer.TOKEN_IDENT)
		if err != nil {
			return nil, err
		}
		fields = append(fields, field.Literal)
		p.skipSemicolons()
	}

	if _, err := p.expect(lexer.TOKEN_RBRACE); err != nil {
		return nil, err
	}

	return &StructDeclStmt{Name: name.Literal, Fields: fields, Line: name.Line, Col: name.Col}, nil
}

func (p *Parser) parseStructInit(varName string) (Stmt, error) {
	line, col := p.current.Line, p.current.Col
	typeName := p.advance() // struct type name
	p.advance()             // consume {
	p.skipSemicolons()

	fields := make(map[string]Expr)
	for p.current.Type != lexer.TOKEN_RBRACE && p.current.Type != lexer.TOKEN_EOF {
		fieldName, err := p.expect(lexer.TOKEN_IDENT)
		if err != nil {
			return nil, err
		}
		val, err := p.parseExpression(PREC_LOWEST)
		if err != nil {
			return nil, err
		}
		fields[fieldName.Literal] = val
		p.skipSemicolons()
	}

	if _, err := p.expect(lexer.TOKEN_RBRACE); err != nil {
		return nil, err
	}

	initExpr := &StructInitExpr{TypeName: typeName.Literal, Fields: fields, Line: line, Col: col}
	return &AssignStmt{Name: varName, Value: initExpr, Line: line, Col: col}, nil
}

// parseExpression is the core Pratt parser.
func (p *Parser) parseExpression(prec int) (Expr, error) {
	left, err := p.parsePrefix()
	if err != nil {
		return nil, err
	}

	for prec < p.peekPrecedence() {
		left, err = p.parseInfix(left)
		if err != nil {
			return nil, err
		}
	}

	return left, nil
}

func (p *Parser) parsePrefix() (Expr, error) {
	switch p.current.Type {
	case lexer.TOKEN_NUMBER:
		tok := p.advance()
		return &NumberLit{Value: tok.Literal, Line: tok.Line, Col: tok.Col}, nil

	case lexer.TOKEN_STRING:
		tok := p.advance()
		return &StringLit{Value: tok.Literal, Line: tok.Line, Col: tok.Col}, nil

	case lexer.TOKEN_RAW_STRING:
		tok := p.advance()
		return &RawStringLit{Value: tok.Literal, Line: tok.Line, Col: tok.Col}, nil

	case lexer.TOKEN_STRING_PART:
		return p.parseStringInterp()

	case lexer.TOKEN_TRUE:
		tok := p.advance()
		return &BoolLit{Value: true, Line: tok.Line, Col: tok.Col}, nil

	case lexer.TOKEN_FALSE:
		tok := p.advance()
		return &BoolLit{Value: false, Line: tok.Line, Col: tok.Col}, nil

	case lexer.TOKEN_UNKNOWN:
		tok := p.advance()
		return &UnknownLit{Line: tok.Line, Col: tok.Col}, nil

	case lexer.TOKEN_IDENT:
		tok := p.advance()
		return &Ident{Name: tok.Literal, Line: tok.Line, Col: tok.Col}, nil

	case lexer.TOKEN_LPAREN:
		p.advance()
		expr, err := p.parseExpression(PREC_LOWEST)
		if err != nil {
			return nil, err
		}
		if _, err := p.expect(lexer.TOKEN_RPAREN); err != nil {
			return nil, err
		}
		return expr, nil

	case lexer.TOKEN_LBRACKET:
		return p.parseArrayLit()

	case lexer.TOKEN_LBRACE:
		return p.parseMapLit()

	case lexer.TOKEN_MINUS:
		tok := p.advance()
		operand, err := p.parseExpression(PREC_UNARY)
		if err != nil {
			return nil, err
		}
		return &UnaryExpr{Op: tok.Type, Operand: operand, Line: tok.Line, Col: tok.Col}, nil

	case lexer.TOKEN_NOT:
		tok := p.advance()
		operand, err := p.parseExpression(PREC_UNARY)
		if err != nil {
			return nil, err
		}
		return &UnaryExpr{Op: tok.Type, Operand: operand, Line: tok.Line, Col: tok.Col}, nil

	case lexer.TOKEN_ARROW:
		return p.parseArrowFn()

	case lexer.TOKEN_IN:
		return p.parseInput()

	case lexer.TOKEN_AWT:
		tok := p.advance()
		val, err := p.parseExpression(PREC_UNARY)
		if err != nil {
			return nil, err
		}
		return &AwaitExpr{Value: val, Line: tok.Line, Col: tok.Col}, nil

	default:
		return nil, p.errorf("E109", "illegal %s", p.current.Literal)
	}
}

func (p *Parser) parseInfix(left Expr) (Expr, error) {
	switch p.current.Type {
	case lexer.TOKEN_PLUS, lexer.TOKEN_MINUS, lexer.TOKEN_STAR, lexer.TOKEN_SLASH,
		lexer.TOKEN_PERCENT, lexer.TOKEN_LT, lexer.TOKEN_GT, lexer.TOKEN_LT_EQ,
		lexer.TOKEN_GT_EQ, lexer.TOKEN_ASSIGN, lexer.TOKEN_NOT_EQ,
		lexer.TOKEN_AND, lexer.TOKEN_OR:
		tok := p.advance()
		prec := precedences[tok.Type]
		right, err := p.parseExpression(prec)
		if err != nil {
			return nil, err
		}
		return &BinaryExpr{Left: left, Op: tok.Type, Right: right, Line: tok.Line, Col: tok.Col}, nil

	case lexer.TOKEN_LPAREN:
		return p.parseCallExpr(left)

	case lexer.TOKEN_LBRACKET:
		return p.parseIndexExpr(left)

	case lexer.TOKEN_DOT:
		return p.parseDotExpr(left)

	default:
		return left, nil
	}
}

func (p *Parser) parseCallExpr(callee Expr) (Expr, error) {
	tok := p.advance() // consume (
	var args []Expr

	if p.current.Type != lexer.TOKEN_RPAREN {
		arg, err := p.parseExpression(PREC_LOWEST)
		if err != nil {
			return nil, err
		}
		args = append(args, arg)

		for p.current.Type == lexer.TOKEN_COMMA {
			p.advance()
			arg, err := p.parseExpression(PREC_LOWEST)
			if err != nil {
				return nil, err
			}
			args = append(args, arg)
		}
	}

	if _, err := p.expect(lexer.TOKEN_RPAREN); err != nil {
		return nil, err
	}

	return &CallExpr{Callee: callee, Args: args, Line: tok.Line, Col: tok.Col}, nil
}

func (p *Parser) parseIndexExpr(obj Expr) (Expr, error) {
	tok := p.advance() // consume [
	idx, err := p.parseExpression(PREC_LOWEST)
	if err != nil {
		return nil, err
	}
	if _, err := p.expect(lexer.TOKEN_RBRACKET); err != nil {
		return nil, err
	}
	return &IndexExpr{Object: obj, Index: idx, Line: tok.Line, Col: tok.Col}, nil
}

func (p *Parser) parseDotExpr(obj Expr) (Expr, error) {
	p.advance() // consume .
	field, err := p.expect(lexer.TOKEN_IDENT)
	if err != nil {
		return nil, err
	}
	return &DotExpr{Object: obj, Field: field.Literal, Line: field.Line, Col: field.Col}, nil
}

func (p *Parser) parseArrayLit() (Expr, error) {
	tok := p.advance() // consume [
	var items []Expr

	p.skipSemicolons()
	if p.current.Type != lexer.TOKEN_RBRACKET {
		item, err := p.parseExpression(PREC_LOWEST)
		if err != nil {
			return nil, err
		}
		items = append(items, item)

		for p.current.Type == lexer.TOKEN_COMMA {
			p.advance()
			p.skipSemicolons()
			if p.current.Type == lexer.TOKEN_RBRACKET {
				break // trailing comma
			}
			item, err := p.parseExpression(PREC_LOWEST)
			if err != nil {
				return nil, err
			}
			items = append(items, item)
		}
	}
	p.skipSemicolons()

	if _, err := p.expect(lexer.TOKEN_RBRACKET); err != nil {
		return nil, err
	}

	return &ArrayLit{Items: items, Line: tok.Line, Col: tok.Col}, nil
}

func (p *Parser) parseMapLit() (Expr, error) {
	tok := p.advance() // consume {
	var keys, values []Expr

	p.skipSemicolons()
	if p.current.Type != lexer.TOKEN_RBRACE {
		k, v, err := p.parseMapEntry()
		if err != nil {
			return nil, err
		}
		keys = append(keys, k)
		values = append(values, v)

		for p.current.Type == lexer.TOKEN_COMMA {
			p.advance()
			p.skipSemicolons()
			if p.current.Type == lexer.TOKEN_RBRACE {
				break
			}
			k, v, err := p.parseMapEntry()
			if err != nil {
				return nil, err
			}
			keys = append(keys, k)
			values = append(values, v)
		}
	}
	p.skipSemicolons()

	if _, err := p.expect(lexer.TOKEN_RBRACE); err != nil {
		return nil, err
	}

	return &MapLit{Keys: keys, Values: values, Line: tok.Line, Col: tok.Col}, nil
}

func (p *Parser) parseMapEntry() (Expr, Expr, error) {
	key, err := p.parseExpression(PREC_LOWEST)
	if err != nil {
		return nil, nil, err
	}
	if _, err := p.expect(lexer.TOKEN_COLON); err != nil {
		return nil, nil, err
	}
	val, err := p.parseExpression(PREC_LOWEST)
	if err != nil {
		return nil, nil, err
	}
	return key, val, nil
}

func (p *Parser) parseArrowFn() (Expr, error) {
	tok := p.advance() // consume ->
	params, err := p.parseParamList()
	if err != nil {
		return nil, err
	}
	body, err := p.parseBlock()
	if err != nil {
		return nil, err
	}
	return &ArrowFnExpr{Params: params, Body: body, Line: tok.Line, Col: tok.Col}, nil
}

func (p *Parser) parseInput() (Expr, error) {
	tok := p.advance() // consume in
	var prompt Expr

	if p.current.Type == lexer.TOKEN_LPAREN {
		p.advance()
		if p.current.Type != lexer.TOKEN_RPAREN {
			var err error
			prompt, err = p.parseExpression(PREC_LOWEST)
			if err != nil {
				return nil, err
			}
		}
		if _, err := p.expect(lexer.TOKEN_RPAREN); err != nil {
			return nil, err
		}
	}

	return &InputExpr{Prompt: prompt, Line: tok.Line, Col: tok.Col}, nil
}

func (p *Parser) parseStringInterp() (Expr, error) {
	line, col := p.current.Line, p.current.Col
	var parts []Expr

	for {
		if p.current.Type == lexer.TOKEN_STRING_PART {
			tok := p.advance()
			if tok.Literal != "" {
				parts = append(parts, &StringLit{Value: tok.Literal, Line: tok.Line, Col: tok.Col})
			}
		} else if p.current.Type == lexer.TOKEN_INTERP_START {
			p.advance() // consume ${
			expr, err := p.parseExpression(PREC_LOWEST)
			if err != nil {
				return nil, err
			}
			parts = append(parts, expr)
			if p.current.Type == lexer.TOKEN_INTERP_END {
				p.advance()
			}
		} else if p.current.Type == lexer.TOKEN_INTERP_END {
			p.advance() // final closing
			break
		} else {
			break
		}
	}

	if len(parts) == 1 {
		return parts[0], nil
	}

	return &StringInterp{Parts: parts, Line: line, Col: col}, nil
}

func (p *Parser) parseParamList() ([]string, error) {
	if _, err := p.expect(lexer.TOKEN_LPAREN); err != nil {
		return nil, err
	}

	var params []string
	if p.current.Type != lexer.TOKEN_RPAREN {
		name, err := p.expect(lexer.TOKEN_IDENT)
		if err != nil {
			return nil, err
		}
		params = append(params, name.Literal)

		for p.current.Type == lexer.TOKEN_COMMA {
			p.advance()
			name, err := p.expect(lexer.TOKEN_IDENT)
			if err != nil {
				return nil, err
			}
			params = append(params, name.Literal)
		}
	}

	if _, err := p.expect(lexer.TOKEN_RPAREN); err != nil {
		return nil, err
	}

	return params, nil
}

func (p *Parser) parseBlock() ([]Stmt, error) {
	if _, err := p.expect(lexer.TOKEN_LBRACE); err != nil {
		return nil, err
	}
	p.skipSemicolons()

	var stmts []Stmt
	for p.current.Type != lexer.TOKEN_RBRACE && p.current.Type != lexer.TOKEN_EOF {
		stmt, err := p.parseStatement()
		if err != nil {
			return nil, err
		}
		if stmt != nil {
			stmts = append(stmts, stmt)
		}
		p.skipSemicolons()
	}

	if _, err := p.expect(lexer.TOKEN_RBRACE); err != nil {
		return nil, err
	}

	return stmts, nil
}

func isAllUpper(s string) bool {
	if len(s) == 0 {
		return false
	}
	for _, r := range s {
		if r == '_' {
			continue
		}
		if r < 'A' || r > 'Z' {
			return false
		}
	}
	return true
}

// Ensure strings import is used
var _ = strings.Builder{}
