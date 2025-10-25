## 🎯 PROMPT 11: REGISTRAR PROMISE EM eval/builtin.go

```
TAREFA: Adicionar funções globais resolved() e rejected() em eval/builtin.go

CONTEXTO:
Helpers para criar Promises já resolvidas/rejeitadas

IMPLEMENTAÇÃO REQUERIDA:

1. ADICIONAR resolved():

name("resolved"): &Builtin{
    Fn: func(line string, scope *Scope, args ...Object) Object {
        if len(args) != 1 {
            return newError(line, "resolved() requires exactly 1 argument")
        }
        
        // Criar promise já resolvida
        promise := &PromiseObject{
            Ch:       &ChanObject{Chan: make(chan Object, 1)},
            Resolved: true,
            Result: map[string]Object{
                "status": &String{Value: "fulfilled"},
                "value":  args[0],
            },
            mutex: sync.Mutex{},
        }
        
        return promise
    },
},

2. ADICIONAR rejected():

name("rejected"): &Builtin{
    Fn: func(line string, scope *Scope, args ...Object) Object {
        if len(args) != 1 {
            return newError(line, "rejected() requires exactly 1 argument")
        }
        
        // Criar promise já rejeitada
        promise := &PromiseObject{
            Ch:       &ChanObject{Chan: make(chan Object, 1)},
            Resolved: true,
            Result: map[string]Object{
                "status": &String{Value: "rejected"},
                "reason": args[0],
            },
            mutex: sync.Mutex{},
        }
        
        return promise
    },
},

VALIDAÇÃO:
- Código deve compilar
- Testes de helpers devem passar:
  * "Promise: resolved helper"
  * "Promise: rejected helper"
```

---
