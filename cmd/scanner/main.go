package main

import (
	"fmt"
	"os"
	"weltraumschaf.de/blackhat/internal/scanner"
)

func main() {
	var app = scanner.Create()
	err := app.Run(os.Args)

	if err != nil {
		fmt.Println("Error:", err.Error())
		os.Exit(1)
	}
}
