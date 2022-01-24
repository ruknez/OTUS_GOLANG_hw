package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	allCommandLineArgs := os.Args[1:]
	if len(allCommandLineArgs) < 2 {
		log.Fatalf("Not enough args = %s", allCommandLineArgs)
	}

	envirment, err := ReadDir(allCommandLineArgs[0])
	if err != nil {
		log.Fatal(fmt.Errorf("cannot read dir = %w", err))
	}

	RunCmd(allCommandLineArgs[1:], envirment)
}
