package main

import "fmt"

/* Com funções closures, eu consigo definir comportamentos pre-definidos para serem executados antes de determinada funçã, atraves de um encadeamento de funções.
assim como também é possível definir e atribuir valores a variáveis que ficam armazenadas dentro da função closure, e que podem ser acessadas
nas enclosured functions.
*/

/* está função clousure recebe uma string, e retorna:
- a quantidade de letras em todas as vezes que foi chamada,
- a quantidade de letras no parametro passado
- o parametro passado
*/

func closure() func(string) (int, int, string) {
	var totalLen int

	funcao := func(text string) (int, int, string) {
		characters := len(text)
		totalLen += characters
		return totalLen, characters, text
	}

	return funcao
}

func main() {
	countText, countOthersText := closure(), closure()

	fmt.Println(countText("alexandre"))     // rtn: 9 9 alexandre
	fmt.Println(countText("acrescentando")) // rtn: 22 13 acrescentando

	fmt.Println(countOthersText("novoAcumulator"))        // 14 14 novoAcumulator
	fmt.Println(countOthersText("acrescentando ao novo")) //  35 21 acrescentando ao novo
}
