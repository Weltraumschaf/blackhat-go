package main

import (
	"fmt"
	"os"
	"weltraumschaf.de/blackhat/internal/proxy"
)

func main() {
	var app = proxy.Create()
	err := app.Run(os.Args)

	if err != nil {
		fmt.Println("Error:", err.Error())
		os.Exit(1)
	}
}
