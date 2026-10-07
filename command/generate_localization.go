package command

import (
	"flag"
	"fmt"

	"github.com/dinmukhamednurkaliyev/toolbox/foundation"
	"github.com/dinmukhamednurkaliyev/toolbox/localization"
)

func runGenerateLocalization(
	arguments []string,
) error {
	commandFlags := flag.NewFlagSet(
		"generate localization",
		flag.ContinueOnError,
	)

	sourceDirectoryPath := commandFlags.String(
		"source",
		"",
		"localization source directory",
	)

	targetName := commandFlags.String(
		"target",
		"",
		"generation target",
	)

	outputFilePath := commandFlags.String(
		"output",
		"",
		"generated output file",
	)

	if parseError := commandFlags.Parse(arguments); parseError != nil {
		return parseError
	}

	if *sourceDirectoryPath == "" {
		return fmt.Errorf(
			"--source is required",
		)
	}

	if *targetName == "" {
		return fmt.Errorf(
			"--target is required",
		)
	}

	if *outputFilePath == "" {
		return fmt.Errorf(
			"--output is required",
		)
	}

	generatedCode, generateError := localization.Generate(
		*sourceDirectoryPath,
		localization.Target(
			*targetName,
		),
	)

	if generateError != nil {
		return generateError
	}

	if writeError := foundation.WriteFile(
		*outputFilePath,
		[]byte(generatedCode),
	); writeError != nil {
		return writeError
	}

	fmt.Printf(
		"generated %s\n",
		*outputFilePath,
	)

	return nil
}
