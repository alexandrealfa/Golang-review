package main

import (
	"fmt"
	"net/http"
)

func main() {
	filePath := http.Dir("./public")
	fileServer := http.FileServer(filePath)

	fmt.Println(filePath)

	mux := http.NewServeMux()
	mux.Handle("/", fileServer)

	if err := http.ListenAndServe(":5000", mux); err != nil {
		fmt.Println(err)
	}
}
