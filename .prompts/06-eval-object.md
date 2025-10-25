## 🎯 PROMPT 6: ADICIONAR CONSTANTES EM eval/object.go

```
TAREFA: Adicionar constantes de tipo para Result e Promise em eval/object.go

LOCALIZAÇÃO:
Procure por:
const (
    INTEGER_OBJ = "INTEGER"
    FLOAT_OBJ = "FLOAT"
    ...
)

IMPLEMENTAÇÃO REQUERIDA:

Adicionar ao bloco const:
    RESULT_OBJ  = "RESULT"
    PROMISE_OBJ = "PROMISE"

VALIDAÇÃO:
- Código deve compilar
- ResultObject.Type() deve retornar RESULT_OBJ corretamente
```

---
