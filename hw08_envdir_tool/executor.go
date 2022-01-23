package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"syscall"
)

// RunCmd runs a command + arguments (cmd) with environment variables from env.
func RunCmd(cmd []string, env Environment) (returnCode int) {
	for envName, envValue := range env {
		if envValue.NeedRemove {
			if err := os.Unsetenv(envName); err != nil {
			}
		}
		if err := os.Setenv(envName, envValue.Value); err != nil {
			log.Print(fmt.Errorf("can not set env %s, value %s %w", envName, envValue.Value, err))
		}
	}

	command := exec.Command(cmd[0], cmd[1:]...)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	command.Stdin = os.Stdin

	if err := command.Start(); err != nil {
		log.Print(fmt.Errorf("command not started %w", err))
	}

	if err := command.Wait(); err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			// The program has exited with an exit code != 0

			// This works on both Unix and Windows. Although package
			// syscall is generally platform dependent, WaitStatus is
			// defined for both Unix and Windows and in both cases has
			// an ExitStatus() method with the same signature.
			if status, okSys := exitErr.Sys().(syscall.WaitStatus); okSys {
				return status.ExitStatus()
			}
		} else {
			log.Print(fmt.Errorf("cannot take exit code command.Wait: %w", err))
			return -1
		}
	}
	return 0
}
