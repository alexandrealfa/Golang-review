package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	/* quando a propriedade timeout é setada ao declarar o cliente, indica o tempo que a
	requisição pode ficar viva, onde, a partir do momento que esgotou o tempo, a request devolve
	Client.Timeout exceeded while awaiting headers
	*/
	client := http.Client{Timeout: time.Second}
	url := "https://google.com"
	res, err := client.Get(url)

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))

	postBody := bytes.NewBuffer([]byte(`{"name": "alexandre"}`))
	res, err = client.Post(url, "application/json", postBody)

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	converted, err := io.ReadAll(res.Body)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(converted))

	io.CopyBuffer(os.Stdout, res.Body, nil)

}
