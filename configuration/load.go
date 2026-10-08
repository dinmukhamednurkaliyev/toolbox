package configuration

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"

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
	decoder := json.NewDecoder(bytes.NewReader(content))
	decoder.DisallowUnknownFields()

	var definition Definition

	if decodeError := decoder.Decode(&definition); decodeError != nil {
		return Definition{}, fmt.Errorf(
			"parse configuration %q: %w",
			FileName,
			decodeError,
		)
	}

	var additionalContent json.RawMessage

	if decodeError := decoder.Decode(&additionalContent); decodeError != io.EOF {
		if decodeError != nil {
			return Definition{}, fmt.Errorf(
				"parse configuration %q: %w",
				FileName,
				decodeError,
			)
		}

		return Definition{}, fmt.Errorf(
			"configuration %q must contain only one JSON object",
			FileName,
		)
	}

	if validationError := Validate(definition); validationError != nil {
		return Definition{}, fmt.Errorf(
			"validate configuration %q: %w",
			FileName,
			validationError,
		)
	}

	return definition, nil
}
