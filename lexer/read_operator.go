package lexer

// readOperator scans operators and delimiters.
func (l *Lexer) readOperator() (Token, error) {
	line, col := l.line, l.col
	ch := l.advance()

	switch ch {
	case '=':
		return l.makeToken(TOKEN_ASSIGN, "=", line, col), nil
	case '+':
		return l.makeToken(TOKEN_PLUS, "+", line, col), nil
	case '*':
		return l.makeToken(TOKEN_STAR, "*", line, col), nil

	case '%':
		return l.makeToken(TOKEN_PERCENT, "%", line, col), nil
	case '&':
		return l.makeToken(TOKEN_AND, "&", line, col), nil
	case '|':
		return l.makeToken(TOKEN_OR, "|", line, col), nil
	case '?':
		if l.pos < len(l.input) && l.peek() == '.' {
			l.advance()
			return l.makeToken(TOKEN_QMARK_DOT, "?.", line, col), nil
		}
		if l.pos < len(l.input) && l.peek() == '?' {
			l.advance()
			return l.makeToken(TOKEN_QQ, "??", line, col), nil
		}
		return l.makeToken(TOKEN_QUESTION, "?", line, col), nil
	case '(':
		return l.makeToken(TOKEN_LPAREN, "(", line, col), nil
	case ')':
		return l.makeToken(TOKEN_RPAREN, ")", line, col), nil
	case '{':
		return l.makeToken(TOKEN_LBRACE, "{", line, col), nil
	case '}':
		return l.makeToken(TOKEN_RBRACE, "}", line, col), nil
	case '[':
		return l.makeToken(TOKEN_LBRACKET, "[", line, col), nil
	case ']':
		return l.makeToken(TOKEN_RBRACKET, "]", line, col), nil
	case ',':
		return l.makeToken(TOKEN_COMMA, ",", line, col), nil
	case ':':
		return l.makeToken(TOKEN_COLON, ":", line, col), nil
	case ';':
		return l.makeToken(TOKEN_SEMICOLON, ";", line, col), nil

	case '-':
		if l.pos < len(l.input) && l.peek() == '>' {
			l.advance()
			return l.makeToken(TOKEN_ARROW, "->", line, col), nil
		}
		return l.makeToken(TOKEN_MINUS, "-", line, col), nil

	case '!':
		if l.pos < len(l.input) && l.peek() == '=' {
			l.advance()
			return l.makeToken(TOKEN_NOT_EQ, "!=", line, col), nil
		}
		return l.makeToken(TOKEN_NOT, "!", line, col), nil

	case '<':
		if l.pos < len(l.input) && l.peek() == '=' {
			l.advance()
			return l.makeToken(TOKEN_LT_EQ, "<=", line, col), nil
		}
		if l.pos < len(l.input) && l.peek() == '-' {
			l.advance()
			return l.makeToken(TOKEN_LARROW, "<-", line, col), nil
		}
		return l.makeToken(TOKEN_LT, "<", line, col), nil

	case '>':
		if l.pos < len(l.input) && l.peek() == '=' {
			l.advance()
			return l.makeToken(TOKEN_GT_EQ, ">=", line, col), nil
		}
		return l.makeToken(TOKEN_GT, ">", line, col), nil
	}

	return Token{}, l.errorf("E109", "illegal %c", ch)
}
