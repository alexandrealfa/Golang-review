package main

import (
	"encoding/json"
	"fmt"
	cepService "github.com/alexandrealfa/cepsearch/cep"
	"log"
	"net/http"
)

func getCep(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		fmt.Println("right route")
	}

	cep := r.URL.Query().Get("cep")

	if cep == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("product-signature", "my-product")

	w.WriteHeader(http.StatusOK)

	cepStruct, err := cepService.GetCepByParam(cep)
	if err != nil {
		_, err = w.Write([]byte("cep Não Encontrado"))
		if err != nil {
			return
		}
	} else {
		if err = json.NewEncoder(w).Encode(cepStruct); err != nil {
			fmt.Println(err)
		}

		fmt.Println(cepStruct)
	}
}

func main() {
	http.HandleFunc("/", getCep)

	if err := http.ListenAndServe(":5000", nil); err != nil {
		log.Fatal(err)
	}

}
