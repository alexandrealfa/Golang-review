package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

/*
 métodos, são ações de um struct, e outro tipo de dado em go, com ele é possível acessar os dados do mesmo,
 	e realizar modificações a partir do mesmo.
*/

func (u usuario) salvar() {
	fmt.Printf("salvando os dados do usuário %s no banco de dados.\n", u.nome)
}

/*
 no metódo, ao definir o * no usuario, é realizado a alteração no ponteiro,
  porém não é necessario utilizar do & para sinalizar que é o valor no ponteiro.
*/

func (u *usuario) fazerAniversario() {
	u.idade++
}

func (u usuario) maiorDeIdade() bool {
	if u.idade >= 18 {
		fmt.Printf("usuario %s é maior de idade\n", u.nome)

		return true
	}

	return false
}

func (u *usuario) atualizarIdade(novaIdade uint8) {
	u.idade = novaIdade
}

func main() {
	var newUser usuario = usuario{nome: "Alexandre", idade: 23}
	newUser.salvar()

	var newUser2 usuario = usuario{nome: "Carlos", idade: 23}
	newUser2.salvar()

	validarIdade := newUser.maiorDeIdade()
	fmt.Println(validarIdade)

	newUser2.fazerAniversario()

	fmt.Println(newUser.idade)

	fmt.Println(newUser2.idade)

	newUser2.atualizarIdade(45)

	fmt.Println("idade Usuario 2: ", newUser2.idade)
}
