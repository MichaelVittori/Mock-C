package parser

import (
	"mockc/ast"
	"mockc/lexer"
	"mockc/token"
	"fmt"
	"strconv"
)

const (
	_ int = iota // Automatically assigns ascending ints to the consts below
	LOWEST		// Default, non operator precedence
	EQUALS		// ==
	LESSGREATER // < >
	SUM			// +
	PRODUCT		// *
	PREFIX		// -X or !X
	CALL		// foobar(x)
)

var precedences = map[token.TokenType]int{
	token.EQ:	  EQUALS,
	token.NEQ:	  EQUALS,
	token.LTHAN:  LESSGREATER,
	token.GTHAN:  LESSGREATER,
	token.PLUS:   SUM,
	token.MINUS:  SUM,
	token.DIVIDE: PRODUCT,
	token.TIMES:  PRODUCT,
}

type (
	prefixParseFn func() ast.Expression
	infixParseFn func(ast.Expression) ast.Expression
)

type Parser struct {
	l *lexer.Lexer
	currToken token.Token
	peekToken token.Token
	errors []string // List of errors

	prefixParseFns map[token.TokenType]prefixParseFn
	infixParseFns map[token.TokenType]infixParseFn
}

func (p *Parser) registerPrefix(tokenType token.TokenType, fn prefixParseFn) {
	p.prefixParseFns[tokenType] = fn
}

func (p *Parser) registerInfix(tokenType token.TokenType, fn infixParseFn) {
	p.infixParseFns[tokenType] = fn
}

func (p *Parser) peekPrecedence() int {
	if p, ok := precedences[p.peekToken.Type]; ok {
		return p
	}
	return LOWEST
}

func (p *Parser) currPrecedence() int {
	if p, ok := precedences[p.currToken.Type]; ok {
		return p
	}
	return LOWEST
}

/*
"Initializes" a new Parser struct
*/
func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l     : l,
		errors: []string{},
	} // Returns the value of the parser being pointed to

	// Make map of prefix token parse functions
	p.prefixParseFns = make(map[token.TokenType]prefixParseFn)
	p.registerPrefix(token.IDENTIFIER, p.parseIdentifier)
	p.registerPrefix(token.INTEGER, p.parseIntegerLiteral)
	p.registerPrefix(token.NOT, p.parsePrefixExpression)
	p.registerPrefix(token.MINUS, p.parsePrefixExpression)

	// Make map of infix token parse functions
	p.infixParseFns = make(map[token.TokenType]infixParseFn)
	p.registerInfix(token.PLUS, p.parseInfixExpression)
	p.registerInfix(token.MINUS, p.parseInfixExpression)
	p.registerInfix(token.DIVIDE, p.parseInfixExpression)
	p.registerInfix(token.TIMES, p.parseInfixExpression)
	p.registerInfix(token.EQ, p.parseInfixExpression)
	p.registerInfix(token.NEQ, p.parseInfixExpression)
	p.registerInfix(token.LTHAN, p.parseInfixExpression)
	p.registerInfix(token.GTHAN, p.parseInfixExpression)

	//Read two tokens to populate curr and next
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) parseIdentifier() ast.Expression {
	return &ast.Identifier{Token: p.currToken, Value: p.currToken.Literal}
}

/*
 Return the errors in the parser
 */
func (p *Parser) Errors() []string {
	return p.errors
}

/*
 Create and add error message to p.errors
 */
func (p *Parser) peekError(t token.TokenType) {
	msg := fmt.Sprintf("Expected next token to be %s, got %s instead", t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}


/*
Move the cursor forward by one token and peek ahead
*/
func (p *Parser) nextToken() {
	p.currToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

/*
 Build AST root node, iterate through tokens and build child nodes
*/
func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{} // Construct AST root node
	program.Statements = []ast.Statement{}

	for p.currToken.Type != token.EOF { // As long as the current token isn't the end of the file...
		stmt := p.parseStatement()
		if stmt != nil { // If the statement isn't null
			program.Statements = append(program.Statements, stmt) // Append that statement to program's statements
		}
		p.nextToken()
	}

	return program
}


/*
 Uses a switch statement to determine how best to parse the given token
 */
func (p *Parser) parseStatement() ast.Statement {
	switch p.currToken.Type {
	case token.LET:
		return p.parseLetStatement()
	case token.RETURN:
		return p.parseReturnStatement()
	default:
		return p.parseExpressionStatement()
	}
}

/*
 Parse out let statements and place them in AST
 */
func (p *Parser) parseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{Token: p.currToken}

	if !p.expectPeek(token.IDENTIFIER) { // If let is not followed up by an identifier, return nil
		return nil
	}

	stmt.Name = &ast.Identifier{Token: p.currToken, Value: p.currToken.Literal}

	if !p.expectPeek (token.ASSIGN){ // If the next token after identifier is not an assign operator, return nil
		return nil
	}

	//TODO
	// Currently this just skips to the semicolon of each line, change that later
	for !p.currTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

/*

 */
func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	stmt := &ast.ReturnStatement{Token: p.currToken}

	p.nextToken()

	//TODO
	// Gonna skip to semicolon here too
	for !p.currTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	defer untrace(trace("parseExpressionStatement"))

	stmt := &ast.ExpressionStatement{Token: p.currToken}
	stmt.Expression = p.parseExpression(LOWEST) // Lowest refers to operator precedence for PEMDAS

	if p.peekTokenIs(token.SEMICOLON) { // Semicolons are optional for expressions
		p.nextToken()
	}

	return stmt
}

func (p *Parser) noPrefixParseFnError(t token.TokenType) {
	msg := fmt.Sprintf("no prefix parse function for %s found", t)
	p.errors = append(p.errors, msg)
}

func (p *Parser) parseExpression(precedence int) ast.Expression {
	defer untrace(trace("parseExpression"))

	prefix := p.prefixParseFns[p.currToken.Type] // If the current token has a parsing function use that
	if prefix == nil { // If the token doesn't have a type, parse function, or just doesn't exist return null
		p.noPrefixParseFnError(p.currToken.Type)
		return nil
	}

	leftExpression := prefix() // Execute the prefixParseFunction from above

	for !p.peekTokenIs(token.SEMICOLON) && precedence < p.peekPrecedence() { // Keep running until finding a semicolon or a lower precedence token
		infix := p.infixParseFns[p.peekToken.Type] // Scan the line for an infix expression
		if infix == nil { // If none exists return the left expression as is
			return leftExpression
		}

		p.nextToken()
		leftExpression = infix(leftExpression)
	}
	return leftExpression
}

/*
 Return whether actual token type of currToken is equal to expected token type
*/
func (p *Parser) currTokenIs(t token.TokenType) bool {
	return p.currToken.Type == t
}

/*
 Return whether actual token type of peekToken is equal to expected token type
*/
func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) { // If the token is a match, advance cursor and return true
		p.nextToken()
		return true
	} else { // Otherwise append an error and return false
		p.peekError(t)
		return false
	}
}

func (p *Parser) parseIntegerLiteral() ast.Expression {
	defer untrace(trace("parseIntegerLiteral"))

	literal := &ast.IntegerLiteral{Token: p.currToken}

	value, err := strconv.ParseInt(p.currToken.Literal, 0, 64) // Convert it to a 64 bit integer
	if err != nil {
		msg := fmt.Sprintf("Could not parse %q as int", p.currToken.Literal)
		p.errors = append(p.errors, msg)
		return nil
	}

	literal.Value = value
	return literal
}

func (p *Parser) parsePrefixExpression() ast.Expression {
	defer untrace(trace("parsePrefixExpression"))

	expression := &ast.PrefixExpression{ // Create an expression for the prefix token
		Token:	  p.currToken,
		Operator: p.currToken.Literal,
	}

	p.nextToken() // Advance one
	expression.Right = p.parseExpression(PREFIX) // Parse the right side of the prefix expression
	return expression
}

func (p *Parser) parseInfixExpression(left ast.Expression) ast.Expression {
	defer untrace(trace("parseInfixExpression"))

	expression := &ast.InfixExpression{
		Token:		p.currToken, // Set token to the operator currently being looked at
		Operator:	p.currToken.Literal,
		Left:		left, // Take left part of the expression
	}

	precedence := p.currPrecedence() // Get precedence of the operator
	p.nextToken() // Advance and assign next token to right
	expression.Right = p.parseExpression(precedence)

	return expression
}