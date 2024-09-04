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

func allowAccess(access *bool) {
	*access = true
}

func changeAccess(access bool) {
	access = !access

	fmt.Println("chamada de dentro da função changeAccess ", access)
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
	fmt.Println(currentNumber) // -20]

	// exemplo mais simples da mudança do valor da variável.

	access := false // variável inicializada com o valor false

	changeAccess(access) // mudança do valor da variável dentro da função, porém, como a função não acessa o ponteiro ...
	// ... memória em que a variável recebida por parametro está alocada, logo, a variável não sofre alteração fora do sistema.

	fmt.Println("chamada da váriavel após a execução da função changeAccess", access) // a váriavel access permanece com o valor false

	allowAccess(&access) // Essa função por sua vez recebe o ponteiro da memória, onde a partir disso ela viabiliza a mudança do valor global da variável.

	fmt.Println(access) // agora a variável access contem o valor true mesmo no escopo global da função.

}
