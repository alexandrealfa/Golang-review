package main

import "fmt"

// Anonimal function

func main() {

	func() {
		fmt.Println("Ola Mundo!")
	}()

	valor_recebido := func(text string) string {
		return fmt.Sprintf("Recebido -> %s", text)
	}("imprima isso")

	fmt.Println(valor_recebido)
}
