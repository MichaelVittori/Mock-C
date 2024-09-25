package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string //Strings don't offer the best performance, but they're more convenient to work with
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers and literals
	IDENTIFIER = "IDENTIFIER" //method and variable names
	INTEGER    = "INTEGER"    //integers, duh (1, 2, 3, 4...)
	FLOAT      = "FLOAT"      //floating point numbers (1.1, 2.1231, 3.14)

	// Operators
	ASSIGN = "="
	PLUS   = "+"
	MINUS  = "-"
	TIMES  = "*"
	DIVIDE = "/"
	MOD    = "%"

	// Delimeters
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"
	LBRACKET  = "["
	RBRACKET  = "]"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)
