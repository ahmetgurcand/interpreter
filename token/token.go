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
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"

	LT = "<"
	GT = ">"

	EQ = "=="
	NOT_EQ = "!="

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
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF  	 = "IF"
	ELSE 	 = "ELSE"
	RETURN 	 = "RETURN"

	STRING   = "STRING"

	LBRACKET = "["
	RBRACKET = "]"

	COLON = ":"
)

var keywords = map[string]TokenType{
	"fn": FUNCTION,
	"let": LET,
	"true": TRUE,
	"false": FALSE,
	"if": IF,
	"else": ELSE,
	"return": RETURN,
}

// LookupIdent checks whether the given identifier is in fact a keyword.
func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}