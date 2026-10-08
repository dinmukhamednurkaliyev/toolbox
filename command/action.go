package command

import (
	"fmt"
	"io"

	"github.com/dinmukhamednurkaliyev/toolbox/foundation"
)

func RunAction(program Program, arguments []string, output io.Writer) error {
	projectConfiguration, exists, loadError := loadProjectConfiguration(program)
	if loadError != nil {
		return loadError
	}

	if !exists {
		return fmt.Errorf("unknown command %q", arguments[0])
	}

	actionName := arguments[0]

	actionDefinition, actionExists := projectConfiguration.Actions[actionName]
	if !actionExists {
		return fmt.Errorf("unknown action %q", actionName)
	}

	if len(arguments) == 1 {
		if actionDefinition.Run != "" {
			return foundation.ExecuteProcess(actionDefinition.Run, output)
		}

		return fmt.Errorf("action %q requires a target", actionName)
	}

	if len(arguments) > 2 {
		return fmt.Errorf("action %q accepts at most one target", actionName)
	}

	if len(actionDefinition.Targets) == 0 {
		return fmt.Errorf("action %q does not accept a target", actionName)
	}

	targetName := arguments[1]

	targetDefinition, targetExists := actionDefinition.Targets[targetName]
	if !targetExists {
		return fmt.Errorf(
			"unknown target %q for action %q",
			targetName,
			actionName,
		)
	}

	return foundation.ExecuteProcess(targetDefinition.Run, output)
}
