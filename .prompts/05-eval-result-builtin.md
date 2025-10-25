## 🎯 PROMPT 5: REGISTRAR RESULT EM eval/builtin.go

```
TAREFA: Adicionar funções globais Ok() e Err() em eval/builtin.go

LOCALIZAÇÃO:
- Procure por 'builtins := map[string]*Builtin{'
- Adicione as novas entradas seguindo o padrão existente

IMPLEMENTAÇÃO REQUERIDA:

1. ADICIONAR Ok():

name("Ok"): &Builtin{
    Fn: func(line string, scope *Scope, args ...Object) Object {
        if len(args) != 1 {
            return newError(line, "Ok() requires exactly 1 argument")
        }
        return &ResultObject{
            IsOk:  true,
            Value: args[0],
            Error: NIL,
        }
    },
},

2. ADICIONAR Err():

name("Err"): &Builtin{
    Fn: func(line string, scope *Scope, args ...Object) Object {
        if len(args) != 1 {
            return newError(line, "Err() requires exactly 1 argument")
        }
        return &ResultObject{
            IsOk:  false,
            Value: NIL,
            Error: args[0],
        }
    },
},

NOTAS:
- Verifique a assinatura exata de Builtin.Fn no código (pode ter variações)
- Certifique-se de que NIL está disponível (constante global)
- Use name() se for uma função helper existente, senão use string direta

VALIDAÇÃO:
- Código deve compilar
- Testes básicos de Result devem passar:
  * "Result: Create Ok value"
  * "Result: Create Err value"
```

---
