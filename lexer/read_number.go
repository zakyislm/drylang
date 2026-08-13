package lexer

import (
	"strings"
	"unicode"
)

// readNumber scans an integer, float, or scientific-notation literal
// (e.g. 42, 3.14, 1e400, 2.5e-3).
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
		l.advance() // consume dot

		for l.pos < len(l.input) && unicode.IsDigit(l.peek()) {
			buf.WriteRune(l.advance())
		}
	}

	// Scientific notation: 1e400, 2.5e-3, 1E+10
	if l.pos < len(l.input) && (l.peek() == 'e' || l.peek() == 'E') {
		// Must be followed by digits (or sign then digits) to be an exponent,
		// otherwise the 'e' starts an identifier (e.g. `1e` is not a number).
		next := l.peekAt(1)
		if unicode.IsDigit(next) || ((next == '+' || next == '-') && unicode.IsDigit(l.peekAt(2))) {
			buf.WriteRune(l.advance()) // consume e/E
			if l.peek() == '+' || l.peek() == '-' {
				buf.WriteRune(l.advance())
			}
			for l.pos < len(l.input) && unicode.IsDigit(l.peek()) {
				buf.WriteRune(l.advance())
			}
		}
	}

	return l.makeToken(TOKEN_NUMBER, buf.String(), line, col), nil
}
