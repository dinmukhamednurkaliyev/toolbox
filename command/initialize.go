package command

import (
	"fmt"
	"io"

	"github.com/dinmukhamednurkaliyev/toolbox/configuration"
)

var Initialize = Definition{
	Name:        "initialize",
	Description: "Initialize Toolbox in the current directory",
}

func runInitialize(arguments []string, output io.Writer) error {
	if len(arguments) > 0 {
		return fmt.Errorf(
			"command %q does not accept arguments",
			Initialize.Name,
		)
	}

	if configurationError := configuration.Initialize(); configurationError != nil {
		return configurationError
	}

	if _, writeError := fmt.Fprintf(
		output,
		"Created %s\n",
		configuration.FileName,
	); writeError != nil {
		return fmt.Errorf("write initialization result: %w", writeError)
	}

	return nil
}
