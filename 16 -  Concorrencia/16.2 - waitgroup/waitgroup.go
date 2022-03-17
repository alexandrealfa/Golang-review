package main

import (
	"fmt"
	"sync"
	"time"
)

// waitGroup é uma forma de sincronizar as goroutines, para que ambas sejam executadas e somente após isso, o programa seja encerrado
func main() {
	var waitGroup sync.WaitGroup

	waitGroup.Add(2) // Criando a fila de execução e sinalizando o número de goroutines que precisam ser executadas.

	go func() {
		escrever("func 1")
		waitGroup.Done() // Sinaliza a finalização da primeira goroutine
	}()

	go func() {
		escrever("func 2")
		waitGroup.Done() // Sinaliza a finalização da segunda goroutine
	}()

	waitGroup.Wait() // Faz o programa aguardar que ambas as rotinas sejam finalizadas.
}

func escrever(text string) {
	for i := 0; i < 5; i++ {
		fmt.Println(text)
		time.Sleep(time.Second)
	}
}
