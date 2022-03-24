package enderecos

import "strings"

func tipoDeEndereco(endereco string) string {
	tiposValidos := []string{"rua", "avenida", "estrada", "rodovia"}

	PrimeiraPalavra := strings.ToLower(strings.Split(endereco, " ")[0])

	for _, value := range tiposValidos {
		if PrimeiraPalavra == value {
			return PrimeiraPalavra
		}
	}

	return "Tipo Invalido"
}
