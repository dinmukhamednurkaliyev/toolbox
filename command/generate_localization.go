package command

import (
	"flag"
	"fmt"
	"io"

	"github.com/dinmukhamednurkaliyev/toolbox/foundation"
	"github.com/dinmukhamednurkaliyev/toolbox/localization"
)

func runGenerateLocalization(
	arguments []string,
) error {
	if len(arguments) == 0 {
		printGenerateLocalizationHelp()
		return nil
	}

	for _, argument := range arguments {
		if isHelpOption(argument) {
			printGenerateLocalizationHelp()
			return nil
		}
	}

	commandFlags := flag.NewFlagSet(
		"generate localization",
		flag.ContinueOnError,
	)

	commandFlags.SetOutput(
		io.Discard,
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
		return generateLocalizationUsageError(
			parseError.Error(),
		)
	}

	if commandFlags.NArg() > 0 {
		return generateLocalizationUsageError(
			fmt.Sprintf(
				"unexpected argument %q",
				commandFlags.Arg(0),
			),
		)
	}

	if *sourceDirectoryPath == "" {
		return generateLocalizationUsageError(
			"--source is required",
		)
	}

	if *targetName == "" {
		return generateLocalizationUsageError(
			"--target is required",
		)
	}

	if *outputFilePath == "" {
		return generateLocalizationUsageError(
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
