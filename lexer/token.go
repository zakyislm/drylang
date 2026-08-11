package lexer

type TokenType int

const (
	// Special
	TOKEN_EOF TokenType = iota
	TOKEN_ILLEGAL

	// Literals
	TOKEN_IDENT        // variable/function names
	TOKEN_NUMBER       // 42, 89,5
	TOKEN_STRING       // "hello", 'hello'
	TOKEN_RAW_STRING   // `raw string`
	TOKEN_STRING_PART  // string segment before/between ${} interpolation
	TOKEN_INTERP_START // ${
	TOKEN_INTERP_END   // }

	// Keywords
	TOKEN_CNS     // cns
	TOKEN_TRUE    // t
	TOKEN_FALSE   // f
	TOKEN_FN      // fn
	TOKEN_REV     // rev
	TOKEN_IF      // if
	TOKEN_ELIF    // elif
	TOKEN_EL      // el
	TOKEN_ON      // on
	TOKEN_LP      // lp
	TOKEN_DONE    // done
	TOKEN_CON     // con
	TOKEN_ASN     // asn
	TOKEN_AWT     // awt
	TOKEN_TRY     // try
	TOKEN_ERR     // err
	TOKEN_PV      // pv
	TOKEN_USE     // use
	TOKEN_CL      // cl
	TOKEN_UNKNOWN // unknown
	TOKEN_MUL     // mul
	TOKEN_UNI     // uni

	// Operators (1-char)
	TOKEN_ASSIGN    // =
	TOKEN_PLUS      // +
	TOKEN_MINUS     // -
	TOKEN_STAR      // *
	TOKEN_SLASH     // /
	TOKEN_PERCENT   // %
	TOKEN_LT        // <
	TOKEN_GT        // >
	TOKEN_AND       // &
	TOKEN_OR        // |
	TOKEN_NOT       // !
	TOKEN_QUESTION  // ?
	TOKEN_QMARK_DOT // ?.
	TOKEN_QQ        // ??
	TOKEN_DOT       // . (also comment delimiter)

	// Operators (2-char)
	TOKEN_NOT_EQ  // !=
	TOKEN_LT_EQ   // <=
	TOKEN_GT_EQ   // >=
	TOKEN_ARROW   // ->
	TOKEN_LARROW  // <-

	// Delimiters
	TOKEN_LPAREN    // (
	TOKEN_RPAREN    // )
	TOKEN_LBRACE    // {
	TOKEN_RBRACE    // }
	TOKEN_LBRACKET  // [
	TOKEN_RBRACKET  // ]
	TOKEN_COMMA     // ,
	TOKEN_COLON     // :
	TOKEN_SEMICOLON // ;
)

var tokenNames = map[TokenType]string{
	TOKEN_EOF:          "EOF",
	TOKEN_ILLEGAL:      "ILLEGAL",
	TOKEN_IDENT:        "IDENT",
	TOKEN_NUMBER:       "NUMBER",
	TOKEN_STRING:       "STRING",
	TOKEN_RAW_STRING:   "RAW_STRING",
	TOKEN_STRING_PART:  "STRING_PART",
	TOKEN_INTERP_START: "INTERP_START",
	TOKEN_INTERP_END:   "INTERP_END",
	TOKEN_CNS:          "cns",
	TOKEN_TRUE:         "t",
	TOKEN_FALSE:        "f",
	TOKEN_FN:           "fn",
	TOKEN_REV:          "rev",
	TOKEN_IF:           "if",
	TOKEN_ELIF:         "elif",
	TOKEN_EL:           "el",
	TOKEN_ON:           "on",
	TOKEN_LP:           "lp",
	TOKEN_DONE:         "done",
	TOKEN_CON:          "con",
	TOKEN_ASN:          "asn",
	TOKEN_AWT:          "awt",
	TOKEN_TRY:          "try",
	TOKEN_ERR:          "err",
	TOKEN_PV:           "pv",
	TOKEN_USE:          "use",
	TOKEN_CL:           "cl",
	TOKEN_UNKNOWN:      "unknown",
	TOKEN_MUL:          "mul",
	TOKEN_UNI:          "uni",
	TOKEN_ASSIGN:       "=",
	TOKEN_PLUS:         "+",
	TOKEN_MINUS:        "-",
	TOKEN_STAR:         "*",
	TOKEN_SLASH:        "/",
	TOKEN_PERCENT:      "%",
	TOKEN_LT:           "<",
	TOKEN_GT:           ">",
	TOKEN_AND:          "&",
	TOKEN_OR:           "|",
	TOKEN_NOT:          "!",
	TOKEN_QUESTION:     "?",
	TOKEN_QMARK_DOT:    "?.",
	TOKEN_QQ:           "??",
	TOKEN_DOT:          ".",
	TOKEN_NOT_EQ:       "!=",
	TOKEN_LT_EQ:        "<=",
	TOKEN_GT_EQ:        ">=",
	TOKEN_ARROW:        "->",
	TOKEN_LARROW:       "<-",
	TOKEN_LPAREN:       "(",
	TOKEN_RPAREN:       ")",
	TOKEN_LBRACE:       "{",
	TOKEN_RBRACE:       "}",
	TOKEN_LBRACKET:     "[",
	TOKEN_RBRACKET:     "]",
	TOKEN_COMMA:        ",",
	TOKEN_COLON:        ":",
	TOKEN_SEMICOLON:    ";",
}

func (t TokenType) String() string {
	if name, ok := tokenNames[t]; ok {
		return name
	}
	return "UNKNOWN_TOKEN"
}

// Token represents a single lexical token.
type Token struct {
	Type    TokenType
	Literal string
	Line    int
	Col     int
}

// keywords maps dryLang keywords to their token types.
var keywords = map[string]TokenType{
	"cns":     TOKEN_CNS,
	"t":       TOKEN_TRUE,
	"f":       TOKEN_FALSE,
	"fn":      TOKEN_FN,
	"rev":     TOKEN_REV,
	"if":      TOKEN_IF,
	"elif":    TOKEN_ELIF,
	"el":      TOKEN_EL,
	"on":      TOKEN_ON,
	"lp":      TOKEN_LP,
	"done":    TOKEN_DONE,
	"con":     TOKEN_CON,
	"asn":     TOKEN_ASN,
	"awt":     TOKEN_AWT,
	"try":     TOKEN_TRY,
	"err":     TOKEN_ERR,
	"pv":      TOKEN_PV,
	"use":     TOKEN_USE,
	"cl":      TOKEN_CL,
	"unknown": TOKEN_UNKNOWN,
	"mul":     TOKEN_MUL,
	"uni":     TOKEN_UNI,
}

// LookupIdent checks if an identifier is a keyword.
func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return TOKEN_IDENT
}
