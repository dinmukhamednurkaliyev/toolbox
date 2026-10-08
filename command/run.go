package command

import (
	"io"
)

func Run(program Program, arguments []string, output io.Writer) error {
	if len(arguments) == 0 {
		return printHelp(program, output)
	}

	for _, registeredCommand := range registeredCommands(program) {
		if arguments[0] == registeredCommand.definition.Name {
			return registeredCommand.execute(arguments[1:], output)
		}
	}

	return runAction(arguments, output)
}
