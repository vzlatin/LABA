package token

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers + literals
	IDENT = "IDENT"
	INT   = "INT"

	// Operators
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"
	LT       = "<"
	GT       = ">"
	EQ       = "=="
	NOT_EQ   = "!="

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"

	// Keywords
	COMBINATIE = "COMBINATIE"
	BAGA       = "BAGA"
	PE_BUNE    = "PE_BUNE"
	VRAJEALA   = "VRAJEALA"
	DACA       = "DACA"
	SAPOI_DACA = "SAPOI_DACA" // elif
	SAU        = "SAU"
	SCOATE     = "SCOATE"
)

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

var keywords = map[string]TokenType{
	"combinație": COMBINATIE,
	"bagă":       BAGA,
	"scoate":     SCOATE,
	"pe bune":    PE_BUNE,
	"vrajeală":   VRAJEALA,
	"șapoi dacă": SAPOI_DACA,
	"dacă":       DACA,
	"sau":        SAU,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
