package scanner

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"net"
	"sort"
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

const (
	portStart  = 1
	portEnd    = 1024
	maxWorkers = 100
)

func Execute(c *cli.Context) error {
	ports := make(chan int, maxWorkers)
	results := make(chan int)
	var openPorts []int

	for i := 0; i < cap(ports); i++ {
		go worker(ports, results)
	}

	go func() {
		for port := portStart; port <= portEnd; port++ {
			ports <- port
		}
	}()

	for port := portStart; port <= portEnd; port++ {
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

func worker(ports, results chan int) {
	for port := range ports {
		address := fmt.Sprintf("scanme.nmap.org:%d", port)
		connection, err := net.Dial("tcp", address)

		if err != nil {
			results <- 0
			continue
		}

		connection.Close()
		results <- port
	}
}
