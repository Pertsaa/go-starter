package main

import (
	"fmt"
	"os"

	"github.com/Pertsaa/go-starter/internal/app/email"
)

func main() {
	if err := email.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
