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

	generatedCode, generateError := generateTypeScript(languages)

	if generateError != nil {
		fmt.Println(generateError)
		return
	}

	fmt.Println(generatedCode)
}
