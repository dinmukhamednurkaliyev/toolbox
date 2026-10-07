package foundation

import (
	"fmt"
	"os"
	"path/filepath"
)

func ReadFile(
	filePath string,
) ([]byte, error) {
	fileContent, readError := os.ReadFile(filePath)

	if readError != nil {
		return nil, fmt.Errorf(
			"read file %q: %w",
			filePath,
			readError,
		)
	}

	return fileContent, nil
}

func WriteFile(
	filePath string,
	content []byte,
) error {
	directoryPath := filepath.Dir(filePath)

	if createDirectoryError := os.MkdirAll(
		directoryPath,
		0o755,
	); createDirectoryError != nil {
		return fmt.Errorf(
			"create directory %q: %w",
			directoryPath,
			createDirectoryError,
		)
	}

	if writeError := os.WriteFile(
		filePath,
		content,
		0o644,
	); writeError != nil {
		return fmt.Errorf(
			"write file %q: %w",
			filePath,
			writeError,
		)
	}

	return nil
}
