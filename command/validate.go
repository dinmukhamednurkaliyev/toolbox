package command

import "fmt"

func runValidate(
	arguments []string,
) error {
	if len(arguments) == 0 {
		printValidateHelp()
		return nil
	}

	if isHelpArgument(arguments[0]) {
		printValidateHelp()
		return nil
	}

	subject := arguments[0]
	commandArguments := arguments[1:]

	switch subject {
	case "localization":
		return runValidateLocalization(
			commandArguments,
		)

	default:
		return fmt.Errorf(
			"unsupported validate subject %q\n\nUse \"toolbox validate --help\" to see available subjects",
			subject,
		)
	}
}
