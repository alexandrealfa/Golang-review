package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

/* Com a notação `json:"key"` é indicado o nome desse campo ao ser convertido em um json, ou seja,
a key Name, ao ser convertida em Json será chamada, "nome".
*/
type person struct {
	Name string `json:"nome"`
	Age  uint   `json:"idade"`
}

func main() {
	marshalExample()
	unmarshalExample()
}

func marshalExample() {
	pessoa1 := person{"Alexandre", 23}

	/* json.Marshal é o método utilizado para converter um struct em um formato JSON, esse método retorna dois valores,
	um o valor da conversão e o outro o valor para erro */
	pessoaJSON, err := json.Marshal(pessoa1)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(pessoaJSON)
	fmt.Println(bytes.NewBuffer(pessoaJSON))

	// Exemplo de conversão de map para o Formato JSON

	pessoa2 := map[string]string{
		"name":      "Alexandre",
		"languages": "pt",
	}

	pessoa2JSON, err := json.Marshal(pessoa2)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(pessoa2JSON)
	fmt.Println(bytes.NewBuffer(pessoa2JSON))
}

func unmarshalExample() {

	pessoaJSON := `{"nome": "Alexandre", "idade": 23}`
	var p person

	/* json.Unmarshal é o método utilizado para converter um JSON em um formato desejado, esse método recebe 2 parametros
	o json em bytes, e o ponteiro da váriavel a ser atribuida, e com isso ele retorna apenas um valor para erro,
	visto que o valor da trasformação já estará na atribuição realizada pelo ponteiro recebido como segundo parametro.
	*/

	if err := json.Unmarshal([]byte(pessoaJSON), &p); err != nil {
		log.Fatal(err)
	}

	fmt.Println(p)

	pessoa2JSON := `{"nome": "Alexandre", "idade": "23"}`
	pessoa2 := make(map[string]string)
	if err := json.Unmarshal([]byte(pessoa2JSON), &pessoa2); err != nil {
		log.Fatal(err)
	}

	fmt.Println(pessoa2)
}
