package lexer

import (
	"strings"
	"unicode"
)

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
