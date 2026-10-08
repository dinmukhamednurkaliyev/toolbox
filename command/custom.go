package command

import (
	"fmt"

	"github.com/dinmukhamednurkaliyev/toolbox/configuration"
)

func loadProjectConfiguration(program Program) (configuration.Definition, bool, error) {
	projectConfiguration, exists, loadError := configuration.LoadIfExists()
	if loadError != nil {
		return configuration.Definition{}, exists, loadError
	}

	if !exists {
		return configuration.Definition{}, false, nil
	}

	for _, registeredCommand := range registeredCommands(program) {
		commandName := registeredCommand.definition.Name

		if _, actionExists := projectConfiguration.Actions[commandName]; actionExists {
			return configuration.Definition{}, true, fmt.Errorf(
				"project action %q conflicts with a built-in command",
				commandName,
			)
		}
	}

	return projectConfiguration, true, nil
}
