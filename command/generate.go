package command

import "fmt"

func runGenerate(
	subject string,
	arguments []string,
) error {
	switch subject {
	case "localization":
		return runGenerateLocalization(
			arguments,
		)

	default:
		return fmt.Errorf(
			"unsupported generate subject %q",
			subject,
		)
	}
}
