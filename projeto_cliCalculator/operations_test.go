package main

import (
	"testing"
)

func TestSum(t *testing.T) {
	result := Sum(2, 4)
	expected := 6.0

	if result != expected {
		t.Errorf("Expected %.2f but got %.2f", expected, result)
	}

}

func TestMinus(t *testing.T) {
	result := Minus(10, 4)
	expected := 6.0

	if result != expected {
		t.Errorf("Expected %.2f but got %.2f", expected, result)
	}

}

func TestMultiply(t *testing.T) {
	result := Multiply(10, 4)
	expected := 40.0

	if result != expected {
		t.Errorf("Expected %.2f but got %.2f", expected, result)
	}
}

func TestDivision(t *testing.T) {
	result, err := Division(10, 5)
	expected := 2.0

	if err != nil{
		t.Errorf("Error: %v ", err)
	}


	if result != expected {
		t.Errorf("Expected %.2f but got %.2f", expected, result)
	}
}

func TestEvaluateExpression(t *testing.T) {
	expr := []string{"2", "+", "3", "*", "4"}
	result, err := EvaluateExpression(expr)

	if err != nil {
		t.Errorf("Erro inesperado: %v", err)
	}

	expected := 20.0 // Avaliação da esquerda: (2 + 3) * 4 = 5 * 4 = 20
	if result != expected {
		t.Errorf("Esperado %.2f, mas retornou %.2f", expected, result)
	}
}