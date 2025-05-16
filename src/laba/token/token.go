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
	PA_ORMA    = "PA_ORMA" // elif
	SAU        = "SAU"
	SCOATE     = "SCOATE"
)

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

var keywords = map[string]TokenType{
	"combinatie": COMBINATIE,
	"baga":       BAGA,
	"scoate":     SCOATE,
	"pe_bune":    PE_BUNE,
	"vrajeala":   VRAJEALA,
	"pa_orma":    PA_ORMA,
	"daca":       DACA,
	"sau":        SAU,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
