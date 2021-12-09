package lib

import "fmt"

type NetworkType int

const (
	Tcp NetworkType = iota
)

func (nt NetworkType) String() string {
	names := [...]string{"tcp"}

	if len(names) < int(nt) {
		return "unknown"
	}

	return names[nt]
}

func CreateAddress(host string, port int) string {
	return fmt.Sprintf("%s:%d", host, port)
}
