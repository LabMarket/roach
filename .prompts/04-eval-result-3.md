## 🎯 PROMPT 4: IMPLEMENTAR eval/result.go (PARTE 3 - MATCH)

```
TAREFA: Adicionar método match() ao ResultObject em eval/result.go

CONTEXTO:
match() é pattern matching: recebe duas funções (uma para Ok, outra para Err)
e executa a apropriada baseado no estado do Result

IMPLEMENTAÇÃO REQUERIDA:

1. ADICIONAR AO CallMethod():
   case "match": return r.resultMatch(line, scope, args)

2. IMPLEMENTAR resultMatch():

func (r *ResultObject) resultMatch(line string, scope *Scope, args []Object) Object {
    if len(args) != 2 {
        return newError(line, "match() requires exactly 2 arguments (okFn, errFn)")
    }
    
    okFn, ok1 := args[0].(*Function)
    errFn, ok2 := args[1].(*Function)
    
    if !ok1 || !ok2 {
        return newError(line, "match() arguments must be functions")
    }
    
    if r.IsOk {
        // Chamar função de sucesso com o valor
        return evalFunctionCall(line, scope, okFn, []Object{r.Value})
    } else {
        // Chamar função de erro com o erro
        return evalFunctionCall(line, scope, errFn, []Object{r.Error})
    }
}

VALIDAÇÃO:
- Código deve compilar
- Teste "Result: match pattern with Ok" deve passar
- Teste "Result: match pattern with Err" deve passar
```

---
