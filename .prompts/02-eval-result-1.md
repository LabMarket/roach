## 🎯 PROMPT 2: IMPLEMENTAR eval/result.go (PARTE 1 - ESTRUTURA BÁSICA)

```
TAREFA: Criar o arquivo eval/result.go com a estrutura básica do ResultObject

REQUISITOS:

1. DEFINIR A STRUCT ResultObject:
   - Campo IsOk (bool): indica se é Ok(true) ou Err(false)
   - Campo Value (Object): valor quando IsOk=true
   - Campo Error (Object): mensagem de erro quando IsOk=false

2. IMPLEMENTAR INTERFACE Object:
   - Type() ObjectType: retornar RESULT_OBJ
   - Inspect() string: 
     * Se IsOk: "Ok(<valor>)"
     * Se !IsOk: "Err(<erro>)"

3. IMPLEMENTAR CallMethod(method string, args ...Object) Object:
   Despachar para métodos privados baseado em 'method':
   - "ok" -> resultOk()
   - "err" -> resultErr()
   - "get" -> resultGet()
   - "getOrElse" -> resultGetOrElse(args)
   - default: newError("Result has no method: " + method)

4. IMPLEMENTAR MÉTODOS PRIVADOS BÁSICOS:
   
   func (r *ResultObject) resultOk() Object {
       return &Boolean{Value: r.IsOk}
   }
   
   func (r *ResultObject) resultErr() Object {
       if !r.IsOk {
           return r.Error
       }
       return NIL
   }
   
   func (r *ResultObject) resultGet(line string) Object {
       if !r.IsOk {
           return newError(line, "Called get() on Err: " + r.Error.Inspect())
       }
       return r.Value
   }
   
   func (r *ResultObject) resultGetOrElse(line string, args []Object) Object {
       if len(args) != 1 {
           return newError(line, "getOrElse() requires 1 argument")
       }
       if r.IsOk {
           return r.Value
       }
       return args[0]
   }

PADRÃO DE REFERÊNCIA:
- Use eval/optional.go como template
- Siga as convenções de nomenclatura (métodos privados com prefixo do tipo)
- Use newError(line, message) para erros

VALIDAÇÃO:
- Código deve compilar sem erros
- Struct deve implementar completamente a interface Object
```

---