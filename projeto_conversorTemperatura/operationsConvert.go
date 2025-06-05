package main

import (
	"fmt"
)

func CelsiusToFahrenheit(number float64) (float64, error) {

	if number == 0 {
		return 0, fmt.Errorf("Valor inválido para conversão!")
	}

	fahrenheit := (number * 9 / 5) + 32

	return fahrenheit, nil

}
