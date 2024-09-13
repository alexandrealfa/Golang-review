package Math

import (
	"fmt"
	"os"
)

type Number interface {
	float64 | int
}

type Operation struct {
	Type      string
	arguments any
	Total     int
}

func (O Operation) Add(value int) {
	O.Total += value
	O.GetTotal()
}

func (O Operation) GetTotal() {
	fmt.Println("toma: ", O.Total)
}

func Sum[T Number](value T, valueToSum T) T {
	operation := value + valueToSum

	LogOperation(operation)

	return operation
}

func LogOperation[T any](value T) {
	info := os.FileInfo.Name
	fmt.Println(info)
	fmt.Println("my custom log off Operation: ", value)
}
