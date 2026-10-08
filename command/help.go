package command

import (
	"fmt"
	"io"
	"strings"
)

var Help = Definition{
	Name:        "help",
	Description: "Show available commands",
}

func runHelp(program Program, arguments []string, output io.Writer) error {
	if len(arguments) > 0 {
		return fmt.Errorf(
			"command %q does not accept arguments",
			Help.Name,
		)
	}

	return printHelp(program, output)
}

func printHelp(program Program, output io.Writer) error {
	var helpContent strings.Builder
	fmt.Fprintf(&helpContent, "%s\n%s\n\nUsage:\n  %s <action> [target]\n\nCommands:\n",
		program.Name,
		program.Description,
		program.Name)

	for _, registeredCommand := range registeredCommands(program) {
		fmt.Fprintf(&helpContent, "  %-12s %s\n",
			registeredCommand.definition.Name,
			registeredCommand.definition.Description)
	}

	if _, writeError := io.WriteString(output, helpContent.String()); writeError != nil {
		return fmt.Errorf("write help: %w", writeError)
	}

	return nil
}
