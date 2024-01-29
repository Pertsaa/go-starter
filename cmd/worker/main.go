package main

import (
	"fmt"
	"os"

	"github.com/Pertsaa/go-starter/internal/app/worker"
)

func main() {
	if err := worker.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
