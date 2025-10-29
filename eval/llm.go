package eval

import (
	"context"
	"strings"

	"github.com/ollama/ollama/api"
)

const (
	LLM_OBJ  = "LLM_OBJ"
	llm_name = "llm"
)

type LLMClientObject struct {
	Client          *api.Client
	Host            string
	model           string
	ctxSize         int64
	numGPU          int64
	numBatch        int64
	temperature     float64
	numThread       int64
	numPredict      int64
	repeatPenalty   float64
	systemCard      string
	conversationCtx []int
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
		return t.setModel(line, args...)
	case "model":
		return NewString(t.model)
	case "setCtxSize":
		return t.setCtxSize(line, args...)
	case "ctxSize":
		return NewInteger(t.ctxSize)
	case "setNumGPU":
		return t.setNumGPU(line, args...)
	case "numGPU":
		return NewInteger(t.numGPU)
	case "setNumBatch":
		return t.setNumBatch(line, args...)
	case "numBatch":
		return NewInteger(t.numBatch)
	case "setTemperature":
		return t.setTemperature(line, args...)
	case "temperature":
		return NewFloat(t.temperature)
	case "setNumThread":
		return t.setNumThread(line, args...)
	case "numThread":
		return NewInteger(t.numThread)
	case "setNumPredict":
		return t.setNumPredict(line, args...)
	case "numPredict":
		return NewInteger(t.numPredict)
	case "setRepeatPenalty":
		return t.setRepeatPenalty(line, args...)
	case "repeatPenalty":
		return NewFloat(t.repeatPenalty)
	case "systemCard":
		return NewString(t.systemCard)
	case "setSystemCard":
		return t.setSystemCard(line, args...)
	case "generate":
		return t.generate(line, args...)
	case "resetContext":
		return t.resetContext(line, args...)
	}
	return NewError(line, NOMETHODERROR, method, t.Type())
}

func (t *LLMClientObject) setModel(line string, args ...Object) Object {
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

func (t *LLMClientObject) setCtxSize(line string, args ...Object) Object {
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

func (t *LLMClientObject) setNumGPU(line string, args ...Object) Object {
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

func (t *LLMClientObject) setNumBatch(line string, args ...Object) Object {
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

func (t *LLMClientObject) setTemperature(line string, args ...Object) Object {
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

func (t *LLMClientObject) setNumThread(line string, args ...Object) Object {
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

func (t *LLMClientObject) setNumPredict(line string, args ...Object) Object {
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

func (t *LLMClientObject) setRepeatPenalty(line string, args ...Object) Object {
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

func (t *LLMClientObject) setSystemCard(line string, args ...Object) Object {
	if len(args) != 1 {
		return NewError(line, ARGUMENTERROR, "1", len(args))
	}

	s, ok := args[0].(*String)
	if !ok {
		return NewError(line, PARAMTYPEERROR, "first", "setSystemCard", "*String", args[0].Type())
	}

	t.systemCard = s.String
	return NIL
}

func (t *LLMClientObject) resetContext(line string, args ...Object) Object {
	if len(args) != 0 {
		return NewError(line, ARGUMENTERROR, "0", len(args))
	}

	t.conversationCtx = nil

	return NIL
}

func (t *LLMClientObject) generate(line string, args ...Object) Object {
	if len(args) != 1 {
		return NewError(line, ARGUMENTERROR, "1", len(args))
	}

	promptStr, ok := args[0].(*String)
	if !ok {
		return NewError(line, PARAMTYPEERROR, "first", "prompt", "*String", args[0].Type())
	}

	ctx := context.Background()

	options := map[string]interface{}{}
	if t.ctxSize != 0 {
		options["num_ctx"] = t.ctxSize
	}
	if t.numGPU != 0 {
		options["num_gpu"] = t.numGPU
	}
	if t.numBatch != 0 {
		options["num_batch"] = t.numBatch
	}
	if t.temperature != 0.0 {
		options["temperature"] = t.temperature
	}
	if t.numThread != 0 {
		options["num_thread"] = t.numThread
	}
	if t.numPredict != 0 {
		options["num_predict"] = t.numPredict
	}
	if t.repeatPenalty != 0.0 {
		options["repeat_penalty"] = t.repeatPenalty
	}
	req := &api.GenerateRequest{
		Model:   t.model,
		Prompt:  promptStr.String,
		Options: options,
		System:  t.systemCard,
		Context: t.conversationCtx,
	}

	var sb strings.Builder
	responseFunc := func(r api.GenerateResponse) error {
		sb.WriteString(r.Response)
		// Se a resposta estiver completa (Done), atualizamos o contexto para a próxima chamada.
		if r.Done {
			t.conversationCtx = r.Context
		}
		return nil
	}

	err := t.Client.Generate(ctx, req, responseFunc)
	if err != nil {
		return NewError(line, GENERICERROR, err.Error())
	}

	return NewString(sb.String())
}
