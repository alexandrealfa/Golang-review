package main

import "fmt"

func fibonacci(posicao uint) uint {
	if posicao <= 1 {
		return posicao
	}

	return fibonacci(posicao-1) + fibonacci(posicao-2)
}

func main() {
	position := uint(16)

	for i := uint(0); i < position; i++ {
		fmt.Println(fibonacci(i))
	}

}
