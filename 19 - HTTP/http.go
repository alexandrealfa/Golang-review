package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

/*
	função declarada para ser utilizada em uma route, onde por padrão recebe dois parametros,
	w http.ResponseWriter, dados da response, http.Request payload da requisição
*/

func home(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(200)
	_, err := w.Write([]byte("get router"))

	if err != nil {
		return
	}
}

func getContent() {
	req, err := http.Get("https://www.google.com")

	defer func(req *http.Response) {
		if err := req.Body.Close(); err != nil {
			log.Fatal(err)
		}
	}(req)

	if err != nil {
		log.Fatal(err)
	}

	res, err := io.ReadAll(req.Body)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(res))

}

func main() {
	getContent()
	/*
		http.HandleFunc é responsável por definir a uri e a função a ser executada nessa uri,
		onde a função tem 2 params obrigatórios para tratamento dos dados recebidos.
	*/

	http.HandleFunc("/home", home)

	/*
		log.Fatal(http.ListenAndServe(":port", null) é reponsável por gerar
		um server que atende pela porta definida, para receber requisições.
	*/
	log.Fatal(http.ListenAndServe(":5000", nil))
}
