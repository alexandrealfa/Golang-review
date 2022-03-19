package main

import "fmt"

func main() {
	// Canal com buffer, permite o envio e recebimento em uma única function.
	canal := make(chan string, 1)

	canal <- "salve"

	message := <-canal
	fmt.Println(message)
}
