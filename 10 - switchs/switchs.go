package main

import "fmt"

func dia_da_semana(numero int) string {
	switch numero {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda-Feira"
	case 3:
		return "Terça-Feira"
	case 4:
		return "Quarta-Feira"
	case 5:
		return "Quinta-Feira"
	case 6:
		return "Sexta-Feira"
	case 7:
		return "Sabado"
	default:
		return "Número Inválido"
	}
}

func dia_da_semana_2(numero int) string {
	var dia_selecionado string

	switch {
	case numero == 1:
		dia_selecionado = "Domingo"
	case numero == 2:
		dia_selecionado = "Segunda-Feira"
	case numero == 3:
		dia_selecionado = "Teça-Feira"
	case numero == 4:
		dia_selecionado = "Quarta-Feira"
	case numero == 5:
		dia_selecionado = "Quinta-Feira"
	case numero == 6:
		dia_selecionado = "Sexta-Feira"
	case numero == 7:
		dia_selecionado = "Sabado"
	default:
		dia_selecionado = "Invalid Number"
	}

	return dia_selecionado
}

func main() {
	fmt.Println("switchs")
	dia_selecionado := dia_da_semana(2)
	fmt.Println(dia_selecionado)

	new_dia := dia_da_semana_2(1)
	fmt.Println(new_dia)
}
