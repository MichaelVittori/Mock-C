package ast

import "mockc/token"

type Node interface { // Define our treenode interface
	TokenLiteral() string
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

type LetStatement struct {
	Token token.Token // token.LET token
	Name *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode() {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

type Identifier struct {
	Token token.Token // Identifier token
	Value string
}

func (i *Identifier) expressionNode() {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

type ReturnStatement struct {
	Token token.Token // Return token
	ReturnValue Expression // Return Value
}

func (rs *ReturnStatement) statementNode() {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }