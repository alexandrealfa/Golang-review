package main

import "fmt"

/*
Função init é executada antes da função main. É possivel ter uma função init por arquivo.
as alterações realizadas nas variaveis globais pela função init, são extraidas e refletidas na função main
*/

var n int

func init() {
	fmt.Println("executando da função init ...")
	n = 10
}

func main() {
	fmt.Println("executando da função main ...")
	fmt.Println("variavel definida no escopo global, e atribuida na função init, e esta sendo exibida na função main", n)
}
