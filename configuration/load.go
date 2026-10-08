package configuration

import (
	"encoding/json"
	"fmt"

	"github.com/dinmukhamednurkaliyev/toolbox/foundation"
)

func Load() (Definition, error) {
	content, readError := foundation.ReadFile(FileName)
	if readError != nil {
		return Definition{}, readError
	}

	return parseDefinition(content)
}

func LoadIfExists() (Definition, bool, error) {
	content, exists, readError := foundation.ReadFileIfExists(FileName)
	if readError != nil {
		return Definition{}, false, readError
	}

	if !exists {
		return Definition{}, false, nil
	}

	definition, parseError := parseDefinition(content)
	if parseError != nil {
		return Definition{}, true, parseError
	}

	return definition, true, nil
}

func parseDefinition(content []byte) (Definition, error) {
	var definition Definition

	if parseError := json.Unmarshal(content, &definition); parseError != nil {
		return Definition{}, fmt.Errorf(
			"parse configuration %q: %w",
			FileName,
			parseError,
		)
	}

	if definition.Actions == nil {
		return Definition{}, fmt.Errorf(
			"configuration %q must contain an actions object",
			FileName,
		)
	}

	return definition, nil
}
