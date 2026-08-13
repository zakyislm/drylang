package lexer

import (
	"strings"
	"unicode"
)

// readString scans a string with ${} interpolation support.
// Returns STRING tokens for simple strings, or STRING_PART + INTERP tokens for interpolated strings.
func (l *Lexer) readString(quote rune) (Token, error) {
	line, col := l.line, l.col
	l.advance() // consume opening quote

	var buf strings.Builder
	hasInterp := false

	for l.pos < len(l.input) {
		ch := l.peek()

		if ch == quote {
			l.advance() // consume closing quote
			if hasInterp {
				// Emit final string part
				if buf.Len() > 0 {
					l.tokens = append(l.tokens, l.makeToken(TOKEN_STRING_PART, buf.String(), line, col))
				}
				return l.nextToken()
			}
			return l.makeToken(TOKEN_STRING, buf.String(), line, col), nil
		}

		if ch == '\\' {
			l.advance() // consume backslash
			esc := l.advance()
			switch esc {
			case 'n':
				buf.WriteRune('\n')
			case 't':
				buf.WriteRune('\t')
			case '\\':
				buf.WriteRune('\\')
			case '\'':
				buf.WriteRune('\'')
			case '"':
				buf.WriteRune('"')
			case '$':
				buf.WriteRune('$')
			case '0':
				buf.WriteRune(0)
			default:
				buf.WriteRune('\\')
				buf.WriteRune(esc)
			}
			continue
		}

		// Interpolation: ${expr}
		if ch == '$' && l.peekAt(1) == '{' {
			// Emit string part before interpolation
			if buf.Len() > 0 || !hasInterp {
				l.tokens = append(l.tokens, l.makeToken(TOKEN_STRING_PART, buf.String(), line, col))
				buf.Reset()
			}
			l.advance() // consume $
			l.advance() // consume {
			l.tokens = append(l.tokens, l.makeToken(TOKEN_INTERP_START, "${", l.line, l.col))

			// Tokenize the interpolation expression until }
			depth := 1
			for l.pos < len(l.input) && depth > 0 {
				interpTok, err := l.nextInterpToken()
				if err != nil {
					return Token{}, err
				}
				if interpTok.Type == TOKEN_LBRACE {
					depth++
				}
				if interpTok.Type == TOKEN_RBRACE {
					depth--
					if depth == 0 {
						l.tokens = append(l.tokens, l.makeToken(TOKEN_INTERP_END, "}", l.line, l.col))
						break
					}
				}
				if interpTok.Type == TOKEN_EOF {
					return Token{}, l.errorf("E108", "needs closing } for interp")
				}
				l.tokens = append(l.tokens, interpTok)
			}

			hasInterp = true
			col = l.col
			continue
		}

		if ch == '\n' {
			return Token{}, l.errorf("E108", "needs closing %c", quote)
		}

		buf.WriteRune(l.advance())
	}

	return Token{}, l.errorf("E108", "needs closing %c", quote)
}

// nextInterpToken scans a single token inside ${} interpolation.
func (l *Lexer) nextInterpToken() (Token, error) {
	l.skipWhitespace()
	if l.pos >= len(l.input) {
		return l.makeToken(TOKEN_EOF, "", l.line, l.col), nil
	}

	ch := l.peek()

	if unicode.IsDigit(ch) {
		return l.readNumber()
	}
	if unicode.IsLetter(ch) || ch == '_' {
		return l.readIdentifier()
	}
	return l.readOperator()
}
