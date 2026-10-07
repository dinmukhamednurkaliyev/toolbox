package localization

import (
	"fmt"
	"os"
	"path/filepath"
)

type Languages map[string]Entries

func readLanguages(
	directoryPath string,
) (Languages, error) {
	directoryEntries, readError := os.ReadDir(directoryPath)

	if readError != nil {
		return nil, fmt.Errorf(
			"read languages directory %q: %w",
			directoryPath,
			readError,
		)
	}

	languages := make(Languages)

	for _, directoryEntry := range directoryEntries {
		if !directoryEntry.IsDir() {
			continue
		}

		languagePath := filepath.Join(
			directoryPath,
			directoryEntry.Name(),
		)

		entries, readLanguageError := readDirectory(languagePath)

		if readLanguageError != nil {
			return nil, fmt.Errorf(
				"read language %q: %w",
				directoryEntry.Name(),
				readLanguageError,
			)
		}

		languages[directoryEntry.Name()] = entries
	}

	if len(languages) == 0 {
		return nil, fmt.Errorf(
			"no languages found in %q",
			directoryPath,
		)
	}

	return languages, nil
}
