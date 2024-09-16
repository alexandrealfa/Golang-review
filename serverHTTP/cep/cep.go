package cep

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func getUrlByCep(cep string) string {
	return "https://viacep.com.br/ws/" + cep + "/json/"
}

type CepSchema struct {
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

// GetCepByParam receive a string param to do a request to viacep site, and process body/*
func GetCepByParam(cep string) (*CepSchema, error) {
	var Response CepSchema

	url := getUrlByCep(cep)

	req, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	defer func(req *http.Response) {
		err := req.Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(req)

	res, err := io.ReadAll(req.Body)

	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(res, &Response); err != nil {
		return nil, err
	}

	return &Response, err
}
