package main

import (
	"fmt"
	"linha-de-comando/app"
	"log"
	"os"
)

func main() {
	fmt.Println("ponto de Partida")
	aplicacao := app.Generate()

	if err := aplicacao.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
