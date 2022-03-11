package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Arrays e Slices")

	// Para declarar um array é necessário definir o tamanho e o tipo do mesmo, conforme exemplo abaixo:
	var array1 [5]int

	// para atribuir um valor ao array é necessário acessar a posição do mesmo e a partir disso atribuir o novo valor:
	array1[0] = 1
	fmt.Println(array1)

	//Para declarar um array com inferência de tipo também é necessário definir um tamanho máximo para o mesmo, e os valores iniciais do mesmo:
	array2 := [5]string{"Posição1", "Posição2"}
	fmt.Println(array2)

	//Caso nenhum valor sejá inserido na criação de um array por inferência de tipo, o mesmo inicia todas as posições definidas, com o valor zero do tipo sinalizado.
	array4 := [4]string{}
	fmt.Println(array4)

	//Utilizando o [...] você cria um array com posições relativas a quantidade de itens. Não é possível criar um array com posições relativas, nesse caso usamos o slice.
	array3 := [...]int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(array3)

	// -- Slice --
	// Slice diferente de um array em go, é um array interno com ponteiro de memória, que permite a inserção de valores , e implementação relativa do mesmo.

	//Declarando um slice vazio:
	var slice0 = []string{}
	fmt.Println(slice0)

	//Declarando um slice com inferência de tipo e valores iniciais
	slice := []int{10, 11, 12, 13, 14, 15, 16, 17}
	fmt.Println(slice)
	fmt.Println(reflect.TypeOf(slice), reflect.TypeOf(array3))

	//O Metodo append permite a criação e inserção em uma nova posição no slice definido:
	slice = append(slice, 18)
	fmt.Println(slice)

	//Declarando um slice apontando ao endereço de memória de algumas posições de um array atraves da inferência de Tipo.
	slice2 := array2[0:3]
	//Redefinindo um valor tanto no slice, quando na posição da mémoria na qual o slice aponta.
	slice2[2] = "Fon"
	fmt.Println(slice2)

	array2[1] = "Posição Alterada"
	fmt.Println(slice2)
	fmt.Println(array2)

}
