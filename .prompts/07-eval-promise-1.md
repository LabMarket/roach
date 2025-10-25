## 🎯 PROMPT 7: IMPLEMENTAR eval/promise.go (PARTE 1 - ESTRUTURA)

```
TAREFA: Criar eval/promise.go com estrutura básica do PromiseObject

CONTEXTO:
Promise usa channels do Go internamente para comunicação entre threads.
Use eval/channel.go como referência.

REQUISITOS:

1. IMPORTS NECESSÁRIOS:
import (
    "sync"
)

2. DEFINIR STRUCT PromiseObject:

type PromiseObject struct {
    Ch       *ChanObject           // Channel interno para comunicação
    Resolved bool                  // Flag se já foi resolvido
    Result   map[string]Object     // Armazena resultado: {"status": "fulfilled/rejected", "value": ...}
    mutex    sync.Mutex            // Proteger acesso concorrente
}

3. IMPLEMENTAR INTERFACE Object:

func (p *PromiseObject) Type() ObjectType {
    return PROMISE_OBJ
}

func (p *PromiseObject) Inspect() string {
    p.mutex.Lock()
    defer p.mutex.Unlock()
    
    if !p.Resolved {
        return "Promise{<pending>}"
    }
    
    status := p.Result["status"].(*String).Value
    if status == "fulfilled" {
        return "Promise{<fulfilled>: " + p.Result["value"].Inspect() + "}"
    }
    return "Promise{<rejected>: " + p.Result["reason"].Inspect() + "}"
}

4. IMPLEMENTAR CallMethod():

func (p *PromiseObject) CallMethod(line string, scope *Scope, method string, args ...Object) Object {
    switch method {
    case "then":
        return p.promiseThen(line, scope, args)
    case "catch":
        return p.promiseCatch(line, scope, args)
    case "await":
        return p.promiseAwait(line, scope)
    default:
        return newError(line, "Promise has no method: " + method)
    }
}

VALIDAÇÃO:
- Código deve compilar
- Struct básica implementa interface Object
```

---
