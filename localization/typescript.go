package main

import (
	"encoding/json"
	"sort"
	"strings"
)

func generateTypeScript(
	languages Languages,
) (string, error) {
	var output strings.Builder

	languageNames := sortedLanguageNames(languages)

	output.WriteString(
		"export const languages = {\n",
	)

	for _, languageName := range languageNames {
		output.WriteString("  ")

		if writeError := writeTypeScriptString(
			&output,
			languageName,
		); writeError != nil {
			return "", writeError
		}

		output.WriteString(": ")

		if writeError := writeTypeScriptEntries(
			&output,
			languages[languageName],
			1,
		); writeError != nil {
			return "", writeError
		}

		output.WriteString(",\n")
	}

	output.WriteString(
		"} as const\n",
	)

	return output.String(), nil
}

func sortedLanguageNames(
	languages Languages,
) []string {
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

	return languageNames
}

func writeTypeScriptString(
	output *strings.Builder,
	value string,
) error {
	encodedValue, encodeError := json.Marshal(value)

	if encodeError != nil {
		return encodeError
	}

	output.Write(encodedValue)

	return nil
}

func writeTypeScriptEntries(
	output *strings.Builder,
	entries Entries,
	depth int,
) error {
	output.WriteString("{\n")

	names := make(
		[]string,
		0,
		len(entries),
	)

	for name := range entries {
		names = append(
			names,
			name,
		)
	}

	sort.Strings(names)

	for _, name := range names {
		entry := entries[name]

		output.WriteString(
			strings.Repeat(
				"  ",
				depth+1,
			),
		)

		if writeError := writeTypeScriptString(
			output,
			name,
		); writeError != nil {
			return writeError
		}

		output.WriteString(": ")

		if entry.Children != nil {
			if writeError := writeTypeScriptEntries(
				output,
				entry.Children,
				depth+1,
			); writeError != nil {
				return writeError
			}
		} else {
			if writeError := writeTypeScriptString(
				output,
				entry.Text,
			); writeError != nil {
				return writeError
			}
		}

		output.WriteString(",\n")
	}

	output.WriteString(
		strings.Repeat(
			"  ",
			depth,
		),
	)

	output.WriteString("}")

	return nil
}
