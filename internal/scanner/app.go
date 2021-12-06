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
	const portStart = 1
	const portEnd = 1024

	for port := portStart; port <= portEnd; port++ {
		address := fmt.Sprintf("scanme.nmap.org:%d", port)
		connection, err := net.Dial("tcp", address)

		if err != nil {
			fmt.Println("closed/filtered:", port)
			continue
		}

		fmt.Print("open: ", port)

		err = connection.Close()

		if err != nil {
			fmt.Println("(Error on close:", err, ")")
		} else {
			fmt.Println()
		}
	}

	return nil
}
