package main

import (
	"fmt"
	"os"
	"os/exec"
)

func runModule(
	moduleName string,
	arguments []string,
) error {
	commandArguments := []string{
		"-C",
		moduleName,
		"run",
		".",
	}

	commandArguments = append(
		commandArguments,
		arguments...,
	)

	command := exec.Command(
		"go",
		commandArguments...,
	)

	command.Stdin = os.Stdin
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	if runError := command.Run(); runError != nil {
		return fmt.Errorf(
			"run module %q: %w",
			moduleName,
			runError,
		)
	}

	return nil
}

func moduleExists(
	moduleName string,
) (bool, error) {
	moduleInformation, inspectError := os.Stat(moduleName)

	if inspectError == nil {
		return moduleInformation.IsDir(), nil
	}

	if os.IsNotExist(inspectError) {
		return false, nil
	}

	return false, fmt.Errorf(
		"inspect module %q: %w",
		moduleName,
		inspectError,
	)
}
