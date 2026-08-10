package lexer

import (
	"drylang/errfmt"
	"fmt"
	"strings"
	"unicode"
)

// Lexer tokenizes dryLang source code.
type Lexer struct {
	input   []rune
	pos     int // current position
	line    int
	col     int
	tokens  []Token
}

// New creates a new Lexer for the given source.
func New(input string) *Lexer {
	return &Lexer{
		input: []rune(input),
		pos:   0,
		line:  1,
		col:   1,
	}
}

// Tokenize scans the entire input and returns all tokens.
func (l *Lexer) Tokenize() ([]Token, error) {
	for {
		tok, err := l.nextToken()
		if err != nil {
			return nil, err
		}
		l.tokens = append(l.tokens, tok)
		if tok.Type == TOKEN_EOF {
			break
		}
	}
	return l.tokens, nil
}

func (l *Lexer) peek() rune {
	if l.pos >= len(l.input) {
		return 0
	}
	return l.input[l.pos]
}

func (l *Lexer) peekAt(offset int) rune {
	p := l.pos + offset
	if p >= len(l.input) {
		return 0
	}
	return l.input[p]
}

func (l *Lexer) advance() rune {
	if l.pos >= len(l.input) {
		return 0
	}
	ch := l.input[l.pos]
	l.pos++
	if ch == '\n' {
		l.line++
		l.col = 1
	} else {
		l.col++
	}
	return ch
}

func (l *Lexer) skipWhitespace() {
	for l.pos < len(l.input) {
		ch := l.input[l.pos]
		if ch == ' ' || ch == '\t' || ch == '\r' {
			l.advance()
		} else {
			break
		}
	}
}

func (l *Lexer) errorf(code, format string, args ...interface{}) error {
	return errfmt.Format(code, l.line, l.col, fmt.Sprintf(format, args...))
}

func (l *Lexer) makeToken(typ TokenType, literal string, line, col int) Token {
	return Token{Type: typ, Literal: literal, Line: line, Col: col}
}

func (l *Lexer) nextToken() (Token, error) {
	l.skipWhitespace()

	if l.pos >= len(l.input) {
		return l.makeToken(TOKEN_EOF, "", l.line, l.col), nil
	}

	ch := l.peek()
	line, col := l.line, l.col

	// Newline → implicit statement separator
	if ch == '\n' {
		l.advance()
		// Skip consecutive newlines
		for l.pos < len(l.input) && l.input[l.pos] == '\n' {
			l.skipWhitespace()
			if l.pos < len(l.input) && l.input[l.pos] == '\n' {
				l.advance()
			} else {
				break
			}
		}
		return l.makeToken(TOKEN_SEMICOLON, "\\n", line, col), nil
	}

	// Slash or Comments
	if ch == '/' {
		if l.peek() == '/' {
			// Single-line comment
			l.advance() // consume second /
			for l.pos < len(l.input) && l.peek() != '\n' {
				l.advance()
			}
			return l.nextToken() // skip comment, get next
		} else if l.peek() == '*' {
			// Multi-line comment
			l.advance() // consume *
			for l.pos < len(l.input) {
				if l.peek() == '*' && l.peekAt(1) == '/' {
					l.advance() // consume *
					l.advance() // consume /
					break
				}
				if l.peek() == '\n' {
					l.line++
					l.col = 1
					l.pos++
				} else {
					l.advance()
				}
			}
			return l.nextToken() // skip comment, get next
		}
		return l.makeToken(TOKEN_SLASH, "/", line, col), nil
	}

	// Dot (now purely for decimal or property access, but here we just emit TOKEN_DOT)
	if ch == '.' {
		l.advance()
		return l.makeToken(TOKEN_DOT, ".", line, col), nil
	}

	// String literals
	if ch == '"' || ch == '\'' {
		return l.readString(ch)
	}

	// Raw string (backtick)
	if ch == '`' {
		return l.readRawString()
	}

	// Numbers
	if unicode.IsDigit(ch) {
		return l.readNumber()
	}

	// Identifiers and keywords
	if unicode.IsLetter(ch) || ch == '_' {
		return l.readIdentifier()
	}

	// Operators and delimiters
	return l.readOperator()
}

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
				return l.makeToken(TOKEN_INTERP_END, "}", line, col), nil
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

// readRawString scans a raw string literal (backtick).
func (l *Lexer) readRawString() (Token, error) {
	line, col := l.line, l.col
	l.advance() // consume opening `

	var buf strings.Builder
	for l.pos < len(l.input) {
		ch := l.peek()
		if ch == '`' {
			l.advance() // consume closing `
			return l.makeToken(TOKEN_RAW_STRING, buf.String(), line, col), nil
		}
		buf.WriteRune(l.advance())
	}

	return Token{}, l.errorf("E108", "needs closing `")
}

// readNumber scans an integer or float (comma decimal: 89,5).
func (l *Lexer) readNumber() (Token, error) {
	line, col := l.line, l.col
	var buf strings.Builder

	// Read integer part
	for l.pos < len(l.input) && unicode.IsDigit(l.peek()) {
		buf.WriteRune(l.advance())
	}

	// Check for dot decimal (e.g. 3.14)
	if l.pos < len(l.input) && l.peek() == '.' && l.pos+1 < len(l.input) && unicode.IsDigit(l.peekAt(1)) {
		buf.WriteRune('.') 
		l.advance()        // consume dot

		for l.pos < len(l.input) && unicode.IsDigit(l.peek()) {
			buf.WriteRune(l.advance())
		}
	}

	return l.makeToken(TOKEN_NUMBER, buf.String(), line, col), nil
}

// readIdentifier scans an identifier or keyword.
func (l *Lexer) readIdentifier() (Token, error) {
	line, col := l.line, l.col
	var buf strings.Builder

	for l.pos < len(l.input) {
		ch := l.peek()
		if unicode.IsLetter(ch) || unicode.IsDigit(ch) || ch == '_' {
			buf.WriteRune(l.advance())
		} else {
			break
		}
	}

	word := buf.String()
	tokType := LookupIdent(word)
	return l.makeToken(tokType, word, line, col), nil
}

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
