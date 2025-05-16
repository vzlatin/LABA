package token

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers + literals
	IDENT = "IDENT"
	INT   = "INT"

	// Operators
	ASSIGN = "="
	PLUS   = "+"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"

	// Keywords
	COMBINATIE = "COMBINATIE"
	SCOATE     = "SCOATE"
	BAGA       = "BAGA"
)

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}
