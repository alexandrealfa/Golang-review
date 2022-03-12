package main

import "fmt"

/* Clausulas Panic, interrompem a execução do programa, diferente de um error que alerta que a aplicação teve um erro mas continua em execução
uma clausula Panic irá apontar que a aplicação entrou em "panico" e não conseguirá executar nada em diante, porém antes de entrar em panic
as aplicações conseguem executar as clausulas defer definidas antes da execução de panic.

Para sair de um Panic, é necessário definir uma clausula recovery.

visto isso, é interessante sempre que definir uma claúsula panic, definir anteriormente uma cláusula recover assinada com um defer antés da Panic.

O Panic não é uma solução interessante para tratar erros, ele é utilizado em situações especificas, em que seja interessante interromper
as execuções do programa caso determinada situação ocorra.
*/

func recoverExec() {
	if r := recover(); r != nil {
		fmt.Println("Execução recuperada com sucesso!.")
	}
}

func calcMedia(n1, n2, validate float64) (result bool) {
	// defer recoverExec()
	defer fmt.Println("Loading...")

	// É possivel definir funcoes recover em funções anonimas.
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Execução recuperada com sucesso!.")
		}
	}()

	media := (n1 + n2) / 2

	result = media >= validate && media != 6

	if media == 6 {
		panic("A média é exatamente 6!.")
	}

	return
}

func main() {
	fmt.Println(calcMedia(6, 6, 6))
}
