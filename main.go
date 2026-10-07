package main

import (
	"fmt"
	"os"

	"github.com/dinmukhamednurkaliyev/toolbox/foundation"
	"github.com/dinmukhamednurkaliyev/toolbox/localization"
)

func main() {
	if len(os.Args) != 5 {
		fmt.Fprintln(
			os.Stderr,
			"usage: toolbox localization <source-directory> <target> <output-file>",
		)

		os.Exit(1)
	}

	commandName := os.Args[1]

	if commandName != "localization" {
		fmt.Fprintf(
			os.Stderr,
			"unsupported command %q\n",
			commandName,
		)

		os.Exit(1)
	}

	sourceDirectoryPath := os.Args[2]
	target := localization.Target(
		os.Args[3],
	)
	outputFilePath := os.Args[4]

	generatedCode, generateError := localization.Generate(
		sourceDirectoryPath,
		target,
	)

	if generateError != nil {
		fmt.Fprintln(
			os.Stderr,
			generateError,
		)

		os.Exit(1)
	}

	if writeError := foundation.WriteFile(
		outputFilePath,
		[]byte(generatedCode),
	); writeError != nil {
		fmt.Fprintln(
			os.Stderr,
			writeError,
		)

		os.Exit(1)
	}

	fmt.Printf(
		"generated %s\n",
		outputFilePath,
	)
}
