package proxy

import (
	"github.com/urfave/cli/v2"
	"io"
	"log"
	"weltraumschaf.de/blackhat/internal/lib"
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
	listener, err := net.Listen(lib.Tcp.String(), lib.CreateAddress("", 80))
	if err != nil {
		log.Fatalln("Unable to bind to port 80")
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln("Unable to accept connection")
		}

		go handle(conn)
	}

	return nil
}

func handle(src net.Conn) {
	address := lib.CreateAddress("neverssl.com", 80)
	dst, err := net.Dial(lib.Tcp.String(), address)
	if err != nil {
		log.Fatalln("Unable to connect")
	}

	defer dst.Close()

	go func() {
		if _, err := io.Copy(dst, src); err != nil {
			log.Fatalln(err)
		}
	}()

	if _,err := io.Copy(src, dst); err != nil {
		log.Fatalln(err)
	}
}
