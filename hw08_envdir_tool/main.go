package main

import (
	"log"
	"os"
)

func main() {
	allCommandLineArgs := os.Args[1:]
	if len(allCommandLineArgs) < 2 {
		log.Fatalf("Not enough args = %s", allCommandLineArgs)
	}

	ReadDir(allCommandLineArgs[0])

}
