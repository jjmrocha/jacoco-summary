package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("No arguments provided.")
		return
	}

	fmt.Println("Arguments:", args)
}
