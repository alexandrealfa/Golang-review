package enderecos_test

import (
	. "introducao-testes/enderecos"
	"testing"
)

type testScenario struct {
	addressToTest   string
	addressExpected string
}

/*
 => Unit Tests:
 O nome da função de test é declarada com a inicial Test e na sequência o nome da função que será testada, em camelCase.
 A função de test recebe o ponteiro de testing, que é importado do repositório de funções go.
*/

func TestTipoDeEndereco(t *testing.T) {

	testsScenario := []testScenario{
		{"Rua Teste", "rua"},
		{"Avenida Mato Grosso", "avenida"},
		{"Praça Teste", "Tipo Invalido"},
		{"Estrada Teste", "estrada"},
		{"Avenida Teste", "avenida"},
	}

	for _, scenario := range testsScenario {
		// Definindo o tipo de endereco recebido, ao executar a função a ser testada.
		addressTypeReceived := TipoDeEndereco(scenario.addressToTest)

		/* Faz a validação se o valor recebido pela função a ser testada, é diferente do retorno esperado pela mesma
		caso sejam diferentes aponta ele reprova o test, caso contrario o test passa.

		Para rodar o test basta rodar o comando ```go test``` no terminal.
		*/
		if addressTypeReceived != scenario.addressExpected {
			t.Errorf("O Formato recebido não é compativel com o esperado. Recebido: %s, Esperado: %s",
				addressTypeReceived,
				scenario.addressExpected)
		}
	}
}
