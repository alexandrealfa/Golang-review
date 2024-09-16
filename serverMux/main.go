package main

import (
	"fmt"
	"net/http"
)

func homeFunc(w http.ResponseWriter, r *http.Request) {
	if infraKey := r.Header.Get("x-infra-key"); infraKey == "3012" {
		fmt.Println(infraKey)
	}

	if _, err := w.Write([]byte("Welcome!")); err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
}

type Book struct {
	Name      string `json:"name,omitempty"`
	Pages     int    `json:"pages,omitempty"`
	PagesRead int    `json:"pages_read,omitempty"`
}

func (b Book) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Println(req)
	if _, err := w.Write([]byte(b.Name)); err != nil {
		fmt.Println("error ao tentar pegar os livros do user")
	}
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", homeFunc)
	mux.Handle("/books", Book{"Python for data Analisys", 634, 234})

	if err := http.ListenAndServe(":5000", mux); err != nil {
		return
	}
}
