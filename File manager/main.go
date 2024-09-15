package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	var filename string

	filename = "document.csv"

	file, err := os.Create(filename)

	if err != nil {
		log.Fatal(err.Error())
	}

	defer func(file *os.File) {
		if err := file.Close(); err != nil {
			log.Fatal(err)
		}
		if err = os.Remove(file.Name()); err != nil {
			log.Fatal(err)
		}
	}(file)

	content := []string{"Alexandre ", "Alfa"}

	for _, text := range content {
		_, err = file.WriteString(text)

		if err != nil {
			log.Fatal("error on save file", err.Error())
		}
	}

	read, err := os.ReadFile(filename)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(read))

	// leitura com buffer

	contentFile, err := os.Open("documentWithLines.txt")

	if err != nil {
		log.Fatal(err.Error())
	}

	reader := bufio.NewReader(contentFile)
	buffer := make([]byte, 20)
	acc, words := "", 0

	for {
		_, err := reader.Read(buffer)

		if err != nil {
			break
		}

		if phrase := strings.Split(string(buffer), " "); len(phrase) > 0 {
			words += len(phrase)
		}

		acc += string(buffer)
	}

	fmt.Println(acc)
	fmt.Println("total de palavras : ", words)
}
