## 🎯 PROMPT 10: IMPLEMENTAR eval/promise.go (PARTE 4 - THEN E CATCH)

```
TAREFA: Implementar métodos then() e catch() em PromiseObject

CONTEXTO:
then(onFulfilled, onRejected) encadeia promises
catch(onRejected) é açúcar sintático para then(nil, onRejected)

IMPLEMENTAÇÃO REQUERIDA:

1. IMPLEMENTAR promiseCatch():

func (p *PromiseObject) promiseCatch(line string, scope *Scope, args []Object) Object {
    if len(args) != 1 {
        return newError(line, "catch() requires exactly 1 argument (function)")
    }
    
    // catch(fn) é equivalente a then(nil, fn)
    return p.promiseThen(line, scope, []Object{NIL, args[0]})
}

2. IMPLEMENTAR promiseThen():

func (p *PromiseObject) promiseThen(line string, scope *Scope, args []Object) Object {
    if len(args) != 2 {
        return newError(line, "then() requires exactly 2 arguments (onFulfilled, onRejected)")
    }
    
    onFulfilled := args[0]
    onRejected := args[1]
    
    // Validar que são funções ou NIL
    if onFulfilled.Type() != FUNCTION_OBJ && onFulfilled.Type() != NIL_OBJ {
        return newError(line, "then() first argument must be a function or nil")
    }
    if onRejected.Type() != FUNCTION_OBJ && onRejected.Type() != NIL_OBJ {
        return newError(line, "then() second argument must be a function or nil")
    }
    
    // Criar nova Promise encadeada
    return newPromise(line, scope, &Function{
        Body: nil, // Será executado no callback abaixo
        Fn: func(line string, scope *Scope, resolve, reject *Builtin) Object {
            // Esperar a promise atual resolver
            result := p.promiseAwait(line, scope)
            
            // Checar se foi fulfilled ou rejected
            if isError(result) {
                // Promise foi rejected
                if onRejected.Type() == FUNCTION_OBJ {
                    // Chamar handler de erro
                    handlerResult := evalFunctionCall(line, scope, onRejected.(*Function), []Object{result})
                    
                    if isError(handlerResult) {
                        reject.Fn(line, scope, handlerResult)
                    } else {
                        resolve.Fn(line, scope, handlerResult)
                    }
                } else {
                    // Propagar erro
                    reject.Fn(line, scope, result)
                }
            } else {
                // Promise foi fulfilled
                if onFulfilled.Type() == FUNCTION_OBJ {
                    // Chamar handler de sucesso
                    handlerResult := evalFunctionCall(line, scope, onFulfilled.(*Function), []Object{result})
                    
                    if isError(handlerResult) {
                        reject.Fn(line, scope, handlerResult)
                    } else {
                        // Se resultado é uma Promise, encadear
                        if handlerResult.Type() == PROMISE_OBJ {
                            nestedResult := handlerResult.(*PromiseObject).promiseAwait(line, scope)
                            if isError(nestedResult) {
                                reject.Fn(line, scope, nestedResult)
                            } else {
                                resolve.Fn(line, scope, nestedResult)
                            }
                        } else {
                            resolve.Fn(line, scope, handlerResult)
                        }
                    }
                } else {
                    // Propagar valor
                    resolve.Fn(line, scope, result)
                }
            }
            
            return NIL
        },
    })
}

NOTAS IMPORTANTES:
- Esta implementação é conceitual. Você precisará ajustar:
  * Como Function é estruturado no Roach
  * A assinatura correta de Function.Fn
  * Como spawn goroutines (dentro de newPromise)
  * Como detectar erros (isError função)

- A lógica é complexa porque then() precisa:
  1. Esperar a promise atual
  2. Executar o handler apropriado
  3. Retornar uma nova promise
  4. Encadear se o handler retornar outra promise

VALIDAÇÃO:
- Código deve compilar
- Testes de encadeamento devem passar:
  * "Promise: then chains success"
  * "Promise: catch handles rejection"
```

---
