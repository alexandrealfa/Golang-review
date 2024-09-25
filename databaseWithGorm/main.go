package main

import (
	"gorm.io/gorm"
)

type Products struct {
	Id    string
	Name  string
	Price float64
	gorm.Model
}

func main() {
	return
}
