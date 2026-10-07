package foundation

import (
	"fmt"
	"os"
)

func ReadDirectory(
	directoryPath string,
) ([]os.DirEntry, error) {
	directoryEntries, readError := os.ReadDir(directoryPath)

	if readError != nil {
		return nil, fmt.Errorf(
			"read directory %q: %w",
			directoryPath,
			readError,
		)
	}

	return directoryEntries, nil
}
