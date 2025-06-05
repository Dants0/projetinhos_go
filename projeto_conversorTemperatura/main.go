package main

import (
	"fmt"
	"strconv"
	"github.com/inancgumus/screen"
)

func main() {

	screen.Clear()


	for {
		screen.MoveTopLeft()

		fmt.Println("\nEscolha uma opção: ")
		fmt.Println("Celsius to Fahrenheit = 1")
		fmt.Println("Celsius to Kelvin = 2")
		fmt.Println("Fahrenheit to Celsius = 3")
		fmt.Println("Fahrenheit to Kelvin = 4")
		fmt.Println("Kelvin to Celsius = 5")
		fmt.Println("Kelvin to Fahrenheit = 6")

		fmt.Println("0. Sair")

		opcao := ""

		fmt.Scanln(&opcao)

		if opcao == "0" {
			fmt.Println("Saindo...")
			break
		}

		valor := ""
		fmt.Println("Digite o valor a ser convertido: ")
		fmt.Scanln(&valor)

		valueToConvert, err := strconv.ParseFloat(valor, 64)

		if err != nil {
			fmt.Println("Error quando converteu o valor.")
			break
		}

		switch opcao {
		case "1":
			result, err := CelsiusToFahrenheit(valueToConvert)

			if err != nil {
				fmt.Println("Error")
			}

			fmt.Printf("Result: %.2f °F", result)

		default:
			fmt.Println("Default case! ", err)
		}

	}
}
