package main

/* uma inteface generica
É necessário ter cuidado com interfaces genericas, pois seu comportamento pode gerar "gambiarras"
no seu códigos, visto que golang é uma linguagem fortemente tipada, oque gera uma segunança maior
no seu código.
a ideia de uma interface generica é para situações que necessitam de uma dinamicidade,
onde a resultante, e os parametros não podem ser estabelecidos de acordo com uma
tipagem fixa, fora essas situações o uso de interfaces genericas não é recomendado.
*/

import "fmt"

func generica(interf interface{}) {
	fmt.Println(interf)
}

func main() {
	generica("salve")
	generica(1)
	generica(false)

	// exemplo de coisa pra não se fazer em go com generics.
	mapa := map[interface{}]interface{}{
		1:            "qqcoisa",
		float32(100): true,
		true:         float64(12),
	}
	fmt.Println(mapa)
}
