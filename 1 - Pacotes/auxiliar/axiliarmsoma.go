package auxiliar

import (
	"fmt"
)

// eu consigo usar essa função em toutro arquivo do mesmo pacote auxiliar, porém fora dele eu não consigo visto que ele foi escrito com a inicial minuscula, sinalizando que é um auxiliar do pacote.
func auxiliarsoma() {
	fmt.Println(5 + 5)
}
