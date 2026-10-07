package command

import "fmt"

const generateLocalizationUsage = "toolbox generate localization --source <directory> --target <target> --output <file>"

const validateLocalizationUsage = "toolbox validate localization --source <directory>"

func printRootHelp() {
	fmt.Print(`Toolbox

Usage:
  toolbox <verb> <subject> [options]

Commands:
  generate localization   Generate localization source code
  validate localization   Validate localization source files

Use "toolbox <verb> --help" for more information.
`)
}

func printGenerateHelp() {
	fmt.Print(`Generate

Usage:
  toolbox generate <subject> [options]

Subjects:
  localization   Generate localization source code

Use "toolbox generate <subject> --help" for more information.
`)
}

func printGenerateLocalizationHelp() {
	fmt.Printf(`Generate localization source code.

Usage:
  %s

Options:
  --source <directory>   Localization source directory
  --target <target>      Generation target: dart or typescript
  --output <file>        Generated output file
`, generateLocalizationUsage)
}

func printValidateHelp() {
	fmt.Print(`Validate

Usage:
  toolbox validate <subject> [options]

Subjects:
  localization   Validate localization source files

Use "toolbox validate <subject> --help" for more information.
`)
}

func printValidateLocalizationHelp() {
	fmt.Printf(`Validate localization source files.

Usage:
  %s

Options:
  --source <directory>   Localization source directory
`, validateLocalizationUsage)
}

func generateLocalizationUsageError(
	message string,
) error {
	return fmt.Errorf(
		"%s\n\nUsage:\n  %s\n\nUse \"toolbox generate localization --help\" for more information.",
		message,
		generateLocalizationUsage,
	)
}

func validateLocalizationUsageError(
	message string,
) error {
	return fmt.Errorf(
		"%s\n\nUsage:\n  %s\n\nUse \"toolbox validate localization --help\" for more information.",
		message,
		validateLocalizationUsage,
	)
}

func isHelpArgument(
	argument string,
) bool {
	return argument == "help" ||
		argument == "--help" ||
		argument == "-h"
}

func isHelpOption(
	argument string,
) bool {
	return argument == "--help" ||
		argument == "-h"
}
