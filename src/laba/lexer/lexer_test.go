package lexer

import (
	"testing"

	"github.com/vzlatin/LABA/token"
)

func TestNextToken(t *testing.T) {
	input := `bagă cinci = 5;
	bagă zece = 10;

	bagă add = combinație(x, y) {
		scoate x + y;
	}
	bagă rezultat = add(cinci, zece);
	!-/*5;
	5 < 10 > 5;
	dacă (5 < 10) {
		scoate pe_bune;
	} sau {
		scoate vrajeală;
	}
	10 == 10;
	10 != 9;
	`
	type testToken struct {
		exectedType     token.TokenType
		expectedLiteral string
	}

	tests := []testToken{
		{token.BAGA, "bagă"},
		{token.IDENT, "cinci"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.BAGA, "bagă"},
		{token.IDENT, "zece"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.BAGA, "bagă"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.COMBINATIE, "combinație"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.SCOATE, "scoate"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.BAGA, "bagă"},
		{token.IDENT, "rezultat"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "cinci"},
		{token.COMMA, ","},
		{token.IDENT, "zece"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.BANG, "!"},
		{token.MINUS, "-"},
		{token.SLASH, "/"},
		{token.ASTERISK, "*"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.GT, ">"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.DACA, "dacă"},
		{token.LPAREN, "("},
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.SCOATE, "scoate"},
		{token.PE_BUNE, "pe_bune"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.SAU, "sau"},
		{token.LBRACE, "{"},
		{token.SCOATE, "scoate"},
		{token.VRAJEALA, "vrajeală"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.INT, "10"},
		{token.EQ, "=="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.INT, "10"},
		{token.NOT_EQ, "!="},
		{token.INT, "9"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()
		if tok.Type != tt.exectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q", i, tt.exectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", i, tt.expectedLiteral, tok.Literal)
		}
	}
}
