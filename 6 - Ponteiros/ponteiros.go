package main

import "fmt"

func main() {
	fmt.Println("Ponteiro")

	var variavel int = 10
	var variavel2 int = variavel

	fmt.Println(variavel, variavel2)

	variavel += 3

	println(variavel, variavel2)

	// Ponteiro é uma referência de memória.
	var variavel3 int
	var ponteiro *int

	fmt.Println(variavel3, ponteiro)

	variavel3 = 100
	ponteiro = &variavel3

	fmt.Println(variavel3, ponteiro)

	//desreferenciação
	fmt.Println(variavel3, *ponteiro)

	variavel3 = 150

	fmt.Println(variavel3, *ponteiro)

}
