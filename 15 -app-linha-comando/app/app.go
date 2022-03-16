package app

import "github.com/urfave/cli"

// Generate irá retornar a aplicação de linha de comando pronta para ser executada.
func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicação de linha de comando"
	app.Usage = "Busca Ips e nomes de servidores na internet"

	return app
}
