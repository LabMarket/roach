## 📋 CONTEXTO GERAL

### Objetivos
Você vai implementar dois novos tipos (Result e Promise) no interpretador Roach, 
uma linguagem de programação escrita em Go.

DOCUMENTAÇÃO DE REFERÊNCIA:
1. Arquitetura do projeto: https://cdn.jsdelivr.net/gh/LabMarket/roach@feat/results/PROJECT.md
2. Documentação da linguagem: https://cdn.jsdelivr.net/gh/LabMarket/roach@feat/results/docs/README.md
3. Testes de validação: [script de validação em *Roach*](../examples/result_promisse.roach)

CONCEITOS:
- Result: Tipo que encapsula sucesso (Ok) ou erro (Err) com informação contextual
- Promise: Tipo que encapsula operações assíncronas com then/catch/await

ARQUIVOS A SEREM CRIADOS:
- eval/result.go
- eval/promise.go

ARQUIVOS A SEREM MODIFICADOS:
- eval/object.go (adicionar constantes de tipo)
- eval/builtin.go (registrar funções globais Ok, Err, resolved, rejected)

PADRÃO A SEGUIR:
- Examine eval/optional.go como referência de implementação
- Use eval/channel.go como base para Promise (usa channels internamente)

---


### Asynchronous Programming with Promise

Roach supports Promises for async operations:

```swift
promise = new Promise(fn(resolve, reject) {
    // async work here
    resolve(result)
})

promise
    .then(fn(data) { processData(data) })
    .catch(fn(err) { handleError(err) })
```

---
