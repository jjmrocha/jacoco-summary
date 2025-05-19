package main

import (
	"fmt"
	"os"

	"github.com/jjmrocha/jacoco-summary/action"
	"github.com/jjmrocha/jacoco-summary/jacoco"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("No arguments provided.")
		return
	}

	report, err := jacoco.ReadReport(args[0])
	if err != nil {
		fmt.Println("Failed to read report.")
		os.Exit(1)
	}

	summary := action.MarkdownReport(report)
	err = action.WriteJobSummary(summary)

	if err != nil {
		fmt.Println("Failed to write summary.")
		fmt.Println(err)
		os.Exit(1)
	}
}
