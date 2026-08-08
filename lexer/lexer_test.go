package lexer

import (
	"testing"
)

func TestBasicTokens(t *testing.T) {
	input := `name "Zaky"
age 17
batteryhealth 89,5`

	lex := New(input)
	tokens, err := lex.Tokenize()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := []struct {
		typ TokenType
		lit string
	}{
		{TOKEN_IDENT, "name"},
		{TOKEN_STRING, "Zaky"},
		{TOKEN_SEMICOLON, "\\n"},
		{TOKEN_IDENT, "age"},
		{TOKEN_NUMBER, "17"},
		{TOKEN_SEMICOLON, "\\n"},
		{TOKEN_IDENT, "batteryhealth"},
		{TOKEN_NUMBER, "89.5"}, // comma → dot internally
		{TOKEN_EOF, ""},
	}

	if len(tokens) != len(expected) {
		t.Fatalf("want %d tokens, got %d", len(expected), len(tokens))
		for i, tok := range tokens {
			t.Logf("  [%d] %s %q", i, tok.Type, tok.Literal)
		}
	}

	for i, exp := range expected {
		if tokens[i].Type != exp.typ || tokens[i].Literal != exp.lit {
			t.Errorf("[%d] want %s %q, got %s %q", i, exp.typ, exp.lit, tokens[i].Type, tokens[i].Literal)
		}
	}
}

func TestKeywords(t *testing.T) {
	input := `fn rev if elif el on lp done con asn awt try err pv use cns t f unknown pt in`

	lex := New(input)
	tokens, err := lex.Tokenize()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := []TokenType{
		TOKEN_FN, TOKEN_REV, TOKEN_IF, TOKEN_ELIF, TOKEN_EL,
		TOKEN_ON, TOKEN_LP, TOKEN_DONE, TOKEN_CON, TOKEN_ASN,
		TOKEN_AWT, TOKEN_TRY, TOKEN_ERR, TOKEN_PV, TOKEN_USE,
		TOKEN_CNS, TOKEN_TRUE, TOKEN_FALSE, TOKEN_UNKNOWN,
		TOKEN_PT, TOKEN_IN, TOKEN_EOF,
	}

	for i, exp := range expected {
		if i >= len(tokens) {
			t.Fatalf("missing token at %d, want %s", i, exp)
		}
		if tokens[i].Type != exp {
			t.Errorf("[%d] want %s, got %s %q", i, exp, tokens[i].Type, tokens[i].Literal)
		}
	}
}

func TestOperators(t *testing.T) {
	input := `= + - * / % < > & | ! ? != <= >= ->`

	lex := New(input)
	tokens, err := lex.Tokenize()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := []TokenType{
		TOKEN_ASSIGN, TOKEN_PLUS, TOKEN_MINUS, TOKEN_STAR, TOKEN_SLASH,
		TOKEN_PERCENT, TOKEN_LT, TOKEN_GT, TOKEN_AND, TOKEN_OR,
		TOKEN_NOT, TOKEN_QUESTION, TOKEN_NOT_EQ, TOKEN_LT_EQ, TOKEN_GT_EQ,
		TOKEN_ARROW, TOKEN_EOF,
	}

	for i, exp := range expected {
		if i >= len(tokens) {
			t.Fatalf("missing token at %d, want %s", i, exp)
		}
		if tokens[i].Type != exp {
			t.Errorf("[%d] want %s, got %s %q", i, exp, tokens[i].Type, tokens[i].Literal)
		}
	}
}

func TestComment(t *testing.T) {
	input := `.this is a comment.
name "Zaky"
. multi
  line
  comment .
age 17`

	lex := New(input)
	tokens, err := lex.Tokenize()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Comments should be skipped
	foundName := false
	foundAge := false
	for _, tok := range tokens {
		if tok.Type == TOKEN_IDENT && tok.Literal == "name" {
			foundName = true
		}
		if tok.Type == TOKEN_IDENT && tok.Literal == "age" {
			foundAge = true
		}
	}

	if !foundName {
		t.Error("missing 'name' token after comment")
	}
	if !foundAge {
		t.Error("missing 'age' token after multi-line comment")
	}
}

func TestCommentWithDotsInside(t *testing.T) {
	input := `.this is a .nested. comment.
name "ok"`

	lex := New(input)
	tokens, err := lex.Tokenize()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	found := false
	for _, tok := range tokens {
		if tok.Type == TOKEN_IDENT && tok.Literal == "name" {
			found = true
		}
	}
	if !found {
		t.Error("missing 'name' after comment with dots inside")
	}
}

func TestRawString(t *testing.T) {
	input := "pattern `\\d+\\.\\w+`"

	lex := New(input)
	tokens, err := lex.Tokenize()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	found := false
	for _, tok := range tokens {
		if tok.Type == TOKEN_RAW_STRING && tok.Literal == `\d+\.\w+` {
			found = true
		}
	}
	if !found {
		t.Error("raw string not tokenized correctly")
	}
}

func TestFloatComma(t *testing.T) {
	input := `pi 3,14`

	lex := New(input)
	tokens, err := lex.Tokenize()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	found := false
	for _, tok := range tokens {
		if tok.Type == TOKEN_NUMBER && tok.Literal == "3.14" {
			found = true
		}
	}
	if !found {
		t.Error("float comma decimal not converted to dot internally")
	}
}

func TestStringInterpolation(t *testing.T) {
	input := `pt "Hello ${name}"`

	lex := New(input)
	tokens, err := lex.Tokenize()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Should have: PT, STRING_PART("Hello "), INTERP_START, IDENT(name), INTERP_END, INTERP_END, EOF
	hasStringPart := false
	hasInterpStart := false
	hasIdent := false
	for _, tok := range tokens {
		if tok.Type == TOKEN_STRING_PART && tok.Literal == "Hello " {
			hasStringPart = true
		}
		if tok.Type == TOKEN_INTERP_START {
			hasInterpStart = true
		}
		if tok.Type == TOKEN_IDENT && tok.Literal == "name" {
			hasIdent = true
		}
	}

	if !hasStringPart {
		t.Error("missing STRING_PART 'Hello '")
	}
	if !hasInterpStart {
		t.Error("missing INTERP_START")
	}
	if !hasIdent {
		t.Error("missing IDENT 'name'")
	}
}

func TestLineNumbers(t *testing.T) {
	input := `name "Zaky"
age 17`

	lex := New(input)
	tokens, err := lex.Tokenize()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// 'name' should be line 1, 'age' should be line 2
	for _, tok := range tokens {
		if tok.Literal == "name" && tok.Line != 1 {
			t.Errorf("'name' want line 1, got %d", tok.Line)
		}
		if tok.Literal == "age" && tok.Line != 2 {
			t.Errorf("'age' want line 2, got %d", tok.Line)
		}
	}
}
