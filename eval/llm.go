package eval

import (
	"github.com/ollama/ollama/api"
)

const (
	LLM_OBJ  = "LLM_OBJ"
	llm_name = "llm"
)

type LLMClientObject struct {
	Client *api.Client
	Host   string

	model         string
	ctxSize       int64
	numGPU        int64
	numBatch      int64
	temperature   float64
	numThread     int64
	numPredict    int64
	repeatPenalty float64
}

func NewClientObjectj() Object {
	ret := &LLMClientObject{}
	SetGlobalObj(llm_name, ret)

	return ret
}

func (t *LLMClientObject) Inspect() string {
	return "<" + llm_name + ">"
}

func (t *LLMClientObject) Type() ObjectType { return LLM_OBJ }

func (t *LLMClientObject) CallMethod(line string, scope *Scope, method string, args ...Object) Object {
	switch method {
	case "host":
		return NewString(t.Host)
	case "setModel":
		return t.SetModel(line, args...)
	case "model":
		return NewString(t.model)
	case "setCtxSize":
		return t.SetCtxSize(line, args...)
	case "ctxSize":
		return NewInteger(t.ctxSize)
	case "setNumGPU":
		return t.SetNumGPU(line, args...)
	case "numGPU":
		return NewInteger(t.numGPU)
	case "setNumBatch":
		return t.SetNumBatch(line, args...)
	case "numBatch":
		return NewInteger(t.numBatch)
	case "setTemperature":
		return t.SetTemperature(line, args...)
	case "temperature":
		return NewFloat(t.temperature)
	case "setNumThread":
		return t.SetNumThread(line, args...)
	case "numThread":
		return NewInteger(t.numThread)
	case "setNumPredict":
		return t.SetNumPredict(line, args...)
	case "numPredict":
		return NewInteger(t.numPredict)
	case "setRepeatPenalty":
		return t.SetRepeatPenalty(line, args...)
	case "repeatPenalty":
		return NewFloat(t.repeatPenalty)
	}
	return NewError(line, NOMETHODERROR, method, t.Type())
}

func (t *LLMClientObject) SetModel(line string, args ...Object) Object {
	if len(args) != 1 {
		return NewError(line, ARGUMENTERROR, "1", len(args))
	}

	s, ok := args[0].(*String)
	if !ok {
		return NewError(line, PARAMTYPEERROR, "first", "setModel", "*String", args[0].Type())
	}

	t.model = s.String
	return NIL
}

func (t *LLMClientObject) SetCtxSize(line string, args ...Object) Object {
	if len(args) != 1 {
		return NewError(line, ARGUMENTERROR, "1", len(args))
	}

	n, ok := args[0].(*Integer)
	if !ok {
		return NewError(line, PARAMTYPEERROR, "first", "setCtxSize", "*Integer", args[0].Type())
	}

	t.ctxSize = n.Int64
	return NIL
}

func (t *LLMClientObject) SetNumGPU(line string, args ...Object) Object {
	if len(args) != 1 {
		return NewError(line, ARGUMENTERROR, "1", len(args))
	}

	n, ok := args[0].(*Integer)
	if !ok {
		return NewError(line, PARAMTYPEERROR, "first", "setCtxSize", "*Integer", args[0].Type())
	}

	t.numGPU = n.Int64
	return NIL
}

func (t *LLMClientObject) SetNumBatch(line string, args ...Object) Object {
	if len(args) != 1 {
		return NewError(line, ARGUMENTERROR, "1", len(args))
	}

	n, ok := args[0].(*Integer)
	if !ok {
		return NewError(line, PARAMTYPEERROR, "first", "setCtxSize", "*Integer", args[0].Type())
	}

	t.numBatch = n.Int64
	return NIL
}

func (t *LLMClientObject) SetTemperature(line string, args ...Object) Object {
	if len(args) != 1 {
		return NewError(line, ARGUMENTERROR, "1", len(args))
	}

	n, ok := args[0].(*Float)
	if !ok {
		return NewError(line, PARAMTYPEERROR, "first", "setCtxSize", "*Integer", args[0].Type())
	}

	t.temperature = n.Float64
	return NIL
}

func (t *LLMClientObject) SetNumThread(line string, args ...Object) Object {
	if len(args) != 1 {
		return NewError(line, ARGUMENTERROR, "1", len(args))
	}

	n, ok := args[0].(*Integer)
	if !ok {
		return NewError(line, PARAMTYPEERROR, "first", "setCtxSize", "*Integer", args[0].Type())
	}

	t.numThread = n.Int64
	return NIL
}

func (t *LLMClientObject) SetNumPredict(line string, args ...Object) Object {
	if len(args) != 1 {
		return NewError(line, ARGUMENTERROR, "1", len(args))
	}

	n, ok := args[0].(*Integer)
	if !ok {
		return NewError(line, PARAMTYPEERROR, "first", "setCtxSize", "*Integer", args[0].Type())
	}

	t.numPredict = n.Int64
	return NIL
}

func (t *LLMClientObject) SetRepeatPenalty(line string, args ...Object) Object {
	if len(args) != 1 {
		return NewError(line, ARGUMENTERROR, "1", len(args))
	}

	n, ok := args[0].(*Float)
	if !ok {
		return NewError(line, PARAMTYPEERROR, "first", "setCtxSize", "*Integer", args[0].Type())
	}

	t.repeatPenalty = n.Float64
	return NIL
}
