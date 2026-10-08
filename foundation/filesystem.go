package foundation

import (
	"fmt"
	"os"
)

func CreateFile(fileName string, content []byte) error {
	createdFile, createError := os.OpenFile(
		fileName,
		os.O_WRONLY|os.O_CREATE|os.O_EXCL,
		0o644,
	)
	if createError != nil {
		return fmt.Errorf("create file %q: %w", fileName, createError)
	}

	_, writeError := createdFile.Write(content)
	closeError := createdFile.Close()

	if writeError != nil {
		return fmt.Errorf("write file %q: %w", fileName, writeError)
	}

	if closeError != nil {
		return fmt.Errorf("close file %q: %w", fileName, closeError)
	}

	return nil
}
