package main

import (
	"log"
	"os"

	"github.com/jjmrocha/jacoco-summary/action"
	"github.com/jjmrocha/jacoco-summary/jacoco"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		log.Fatalln("No arguments provided.")
	}

	fileName := action.GetFileName(args[0])

	report, err := jacoco.ReadReport(fileName)
	if err != nil {
		log.Fatalf("%v\n", err)
	}

	summary := action.MarkdownReport(report)

	err = action.WriteJobSummary(summary)
	if err != nil {
		log.Fatalf("%v\n", err)
	}
}
