package email

import (
	"fmt"

	"github.com/Pertsaa/go-starter/internal/pkg/report"
)

func Run() error {
	fmt.Println("Running Email app...")

	r := report.GenerateReport()

	fmt.Println(r)

	return nil
}
