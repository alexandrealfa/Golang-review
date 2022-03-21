package main

import (
	"fmt"
	"time"
)

/*
	Select
	é utilizado para resolver o problema de sincronização de execução em casos de diferentes tempos de recebimento de dados em
	um canal, em uma execução em que ele precise aguardar o recebimento para liberar o outro.
*/

func main() {
	channel1, channel2 := make(chan string), make(chan int)
	go func() {
		for {
			time.Sleep(time.Second * 2)
			channel1 <- "Save Channel 1"
		}
	}()

	go func() {
		for {
			time.Sleep(time.Millisecond * 500)
			channel2 <- 1
		}
	}()

	for {
		select {
		case messageChannel1 := <-channel1:
			fmt.Println(messageChannel1)
		case messageChannel2 := <-channel2:
			fmt.Println(messageChannel2)
		}
	}
}
