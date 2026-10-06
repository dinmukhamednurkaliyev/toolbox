package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func writeGeneratedFile(
	filePath string,
	content string,
) error {
	directoryPath := filepath.Dir(filePath)

	createDirectoryError := os.MkdirAll(
		directoryPath,
		0o755,
	)

	if createDirectoryError != nil {
		return fmt.Errorf(
			"create output directory %q: %w",
			directoryPath,
			createDirectoryError,
		)
	}

	writeError := os.WriteFile(
		filePath,
		[]byte(content),
		0o644,
	)

	if writeError != nil {
		return fmt.Errorf(
			"write generated file %q: %w",
			filePath,
			writeError,
		)
	}

	return nil
}
