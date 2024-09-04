package main

import "fmt"

// Função com valores propriedades variaveis
func soma(numeros ...int) int {

	total := 0

	for _, value := range numeros {
		total += value
	}

	return total
}

// Função com propriedades variaveis, e fixas
func escrever_algo(text string, numbers ...int) int {
	fmt.Println(text, numbers)
	fmt.Printf("%T", numbers)

	return 0
}

func main() {
	fmt.Println(soma(1, 2, 3, 4, 5, 6, 7))
	escrever_algo("propriedades fixas e variaveis", 1, 3, 5, 7)
}
