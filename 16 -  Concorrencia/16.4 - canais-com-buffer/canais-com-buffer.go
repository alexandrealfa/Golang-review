package main

import "fmt"

func main() {
	// Canal com buffer, permite o envio e recebimento de dados de um canal em uma única function.
	canal := make(chan string, 1)

	canal <- "salve"

	message := <-canal
	fmt.Println(message)
}
