package main

import (
	"fmt"
)

func main() {
	// for with global scope increment
	// i := 0

	// for i < 10 {
	// 	fmt.Println("Incrementando I")
	// 	time.Sleep(time.Second)
	// 	i++
	// }
	// fmt.Println(i)

	//For init with local scope increment
	for j := 0; j < 10; j++ {
		fmt.Println("Incrementando J", j)
	}

	//For in arrays
	nomes := [4]string{"Alexandre", "Milena", "Moises", "Arthur"}

	for index, value := range nomes {
		fmt.Println(index, ":", value)
	}

	// For in Strings (If use range in one string, you will get ascII value off character, for usage you need convert using string())

	for _, letra := range "Palavra" {
		fmt.Println("asci_value:", letra, "converted_value:", string(letra))
	}

	// For in Map

	usuario := map[string]string{
		"name": "Alexandre",
		"age":  "22",
	}

	for key, value := range usuario {
		fmt.Println(key, value)
	}

	// Não é possível usar range in struct

}
