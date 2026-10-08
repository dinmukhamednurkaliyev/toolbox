package main

import (
	"fmt"
	"os"

	"github.com/dinmukhamednurkaliyev/toolbox/command"
)

func main() {
	if executionError := command.Run(
		Toolbox,
		os.Args[1:],
		os.Stdout,
	); executionError != nil {
		_, _ = fmt.Fprintln(os.Stderr, executionError)
		os.Exit(1)
	}
}
