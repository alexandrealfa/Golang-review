package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func getUrlByCep(cep string) string {
	return "http://viacep.com.br/ws/" + cep + "/json/"
}

type cepSchema struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Unidade     string `json:"unidade"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Estado      string `json:"estado"`
	Regiao      string `json:"regiao"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main() {
	client := http.Client{}

	for _, cep := range os.Args[1:] {
		url := getUrlByCep(cep)

		req, err := client.Get(url)

		if err != nil {
			log.Fatal(err.Error())
		}

		res, err := io.ReadAll(req.Body)

		if err != nil {
			log.Fatal(err)
		}

		var response cepSchema

		if err = json.Unmarshal(res, &response); err != nil {
			log.Fatal(err)
		}

		if err := req.Body.Close(); err != nil {
			log.Fatal(err)
		}

		fmt.Println(response.Localidade, "_", response.Estado)
	}
}
