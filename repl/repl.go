package repl

import (
	"bufio" // Buffered io library
	"fmt" // Formatted i/o, similar to C's printf/scanf
	"io" // Go input/output lib
	"mockc/lexer" // our custom lexer
	"mockc/token" // our token definitions
)

const PROMPT = ">> " // Prompt at the beginning of each newline for users to know when to input

/*
Basically the REPL engine. Called once and runs in a loop until broken by the user.
 */
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for { // Perpetual for loop... I guess Go has those
		fmt.Fprintf(out, PROMPT) // Formats string and writes to out
		scanned := scanner.Scan()
		if !scanned { // If nothing was entered, break the loop
			return
		}

		line := scanner.Text()
		l := lexer.New(line) // Tokenize the user input

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() { // Looks like a java-style for loop
			//tok is declared as the next token (first in l), continues until tok.Type = EOF, and the loop advances the lexer cursor forward after each iteration
			fmt.Fprintf(out, "%+v\n", tok) // %+v shows fields of a struct by name ex. {TokenType: x, Literal: y}
		}
	}
}