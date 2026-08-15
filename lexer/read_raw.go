package lexer

import (
	"strings"
)

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
