package main

import "fmt"

type user struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	rua    string
	numero uint16
}

type userCompleted struct {
	user
	endereco
}

func main() {
	var usuario user
	usuario.nome = "alexandre"
	usuario.idade = 22
	fmt.Println(usuario)

	enderecoNovo := endereco{"Rua Turquesa", 655}

	usuario2 := user{"Carvalho", 22, enderecoNovo}
	fmt.Println(usuario2)

	usuario3 := user{nome: "Alexandre Alfa"}
	usuario3.idade = 22
	fmt.Println(usuario3)

	new_user := userCompleted{
		user:     usuario3,
		endereco: enderecoNovo,
	}

	fmt.Println(new_user.rua)

}
