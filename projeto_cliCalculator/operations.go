package main

import (
	"fmt"
	"strconv"
)

func Sum(num1 float64, num2 float64) float64 {
	return num1 + num2
}

func Minus(num1 float64, num2 float64) float64 {
	return num1 - num2
}

func Multiply(num1 float64, num2 float64) float64 {
	return num1 * num2
}

func Division(num1 float64, num2 float64) (float64, error) {

	if num2 == 0 {
		return 0, fmt.Errorf("Expressão inválida")
	}

	return num1 / num2, nil
}

func EvaluateExpression(tokens []string) (float64, error) {
	if len(tokens)%2 == 0 {
		return 0, fmt.Errorf("Expressão inválida")
	}

	result, err := strconv.ParseFloat(tokens[0], 64)
	if err != nil {
		return 0, fmt.Errorf("Erro ao converter: %v", tokens[0])
	}

	for i := 1; i < len(tokens); i += 2 {
		op := tokens[i]
		num, err := strconv.ParseFloat(tokens[i+1], 64)

		if err != nil {
			return 0, fmt.Errorf("Erro ao converter: %v", tokens[i+1])
		}

		switch op {
		case "+":
			result = Sum(result, num)

		case "-":
			result = Minus(result, num)

		case "*":
			result = Multiply(result, num)

		case "/":

			res, err := Division(result, num)
			if err != nil {
				return 0, err
			}
			result = res

		default:
			return 0, fmt.Errorf("Operador inválido: %s", op)
		}
	}

	return result, nil
}
