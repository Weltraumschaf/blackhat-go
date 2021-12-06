package scanner

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"net"
)

func Create() *cli.App {
	return &cli.App{
		Name:    "scanner",
		Version: "1.0.0",
		Authors: []*cli.Author{
			{
				Name:  "Sven Strittmatter",
				Email: "ich@weltraumschaf.de",
			},
		},
		Action: Execute,
	}
}

func Execute(c *cli.Context) error {
	_, err := net.Dial("tcp", "scanme.nmap.org:80")

	if err == nil {
		fmt.Println("Connection successful")
	}
	return nil
}
