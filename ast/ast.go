package ast

import (
	"mockc/token"
	"bytes"
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
	Name *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode() {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }
func (ls *LetStatement) String() string {
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

func (i *Identifier) expressionNode() {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
func (i *Identifier) String() string { return i.Value } // Helper method used by statement String methods to get name

type ReturnStatement struct {
	Token token.Token // Return token
	ReturnValue Expression // Return Value
}

func (rs *ReturnStatement) statementNode() {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }
func (rs *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(rs.TokenLiteral() + " ") // Space after return statement
	if rs.ReturnValue != nil { // If there's a value being returned, add it to buffer
		out.WriteString(rs.ReturnValue.String())
	}
	out.WriteString(";") // Add semicolon to end

	return out.String() // Return as string
}


type ExpressionStatement struct {
	Token token.Token
	Expression Expression
}

func (es *ExpressionStatement) statementNode() {}
func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }
func (es *ExpressionStatement) String() string {
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

func (il *IntegerLiteral) expressionNode() {}
func (il *IntegerLiteral) TokenLiteral() string { return il.Token.Literal }
func (il *IntegerLiteral) String() string 		{ return il.Token.Literal }