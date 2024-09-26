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

	// Arithmetic Operators
	ASSIGN = "="
	PLUS   = "+"
	MINUS  = "-"
	TIMES  = "*"
	DIVIDE = "/"
	MOD    = "%"

	// Logical Operators
	NOT   = "!"
	LTHAN = "<"
	GTHAN = ">"

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
	IF 		 = "IF"
	ELSE 	 = "ELSE"
	TRUE	 = "TRUE"
	FALSE	 = "FALSE"
	RETURN	 = "RETURN"
)

var keywords = map[string] TokenType {
	"fn": FUNCTION,
	"let": LET,
	"if": IF,
	"else": ELSE,
	"true": TRUE,
	"false": FALSE,
	"return": RETURN,
}

/*
Uses the keywords map to lookup language keywords to differentiate them and var/func names
*/
func LookupIdentity(identity string) TokenType {
	if token, ok := keywords[identity]; ok {
		return token
	}
	return IDENTIFIER
}