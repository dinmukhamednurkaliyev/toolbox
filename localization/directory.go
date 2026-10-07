package localization

import (
	"fmt"
	"os"
	"path/filepath"
)

func readDirectory(
	directoryPath string,
) (Entries, error) {
	directoryEntries, readError := os.ReadDir(directoryPath)

	if readError != nil {
		return nil, fmt.Errorf(
			"read directory %q: %w",
			directoryPath,
			readError,
		)
	}

	entries := make(Entries)

	for _, directoryEntry := range directoryEntries {
		entryPath := filepath.Join(
			directoryPath,
			directoryEntry.Name(),
		)

		if directoryEntry.IsDir() {
			children, readDirectoryError := readDirectory(entryPath)

			if readDirectoryError != nil {
				return nil, readDirectoryError
			}

			directoryEntries := Entries{
				directoryEntry.Name(): Entry{
					Children: children,
				},
			}

			if mergeError := mergeEntries(
				entries,
				directoryEntries,
			); mergeError != nil {
				return nil, mergeError
			}

			continue
		}

		if filepath.Ext(
			directoryEntry.Name(),
		) != ".json" {
			continue
		}

		fileEntries, readFileError := readEntries(entryPath)

		if readFileError != nil {
			return nil, readFileError
		}

		if mergeError := mergeEntries(
			entries,
			fileEntries,
		); mergeError != nil {
			return nil, mergeError
		}
	}

	return entries, nil
}

func mergeEntries(
	entries Entries,
	incomingEntries Entries,
) error {
	for name, incomingEntry := range incomingEntries {
		currentEntry, exists := entries[name]

		if !exists {
			entries[name] = incomingEntry
			continue
		}

		if currentEntry.Children != nil &&
			incomingEntry.Children != nil {
			if mergeError := mergeEntries(
				currentEntry.Children,
				incomingEntry.Children,
			); mergeError != nil {
				return mergeError
			}

			continue
		}

		return fmt.Errorf(
			"localization entry %q is defined more than once",
			name,
		)
	}

	return nil
}
