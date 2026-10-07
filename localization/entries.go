package localization

import (
	"fmt"

	"github.com/dinmukhamednurkaliyev/toolbox/foundation"
)

type Entries map[string]Entry

type Entry struct {
	Text     string
	Children Entries
}

func readEntries(
	filePath string,
) (Entries, error) {
	content, readError := foundation.ReadJSON(filePath)

	if readError != nil {
		return nil, readError
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
