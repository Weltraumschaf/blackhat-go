package scanner

import (
	"fmt"
	"net"
	"weltraumschaf.de/blackhat/internal/lib"
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
		address := lib.CreateAddress(target, port)
		connection, err := net.Dial(lib.Tcp.String(), address)

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
