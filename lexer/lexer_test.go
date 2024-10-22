package lexer

// This lexer is built using the TDD methodology so these tests are developed before the functionality to pass them.

import (
	"mockc/token";
	"testing";
)

func TestNextToken(t *testing.T) {
	input := `let x = 5;
	let y = 10;
	
	let add = fn(x, y){ x + y; };
	let sum = add(x, y);
	!-/*5;
	142<7>31;
	if (5 = 10) {
		return true;
	} else {
		return false;
	}
	10 == 10;
	10 != 9;
	100 >= 99;
	99 <= 100;
	"foobar"
	"foo bar"
	[1, 2];
	:
	`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		// Declare and assign X
		{token.LET, "let"},
		{token.IDENTIFIER, "x"},
		{token.ASSIGN, "="},
		{token.INTEGER, "5"},
		{token.SEMICOLON, ";"},

		// Declare and assign y
		{token.LET, "let"},
		{token.IDENTIFIER, "y"},
		{token.ASSIGN, "="},
		{token.INTEGER, "10"},
		{token.SEMICOLON, ";"},

		// Declare and assign add
		{token.LET, "let"},
		{token.IDENTIFIER, "add"},
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

		// Sum definition and add function call
		{token.LET, "let"},
		{token.IDENTIFIER, "sum"},
		{token.ASSIGN, "="},
		{token.IDENTIFIER, "add"},
		{token.LPAREN, "("},
		{token.IDENTIFIER, "x"},
		{token.COMMA, ","},
		{token.IDENTIFIER, "y"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},

		// Giberish line 1
		{token.NOT, "!"},
		{token.MINUS, "-"},
		{token.DIVIDE, "/"},
		{token.TIMES, "*"},
		{token.INTEGER, "5"},
		{token.SEMICOLON, ";"},

		// Giberish line 2
		{token.INTEGER, "142"},
		{token.LTHAN, "<"},
		{token.INTEGER, "7"},
		{token.GTHAN, ">"},
		{token.INTEGER, "31"},
		{token.SEMICOLON, ";"},

		// If statement
		{token.IF, "if"},
		{token.LPAREN, "("},
		{token.INTEGER, "5"},
		{token.ASSIGN, "="},
		{token.INTEGER, "10"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.TRUE, "true"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},

		// Else statement
		{token.ELSE, "else"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.FALSE, "false"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},

		// Logical equals
		{token.INTEGER, "10"},
		{token.EQ, "=="},
		{token.INTEGER, "10"},
		{token.SEMICOLON, ";"},

		// Not equals
		{token.INTEGER, "10"},
		{token.NEQ, "!="},
		{token.INTEGER, "9"},
		{token.SEMICOLON, ";"},

		// Greater than
		{token.INTEGER, "100"},
		{token.GEQ, ">="},
		{token.INTEGER, "99"},
		{token.SEMICOLON, ";"},

		// Less than
		{token.INTEGER, "99"},
		{token.LEQ, "<="},
		{token.INTEGER, "100"},
		{token.SEMICOLON, ";"},

		// Strings
		{token.STRING, "foobar"},
		{token.STRING, "foo bar"},

		// Array
		{token.LBRACKET, "["},
		{token.INTEGER, "1"},
		{token.COMMA, ","},
		{token.INTEGER, "2"},
		{token.RBRACKET, "]"},
		{token.SEMICOLON, ";"},

		// Colon
		{token.COLON, ":"},

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
