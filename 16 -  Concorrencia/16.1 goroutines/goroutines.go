package main

import (
	"fmt"
	"time"
)

/*
	Goroutines permitem a execução assincrona, onde não limita o processamento a uma única tareffa por execução.
	é necessário mapear todas as execuções assincronas para que não tenha problema com execuções na aplicação, pois
	goroutines encadeadas, sem execução posterior, pode gerar um comportamento não esperado.
*/

func main() {
	go escrever("async text") // Goroutine
	escrever("Ola Mundo")
}

func escrever(text string) {
	for {
		fmt.Println(text)
		time.Sleep(time.Second)
	}
}
