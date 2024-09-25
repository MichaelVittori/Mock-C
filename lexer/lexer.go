package lexer

import "mockc/token"

type Lexer struct {
	input        string
	position     int  //Current position in input (points to curr char)
	readPosition int  //Current reading position in input (after current char)
	ch           byte //Current char being examined
}

/*
Basically a constructor
*/
func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

/*
Read the current char in input and scoot position and readPosition forward by one
*/
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) { // Check if we're at the end of our input
		l.ch = 0 // ASCII for NULL
	} else {
		l.ch = l.input[l.readPosition] // Set current character to the character at readPosition
	}
	l.position = l.readPosition // Advance both positions by one
	l.readPosition += 1
}

/*
Tokenizes the current character
*/
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case '[':
		tok = newToken(token.LBRACKET, l.ch)
	case ']':
		tok = newToken(token.RBRACKET, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '-':
		tok = newToken(token.MINUS, l.ch)
	case '*':
		tok = newToken(token.TIMES, l.ch)
	case '/':
		tok = newToken(token.DIVIDE, l.ch)
	case '%':
		tok = newToken(token.MOD, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	}

	l.readChar()
	return tok
}

/*
Creates a new token struct using the information passed in
*/
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}
