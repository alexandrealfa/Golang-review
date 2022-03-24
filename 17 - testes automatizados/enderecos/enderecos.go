package enderecos

import "strings"

// TipoDeEndereco : Verifica se o endereço tem um tipo valido e retorna o tipo caso válido.
func TipoDeEndereco(endereco string) string {
	tiposValidos := []string{"rua", "avenida", "estrada", "rodovia"}

	PrimeiraPalavra := strings.ToLower(strings.Split(endereco, " ")[0])

	for _, value := range tiposValidos {
		if PrimeiraPalavra == value {
			return PrimeiraPalavra
		}
	}

	return "Tipo Invalido"
}
