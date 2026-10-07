package command

import (
	"fmt"
)

func Run(
	arguments []string,
) error {
	if len(arguments) < 2 {
		return fmt.Errorf(
			"expected command in format: toolbox <verb> <subject>",
		)
	}

	verb := arguments[0]
	subject := arguments[1]
	commandArguments := arguments[2:]

	switch verb {
	case "generate":
		return runGenerate(
			subject,
			commandArguments,
		)

	default:
		return fmt.Errorf(
			"unsupported command %q",
			verb,
		)
	}
}
