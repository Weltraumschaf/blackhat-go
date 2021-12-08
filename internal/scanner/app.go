package scanner

import (
	"fmt"
	"github.com/urfave/cli/v2"
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
				Name:    "start",
				Aliases: []string{"s"},
				Usage:   "Start port.",
				Value:   portStart,
			},
			&cli.IntFlag{
				Name:    "end",
				Aliases: []string{"e"},
				Usage:   "End port.",
				Value:   portEnd,
			},
			&cli.BoolFlag{
				Name:    "closed",
				Aliases: []string{"c"},
				Usage:   "Show closed or filtered ports in result.",
				Value:   false,
			},
		},
	}
}

func Execute(c *cli.Context) error {
	opts := newOptions(c)
	ports := make(chan int, maxWorkers)
	results := make(chan *portResult)

	createScanWorkers(ports, results, opts)
	go submitPortScans(ports, opts)
	scannedPorts := collectScanResults(results, opts)

	close(ports)
	close(results)

	sortResult(scannedPorts)
	printResult(scannedPorts, opts)

	return nil
}

func createScanWorkers(ports chan int, results chan *portResult, opts *options) {
	for i := 0; i < cap(ports); i++ {
		go scan(ports, results, opts.getTargetHost())
	}
}

func submitPortScans(ports chan int, opts *options) {
	for port := opts.getStartPort(); port <= opts.getEndPort(); port++ {
		ports <- port
	}
}

func collectScanResults(results chan *portResult, opts *options) []*portResult {
	var ports []*portResult

	for port := opts.getStartPort(); port <= opts.getEndPort(); port++ {
		result := <-results
		ports = append(ports, result)
	}

	return ports
}

func sortResult(scannedPorts []*portResult) {
	sort.Slice(scannedPorts, func(i, j int) bool {
		return scannedPorts[i].port < scannedPorts[j].port
	})
}

func printResult(scannedPorts []*portResult, opts *options) {
	for _, port := range scannedPorts {
		if port.state == closed && opts.isHideClosedPorts() {
			continue
		}

		fmt.Println(port)
	}
}
