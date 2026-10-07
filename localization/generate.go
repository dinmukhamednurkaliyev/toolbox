package localization

import "fmt"

type Target string

const (
	TargetDart       Target = "dart"
	TargetTypeScript Target = "typescript"
)

func Generate(
	sourceDirectoryPath string,
	target Target,
) (string, error) {
	languages, readError := readLanguages(sourceDirectoryPath)

	if readError != nil {
		return "", readError
	}

	if validationError := validateLanguages(languages); validationError != nil {
		return "", validationError
	}

	switch target {
	case TargetDart:
		return generateDart(languages)

	case TargetTypeScript:
		return generateTypeScript(languages)

	default:
		return "", fmt.Errorf(
			"unsupported localization target %q",
			target,
		)
	}
}
