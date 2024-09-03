package main

import (
    "fmt"
    "testing"
)

// Funções da Calculadora

// Adiciona retorna a soma de dois números
func Adiciona(a, b float64) float64 {
    return a + b
}

// Subtrai retorna a subtração de dois números
func Subtrai(a, b float64) float64 {
    return a - b
}

// Multiplica retorna a multiplicação de dois números
func Multiplica(a, b float64) float64 {
    return a * b
}

// Divide retorna a divisão de dois números. Retorna um erro se o divisor for zero.
func Divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("não é possível dividir por zero")
    }
    return a / b, nil
}

// Funções de Teste

func TestAdiciona(t *testing.T) {
    resultado := Adiciona(1, 2)
    esperado := 3.0
    if resultado != esperado {
        t.Errorf("Adiciona(1, 2) = %f; esperado %f", resultado, esperado)
    }
}

func TestSubtrai(t *testing.T) {
    resultado := Subtrai(5, 3)
    esperado := 2.0
    if resultado != esperado {
        t.Errorf("Subtrai(5, 3) = %f; esperado %f", resultado, esperado)
    }
}

func TestMultiplica(t *testing.T) {
    resultado := Multiplica(4, 3)
    esperado := 12.0
    if resultado != esperado {
        t.Errorf("Multiplica(4, 3) = %f; esperado %f", resultado, esperado)
    }
}

func TestDivide(t *testing.T) {
    resultado, err := Divide(10, 2)
    if err != nil {
        t.Fatalf("Divide(10, 2) retornou erro: %v", err)
    }
    esperado := 5.0
    if resultado != esperado {
        t.Errorf("Divide(10, 2) = %f; esperado %f", resultado, esperado)
    }
}

func TestDividePorZero(t *testing.T) {
    _, err := Divide(10, 0)
    if err == nil {
        t.Errorf("Divide(10, 0) deveria retornar um erro")
    }
}

func main() {
    // Executar os testes
    testing.MainStart(testing.NewFakeRunner(), []testing.InternalTest{
        {"TestAdiciona", TestAdiciona},
        {"TestSubtrai", TestSubtrai},
        {"TestMultiplica", TestMultiplica},
        {"TestDivide", TestDivide},
        {"TestDividePorZero", TestDividePorZero},
    }, nil, nil).Run()
}