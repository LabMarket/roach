package evaluator

import (
	"math"
	"math/rand"
	"time"

	"github.com/LabMarket/roach/object"
)

func mathAbs(args ...object.Object) object.Object {
	if len(args) != 1 {
		return newError("wrong number of arguments. got=%d, want=1",
			len(args))
	}
	switch arg := args[0].(type) {
	case *object.Integer:
		v := arg.Value
		if v < 0 {
			v = v * -1
		}
		return &object.Integer{Value: v}
	case *object.Float:
		v := arg.Value
		if v < 0 {
			v = v * -1
		}
		return &object.Float{Value: v}
	default:
		return newError("argument to `math.abs` not supported, got=%s",
			args[0].Type())
	}

}

// val = math.random()
func mathRandom(args ...object.Object) object.Object {
	return &object.Float{Value: rand.Float64()}
}

// val = math.sqrt(int);
func mathSqrt(args ...object.Object) object.Object {
	if len(args) != 1 {
		return newError("wrong number of arguments. got=%d, want=1",
			len(args))
	}
	switch arg := args[0].(type) {
	case *object.Integer:
		v := arg.Value
		return &object.Float{Value: math.Sqrt(float64(v))}
	case *object.Float:
		v := arg.Value
		return &object.Float{Value: math.Sqrt(v)}
	default:
		return newError("argument to `math.sqrt` not supported, got=%s",
			args[0].Type())
	}

}

// val = math.sin(int);
func mathSin(args ...object.Object) object.Object {
	if len(args) != 1 {
		return newError("wrong number of arguments. got=%d, want=1",
			len(args))
	}
	switch arg := args[0].(type) {
	case *object.Integer:
		v := arg.Value
		return &object.Float{Value: math.Sin(float64(v))}
	case *object.Float:
		v := arg.Value
		return &object.Float{Value: math.Sin(v)}
	default:
		return newError("argument to `math.sin` not supported, got=%s",
			args[0].Type())	
	}
}

// val = math.cos(int);
func mathCos(args ...object.Object) object.Object {
	if len(args) != 1 {
		return newError("wrong number of arguments. got=%d, want=1",
			len(args))
	}
	switch arg := args[0].(type) {
	case *object.Integer:
		v := arg.Value
		return &object.Float{Value: math.Cos(float64(v))}
	case *object.Float:
		v := arg.Value
		return &object.Float{Value: math.Cos(v)}
	default:
		return newError("argument to `math.cos` not supported, got=%s",
			args[0].Type())	
	}
}

// val = math.tan(int);
func mathTan(args ...object.Object) object.Object {
	if len(args) != 1 {
		return newError("wrong number of arguments. got=%d, want=1",
			len(args))
	}
	switch arg := args[0].(type) {
	case *object.Integer:
		v := arg.Value
		return &object.Float{Value: math.Tan(float64(v))}
	case *object.Float:
		v := arg.Value
		return &object.Float{Value: math.Tan(v)}
	default:
		return newError("argument to `math.tan` not supported, got=%s",
			args[0].Type())	
	}
}

// val = math.log(int);
func mathLog(args ...object.Object) object.Object {
	if len(args) != 1 {
		return newError("wrong number of arguments. got=%d, want=1",
			len(args))
	}
	switch arg := args[0].(type) {
	case *object.Integer:
		v := arg.Value
		if v <= 0 {
			return newError("argument to `math.log` must be greater than zero, got=%d", v)
		}
		return &object.Float{Value: math.Log(float64(v))}
	case *object.Float:
		v := arg.Value
		if v <= 0 {
			return newError("argument to `math.log` must be greater than zero, got=%f", v)
		}
		return &object.Float{Value: math.Log(v)}
	default:
		return newError("argument to `math.log` not supported, got=%s",
			args[0].Type())
	}
}

// val = math.ln(int);
func mathLn(args ...object.Object) object.Object {
	if len(args) != 1 {
		return newError("wrong number of arguments. got=%d, want=1",
			len(args))
	}
	switch arg := args[0].(type) {
	case *object.Integer:
		v := arg.Value
		if v <= 0 {
			return newError("argument to `math.ln` must be greater than zero, got=%d", v)
		}
		return &object.Float{Value: math.Log(float64(v))}
	case *object.Float:
		v := arg.Value
		if v <= 0 {
			return newError("argument to `math.ln` must be greater than zero, got=%f", v)
		}
		return &object.Float{Value: math.Log(v)}
	default:
		return newError("argument to `math.ln` not supported, got=%s",
			args[0].Type())
	}
}

// val = math.exp(int);
func mathExp(args ...object.Object) object.Object {
	if len(args) != 1 {
		return newError("wrong number of arguments. got=%d, want=1",
			len(args))
	}
	switch arg := args[0].(type) {
	case *object.Integer:
		v := arg.Value
		return &object.Float{Value: math.Exp(float64(v))}
	case *object.Float:
		v := arg.Value
		return &object.Float{Value: math.Exp(v)}
	default:
		return newError("argument to `math.exp` not supported, got=%s",
			args[0].Type())
	}
}

func init() {
	//
	// Setup our random seed.
	//
	rand.Seed(time.Now().UnixNano())
	RegisterBuiltin("math.abs",
		func(env *object.Environment, args ...object.Object) object.Object {
			return (mathAbs(args...))
		})
	RegisterBuiltin("math.sin",
		func(env *object.Environment, args ...object.Object) object.Object {
			return (mathSin(args...))
		})
	RegisterBuiltin("math.cos",
		func(env *object.Environment, args ...object.Object) object.Object {
			return (mathCos(args...))
		})
	RegisterBuiltin("math.tan",
		func(env *object.Environment, args ...object.Object) object.Object {
			return (mathTan(args...))
		})
	RegisterBuiltin("math.random",
		func(env *object.Environment, args ...object.Object) object.Object {
			return (mathRandom(args...))
		})
	RegisterBuiltin("math.sqrt",
		func(env *object.Environment, args ...object.Object) object.Object {
			return (mathSqrt(args...))
		})
	RegisterBuiltin("math.log",
		func(env *object.Environment, args ...object.Object) object.Object {
			return (mathLog(args...))
		})
	RegisterBuiltin("math.ln",
		func(env *object.Environment, args ...object.Object) object.Object {
			return (mathLn(args...))
		})
	RegisterBuiltin("math.exp",
		func(env *object.Environment, args ...object.Object) object.Object {
			return (mathExp(args...))
		})
	
}
