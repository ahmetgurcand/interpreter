package  token

type TokenType string

type Token struct {
	Type TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF 	= "EOF"

	// Identifiers + literals
	IDENT = "IDENT"
	INT = "INT"

	// Operators
	ASSIGN = "="
	PLUS   = "+"

	// Delimiters
	COMMA 	  = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET 	 = "LET"
)

var keywords = map[string]TokenType{
	"fn": FUNCTION,
	"let": LET,
}

// LookupIdent checks whether the given identifier is in fact a keyword.
func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}