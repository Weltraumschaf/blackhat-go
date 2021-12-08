package scanner

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"net"
	"sort"
)

const (
	portStart  = 1
	portEnd    = 1024
	maxWorkers = 100
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
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "target",
				Aliases:  []string{"t"},
				Usage:    "The target host to scan.",
				Required: true,
			},
			&cli.IntFlag{
				Name: "start",
				Aliases: []string{"s"},
				Usage: "Start port.",
				Value: portStart,
			},
			&cli.IntFlag{
				Name: "end",
				Aliases: []string{"e"},
				Usage: "End port.",
				Value: portEnd,
			},
		},
	}
}

func Execute(c *cli.Context) error {
	opts := newOptions(c)
	ports := make(chan int, maxWorkers)
	results := make(chan int)
	var openPorts []int

	for i := 0; i < cap(ports); i++ {
		go scan(ports, results, opts.getTarget())
	}

	go func() {
		for port := opts.getStart(); port <= opts.getEnd(); port++ {
			ports <- port
		}
	}()

	for port := opts.getStart(); port <= opts.getEnd(); port++ {
		port := <-results

		if port != 0 {
			openPorts = append(openPorts, port)
		}
	}

	close(ports)
	close(results)
	sort.Ints(openPorts)

	for _, port := range openPorts {
		fmt.Printf("%d open\n", port)
	}

	return nil
}

func scan(ports, results chan int, target string) {
	for port := range ports {
		address := fmt.Sprintf("%s:%d", target, port)
		connection, err := net.Dial("tcp", address)

		if err != nil {
			results <- 0
			continue
		}

		connection.Close()
		results <- port
	}
}
