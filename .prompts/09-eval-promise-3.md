## 🎯 PROMPT 9: IMPLEMENTAR eval/promise.go (PARTE 3 - AWAIT)

```
TAREFA: Implementar método await() em PromiseObject

CONTEXTO:
await() bloqueia até a promise resolver, então retorna o valor ou lança exceção

IMPLEMENTAÇÃO REQUERIDA:

func (p *PromiseObject) promiseAwait(line string, scope *Scope) Object {
    p.mutex.Lock()
    
    // Se já resolvido, retornar resultado cacheado
    if p.Resolved {
        p.mutex.Unlock()
        status := p.Result["status"].(*String).Value
        
        if status == "fulfilled" {
            return p.Result["value"]
        } else {
            // Lançar erro
            return newError(line, p.Result["reason"].Inspect())
        }
    }
    
    p.mutex.Unlock()
    
    // Bloquear esperando resultado do channel
    result := <-p.Ch.Chan
    
    p.mutex.Lock()
    p.Resolved = true
    
    // result é um Hash com {"status": ..., "value"/"reason": ...}
    resultHash, ok := result.(*Hash)
    if !ok {
        p.mutex.Unlock()
        return newError(line, "Invalid promise result")
    }
    
    // Converter Hash para map[string]Object para acesso mais fácil
    p.Result = make(map[string]Object)
    for _, pair := range resultHash.Pairs {
        key := pair.Key.(*String).Value
        p.Result[key] = pair.Value
    }
    
    status := p.Result["status"].(*String).Value
    p.mutex.Unlock()
    
    if status == "fulfilled" {
        return p.Result["value"]
    } else {
        return newError(line, p.Result["reason"].Inspect())
    }
}

NOTAS:
- Você precisará ajustar como acessar pares do Hash (verifique Hash.Pairs)
- Proteção de mutex é crítica para concorrência
- Certifique-se de unlock antes de returns

VALIDAÇÃO:
- Código deve compilar
- Testes básicos de Promise devem passar:
  * "Promise: Resolve with value"
  * "Promise: Reject with error"
```

---
