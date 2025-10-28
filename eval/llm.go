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
}

func NewClientObjectj() Object {
	ret := &LLMClientObject{}
	SetGlobalObj(llm_name, ret)

	SetGlobalObj(llm_name+".Host", NewString(GetEnvStr("OLLAMA_HOST", "http://localhost:11434")))
	SetGlobalObj(llm_name+".Model", NewString(GetEnvStr("OLLAMA_MODEL", "llama2")))

	// Context window
	SetGlobalObj(llm_name+".CtxSize", NewInteger(GetEnvInt("LLM_NUM_CTX", 8192)))   // num_ctx
	SetGlobalObj(llm_name+".NumGPU", NewInteger(GetEnvInt("LLM_NUM_GPU", 99)))      // num_gpu
	SetGlobalObj(llm_name+".NumBatch", NewInteger(GetEnvInt("LLM_NUM_BATCH", 512))) // num_batch

	// Inference tunning
	SetGlobalObj(llm_name+".Temperature", NewFloat(GetEnvFloat("LLM_TEMPERATURE", 0.7)))      // temperature
	SetGlobalObj(llm_name+".NumThread", NewInteger(GetEnvInt("LLM_NUM_THREAD", 16)))          // num_thread
	SetGlobalObj(llm_name+".NumPredict", NewInteger(GetEnvInt("LLM_NUM_PREDICT", -1)))        // num_predict
	SetGlobalObj(llm_name+".RepeatPenalty", NewFloat(GetEnvFloat("LLM_REPEAT_PENALTY", 1.1))) // repeat_penalty

	return ret
}

func (t *LLMClientObject) Inspect() string {
	return "<" + llm_name + ">"
}

func (t *LLMClientObject) Type() ObjectType { return LLM_OBJ }

func (t *LLMClientObject) CallMethod(line string, scope *Scope, method string, args ...Object) Object {
	switch method {
	}
	return NewError(line, NOMETHODERROR, method, t.Type())
}
