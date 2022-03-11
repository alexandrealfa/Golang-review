package main

import (
	"errors"
	"fmt"
)

func main() {
	var numero int64 = -1000000000
	fmt.Println(numero)

	//uint segue a mesma convensão do int podendo ter o uint8, uint16, uint 32, uint64, porém sem sinal
	var numero2 uint = 1000
	fmt.Println(numero2)

	// Alias:

	// INT32 = RUNE
	var numero3 rune = 123456
	fmt.Println(numero3)

	//BYTE = uint8
	var numero4 byte = 123
	fmt.Println(numero4)

	// Números Reais

	var numeroReal1 float32 = 123.23
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 1231242141425.42
	fmt.Println(numeroReal2)

	// Strings
	var str string = "Texto"
	fmt.Println(str)

	str2 := "Texto 2"
	fmt.Println(str2)

	// Boolean
	var boleano1 bool = true
	fmt.Println(boleano1)

	var boleano2 bool
	fmt.Println(boleano2)

	// Error

	var erro error
	fmt.Println(erro)

	// nesse caso eu declarei uma variável do tipo erro, e atribui um erro a ele com o pacote errors que é um pacote nativo da linguagem go
	var erro2 error = errors.New("Erro Interno")
	fmt.Println(erro2)

	// A Variável declarada com '' aspas simples irá mostrar o número daquele caracter na tabela asq.
	camelo := '%'
	fmt.Println(camelo)

	// Valor Zero
	// para string o valor zero é vazio
	// para int o valor zero é 0
	// para o nulo o valor zero é nil
	// para bool o valor zero é false
}
