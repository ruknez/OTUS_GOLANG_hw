package main

import (
	"fmt"
	"go.uber.org/zap"
	"log"
	"os"
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}
	defer logger.Sync()

	allCommandLineArgs := os.Args[1:]
	if len(allCommandLineArgs) < 2 {
		log.Fatalf("Not enough args = %s", allCommandLineArgs)
	}

	envirment, err := ReadDir(allCommandLineArgs[0])
	if err != nil {
		log.Fatal(fmt.Errorf("cannot read dir = %w", err))
	}

	RunCmd(allCommandLineArgs[1:], envirment)
	//fmt.Println("exit code LOLO = ", exitCode)
}
