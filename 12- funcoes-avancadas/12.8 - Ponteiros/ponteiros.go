package main

import "fmt"

// função que recebe um valor e retorna esse valor com o sinal invertido
func invertSignal(value int) int {
	return value * -1
}

// função que recebe um ponteiro de memoria, e altera o valor contido nesse ponteiro de memória para um valor com sinal invertido.
func invertSignalWithPointer(numero *int) {
	*numero = *numero * -1
}

func main() {

	currentNumber := 20
	invertedNumber := invertSignal(currentNumber)

	// sem afetar o ponteiro do parametro recebido.

	fmt.Println(invertedNumber) // -20
	fmt.Println(currentNumber)  // 20

	/*
		chamando a função que altera o valor do parametro recebido
		- para acessar o endereço de memoria de uma variavel basta por um & comercial na frente
	*/

	invertSignalWithPointer(&currentNumber)
	fmt.Println(currentNumber) // -20

}
