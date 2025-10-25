## 🎯 PROMPT 16: INTEGRAÇÃO COM FEATURES EXISTENTES

```
TAREFA: Garantir que Result e Promise funcionam com features existentes do Roach

FEATURES PARA TESTAR:

1. LINQ MODULE:
   - Result em collections do LINQ
   - Promise como fonte de dados
   
   Exemplo:
   results = [Ok(1), Err("failed"), Ok(3)]
   successes = linq.from(results).where(fn(r) { r.ok() }).toSlice()

2. TRY-CATCH-FINALLY:
   - Result dentro de try-catch
   - Promise.await() em try-catch
   
   Exemplo:
   try {
       value = somePromise.await()
       println(value)
   } catch e {
       println("Error:", e)
   }

3. PATTERN MATCHING (case-in):
   - Match em Result.ok()
   
   Exemplo:
   result = safeDivide(10, 0)
   case result.ok() in {
       true  { println("Success:", result.get()) }
       false { println("Error:", result.err()) }
   }

4. COMPREHENSIONS:
   - Result em list comprehensions
   
   Exemplo:
   values = [r.get() for r in results where r.ok()]

5. OPERATOR OVERLOADING:
   - Considerar se Result/Promise devem ter operadores

6. SERIALIZATION (JSON):
   - Como Result é serializado?
   - Como Promise é serializado (pode não fazer sentido)?

TESTES DE INTEGRAÇÃO:

Criar arquivo integration_tests.roach com casos reais:

// LINQ + Result
fn test_linq_with_result() {
    results = [
        Ok(10),
        Err("error1"),
        Ok(20),
        Err("error2"),
        Ok(30)
    ]
    
    values = linq.from(results)
        .where(fn(r) { r.ok() })
        .select(fn(r) { r.get() })
        .toSlice()
    
    println("Extracted values:", values) // [10, 20, 30]
}

// Promise + try-catch
fn test_promise_with_error_handling() {
    try {
        result = fetchData().await()
        println("Got data:", result)
    } catch e {
        println("Failed to fetch:", e)
    } finally {
        println("Cleanup done")
    }
}

// Result + pattern matching
fn test_result_with_pattern_matching() {
    result = parseConfig(data)
    
    case result.ok() in {
        true  { 
            config = result.get()
            println("Valid config:", config) 
        }
        false { 
            println("Invalid config:", result.err())
        }
    }
}

VALIDAÇÃO:
□ Todas as integrações funcionam
□ Sem conflitos com features existentes
□ Comportamento é intuitivo
```

---
