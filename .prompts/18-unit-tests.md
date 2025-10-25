## 🎯 PROMPT 18: TESTES UNITÁRIOS EM GO

```
TAREFA: Criar testes unitários em Go para Result e Promise

ARQUIVOS A CRIAR:
- eval/result_test.go
- eval/promise_test.go

ESTRUTURA DE result_test.go:

package eval

import (
    "testing"
)

func TestResultOk(t *testing.T) {
    result := &ResultObject{
        IsOk:  true,
        Value: &Integer{Value: 42},
        Error: NIL,
    }
    
    if !result.IsOk {
        t.Error("Expected IsOk to be true")
    }
    
    if result.Value.(*Integer).Value != 42 {
        t.Error("Expected value to be 42")
    }
}

func TestResultErr(t *testing.T) {
    result := &ResultObject{
        IsOk:  false,
        Value: NIL,
        Error: &String{Value: "error message"},
    }
    
    if result.IsOk {
        t.Error("Expected IsOk to be false")
    }
    
    if result.Error.(*String).Value != "error message" {
        t.Error("Expected error message to match")
    }
}

func TestResultMap(t *testing.T) {
    // Setup
    scope := NewScope(nil, "")
    result := &ResultObject{
        IsOk:  true,
        Value: &Integer{Value: 5},
        Error: NIL,
    }
    
    mapFn := &Function{
        // ... criar função que multiplica por 2
    }
    
    // Execute
    mapped := result.resultMap("test", scope, []Object{mapFn})
    
    // Assert
    if !mapped.(*ResultObject).IsOk {
        t.Error("Expected mapped result to be Ok")
    }
    
    if mapped.(*ResultObject).Value.(*Integer).Value != 10 {
        t.Error("Expected mapped value to be 10")
    }
}

// Adicionar mais testes para:
// - TestResultFlatMap
// - TestResultMatch
// - TestResultGetOrElse
// - TestResultErrorPropagation

ESTRUTURA DE promise_test.go:

package eval

import (
    "testing"
    "time"
)

func TestPromiseResolve(t *testing.T) {
    scope := NewScope(nil, "")
    
    executor := &Function{
        // ... função que chama resolve(42)
    }
    
    promise := newPromise("test", scope, executor)
    
    // Dar tempo para goroutine executar
    time.Sleep(10 * time.Millisecond)
    
    result := promise.(*PromiseObject).promiseAwait("test", scope)
    
    if result.Type() == ERROR_OBJ {
        t.Error("Expected promise to resolve, got error:", result.Inspect())
    }
    
    if result.(*Integer).Value != 42 {
        t.Error("Expected resolved value to be 42")
    }
}

func TestPromiseReject(t *testing.T) {
    scope := NewScope(nil, "")
    
    executor := &Function{
        // ... função que chama reject("error")
    }
    
    promise := newPromise("test", scope, executor)
    
    time.Sleep(10 * time.Millisecond)
    
    result := promise.(*PromiseObject).promiseAwait("test", scope)
    
    if result.Type() != ERROR_OBJ {
        t.Error("Expected promise to reject")
    }
}

func TestPromiseThen(t *testing.T) {
    // Test promise chaining
    // ...
}

func TestPromiseConcurrency(t *testing.T) {
    // Test multiple concurrent awaits
    // ...
}

// Adicionar mais testes

EXECUTAR TESTES:
go test ./eval -v

VALIDAÇÃO:
□ Todos os testes passam
□ Coverage > 80%
□ Sem race conditions (go test -race)
```

---
