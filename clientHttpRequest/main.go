package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	// Defino o Client isoladamente
	client := http.Client{Timeout: time.Minute * 2}

	body := bytes.NewBuffer([]byte(`{"name": "alfa"}`))

	// neste caso é onde monto o conteúdo da request isoladamente.
	req, err := http.NewRequest(http.MethodPost, "https://google.com", body)

	if err != nil {
		panic(err)
	}

	// o Client realiza a chamada da requisição.
	res, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	read, err := io.ReadAll(res.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(read))
	//io.CopyBuffer(os.Stdout, res.Body, nil)

}
