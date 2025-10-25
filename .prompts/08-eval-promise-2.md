## 🎯 PROMPT 8: IMPLEMENTAR eval/promise.go (PARTE 2 - CONSTRUTOR)

```
TAREFA: Criar função para construir novos Promises em eval/promise.go

CONTEXTO:
Em Roach, Promises são criados via: new Promise(fn(resolve, reject) { ... })
Precisamos interceptar isso no avaliador.

IMPLEMENTAÇÃO REQUERIDA:

1. CRIAR FUNÇÃO newPromise():

func newPromise(line string, scope *Scope, executor *Function) Object {
    promise := &PromiseObject{
        Ch:       &ChanObject{Chan: make(chan Object, 1)},
        Resolved: false,
        Result:   nil,
        mutex:    sync.Mutex{},
    }
    
    // Criar funções resolve e reject
    resolveFn := &Builtin{
        Fn: func(line string, scope *Scope, args ...Object) Object {
            if len(args) != 1 {
                return newError(line, "resolve() requires 1 argument")
            }
            
            // Enviar resultado fulfilled para o channel
            result := &Hash{
                Pairs: map[HashKey]HashPair{
                    (&String{Value: "status"}).HashKey(): {
                        Key:   &String{Value: "status"},
                        Value: &String{Value: "fulfilled"},
                    },
                    (&String{Value: "value"}).HashKey(): {
                        Key:   &String{Value: "value"},
                        Value: args[0],
                    },
                },
            }
            
            promise.Ch.Chan <- result
            return NIL
        },
    }
    
    rejectFn := &Builtin{
        Fn: func(line string, scope *Scope, args ...Object) Object {
            if len(args) != 1 {
                return newError(line, "reject() requires 1 argument")
            }
            
            // Enviar resultado rejected para o channel
            result := &Hash{
                Pairs: map[HashKey]HashPair{
                    (&String{Value: "status"}).HashKey(): {
                        Key:   &String{Value: "status"},
                        Value: &String{Value: "rejected"},
                    },
                    (&String{Value: "reason"}).HashKey(): {
                        Key:   &String{Value: "reason"},
                        Value: args[0],
                    },
                },
            }
            
            promise.Ch.Chan <- result
            return NIL
        },
    }
    
    // Executar o executor em goroutine
    go func() {
        // Chamar executor(resolve, reject)
        result := evalFunctionCall(line, scope, executor, []Object{resolveFn, rejectFn})
        
        // Se executor lançou exceção, rejeitar promise
        if isError(result) {
            rejectFn.Fn(line, scope, result)
        }
    }()
    
    return promise
}

NOTAS:
- Esta é uma versão simplificada. Você precisará ajustar:
  * Como Hash é construído no Roach (verifique eval/hash.go)
  * Como spawnar goroutines (pode ter função específica)
  * Como integrar com o sistema de spawn existente do Roach

VALIDAÇÃO:
- Código deve compilar
- Função newPromise pode ser chamada
```

---
