package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
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

	resultEnv := make(Environment, len(files))
	for _, file := range files {
		if file.IsDir() {
			return nil, fmt.Errorf("it is  dir = %s", file.Name())
		}
		if checkFileName(file.Name()) {
			return nil, fmt.Errorf("invalid file Name = %s", file.Name())
		}

		info, errInfo := file.Info()
		if errInfo != nil {
			return nil, fmt.Errorf("cannot get info about %s, %w", file.Name(), errInfo)
		}
		if info.Size() == 0 {
			resultEnv[info.Name()] = EnvValue{NeedRemove: true}
			continue
		}
		value, errCheck := checkEnvValue(dir + "/" + file.Name())
		if errCheck != nil {
			return nil, fmt.Errorf("error env value %s, %w", file.Name(), errCheck)
		}
		resultEnv[info.Name()] = EnvValue{Value: value}
	}

	return resultEnv, nil
}

func checkEnvValue(fileName string) (string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return "", fmt.Errorf("cannot open file %s %w", fileName, err)
	}
	reader := bufio.NewReader(file)

	lineByte, _, errRead := reader.ReadLine()
	if !(errors.Is(errRead, io.EOF) || errRead == nil) {
		return "", fmt.Errorf("more then one line in file %s %w", fileName, err)
	}

	line := strings.ReplaceAll(string(lineByte), "\x00", "\n")
	line = strings.TrimRight(line, " \t")
	return line, nil
}

func checkFileName(fileName string) bool {
	return strings.ContainsAny(fileName, "=;")
}
