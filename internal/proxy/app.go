package proxy

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"io"
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

	if _, err := io.Copy(&writer, &reader); err != nil {
		log.Fatalln("Unable to read/write data")
		return err
	}
	
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
