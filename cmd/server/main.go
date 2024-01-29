package main

import (
	"fmt"
	"os"

	"github.com/Pertsaa/go-starter/internal/app/server"
)

func main() {
	if err := server.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
