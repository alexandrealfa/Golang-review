package main

import (
	"encoding/json"
	"fmt"
	"github.com/valyala/fastjson"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	var J fastjson.Parser
	jsonData := `{"user": {"name": "Alexandre", "age": 25}, "date": "202402", "books": ["learn Python", "learn GO"]}`
	value, err := J.Parse(jsonData)

	if err != nil {
		fmt.Println(err)
	}
	date := value.Get("date")
	fmt.Println(date)

	books := value.GetArray("books")

	fmt.Println(books)

	var user User

	userJson := value.GetObject("user").String()
	if err := json.Unmarshal([]byte(userJson), &user); err != nil {
		panic(err)
	}

	fmt.Println(user.Name)
}
