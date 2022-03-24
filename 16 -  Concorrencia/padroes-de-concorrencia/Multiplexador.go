package main

import (
	"fmt"
	"time"
)

/*
	Multiplexador: recebe N canais por parametros e retorna um único canal que filtra os dados recebidos
	por ambos os canais, e compila tudo em um único canal de retorno, é excelente para tratamento de varios canais e
	agrupando o retorno de multiplos canais tratados em um único canal.
*/

func main() {
	canal, canal2 := writing("Ola Mundo"), writing("Ola Novo Channel")

	canalCompilado := multiplexar(canal, canal2) // canalCompilado := poderia ser assim também.

	for i := 0; i < 10; i++ {
		fmt.Println(<-canalCompilado)
	}
}

// Recebe dois canas e retornar um único canal reponsável por tratar o recebimento dos canais recebidos.

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
