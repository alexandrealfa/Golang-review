package main

import "fmt"

type Account struct {
	bankBalance float64
}

func (ac *Account) pay(value float64) {
	ac.bankBalance -= value

	fmt.Println(ac.bankBalance)
}

type AccountNew struct {
	bankBalance float64
}

func (ac AccountNew) simulate(value float64) {
	ac.bankBalance -= value

	fmt.Println(ac.bankBalance)
}

type Client struct {
	Name   string
	typeOf string
}

func generateClient(name string) *Client {
	return &Client{name, "default"}
}

func main() {

	accNew := AccountNew{300}

	accNew.simulate(200)            // 100
	fmt.Println(accNew.bankBalance) // 300

	acc := Account{1000}

	acc.pay(200)                 // 800
	fmt.Println(acc.bankBalance) // 800
	alfa := generateClient("alexandre alfa")

	fmt.Println(alfa.Name)

	geo := generateClient("geo hotz")

	fmt.Println(geo.Name)
}
