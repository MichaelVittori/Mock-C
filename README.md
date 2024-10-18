# Moxie Interpreter

Moxie is a simple programming language interpreter written in Go. This project follows along with the book *[Writing an Interpreter in Go](https://interpreterbook.com/)* by Thorsten Ball, which teaches the fundamentals of creating an interpreter for a programming language.

Note: this project is still in early stages and I intend to differentiate it more substantially from the book once I have laid the proper foundation.

## Features

- **Lexer**: Tokenizes input source code into distinct lexical tokens.
- **Parser**: Constructs an abstract syntax tree (AST) from the tokens.
- **Evaluator**: Evaluates the AST, handling variables, functions, and control structures.
- **REPL**: A Read-Eval-Print Loop (REPL) for experimenting with Moxie interactively.

## Getting Started

### Prerequisites

To build and run the interpreter, you'll need:

- [Go](https://golang.org/dl/) version 1.16 or higher.

### Installation

1. Clone the repository:

    ```bash
    git clone https://github.com/MichaelVittori/Mock-C
    ```

2. Navigate into the project directory:

    ```bash
    cd Mock-C
    ```

3. Build the project:

    ```bash
    go build
    ```

4. Run the REPL: **Evaluation step is not complete, REPL currently just returns Abstract Syntax Tree nodes. This step will be completed in the coming days.**
    ```bash
    go run main.go
    ```

## Usage

Once you've started the REPL, you can enter Moxie code and get instant feedback.

### Example

```bash
>> let x = 10;
>> let y = x * 2;
>> y;
20
```
```bash
>> let add = fn(a, b) { a + b };
>> add(3, 4);
7
```
## Project Structure
- **lexer/:** Responsible for tokenizing input.
- **parser/:** Turns tokens into an AST.
- **ast/:** Defines the structure of the AST.
- **evaluator/:** Evaluates the AST to produce results.
- **object/:** Contains definitions of all runtime objects (integers, booleans, etc.).
- **repl/:** Implements the REPL (Read-Eval-Print-Loop).

## Testing
Tests can be run using the ```go test``` command:
```bash
go test ./...
```

## Credit
Once again, this project was made using Writing an Interpreter in Go by Thorsten Ball.
