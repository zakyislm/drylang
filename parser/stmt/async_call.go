package stmt

import (
	"drylang/ast"
	"drylang/core"
	"drylang/lexer"
	"strconv"
)

func ParseMulCall(p core.ParserCore) (ast.Stmt, error) {
	tok := p.Advance() // Consume 'mul'

	workers := 2
	if p.Current().Type == lexer.TOKEN_NUMBER {
		wTok := p.Advance()
		w, err := strconv.Atoi(wTok.Literal)
		if err == nil && w > 0 {
			workers = w
		} else {
			return nil, p.Errorf("E110", "bad number for mul workers")
		}
	}

	expr, err := p.ParseExpression(core.PREC_LOWEST)
	if err != nil {
		return nil, err
	}

	callExpr, ok := expr.(*ast.CallExpr)
	if !ok {
		return nil, p.Errorf("E105", "mul needs function call")
	}

	return &ast.MulCallStmt{
		Call:    callExpr,
		Workers: workers,
		Line:    tok.Line,
		Col:     tok.Col,
	}, nil
}

func ParseUniCall(p core.ParserCore) (ast.Stmt, error) {
	tok := p.Advance() // Consume 'uni'

	expr, err := p.ParseExpression(core.PREC_LOWEST)
	if err != nil {
		return nil, err
	}

	callExpr, ok := expr.(*ast.CallExpr)
	if !ok {
		return nil, p.Errorf("E105", "uni needs function call")
	}

	return &ast.MulCallStmt{
		Call:    callExpr,
		Workers: 1,
		Line:    tok.Line,
		Col:     tok.Col,
	}, nil
}
