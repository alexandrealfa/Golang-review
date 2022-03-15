package main

import (
	"fmt"
	"math"
)

/*
 - Interfaces só aceita assinaturas de métodos.
 - Integração das interfaces é feita de forma implicita.
*/

/*
 qualquer struct que eu criar, que contenha o mesmo metodo definido na interface,
  juntamente dos seu métodos e retornos, poderá receber essa interface como parametro
*/

type form interface {
	area() float64
}

type rectangle struct {
	height float64
	width  float64
}

type circle struct {
	raio float64
}

// Criando um método chamado "area" para o struct rectangle, e com isso habilito essa struct para a implementação da interface form.
func (r rectangle) area() float64 {
	return r.height * r.width
}

// Criando um método chamado "circle" para o struct rectangle, e com isso habilito essa struct para a implementação da interface form.
func (c circle) area() float64 {
	return math.Pi * math.Pow(c.raio, 2)
}

// aqui estou criando uma função que recebe uma interface do tipo form que pode ser implementada por varios tipos diferentes
func writeArea(f form) {
	fmt.Printf("A área da forma é %0.2f\n", f.area())
}

func main() {
	r := rectangle{10, 15}
	writeArea(r)

	c := circle{10}
	writeArea(c)
}
