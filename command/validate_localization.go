package command

import (
	"flag"
	"fmt"
	"io"

	"github.com/dinmukhamednurkaliyev/toolbox/localization"
)

func runValidateLocalization(
	arguments []string,
) error {
	if len(arguments) == 0 {
		printValidateLocalizationHelp()
		return nil
	}

	for _, argument := range arguments {
		if isHelpOption(argument) {
			printValidateLocalizationHelp()
			return nil
		}
	}

	commandFlags := flag.NewFlagSet(
		"validate localization",
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

	if parseError := commandFlags.Parse(arguments); parseError != nil {
		return validateLocalizationUsageError(
			parseError.Error(),
		)
	}

	if commandFlags.NArg() > 0 {
		return validateLocalizationUsageError(
			fmt.Sprintf(
				"unexpected argument %q",
				commandFlags.Arg(0),
			),
		)
	}

	if *sourceDirectoryPath == "" {
		return validateLocalizationUsageError(
			"--source is required",
		)
	}

	if validationError := localization.Validate(
		*sourceDirectoryPath,
	); validationError != nil {
		return validationError
	}

	fmt.Println(
		"localization is valid",
	)

	return nil
}
