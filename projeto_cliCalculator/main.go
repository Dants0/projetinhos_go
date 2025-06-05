package main

import (
	"fmt"
	"os"
)

func main() {

	args := os.Args[1:]

	if len(args) < 3 {
		fmt.Println("Uso: go run main.go <expr> Ex: 2 + 3 * 4")
		return
	}

	result, err := EvaluateExpression(args)
	if err != nil {
		fmt.Println("Erro:", err)
		return
	}

	fmt.Printf("Resultado: %.2f\n", result)

}
