package main

import (
	"fmt"
	"github.com/alexandrealfa/math/Math"
)

func main() {
	subTotal := Math.Sum(19, 29)
	fmt.Println("new package ", subTotal)

	OP := Math.Operation{Type: "math"}
	value := 10

	OP.Add(value)
}
