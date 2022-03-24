package main

import (
	"fmt"
	"time"
)

func main() {
	canal, canal2 := writing("Ola Mundo"), writing("Ola Novo Channel")

	canalCompilado := multiplexar(canal, canal2)

	// canalCompilado := poderia ser assim também.

	for i := 0; i < 10; i++ {
		fmt.Println(<-canalCompilado)
	}
}

func multiplexar(canalEntrada1, canalEntrada2 <-chan string) <-chan string {
	canalDeSaida := make(chan string)

	go func() {
		for {
			select {
			case message := <-canalEntrada1:
				canalDeSaida <- message
			case message := <-canalEntrada2:
				canalDeSaida <- message
			}
		}
	}()

	return canalDeSaida
}

func writing(text string) <-chan string {
	channel := make(chan string)
	go func() {
		for {
			channel <- fmt.Sprintf("Sua Mensagem: %s", text)
			time.Sleep(time.Millisecond * 500)
		}

	}()
	return channel
}
