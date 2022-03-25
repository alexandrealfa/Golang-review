package enderecos

import "testing"

/*
 => Unit Tests:
 O nome da função de test é declarada com a inicial Test e na sequência o nome da função que será testada, em camelCase.
 A função de test recebe o ponteiro de testing, que é importado do repositório de funções go.
*/

func TestTipoDeEndereco(t *testing.T) {
	addressToTest := "Avenida Paulista"

	// Definindo o tipo de endereco esperado
	addressTypeExpected := "avenida"

	// Definindo o tipo de endereco recebido, ao executar a função a ser testada.
	addressTypeReceived := TipoDeEndereco(addressToTest)

	/* Faz a validação se o valor recebido pela função a ser testada, é diferente do retorno esperado pela mesma
	caso sejam diferentes aponta ele reprova o test, caso contrario o test passa.

	Para rodar o test basta rodar o comando ```go test``` no terminal.
	*/
	if addressTypeReceived != addressTypeExpected {
		t.Errorf("O tipo de endereço recebido é diferente do tipo de endereço esperado!, esperava: %s, recebido: %s.",
			addressTypeExpected,
			addressTypeReceived,
		)
	}
}
