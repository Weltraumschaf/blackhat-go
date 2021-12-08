package scanner

import (
	"fmt"
	"net"
)

type portState string

const (
	open   portState = "open"
	closed portState = "closed"
)

type portResult struct {
	port int
	state portState
}

func (pr portResult) String() string {
	return fmt.Sprintf("%d %s", pr.port, pr.state)
}


func scan(ports chan int, results chan *portResult, target string) {
	for port := range ports {
		address := fmt.Sprintf("%s:%d", target, port)
		connection, err := net.Dial("tcp", address)

		if err != nil {
			results <- &portResult{
				port: port,
				state: closed,
			}
			continue
		}

		connection.Close()
		results <- &portResult{
			port: port,
			state: open,
		}
	}
}
