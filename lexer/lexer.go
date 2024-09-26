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

	l.skipWhitespace()

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
	case '!':
		tok = newToken(token.NOT, l.ch)
	case '<':
		tok = newToken(token.LTHAN, l.ch)
	case '>':
		tok = newToken(token.GTHAN, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier() //Uses readIdentifier to read all adjacent characters to check if they're letters
			tok.Type = token.LookupIdentity(tok.Literal)
			return tok // If it is a legal identifier, return it as a token
		} else if isDigit(l.ch){
			tok.Type = token.INTEGER
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return tok
}

/*
Whitespace doesn't mean anything in this new language, so this helper function skips it
*/
func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' { // If the current char is whitespace
		l.readChar() // Consume it.
	}
}

/*
Similar to readIdentifier, but with numbers this time
*/
func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) { // Scroll through and collect all digits of the number
		l.readChar()
	}
	return l.input[position:l.position]
}

/*
Similar to isLetter but again, with numbers this time.
*/
func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9' || ch == '.'
}

/*
Scrolls through a multicharacter identifier and returns the whole thing
*/
func (l *Lexer) readIdentifier() string {
	position := l.position // Set position to the start of the word
	for isLetter(l.ch){ // While the cursor is on a letter, scroll through char by char
		l.readChar()
	}
	return l.input[position:l.position] // Return the input from index position to l.position, this is equal to a full identifier
	//EX. if input is "let" position = l, l.position cycles through until t, this function returns "let"
}

/*
Since this lexer can't read the full spectrum of unicode, this helper function parses all ASCII compliant letters
and underscores.
*/
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

/*
Creates a new token struct using the information passed in
*/
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}
