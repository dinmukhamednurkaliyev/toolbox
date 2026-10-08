package command

import (
	"fmt"
	"io"
	"sort"
	"strings"

	"github.com/dinmukhamednurkaliyev/toolbox/configuration"
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
		fmt.Fprintf(&helpContent, "  %-24s %s\n",
			registeredCommand.definition.Name,
			registeredCommand.definition.Description)
	}

	projectConfiguration, exists, loadError := configuration.LoadIfExists()
	if loadError != nil {
		return loadError
	}

	if exists && len(projectConfiguration.Actions) > 0 {
		helpContent.WriteString("\nProject actions:\n")

		actionNames := make([]string, 0, len(projectConfiguration.Actions))

		for actionName := range projectConfiguration.Actions {
			actionNames = append(actionNames, actionName)
		}

		sort.Strings(actionNames)

		for _, actionName := range actionNames {
			action := projectConfiguration.Actions[actionName]

			if len(action.Targets) == 0 {
				fmt.Fprintf(&helpContent, "  %-24s %s\n",
					actionName,
					action.Description)

				continue
			}

			targetNames := make([]string, 0, len(action.Targets))

			for targetName := range action.Targets {
				targetNames = append(targetNames, targetName)
			}

			sort.Strings(targetNames)

			for _, targetName := range targetNames {
				target := action.Targets[targetName]

				fmt.Fprintf(&helpContent, "  %-24s %s\n",
					actionName+" "+targetName,
					target.Description)
			}
		}
	}

	if _, writeError := io.WriteString(output, helpContent.String()); writeError != nil {
		return fmt.Errorf("write help: %w", writeError)
	}

	return nil
}
