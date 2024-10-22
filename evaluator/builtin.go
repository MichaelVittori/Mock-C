package evaluator

import (
	"mockc/object"
	"fmt"
)

// Basically a second environment but for our builtin functions
var builtins = map[string]*object.BuiltIn {
	"len": &object.BuiltIn{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 { return newError("Wrong number of arguments. got=%d, want=1", len(args))}

			switch arg := args[0].(type) {
			case *object.Array: return &object.Integer{Value: int64(len(arg.Elements))}
			case *object.String:
				return &object.Integer{Value: int64(len(arg.Value))}
			default:
				return newError("Argument to `len` not supported, got %s", args[0].Type())
			}
		},
	},

	"first": &object.BuiltIn {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 { return newError("Wrong number of arguments. got=%d, want=1", len(args)) }
			if args[0].Type() != object.ARRAY_OBJECT { return newError("Argument to 'first' must be ARRAY, got %s", args[0].Type()) }
			arr := args[0].(*object.Array)
			if len(arr.Elements) > 0 { return arr.Elements[0] }

			return NULL // If the array is empty, return null
		},
	},

	"last": &object.BuiltIn {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 { return newError("Wrong number of arguments. got=%d, want=1", len(args)) }
			if args[0].Type() != object.ARRAY_OBJECT { return newError("Argument to 'last' must be ARRAY, got %s", args[0].Type()) }
			arr := args[0].(*object.Array)
			length := len(arr.Elements)
			if length > 0 { return arr.Elements[length - 1] }

			return NULL // If the array is empty, return null
		},
	},

	"rest": &object.BuiltIn {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 { return newError("Wrong number of arguments. got=%d, want=1", len(args)) }
			if args[0].Type() != object.ARRAY_OBJECT { return newError("Argument to 'rest' must be ARRAY, got %s", args[0].Type()) }
			arr := args[0].(*object.Array)
			length := len(arr.Elements)

			if length > 0 {
				newElements := make([]object.Object, length - 1, length - 1)
				copy(newElements, arr.Elements[1 : length]) // Take the array slice starting at index 1 to the end
				return &object.Array{Elements: newElements}
			}

			return NULL // If the array is empty, return null
		},
	},

	"push": &object.BuiltIn {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 { return newError("Wrong number of arguments. got=%d, want=1", len(args)) }
			if args[0].Type() != object.ARRAY_OBJECT { return newError("Argument to 'push' must be ARRAY, got %s", args[0].Type()) }
			arr := args[0].(*object.Array)
			length := len(arr.Elements)

			newElements := make([]object.Object, length + 1, length + 1)
			copy(newElements, arr.Elements) // Take the array slice starting at index 1 to the end
			newElements[length] = args[1]

			return &object.Array{Elements: newElements}
		},
	},

	"print": &object.BuiltIn {
		Fn: func(args ...object.Object) object.Object {
			for _, arg := range args { fmt.Println(arg.Inspect()) }
			return NEWLINE // The null return looked bad so I added a new constant to evaluator
		},
	},
}
