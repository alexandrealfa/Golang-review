package main

import (
	"fmt"
)

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func calculosMatematicos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	subtracao := n1 - n2
	return soma, subtracao
}

func printProperties(a string, b int) {
	fmt.Println(a)
	fmt.Println(b)
}

func myStatements() (string, int) {
	return "salve", 10
}
func main() {
	soma := somar(2, 4)
	fmt.Println(soma)

	var f = func(txt string) string {
		fmt.Println(txt)
		return txt
	}

	result := f("Texto da função 1")
	fmt.Println(result)

	resultadoSoma, resultadoSubtracao := calculosMatematicos(10, 14)
	fmt.Println(resultadoSoma, resultadoSubtracao)

	// neste caso a funcao me retorna 2 valores porém o underline oculta o valor selecionado.
	resultadoSoma2, _ := calculosMatematicos(10, 14)
	fmt.Println(resultadoSoma2)

	// neste caso a função recebe suas propriedades dinamicamente a partir do retorno de outra função.
	printProperties(myStatements())
}
