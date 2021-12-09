package proxy

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"os"
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
	var (
		reader FooReader
		writer FooWriter
	)

	input := make([]byte, 4096)
	s, err := reader.Read(input)
	if err != nil {
		log.Fatalln("Unable to read data")
	}

	fmt.Printf("Read %d bytes from stdin\n", s)

	s, err = writer.Write(input)
	if err != nil {
		log.Fatalln("Unable to write data")
	}

	fmt.Printf("Wrote %d bytes to stdout\n", s)
	return nil
}

type FooReader struct {}

func (fr *FooReader) Read(b []byte) (int, error) {
	fmt.Print("in < ")
	return os.Stdin.Read(b)
}

type FooWriter struct {}

func (fw *FooWriter) Write(b []byte) (int, error)  {
	fmt.Print("out>")
	return os.Stdout.Write(b)
}
