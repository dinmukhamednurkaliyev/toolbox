package main

import (
	"fmt"
	"os"

	"github.com/dinmukhamednurkaliyev/toolbox/command"
)

func main() {
	if runError := command.Run(
		os.Args[1:],
	); runError != nil {
		fmt.Fprintln(
			os.Stderr,
			runError,
		)

		os.Exit(1)
	}
}
