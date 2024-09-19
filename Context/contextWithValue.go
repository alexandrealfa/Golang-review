package main

import (
	"context"
	"fmt"
)

func HandleContext(ctx context.Context) {
	if value := ctx.Value("pass"); value == "123Secret" {
		fmt.Println("Access Allowed!")
	} else {
		fmt.Println("Access Denied!")
	}

}
func main() {
	/*
		context.WithValue permite que seja setado um valor ao contexto e com isso o mesmo pode realizar tratamentos
		a pertir disso.

		existe uma convensão a respeito dos contextos, ao serem utilizados como parametros de uma função, os mesmos
		devem ser passados por primeiro.
	*/

	ctx := context.WithValue(context.Background(), "pass", "123Secret")
	HandleContext(ctx)
}
