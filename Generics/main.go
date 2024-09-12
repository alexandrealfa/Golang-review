package main

import (
	"fmt"
)

func squareNumber[T ~int | float64](number T) T {
	return number * number
}

type Number interface {
	int | float64
}

type customtype int

func processNumbers[T Number](numbers map[string]T) T {
	var acc T

	for _, number := range numbers {
		acc += number
	}

	return acc
}

func compare[T comparable](number T, numberToCompare T) bool {
	return number == numberToCompare
}

func main() {
	floatNumber := 10.10
	fmt.Println(squareNumber(floatNumber))

	intNumber := 3
	fmt.Println(squareNumber(intNumber))

	numbers := map[string]int{"alana": 39, "beatriz": 40, "carla": 50}
	floatNumbers := map[string]float64{"alan": 29.00, "bruno": 30.02, "carlos": 70.3}

	fmt.Println(processNumbers(numbers))
	fmt.Println(processNumbers(floatNumbers))

	var customNumber customtype = 5
	var currentNumber int = 10

	fmt.Println("customNumber", squareNumber(customNumber))
	fmt.Println(squareNumber(currentNumber))

	compare(intNumber, 10)
	compare(10.10, floatNumber)
}
