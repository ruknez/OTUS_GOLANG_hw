package main

import (
	"fmt"
	"log"
	"os"
)

type Environment map[string]EnvValue

// EnvValue helps to distinguish between empty files and files with the first empty line.
type EnvValue struct {
	Value      string
	NeedRemove bool
}

// ReadDir reads a specified directory and returns map of env variables.
// Variables represented as files where filename is name of variable, file first line is a value.
func ReadDir(dir string) (Environment, error) {
	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	for  _, file := range files{

	}


	resultEnv := make(Environment, len(files))
	for _, file := range files {
		if file.IsDir() {
			return nil, fmt.Errorf("it is  dir = %s", file.Name())
		}

	}

	return nil, nil
}

