package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"log"
)

type Product struct {
	Id    string
	Name  string
	Price float64
}

func newProduct(name string, price float64) *Product {
	return &Product{
		uuid.New().String(),
		name,
		price,
	}
}

func insertInTable(db *sql.DB, tableName string, product *Product) error {
	statement, err := db.Prepare("insert into " + tableName + "(id, name, price) values(?, ?, ?)")

	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(product.Id, product.Name, product.Price); err != nil {
		return err
	}

	return err
}

func updateProduct(db *sql.DB, product *Product) error {
	stmt, err := db.Prepare("update products set name = ?, price = ? where id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err = stmt.Exec(product.Name, product.Price, product.Id); err != nil {
		return err
	}
	return err
}

func selectProduct(db *sql.DB, id string) (*Product, error) {
	stmt, err := db.Prepare("select name, price from products where id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var p Product

	row := stmt.QueryRow(id)
	if err := row.Scan(&p.Name, &p.Price); err != nil {
		return nil, err
	}
	return &p, err
}

func selectAllProducts(db *sql.DB) ([]Product, error) {
	rows, err := db.Query("select id, name, price from products")
	if err != nil {
		return nil, err
	}
	var products []Product

	for rows.Next() {
		var p Product
		if err = rows.Scan(&p.Id, &p.Name, &p.Price); err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	return products, err
}

func deleteProduct(db *sql.DB, id string) error {
	stmt, err := db.Prepare("delete from products where id = ?")
	if err != nil {
		return err
	}
	if _, err = stmt.Exec(id); err != nil {
		return err
	}

	return nil
}

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/goreview")

	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	product := newProduct("Camisa", 240.0)

	err = insertInTable(db, "products", product)
	if err != nil {
		log.Fatal(err)
	}
	product.Price = 150.0
	if err = updateProduct(db, product); err != nil {
		log.Fatal(err)
	}
	res, err := selectProduct(db, product.Id)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("O produto %s esta custando %.2f", res.Name, res.Price)

	products, err := selectAllProducts(db)
	if err != nil {
		log.Fatal(err)
	}
	for _, product := range products {
		fmt.Println(product)
	}

	fmt.Println("removing :", product.Id)
	if err = deleteProduct(db, product.Id); err != nil {
		log.Fatal(err)
	}
	return
}
