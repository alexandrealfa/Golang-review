package main

import (
	"context"
	"fmt"
	"time"
)

func Orchestrate(ctx context.Context) {
	// use select because context operates with channels, and with selects we can manipulate the operations of channels.
	select {
	case <-ctx.Done(): // in case of done, we show some information, but it's relative, we can call functions or do other things.
		fmt.Println("contexto expirado")
	case <-time.After(time.Second * 4): // in case of time exceed 3 seconds of the channel is active, print some message.
		fmt.Println("Transação realizada com sucesso!")
	}
}

func main() {
	ctx := context.Background()                            // create context background to initialize a new context in application
	ctx, cancel := context.WithTimeout(ctx, time.Second*3) // define a timeout for the context created

	defer cancel() // cancel at the end of operation

	Orchestrate(ctx) // change operations by context states
}
