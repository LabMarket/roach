## 🎯 PROMPT 3: IMPLEMENTAR eval/result.go (PARTE 2 - MAP E FLATMAP)

```
TAREFA: Adicionar métodos map e flatMap ao ResultObject em eval/result.go

CONTEXTO:
map() transforma o valor se Ok, propaga erro se Err
flatMap() encadeia operações que retornam Result

IMPLEMENTAÇÃO REQUERIDA:

1. ADICIONAR AO CallMethod():
   case "map": return r.resultMap(line, scope, args)
   case "flatMap": return r.resultFlatMap(line, scope, args)

2. IMPLEMENTAR resultMap():

func (r *ResultObject) resultMap(line string, scope *Scope, args []Object) Object {
    if len(args) != 1 {
        return newError(line, "map() requires exactly 1 argument (function)")
    }
    
    fn, ok := args[0].(*Function)
    if !ok {
        return newError(line, "map() argument must be a function")
    }
    
    if !r.IsOk {
        // Propagar erro sem modificar
        return r
    }
    
    // Aplicar função ao valor
    result := evalFunctionCall(line, scope, fn, []Object{r.Value})
    
    // Checar se resultado é um erro
    if isError(result) {
        return &ResultObject{IsOk: false, Value: NIL, Error: result}
    }
    
    return &ResultObject{IsOk: true, Value: result, Error: NIL}
}

3. IMPLEMENTAR resultFlatMap():

func (r *ResultObject) resultFlatMap(line string, scope *Scope, args []Object) Object {
    if len(args) != 1 {
        return newError(line, "flatMap() requires exactly 1 argument (function)")
    }
    
    fn, ok := args[0].(*Function)
    if !ok {
        return newError(line, "flatMap() argument must be a function")
    }
    
    if !r.IsOk {
        // Propagar erro sem modificar
        return r
    }
    
    // Aplicar função ao valor
    result := evalFunctionCall(line, scope, fn, []Object{r.Value})
    
    // Resultado DEVE ser um ResultObject
    resultObj, ok := result.(*ResultObject)
    if !ok {
        return newError(line, "flatMap() function must return a Result")
    }
    
    return resultObj
}

NOTAS IMPORTANTES:
- Você precisará encontrar a função correta para chamar funções no Roach
  (pode ser Apply, Call, ou algo similar - verifique como Optional faz)
- Scope precisa ser passado para avaliação de funções
- Use isError() para verificar se resultado é erro
- flatMap espera que a função retorne um ResultObject

VALIDAÇÃO:
- Código deve compilar
- Testes de map e flatMap em result_promise_tests.roach devem passar
```

---