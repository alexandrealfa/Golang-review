package main

import "fmt"

func main() {
	slice3 := make([]float32, 10, 11)
	slice3 = append(slice3, 5)
	fmt.Println("-------------")
	fmt.Println(slice3)
	fmt.Println(len(slice3)) //length
	fmt.Println(cap(slice3)) // capacidade

	slice3 = append(slice3, 5)
	fmt.Println("-------------")
	fmt.Println(slice3)
	fmt.Println(len(slice3)) //length
	fmt.Println(cap(slice3)) // capacidade

}
