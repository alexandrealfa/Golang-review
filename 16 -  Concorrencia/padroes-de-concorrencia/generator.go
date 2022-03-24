package main

import (
	"fmt"
	"time"
)

/*
	O padrão Generators: permite o encapsulamento de um goroutine em uma função, onde é possível
	retornar um canal com uma série de operações que são retornadas no mesmo.
*/

func main() {
	canal := escrever("Ola Mundo")

	for i := 0; i < 10; i++ {
		fmt.Println(<-canal)
	}
}

// Basicamente encapsula uma go routine e retorna um canal com o valor passado.
func escrever(text string) <-chan string {
	canal := make(chan string)

	go func() {
		for {
			canal <- fmt.Sprintf("Valor recebido: %s", text)
			time.Sleep(time.Millisecond * 500)
		}
	}()

	return canal
}
