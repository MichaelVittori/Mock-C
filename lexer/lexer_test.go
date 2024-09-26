package lexer

// This lexer is built using the TDD methodology so these tests are developed before the functionality to pass them.

import (
	"mockc/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `let x = 5;
	let y = 10;
	
	let add = fn(x, y){ x + y; };
	let sum = add(x, y);
	`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		// Declare and assign X
		{token.LET, = "let"},
		{token.IDENTIFIER, "x"}
		{token.ASSIGN, "="},
		{token.INTEGER, "5"},
		{token.SEMICOLON, ";"},

		// Declare and assign y
		{token.LET, = "let"},
		{token.IDENTIFIER, "y"}
		{token.ASSIGN, "="},
		{token.INTEGER, "10"},
		{token.SEMICOLON, ";"},

		// Declare and assign add
		{token.LET, = "let"},
		{token.IDENTIFIER, "add"}
		{token.ASSIGN, "="},
		//Func and params of add
		{token.FUNCTION, "fn"},
		{token.LPAREN, "("},
		{token.IDENTIFIER, "x"},
		{token.COMMA, ","},
		{token.IDENTIFIER, "y"},
		{token.RPAREN, ")"},
		// Body definition
		{token.LBRACE, "{"},
		{token.IDENTIFIER, "x"},
		{token.PLUS, "+"},
		{token.IDENTIFIER, "y"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},

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
