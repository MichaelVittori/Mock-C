package lexer

// This lexer is built using the TDD methodology so these tests are developed before the functionality to pass them.

import (
	"mockc/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `=+*-/%(){}[],;`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.TIMES, "*"},
		{token.MINUS, "-"},
		{token.DIVIDE, "/"},
		{token.MOD, "%"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.LBRACKET, "["},
		{token.RBRACKET, "]"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q", i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - tokenliteral wrong. expected=%q, got=%q", i, tt.expectedLiteral, tok.Literal)
		}
	}
}
