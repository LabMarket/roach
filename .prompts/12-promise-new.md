## 🎯 PROMPT 12: INTEGRAR PROMISE COM CONSTRUTOR 'new'

```
TAREFA: Modificar eval/eval.go para interceptar 'new Promise(...)'

CONTEXTO:
Promises são criados via: new Promise(fn(resolve, reject) { ... })
Precisamos detectar isso no evaluator e chamar newPromise()

LOCALIZAÇÃO:
Procure pela avaliação de 'new' expressions (provavelmente NewExpression ou similar)

IMPLEMENTAÇÃO REQUERIDA:

No handler de 'new' expressions, adicionar caso especial:

// Dentro de evalNewExpression() ou similar
if className == "Promise" {
    if len(args) != 1 {
        return newError(line, "Promise constructor requires 1 argument (executor function)")
    }
    
    executor, ok := args[0].(*Function)
    if !ok {
        return newError(line, "Promise constructor argument must be a function")
    }
    
    return newPromise(line, scope, executor)
}

NOTAS:
- Você precisará encontrar onde 'new' é processado
- Verifique como outras classes built-in (File, Optional, etc) são tratadas
- Pode ser que precise registrar Promise como classe especial

VALIDAÇÃO:
- Construtor de Promise funciona
- new Promise(...) cria PromiseObject corretamente
```

---
