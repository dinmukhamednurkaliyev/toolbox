package localization

import (
	"fmt"
	"sort"
)

func validateLanguages(
	languages Languages,
) error {
	pathsByLanguage := make(
		map[string]map[string]bool,
		len(languages),
	)

	allPaths := make(map[string]bool)

	for languageName, entries := range languages {
		languagePaths := make(map[string]bool)

		collectEntryPaths(
			entries,
			"",
			languagePaths,
		)

		pathsByLanguage[languageName] = languagePaths

		for path := range languagePaths {
			allPaths[path] = true
		}
	}

	languageNames := make(
		[]string,
		0,
		len(languages),
	)

	for languageName := range languages {
		languageNames = append(
			languageNames,
			languageName,
		)
	}

	sort.Strings(languageNames)

	paths := make(
		[]string,
		0,
		len(allPaths),
	)

	for path := range allPaths {
		paths = append(
			paths,
			path,
		)
	}

	sort.Strings(paths)

	for _, path := range paths {
		for _, languageName := range languageNames {
			if pathsByLanguage[languageName][path] {
				continue
			}

			return fmt.Errorf(
				"localization entry %q is missing in language %q",
				path,
				languageName,
			)
		}
	}

	return nil
}

func collectEntryPaths(
	entries Entries,
	parentPath string,
	paths map[string]bool,
) {
	for name, entry := range entries {
		path := name

		if parentPath != "" {
			path =
				parentPath + "." + name
		}

		if entry.Children != nil {
			collectEntryPaths(
				entry.Children,
				path,
				paths,
			)

			continue
		}

		paths[path] = true
	}
}
