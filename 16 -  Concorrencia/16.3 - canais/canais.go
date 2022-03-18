package main

import (
	"fmt"
	"time"
)

/*
	===========
	Channels:
	===========
	Trabalhar com channels é uma forma muito boa de lidar com goroutines.
	Quando trabalhamos com channels em go, estamos trabalhando com um canal para envio de dados, e outro para recebimento.
	Ao trabalhar com canais em golang é muito importante definir o encerramento do canal apos todos os envios serem realizados, e recebidos
	para evitar deadlocks.

	- sintaxe de criação de um canal:
		```canal := make(chan string) ``` usamos 'make', como palavra chave para criação de um canal, dentro dela passamos 'chan', para apontar
		que sera um canal, e na sequência tipamos o tipo de dado que será recebido 'string'.

	- sintaxe de recebimento de dados em um canal:
		```<-canal ```, com essa sintaxe definimos que o canal está recebendo informação, e podemos atribuir o valor recebido
		a uma variavel like ``` recebimento := <-canal``

	- sintaxe de envio de dados por um canal:
		```canal <- dado ```, com essa sintaxe definimos que o canal está enviando uma informação, e podemos definir o valor a
		enviado

	===========
	DEADLOCKS:
	===========
	As deadlocks ocorrem quando um canal continua aberto, mesmo após todos os itens serem enviados para este canal.
	As deadlocks não são capturadas em compilamento, apenas em execução, então é de extrema necessidade tratar todos
    os channels, para evitar este problema.
*/

func main() {
	canal := make(chan string)
	go escrever("goroutine 1", canal)
	/*
		for {
			messages, open := <-canal

			if !open {
				break
			}

			fmt.Println(messages)
		}
	*/

	for message := range canal {
		fmt.Println(message)
	}
}

func escrever(texto string, canal chan string) {
	for i := 0; i < 5; i++ {
		canal <- texto
		time.Sleep(time.Second)
	}
	close(canal)
}
