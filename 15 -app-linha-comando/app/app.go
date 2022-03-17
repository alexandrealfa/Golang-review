package app

import (
	"fmt"
	"github.com/urfave/cli"
	"log"
	"net"
)

// Generate é uma função que irá retornar uma instancia de uma aplicação de linha de comando pronta para ser usada.
func Generate() *cli.App {
	/*
		- app : = cli.NewApp(), app está instanciando uma nova aplicação.
		- app.Name , está setando o nome da aplicação.
		- app.Usage, está sendo setado a mensagem de usabilidade da aplicação;
		- app.Commands, é onde se registra os comandos da aplicação.
	*/

	app := cli.NewApp()
	app.Name = "Aplicação de linha de comando"
	app.Usage = "Busca Ips e nomes de servidores na internet"
	app.Commands = []cli.Command{
		{
			Name:  "ip",
			Usage: "Busca Ips de endereços na internet",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "host",
					Value: "devbook.com.br",
				},
			},
			Action: searchIps,
		},
	}

	return app
}

// SearchIps é uma função que captura os dados de host do cli, e lista todos os ips contidos nele.
func searchIps(c *cli.Context) {
	host := c.String("host")
	ips, err := net.LookupIP(host)

	if err != nil {
		log.Fatal(err)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}
