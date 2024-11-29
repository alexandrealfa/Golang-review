package main

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"
)

func getURL() (url string) {
	url = "https://economia.awesomeapi.com.br/json/last/USD-BRL"
	return
}

type Price struct {
	url string
}

type serializedJson struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type ResponseSchema struct {
	USDBRL struct {
		Code       string `json:"code"`
		Codein     string `json:"codein"`
		Name       string `json:"name"`
		High       string `json:"high"`
		Low        string `json:"low"`
		VarBid     string `json:"varBid"`
		PctChange  string `json:"pctChange"`
		Bid        string `json:"bid"`
		Ask        string `json:"ask"`
		Timestamp  string `json:"timestamp"`
		CreateDate string `json:"create_date"`
	} `json:"USDBRL"`
}

func (p Price) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
	price, err := p.getPrice()
	if err != nil {
		log.Fatal("aquii", err.Error())
	}
	response := serializedJson{"Dólar", price.USDBRL.Bid}

	responseJson, marshalError := json.Marshal(response)
	if marshalError != nil {
		log.Fatal("error to Marshal response: ", marshalError.Error())
	}

	w.WriteHeader(http.StatusOK)

	if _, err := w.Write(responseJson); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Fatal("error to write response: ", err.Error())
	}
}

func (p Price) getPrice() (ResponseSchema, error) {
	var data ResponseSchema
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*200)
	defer cancel()

	requestCtx, err := http.NewRequestWithContext(ctx, http.MethodGet, p.url, nil)
	if err != nil {
		log.Println("Error in request process: ", err.Error())

		return data, err
	}

	req, err := http.DefaultClient.Do(requestCtx)
	if err != nil {
		log.Println("Error to get request data: ", err.Error())
	}

	res, parseError := io.ReadAll(req.Body)

	if parseError != nil {
		log.Println("error to read Json: ", res)
		return data, parseError
	}

	if MarshalError := json.Unmarshal(res, &data); MarshalError != nil {
		log.Println("error in Json parser process: ", MarshalError.Error())
		return data, MarshalError
	}

	return data, err
}

func (p Price) logData() {
	return
}

func main() {
	mutex := http.NewServeMux()
	mutex.Handle("/cotacao", Price{getURL()})

	if err := http.ListenAndServe(":8080", mutex); err != nil {
		log.Fatal("error to run server: ", err.Error())
	}
}
