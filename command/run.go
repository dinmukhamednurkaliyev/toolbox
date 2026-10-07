package command

import "fmt"

// Run parses and executes a Toolbox command.
func Run(
	arguments []string,
) error {
	if len(arguments) == 0 {
		printRootHelp()
		return nil
	}

	if isHelpArgument(arguments[0]) {
		printRootHelp()
		return nil
	}

	verb := arguments[0]
	commandArguments := arguments[1:]

	switch verb {
	case "generate":
		return runGenerate(
			commandArguments,
		)

	case "validate":
		return runValidate(
			commandArguments,
		)

	default:
		return fmt.Errorf(
			"unsupported command %q\n\nUse \"toolbox --help\" to see available commands",
			verb,
		)
	}
}
