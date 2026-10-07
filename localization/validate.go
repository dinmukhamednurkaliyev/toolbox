package localization

// Validate validates localization source files.
func Validate(
	sourceDirectoryPath string,
) error {
	languages, readError := readLanguages(sourceDirectoryPath)

	if readError != nil {
		return readError
	}

	return validateLanguages(languages)
}
