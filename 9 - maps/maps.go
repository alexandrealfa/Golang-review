package main

import "fmt"

func main() {
	fmt.Printf("Maps\n")

	usuario := map[string]string{
		"name": "Alexandre Alfa",
	}

	fmt.Println(usuario["name"])

	// Maps aninhados

	user := map[string]map[string]string{
		"hobbies": {
			"primario":   "Estudar",
			"secundario": "Academia",
		},
		"formacao": {
			"curso": "Gestão da tecnologia da Informação",
		},
	}

	user["signo"] = map[string]string{
		"nome": "Capricornio",
	}

	fmt.Println(user["hobbies"]["primario"])

	delete(user, "formacao")

	fmt.Println(user)
}
