package ast

import (
	"mockc/token"
	"bytes"
	"strings"
)

type Node interface { // Define our treenode interface
	TokenLiteral() string
	String() string
}

type Statement interface { // Statement interface
	Node // Contains a node
	statementNode()
}

type Expression interface { // Expression interface
	Node
	expressionNode() // Including these dummy methods helps determine if we used a statement/expression in the wrong place
}

type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral() // Program node will be the root of the AST
	} else {
		return ""
	}
}

func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements { // For each string, add that string to buffer
		out.WriteString(s.String())
	}

	return out.String() // Return the string form of the buffer
}

type LetStatement struct {
	Token token.Token // token.LET token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode() 	  {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }
func (ls *LetStatement) String() string  	  {
	var out bytes.Buffer

	out.WriteString(ls.TokenLiteral() + " ") // Space after let
	out.WriteString(ls.Name.String()) // Identifier
	out.WriteString(" = ") // Add equals to buffer

	if ls.Value != nil { // If the let statement has a second side, add to buffer
		out.WriteString(ls.Value.String())
	}

	out.WriteString(";")
	return out.String() // Return the whole let statement as a string
}

type Identifier struct {
	Token token.Token // Identifier token
	Value string
}

func (i *Identifier) expressionNode() 	   {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
func (i *Identifier) String() string 	   { return i.Value } // Helper method used by statement String methods to get name

type ReturnStatement struct {
	Token token.Token // Return token
	ReturnValue Expression // Return Value
}

func (rs *ReturnStatement) statementNode() 		 {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }
func (rs *ReturnStatement) String() string 		 {
	var out bytes.Buffer

	out.WriteString(rs.TokenLiteral() + " ") // Space after return statement
	if rs.ReturnValue != nil { // If there's a value being returned, add it to buffer
		out.WriteString(rs.ReturnValue.String())
	}
	out.WriteString(";") // Add semicolon to end

	return out.String() // Return as string
}


type ExpressionStatement struct {
	Token 		token.Token
	Expression  Expression
}

func (es *ExpressionStatement) statementNode() 		 {}
func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }
func (es *ExpressionStatement) String() string 		 {
	if es.Expression == nil { // If there's no expression just return a blank string
		return ""
	} else { // Otherwise return the whole line
		return es.Expression.String()
	}
}

type IntegerLiteral struct {
	Token token.Token
	Value int64
}

func (il *IntegerLiteral) expressionNode() 		{}
func (il *IntegerLiteral) TokenLiteral() string { return il.Token.Literal }
func (il *IntegerLiteral) String() string 		{ return il.Token.Literal }

type PrefixExpression struct {
	Token 	 token.Token // ex. !, -
	Operator string
	Right 	 Expression
}

func (pe *PrefixExpression) expressionNode() 	  {}
func (pe *PrefixExpression) TokenLiteral() string { return pe.Token.Literal }
func (pe *PrefixExpression) String() string 	  {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(pe.Operator)
	out.WriteString(pe.Right.String())
	out.WriteString(")")
	// These lines capture the prefix and right side of the expression in parenthesis
	// ex. -5 becomes (-5)
	return out.String()
}

type InfixExpression struct {
	Token 	 	token.Token // Operator ex. +, -, *, /
	Left 	 	Expression
	Operator 	string
	Right 	 	Expression
}

func (ie *InfixExpression) expressionNode() 	 {}
func (ie *InfixExpression) TokenLiteral() string { return ie.Token.Literal }
func (ie *InfixExpression) String() string 		 {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(ie.Left.String())
	out.WriteString(" " + ie.Operator + " ")
	out.WriteString(ie.Right.String())
	out.WriteString(")")
	// Capture the left, operator, and right sides of the expression
	// ex. 5 + 5 becomes (5 + 5)
	return out.String()
}

type Boolean struct {
	Token	token.Token
	Value 	bool
}

func (b *Boolean) expressionNode() 		{}
func (b *Boolean) TokenLiteral() string { return b.Token.Literal }
func (b *Boolean) String() string 		{ return b.Token.Literal }

type IfExpression struct {
	Token 		token.Token
	Condition 	Expression
	Consequence *BlockStatement // If branch
	Alternative *BlockStatement // Else branch
}

func (ie *IfExpression) expressionNode() 	  {}
func (ie *IfExpression) TokenLiteral() string { return ie.Token.Literal }
func (ie *IfExpression) String() string 	  {
	var out bytes.Buffer

	out.WriteString("if") // Build the main branch ex. if CONDITION { CONSEQUENCE }
	out.WriteString(ie.Condition.String())
	out.WriteString(" ")
	out.WriteString(ie.Consequence.String())

	if ie.Alternative != nil { // If an else branch exists, build it as well
		out.WriteString("else ")
		out.WriteString(ie.Alternative.String())
	}

	return out.String()
}

type BlockStatement struct {
	Token 	   token.Token
	Statements []Statement
}

func (bs *BlockStatement) statementNode() 	  	{}
func (bs *BlockStatement) TokenLiteral() string { return bs.Token.Literal }
func (bs *BlockStatement) String() string 		{
	var out bytes.Buffer

	for _, s := range bs.Statements { // For each statement in bs.Statements
		out.WriteString(s.String())
	}

	return out.String()
}

type FunctionLiteral struct {
	Token 		token.Token // fn
	Parameters  []*Identifier
	Body 		*BlockStatement
}

func (fl *FunctionLiteral) expressionNode() 	  {}
func (fl *FunctionLiteral) TokenLiteral() string  { return fl.Token.Literal }
func (fl *FunctionLiteral) String() string		  {
	var out bytes.Buffer

	params := []string{}
	for _, p := range fl.Parameters {
		params = append(params, p.String())
	}

	out.WriteString(fl.TokenLiteral())
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") ")
	out.WriteString(fl.Body.String())

	return out.String()
}

type CallExpression struct {
	Token	  token.Token // '('
	Function  Expression // Identifier
	Arguments []Expression // List of arguments
}

func (ce *CallExpression) expressionNode() 	  	{}
func (ce *CallExpression) TokenLiteral() string { return ce.Token.Literal }
func (ce *CallExpression) String() string 		{
	var out bytes.Buffer

	args := []string{}
	for _, a := range ce.Arguments { // For each argument, append to array args
		args = append(args, a.String())
	}

	out.WriteString(ce.Function.String())
	out.WriteString("(")
	out.WriteString(strings.Join(args, ", "))
	out.WriteString(")")

	return out.String()
}
