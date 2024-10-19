package object

import (
	"bytes"
	"fmt"
	"mockc/ast"
	"strings"
)

type ObjectType string
const (
	INTEGER_OBJECT  = "INTEGER"
	BOOLEAN_OBJECT  = "BOOLEAN"
	NULL_OBJECT     = "NULL"
	RETURN_OBJECT   = "RETURN"
	ERROR_OBJECT    = "ERROR"
	FUNCTION_OBJECT = "FUNCTION"
)

// All values encountered when evaluating Moxie source code will be wrapped in a struct fulfilling the Object interface
type Object interface { // Note for myself: interfaces abstractly define behavior, structs define states/values
	Type() ObjectType
	Inspect() string
}

type Integer struct {
	Value int64
}

func (i *Integer) Type() ObjectType { return INTEGER_OBJECT }
func (i *Integer) Inspect() string { return fmt.Sprintf("%d", i.Value) }

type Boolean struct {
	Value bool
}

func (b *Boolean) Type() ObjectType { return BOOLEAN_OBJECT }
func (b *Boolean) Inspect() string { return fmt.Sprintf("%t", b.Value)}

type Null struct {} // No value field because null has no value

func (n *Null) Type() ObjectType { return NULL_OBJECT }
func (n *Null) Inspect() string { return fmt.Sprintf("null")}

type ReturnValue struct {
	Value Object
}

func (rv *ReturnValue) Type() ObjectType { return RETURN_OBJECT }
func (rv *ReturnValue) Inspect() string { return rv.Value.Inspect() }

type Error struct { // If this were a really real language there would be a stack trace in here too
	Message string
}

func (e *Error) Type() ObjectType { return ERROR_OBJECT }
func (e *Error) Inspect() string { return e.Message }

type Function struct {
	Parameters []*ast.Identifier
	Body       *ast.BlockStatement
	Env 	   *Environment
}

func (f *Function) Type() ObjectType { return FUNCTION_OBJECT }
func (f *Function) Inspect() string {
	// Similar to the ast structs, we morph the raw function def into a formatted string ex. fn(x) {\n bodystuff \n}
	var out bytes.Buffer
	params := []string{}

	for _, p := range f.Parameters { params = append(params, p.String())}

	out.WriteString("fn")
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") {\n")
	out.WriteString(f.Body.String())
	out.WriteString("\n}")

	return out.String()
}