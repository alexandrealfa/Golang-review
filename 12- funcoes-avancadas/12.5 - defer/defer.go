package main

import "fmt"

// Defer Posterga a execução do item setado até antes do return do bloco de código em que está inserido.

func visualization() {
	fmt.Println("Visualizing first value")
}

func visualization2() {
	fmt.Println("Visualizing second value")
}

func calc_media(nota1 int, nota2 int) bool {
	defer fmt.Println("Média Cauculada, Resultado será retornado!.")
	fmt.Println("Entrando na Função Para Verificação se o aluno está sendo aprovado")

	media := (nota1 + nota2) / 2

	if media >= 6 {
		return true
	}

	return false
}

func main() {
	defer visualization()
	visualization2()

	fmt.Println(calc_media(5, 5))

}
