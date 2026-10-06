package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Entries map[string]Entry

type Entry struct {
	Text     string
	Children Entries
}

func readEntries(
	filePath string,
) (Entries, error) {
	fileContent, readError := os.ReadFile(filePath)

	if readError != nil {
		return nil, fmt.Errorf(
			"read file %q: %w",
			filePath,
			readError,
		)
	}

	var content map[string]any

	parseError := json.Unmarshal(
		fileContent,
		&content,
	)

	if parseError != nil {
		return nil, fmt.Errorf(
			"parse localization file %q: %w",
			filePath,
			parseError,
		)
	}

	return createEntries(content)
}

func createEntries(
	content map[string]any,
) (Entries, error) {
	entries := make(Entries)

	for name, value := range content {
		switch value := value.(type) {
		case string:
			entries[name] = Entry{
				Text: value,
			}

		case map[string]any:
			children, createError := createEntries(value)

			if createError != nil {
				return nil, createError
			}

			entries[name] = Entry{
				Children: children,
			}

		default:
			return nil, fmt.Errorf(
				"localization entry %q must be a string or object",
				name,
			)
		}
	}

	return entries, nil
}
