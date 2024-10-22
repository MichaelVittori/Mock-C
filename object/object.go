package object

import (
	"bytes"
	"fmt"
	"mockc/ast"
	"strings"
	"hash/fnv"
)

type ObjectType string
const (
	INTEGER_OBJECT  = "INTEGER"
	BOOLEAN_OBJECT  = "BOOLEAN"
	NULL_OBJECT     = "NULL"
	RETURN_OBJECT   = "RETURN"
	ERROR_OBJECT    = "ERROR"
	FUNCTION_OBJECT = "FUNCTION"
	STRING_OBJECT   = "STRING"
	BUILTIN_OBJECT  = "BUILTIN"
	ARRAY_OBJECT    = "ARRAY"
	HASH_OBJECT     = "HASH"
)

// All values encountered when evaluating Moxie source code will be wrapped in a struct fulfilling the Object interface
type Object interface { // Note for myself: interfaces abstractly define behavior, structs define states/values
	Type() ObjectType
	Inspect() string
}

type Hashable interface {
	HashKey() HashKey
}

type HashKey struct { // Possible hashing improvements: separate chaining for collisions, caching hash keys for improved performance
	Type  ObjectType
	Value uint64
}

type Integer struct {
	Value int64
}

func (i *Integer) Type() ObjectType { return INTEGER_OBJECT }
func (i *Integer) Inspect() string { return fmt.Sprintf("%d", i.Value) }
func (i *Integer) HashKey() HashKey {
	return HashKey{Type: i.Type(), Value: uint64(i.Value)}
}

type Boolean struct {
	Value bool
}

func (b *Boolean) Type() ObjectType { return BOOLEAN_OBJECT }
func (b *Boolean) Inspect() string { return fmt.Sprintf("%t", b.Value)}
func (b *Boolean) HashKey() HashKey {
	var value uint64

	if b.Value {
		value = 1
	} else {
		value = 0
	}

	return HashKey{Type: b.Type(), Value: value}
}

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

type String struct {
	Value string
}

func (s *String) Type() ObjectType { return STRING_OBJECT }
func (s *String) Inspect() string { return s.Value }
func (s *String) HashKey() HashKey {
	hash := fnv.New64a()
	hash.Write([]byte(s.Value))

	return HashKey{Type: s.Type(), Value: hash.Sum64()}
}

type BuiltInFunction func(args ...Object) Object // Basic built in abstract signature, accepts 0 or more objects as args and returns an object

type BuiltIn struct {
	Fn BuiltInFunction
}

func (b *BuiltIn) Type() ObjectType { return BUILTIN_OBJECT }
func (b *BuiltIn) Inspect() string { return "builtin function" }

type Array struct {
	Elements []Object
}

func (a *Array) Type() ObjectType { return ARRAY_OBJECT }
func (a *Array) Inspect() string {
	var out bytes.Buffer

	elements := []string{}
	for _, e := range a.Elements {
		elements = append(elements, e.Inspect())
	}

	out.WriteString("[")
	out.WriteString(strings.Join(elements, ", "))
	out.WriteString("]")

	return out.String()
}

type HashPair struct {
	Key   Object
	Value Object
}

type Hash struct {
	Pairs map[HashKey]HashPair
}

func (h *Hash) Type() ObjectType { return HASH_OBJECT }
func (h *Hash) Inspect() string {
	var out bytes.Buffer

	pairs := []string{}
	for _, pair := range h.Pairs {
		pairs = append(pairs, fmt.Sprintf("%s: %s", pair.Key.Inspect(), pair.Value.Inspect()))
	}

	out.WriteString("{")
	out.WriteString(strings.Join(pairs, ", "))
	out.WriteString("}")

	return out.String()
}