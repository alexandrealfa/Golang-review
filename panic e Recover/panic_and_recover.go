package main

import "fmt"

func calcMedia(n1, n2, validate float64) (result bool) {
	media := (n1 + n2) / 2

	result = media >= validate
	defer fmt.Println("loading ...")
	return
}

func main() {
	fmt.Println(calcMedia(10, 5, 7))
}
