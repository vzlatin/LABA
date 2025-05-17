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
	COMBINAȚIE = "COMBINAȚIE"
	BAGĂ       = "BAGĂ"
	PE_BUNE    = "PE_BUNE"
	VRAJEALĂ   = "VRAJEALĂ"
	DACĂ       = "DACĂ"
	ȘAPOI      = "ȘAPOI" // elif
	SAU        = "SAU"
	SCOATE     = "SCOATE"
)

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

var keywords = map[string]TokenType{
	"combinație": COMBINAȚIE,
	"bagă":       BAGĂ,
	"scoate":     SCOATE,
	"pe_bune":    PE_BUNE,
	"vrajeală":   VRAJEALĂ,
	"șapoi":      ȘAPOI,
	"dacă":       DACĂ,
	"sau":        SAU,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
