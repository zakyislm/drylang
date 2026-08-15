package lexer

import (
	"drylang/errfmt"
	"fmt"
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
		if l.peekAt(1) == '/' {
			// Single-line comment
			l.advance() // consume first /
			l.advance() // consume second /
			for l.pos < len(l.input) && l.peek() != '\n' {
				l.advance()
			}
			return l.nextToken() // skip comment, get next
		} else if l.peekAt(1) == '*' {
			// Multi-line comment
			l.advance() // consume first /
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
		
		l.advance() // consume / for division
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

