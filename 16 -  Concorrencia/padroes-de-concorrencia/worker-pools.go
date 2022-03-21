package main

import (
	"fmt"
)

func main() {
	tarefas, resultados := make(chan int, 45), make(chan int, 45)

	go worker(tarefas, resultados, "worker-1: ")
	go worker(tarefas, resultados, "worker-2: ")
	go worker(tarefas, resultados, "worker-3: ")
	go worker(tarefas, resultados, "worker-4: ")

	for value := 0; value < 45; value++ {
		tarefas <- value
	}

	close(tarefas)

	for i := 0; i < 45; i++ {
		select {
		case resultado := <-resultados:
			fmt.Println(resultado)
		}
	}

}

func worker(tarefas <-chan int, resultados chan<- int, workName string) {
	for numero := range tarefas {
		fmt.Println(workName, numero)
		resultados <- fibonacci(numero)
	}
}

func fibonacci(posicao int) int {
	if posicao <= 1 {
		return posicao
	}

	return fibonacci(posicao-1) + fibonacci(posicao-2)
}
