package foundation

import (
	"encoding/json"
	"fmt"
)

type JSONObject map[string]any

func ReadJSON(
	filePath string,
) (JSONObject, error) {
	fileContent, readError := ReadFile(filePath)

	if readError != nil {
		return nil, readError
	}

	var object JSONObject

	if parseError := json.Unmarshal(
		fileContent,
		&object,
	); parseError != nil {
		return nil, fmt.Errorf(
			"parse JSON file %q: %w",
			filePath,
			parseError,
		)
	}

	return object, nil
}
