package foundation

import (
	"fmt"
	"os"
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
