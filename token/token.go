package token

// TokenType is a custom type for token
type TokenType string

// Token struct describes a token with a given Type and holds its literal
type Token struct {
	Type    TokenType
	Literal string
}

// Tokens of the monkey language
const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	IDENT = "IDENT" // variables, functions
	INT   = "INT"   // integer literal

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"

	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"

	EQ  = "=="
	NEQ = "!="

	LT = "<"
	GT = ">"

	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"
)

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

// LookupType returns the type of the given token hold in string
func LookupType(s string) TokenType {
	if tokType, ok := keywords[s]; ok {
		return tokType
	}
	return IDENT
}
