package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type Category struct {
	Id   int `gorm:"primaryKey"`
	Name string
}

type Products struct {
	Id         int `gorm:"primaryKey"`
	Name       string
	Price      float64
	CategoryId int
	Category   Category
	gorm.Model
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goreview?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	if err := db.AutoMigrate(&Products{}, &Category{}); err != nil {
		log.Fatal(err)
	}

	category := Category{
		Name: "Acessórios",
	}

	db.Create(&category)

	db.Create(&Products{
		Name:       "Camisa",
		Price:      50,
		CategoryId: category.Id,
	})

	var products []Products

	db.Preload("Category").Find(&products)

	for _, product := range products {
		fmt.Println(product.Id, product.Category.Name)
	}
}
