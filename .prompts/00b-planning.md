# Prompts Estruturados para Agente de Codificação

## 📊 RESUMO EXECUTIVO PARA O AGENTE

ORDEM DE EXECUÇÃO DOS PROMPTS:

FASE 1 - ANÁLISE (Prompts 1):
→ Entender a estrutura existente
* [Análise](01-analise.md)

FASE 2 - RESULT (Prompts 2-6):
→ Implementar Result completo
→ Registrar em builtins
→ Adicionar constantes
* [Eval / Result - Parte 1](02-eval-result-1.md)
* [Eval / Result - Parte 2](03-eval-result-2.md)
* [Eval / Result - Parte 3](04-eval-result-3.md)
* [Eval / Result - BuiltIn](05-eval-result-builtin.md)
* [Eval / Result - Constants](06-eval-object.md)

FASE 3 - PROMISE (Prompts 7-12):
→ Implementar Promise completo
→ Registrar em builtins
→ Integrar com 'new'
* [Eval / Promise - Parte 1](07-eval-promise-1.md)
* [Eval / Promise - Parte 2](08-eval-promise-2.md)
* [Eval / Promise - Parte 3](09-eval-promise-3.md)
* [Eval / Promise - Parte 4](10-eval-promise-4.md)
* [Eval / Promise - BuiltIn](11-eval-promise-builtin.md)
* [Eval / Promise - New](12-promise-new.md)

FASE 4 - VALIDAÇÃO (Prompts 13-17):
→ Rodar testes
→ Corrigir bugs
→ Testar integrações
→ Casos edge
* [Testes](13-test-validation.md)
* [Documentação](14-doc-examples.md)
* [Otimização](15-optimization.md)
* [Integração](16-integration.md)
* [Casos limites](17-edge-cases.md)

FASE 5 - QUALIDADE (Prompts 18-19):
→ Testes unitários Go
→ Benchmarks
→ Otimizações
* [Unit Tests](18-unit-tests.md)
* [Benchmarks](19-benchmarks.md)

FASE 6 - RELEASE (Prompt 20):
→ Documentação final
→ Exemplos
→ Changelog
→ Ready for merge
* [Documentação](20-docs.md)

ARQUIVOS ESPERADOS NO FIM:
```
├── eval/
│   ├── result.go         ← NOVO
│   ├── promise.go        ← NOVO
│   ├── result_test.go    ← NOVO
│   ├── promise_test.go   ← NOVO
│   ├── result_promise_bench_test.go  ← NOVO
│   ├── object.go         ← MODIFICADO
│   ├── builtin.go        ← MODIFICADO
│   └── eval.go           ← MODIFICADO (integração new Promise)
├── examples/
│   ├── result_basic.roach
│   ├── result_advanced.roach
│   ├── promise_basic.roach
│   ├── promise_advanced.roach
│   └── integration_tests.roach
├── docs/
│   └── RESULT_PROMISE.md
├── result_promise_tests.roach  (arquivo de teste fornecido)
└── CHANGELOG.md          ← ATUALIZADO

```

---

## 🚀 COMANDOS RÁPIDOS PARA O AGENTE

```bash
# Compilar
go build -o roach

# Rodar testes Roach
./roach result_promise_tests.roach

# Rodar testes Go
go test ./eval -v

# Rodar testes com race detection
go test -race ./eval

# Rodar benchmarks
go test -bench=. -benchmem ./eval

# Verificar cobertura
go test -cover ./eval

# Lint
golint ./eval

# Vet
go vet ./eval
```
