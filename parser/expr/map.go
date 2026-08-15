package expr

import (
	"drylang/ast"
	"drylang/core"
	"drylang/lexer"
)

func ParseMapLit(p core.ParserCore) (ast.Expr, error) {
	tok := p.Advance() // Consume '{'
	p.SkipSemicolons()
	keys := []ast.Expr{}
	values := []ast.Expr{}

	if p.Current().Type != lexer.TOKEN_RBRACE {
		key, err := p.ParseExpression(core.PREC_LOWEST)
		if err != nil {
			return nil, err
		}
		
		var val ast.Expr
		if p.Current().Type == lexer.TOKEN_COLON {
			p.Advance() // Consume ':'
			val, err = p.ParseExpression(core.PREC_LOWEST)
			if err != nil {
				return nil, err
			}
		} else {
			// Shorthand syntax: {x} means {x: x} (or just for destructuring)
			val = key
		}
		
		keys = append(keys, key)
		values = append(values, val)

		for p.Current().Type == lexer.TOKEN_COMMA {
			p.Advance() // Consume ','
			p.SkipSemicolons()
			if p.Current().Type == lexer.TOKEN_RBRACE {
				break
			}
			key, err := p.ParseExpression(core.PREC_LOWEST)
			if err != nil {
				return nil, err
			}
			if p.Current().Type == lexer.TOKEN_COLON {
				p.Advance() // Consume ':'
				val, err = p.ParseExpression(core.PREC_LOWEST)
				if err != nil {
					return nil, err
				}
			} else {
				val = key
			}
			keys = append(keys, key)
			values = append(values, val)
		}
	}

	if _, err := p.Expect(lexer.TOKEN_RBRACE); err != nil {
		return nil, err
	}

	return &ast.MapLit{Keys: keys, Values: values, Line: tok.Line, Col: tok.Col}, nil
}
