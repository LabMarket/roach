## 🎯 PROMPT 1: ANÁLISE INICIAL

TAREFA: Analisar a estrutura existente do Roach para entender o padrão de implementação

ARQUIVOS PARA EXAMINAR:
1. [eval/optional.go](../eval/optional.go)
2. [eval/object.go](../eval/object.go)
3. [eval/builtin.go](../eval/builtin.go)
4. [eval/channel.go](../eval/channel.go)

QUESTÕES A RESPONDER:
1. Como Optional implementa a interface Object?
2. Quais métodos Object requer (Type(), Inspect(), CallMethod())?
3. Como Optional.CallMethod() despacha para métodos específicos (isPresent, get, etc)?
4. Como builtins são registrados em builtin.go (estrutura do map, assinatura de função)?
5. Como ChanObject usa channels do Go internamente?
6. Qual o padrão de tratamento de erros usado (newError())?

OUTPUT ESPERADO:
- Documento resumindo os padrões encontrados
- Assinaturas de funções-chave que precisaremos imitar
- Lista de constantes ObjectType que precisam ser adicionadas


---