package evaluator

import (
	"Goslang/object"
	"math/rand"
	"sort"
	"strconv"
)

var builtins = map[string]*object.Builtin{
	"len": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("Compile Error: len() function can only have 1 argument")
			}

			switch arg := args[0].(type) {
			case *object.Array:
				return &object.Integer{Value: int64(len(arg.Elements))}
			case *object.String:
				return &object.Integer{Value: int64(len(arg.Value))}

			default:
				return newError("argument to `len` not supported, got %s", args[0].Type())
			}
		},
	},

	"Atoi": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("Compile Error: `Atoi` function can only have 1 argument")
			}

			if arg, ok := args[0].(*object.String); ok {
				str := arg.Value
				int64Value, err := strconv.ParseInt(str, 10, 64)
				if err != nil {
					return newError("Error Parsing string to int64: %s", err.Error())
				}
				return &object.Integer{Value: int64Value}
			}

			return newError("Compile Error: Argument to `Atoi` must be a STRING, got %s", args[0].Type())

		},
	},

	"first": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("Compile error: `first` function can only have 1 argument")
			}
			switch arg := args[0].(type) {
			case *object.Array:
				if len(arg.Elements) > 0 {
					return arg.Elements[0]
				}
			case *object.String:
				str := arg.Value
				if len(str) > 0 {
					return &object.String{Value: string(str[0])}
				}
			default:
				return newError("argument to `first` must be ARRAY or STRING, got %s", args[0].Type())
			}

			return NULL

		},
	},

	"last": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("Compile error: `last` function can only have 1 argument")
			}

			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argument to `last` must be ARRAY, got %s", args[0].Type())
			}

			arr := args[0].(*object.Array)
			length := len(arr.Elements)
			if length > 0 {
				return arr.Elements[length-1]
			}
			return NULL
		},
	},

	"rest": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("Compile error: `rest` function can only have 1 argument")
			}

			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argument to `rest` must be ARRAY, got %s", args[0].Type())
			}
			arr := args[0].(*object.Array)
			length := len(arr.Elements)
			if length > 0 {
				newElements := make([]object.Object, length-1, length-1)
				copy(newElements, arr.Elements[1:length])
				return &object.Array{Elements: newElements}
			}
			return NULL
		},
	},

	"push": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return newError("Compile error: `push` function must have 2 arguments")
			}
			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argument to `push` must be ARRAY, got %s", args[0].Type())
			}
			arr := args[0].(*object.Array)
			/*length := len(arr.Elements)
			newElements := make([]object.Object, length+1, length+1)
			copy(newElements, arr.Elements)
			newElements[length] = args[1]
			return &object.Array{Elements: newElements}*/
			arr.Elements = append(arr.Elements, args[1])
			return arr
		},
	},

	"randInt": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return newError("Compile error: `randInt` function must have 2 arguments")
			}

			if args[0].Type() != object.INTEGER_OBJ || args[1].Type() != object.INTEGER_OBJ {
				return newError("arguments to `random` must be INTEGER, got %s and %s", args[0].Type(), args[1].Type())
			}

			min := args[0].(*object.Integer).Value
			max := args[1].(*object.Integer).Value

			if min >= max {
				return newError("min value must be less than max value")
			}

			// Generate a random number between min and max
			randomValue := min + rand.Int63n(max-min+1)

			return &object.Integer{Value: randomValue}
		},
	},

	"randPick": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("Compile error: `randomElement` function must have 1 argument")
			}

			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argument to `randomElement` must be ARRAY, got %s", args[0].Type())
			}

			arr := args[0].(*object.Array)
			length := len(arr.Elements)

			if length == 0 {
				return NULL // Return NULL if the array is empty
			}

			randomIndex := rand.Intn(length)

			return arr.Elements[randomIndex]
		},
	},

	"sort": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("Compile Error: `sort` function must have 1 argument")
			}

			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argument to `sort` must be ARRAY, got %s", args[0].Type())
			}

			arr := args[0].(*object.Array)
			elements := arr.Elements

			// Check if the elements are either integers or strings
			for _, element := range elements {
				if element.Type() != object.INTEGER_OBJ && element.Type() != object.STRING_OBJ {
					return newError("elements in the array must be INTEGER or STRING, got %s", element.Type())
				}
			}

			// Sort the elements in-place based on their type
			sort.Slice(elements, func(i, j int) bool {
				if elements[i].Type() == object.INTEGER_OBJ && elements[j].Type() == object.INTEGER_OBJ {
					return elements[i].(*object.Integer).Value < elements[j].(*object.Integer).Value
				} else if elements[i].Type() == object.STRING_OBJ && elements[j].Type() == object.STRING_OBJ {
					return elements[i].(*object.String).Value < elements[j].(*object.String).Value
				}
				return false
			})

			return arr // Sorted array
		},
	},
}
