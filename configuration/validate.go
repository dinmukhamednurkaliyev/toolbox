package configuration

import (
	"fmt"
	"sort"
	"strings"
	"unicode"
)

func Validate(definition Definition) error {
	if definition.Actions == nil {
		return fmt.Errorf("actions must be an object")
	}

	actionNames := make([]string, 0, len(definition.Actions))

	for actionName := range definition.Actions {
		actionNames = append(actionNames, actionName)
	}

	sort.Strings(actionNames)

	for _, actionName := range actionNames {
		if nameError := validateName(actionName); nameError != nil {
			return fmt.Errorf("invalid action name %q: %w", actionName, nameError)
		}

		actionDefinition := definition.Actions[actionName]

		if actionDefinition.Run != "" && strings.TrimSpace(actionDefinition.Run) == "" {
			return fmt.Errorf("action %q has an empty run command", actionName)
		}

		if actionDefinition.Run == "" && len(actionDefinition.Targets) == 0 {
			return fmt.Errorf("action %q requires a run command or targets", actionName)
		}

		targetNames := make([]string, 0, len(actionDefinition.Targets))

		for targetName := range actionDefinition.Targets {
			targetNames = append(targetNames, targetName)
		}

		sort.Strings(targetNames)

		for _, targetName := range targetNames {
			if nameError := validateName(targetName); nameError != nil {
				return fmt.Errorf(
					"invalid target name %q in action %q: %w",
					targetName,
					actionName,
					nameError,
				)
			}

			targetDefinition := actionDefinition.Targets[targetName]

			if strings.TrimSpace(targetDefinition.Run) == "" {
				return fmt.Errorf(
					"target %q in action %q requires a run command",
					targetName,
					actionName,
				)
			}
		}
	}

	return nil
}

func validateName(name string) error {
	if name == "" {
		return fmt.Errorf("name cannot be empty")
	}

	if strings.ContainsFunc(name, unicode.IsSpace) {
		return fmt.Errorf("name cannot contain whitespace")
	}

	if strings.HasPrefix(name, "-") {
		return fmt.Errorf("name cannot start with a hyphen")
	}

	return nil
}
