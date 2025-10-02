package evaluator

import (
	"strings"
	"github.com/LabMarket/roach/object"
)

func stringToUpper(args ...object.Object) object.Object {
	if len(args) != 1 {
		return newError("wrong number of arguments. got=%d, want=1", len(args))
	}
	str, ok := args[0].(*object.String)
	if !ok {
		return newError("argument to `toupper` must be STRING, got %s", args[0].Type())
	}
	return &object.String{Value: strings.ToUpper(str.Value)}
}

func init() {
	// Create the module
	stringModule := &object.Module{
		Name: "string",
		Members: make(map[string]object.Object),
	}

	// Register the functions
	stringModule.Members["toupper"] = &object.Builtin{Fn: func(env *object.Environment, args ...object.Object) object.Object {
		return (stringToUpper(args...))
	}}

	// Register the module itself.
	RegisterBuiltin("string", stringModule)
}
