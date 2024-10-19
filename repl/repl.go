package repl

import (
	"bufio" // Buffered io library
	"fmt" // Formatted i/o, similar to C's printf/scanf
	"io" // Go input/output lib
	"mockc/lexer" // our custom lexer
	"mockc/parser"
	"mockc/evaluator"
	"mockc/object"
)

const PROMPT = ">> " // Prompt at the beginning of each newline for users to know when to input

/*
Basically the REPL engine. Called once and runs in a loop until broken by the user.
 */
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()

	for { // Perpetual for loop... I guess Go has those
		fmt.Fprintf(out, PROMPT) // Formats string and writes to out
		scanned := scanner.Scan()

		if !scanned { // If nothing was entered, break the loop
			return
		}

		line := scanner.Text()
		l := lexer.New(line) // Tokenize the user input
		p := parser.New(l) // Parse the tokens

		program := p.ParseProgram()
		if len(p.Errors()) != 0 { // If there are any errors in the parser, print them
			printParserErrors(out, p.Errors())
			continue
		}

		eval := evaluator.Eval(program, env)
		if eval != nil {
			io.WriteString(out, eval.Inspect())
			io.WriteString(out, "\n")
		}
	}
}

const MONKEY_FACE = `
            __,__
   .--.  .-"     "-.  .--.
  / .. \/  .-. .-.  \/ .. \
 | |  '|  /   Y   \  |'  | |
 | \   \  \ 0 | 0 /  /   / |
  \ '- ,\.-"""""""-./, -' /
   ''-' /_   ^ ^   _\ '-''
       |  \._   _./  |
       \   \ '~' /   /
        '._ '-=-' _.'
           '-----'
`

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, MONKEY_FACE)
	io.WriteString(out, "Looks like there's some monkey business over here!\n")
	io.WriteString(out, " parser errors:\n")
	for _, msg := range errors { // Iterate through all error messages and print them
		io.WriteString(out, "\t"+msg+"\n")
	}
}