package main

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type responseJson struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

func URL() (url string) {
	url = "http://localhost:8080/cotacao"
	return
}

func createFileIfNotExist(filename string) (file *os.File, err error) {
	file, err = os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, os.ModeAppend)

	if errors.Is(err, os.ErrNotExist) {
		file, err = os.Create(filename)
	}

	return
}

func save(content string) {
	var filename = "cotacao.txt"

	file, err := createFileIfNotExist(filename)
	if err != nil {
		log.Fatal("error in create file: ", err.Error())
	}
	defer func(file *os.File) {
		if err := file.Close(); err != nil {
			log.Fatal("error closing file: ", err.Error())
		}
	}(file)

	if _, err := file.WriteString(content); err != nil {
		log.Fatal("error to write file", err.Error())
	}
}

func main() {
	client := http.DefaultClient

	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*300)
	defer cancel()

	req, reqError := http.NewRequestWithContext(ctx, http.MethodGet, URL(), nil)
	if reqError != nil {
		log.Fatal(reqError)
	}

	res, clientError := client.Do(req)
	if clientError != nil {
		log.Fatal(clientError)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal("Error to read response body: ", err)
	}

	var responseJson responseJson

	if err := json.Unmarshal(data, &responseJson); err != nil {
		log.Fatal(err)
	}

	save(responseJson.Name + ": " + responseJson.Value + "\n")
}
