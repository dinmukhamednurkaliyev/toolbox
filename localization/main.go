package main

import "fmt"

func main() {
	languages, readError := readLanguages("example")

	if readError != nil {
		fmt.Println(readError)
		return
	}

	validationError := validateLanguages(languages)

	if validationError != nil {
		fmt.Println(validationError)
		return
	}

	generatedCode, generateError := generateDart(languages)

	if generateError != nil {
		fmt.Println(generateError)
		return
	}

	writeError := writeGeneratedFile(
		"generated/localization.dart",
		generatedCode,
	)

	if writeError != nil {
		fmt.Println(writeError)
		return
	}

	fmt.Println("generated localization.dart")
}
