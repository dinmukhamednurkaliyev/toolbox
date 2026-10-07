package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println(
			"usage: toolbox <module> [arguments...]",
		)
		return
	}

	moduleName := os.Args[1]
	moduleArguments := os.Args[2:]

	runError := runModule(
		moduleName,
		moduleArguments,
	)

	if runError != nil {
		fmt.Fprintln(
			os.Stderr,
			runError,
		)
		os.Exit(1)
	}
}
