package app

import (
	"fmt"
	"github.com/urfave/cli"
	"log"
	"net"
)

// Generate irá retornar a aplicação de linha de comando pronta para ser executada.
func Generate() *cli.App {
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
			Action: buscarIps,
		},
	}

	return app
}

func buscarIps(c *cli.Context) {
	host := c.String("host") //c.Command.Flags.value
	ips, err := net.LookupIP(host)
	if err != nil {
		log.Fatal(err)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}
