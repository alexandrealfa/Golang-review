package main

import "fmt"

// Funções com valores nomeados, permitem nomear as variáveis de retorno e com isso,  basta apenas dar um retorn, que está utilizara dos valores passados como retorno
func operacoes(valor1 int, valor2 int) (subtracao int, soma int) {
	soma = valor1 + valor2
	subtracao = valor1 - valor2

	return
}

// if not numbers is passed, return the variabeles with zero value
func operacoes_zero_value() (subtracao int, soma int) {
	return
}

func main() {
	fmt.Println("Funções nomeadas: ")
	soma, sub := operacoes(5, 5)
	fmt.Println(soma, sub)
	fmt.Println(operacoes_zero_value())

}
