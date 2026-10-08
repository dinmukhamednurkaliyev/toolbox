package foundation

import (
	"errors"
	"fmt"
	"os"
)

func CreateFile(fileName string, content []byte) error {
	file, createError := os.OpenFile(
		fileName,
		os.O_WRONLY|os.O_CREATE|os.O_EXCL,
		0o644,
	)
	if createError != nil {
		return fmt.Errorf("create file %q: %w", fileName, createError)
	}

	_, writeError := file.Write(content)
	closeError := file.Close()

	if writeError != nil {
		return fmt.Errorf("write file %q: %w", fileName, writeError)
	}

	if closeError != nil {
		return fmt.Errorf("close file %q: %w", fileName, closeError)
	}

	return nil
}

func ReadFile(fileName string) ([]byte, error) {
	content, readError := os.ReadFile(fileName)
	if readError != nil {
		return nil, fmt.Errorf("read file %q: %w", fileName, readError)
	}

	return content, nil
}

func ReadFileIfExists(fileName string) ([]byte, bool, error) {
	content, readError := ReadFile(fileName)

	if errors.Is(readError, os.ErrNotExist) {
		return nil, false, nil
	}

	if readError != nil {
		return nil, false, readError
	}

	return content, true, nil
}
