package main

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"fmt"
	"sort"
	"strings"
	"text/template"
	"unicode"
)

//go:embed templates/dart.dart
var dartTemplate string

type DartTemplateData struct {
	Classes       string
	Languages     string
	Localizations string
}

func generateDart(
	languages Languages,
) (string, error) {
	languageNames := sortedLanguageNames(languages)

	if len(languageNames) == 0 {
		return "", fmt.Errorf(
			"no languages to generate",
		)
	}

	classesCode := generateDartClasses(
		languages[languageNames[0]],
	)

	languagesCode := generateDartLanguages(languages)

	localizationsCode, generateError := generateDartLocalizations(languages)

	if generateError != nil {
		return "", generateError
	}

	parsedTemplate, parseError := template.New("dart").
		Parse(dartTemplate)

	if parseError != nil {
		return "", fmt.Errorf(
			"parse Dart template: %w",
			parseError,
		)
	}

	var output bytes.Buffer

	executeError := parsedTemplate.Execute(
		&output,
		DartTemplateData{
			Classes:       classesCode,
			Languages:     languagesCode,
			Localizations: localizationsCode,
		},
	)

	if executeError != nil {
		return "", fmt.Errorf(
			"execute Dart template: %w",
			executeError,
		)
	}

	return output.String(), nil
}

func generateDartLanguages(
	languages Languages,
) string {
	var output strings.Builder

	for _, languageName := range sortedLanguageNames(languages) {
		output.WriteString("  ")
		output.WriteString(languageName)
		output.WriteString(",\n")
	}

	return output.String()
}

func generateDartClasses(
	entries Entries,
) string {
	var output strings.Builder

	writeDartClass(
		&output,
		"Localization",
		entries,
	)

	return output.String()
}

func writeDartClass(
	output *strings.Builder,
	className string,
	entries Entries,
) {
	names := sortedEntryNames(entries)

	output.WriteString("class ")
	output.WriteString(className)
	output.WriteString(" {\n")

	output.WriteString("  const ")
	output.WriteString(className)
	output.WriteString("({\n")

	for _, name := range names {
		output.WriteString(
			"    required this.",
		)

		output.WriteString(name)
		output.WriteString(",\n")
	}

	output.WriteString("  });\n")

	if len(names) > 0 {
		output.WriteString("\n")
	}

	for _, name := range names {
		entry := entries[name]

		output.WriteString("  final ")

		if entry.Children != nil {
			output.WriteString(
				dartChildClassName(
					className,
					name,
				),
			)
		} else {
			output.WriteString("String")
		}

		output.WriteString(" ")
		output.WriteString(name)
		output.WriteString(";\n")
	}

	output.WriteString("}")

	for _, name := range names {
		entry := entries[name]

		if entry.Children == nil {
			continue
		}

		output.WriteString("\n\n")

		writeDartClass(
			output,
			dartChildClassName(
				className,
				name,
			),
			entry.Children,
		)
	}
}

func generateDartLocalizations(
	languages Languages,
) (string, error) {
	var output strings.Builder

	for _, languageName := range sortedLanguageNames(languages) {
		output.WriteString("  Language.")
		output.WriteString(languageName)
		output.WriteString(": ")

		writeError := writeDartLocalization(
			&output,
			"Localization",
			languages[languageName],
			1,
		)

		if writeError != nil {
			return "", writeError
		}

		output.WriteString(",\n")
	}

	return output.String(), nil
}

func writeDartLocalization(
	output *strings.Builder,
	className string,
	entries Entries,
	depth int,
) error {
	output.WriteString(className)
	output.WriteString("(\n")

	for _, name := range sortedEntryNames(entries) {
		entry := entries[name]

		output.WriteString(
			strings.Repeat(
				"  ",
				depth+1,
			),
		)

		output.WriteString(name)
		output.WriteString(": ")

		if entry.Children != nil {
			writeError := writeDartLocalization(
				output,
				dartChildClassName(
					className,
					name,
				),
				entry.Children,
				depth+1,
			)

			if writeError != nil {
				return writeError
			}
		} else {
			if writeError := writeDartString(
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

	output.WriteString(")")

	return nil
}

func writeDartString(
	output *strings.Builder,
	value string,
) error {
	encodedValue, encodeError := json.Marshal(value)

	if encodeError != nil {
		return encodeError
	}

	escapedValue := strings.ReplaceAll(
		string(encodedValue),
		"$",
		`\$`,
	)

	output.WriteString(escapedValue)

	return nil
}

func sortedEntryNames(
	entries Entries,
) []string {
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

	return names
}

func dartChildClassName(
	parentClassName string,
	entryName string,
) string {
	parentName := strings.TrimSuffix(
		parentClassName,
		"Localization",
	)

	return parentName +
		toPascalCase(entryName) +
		"Localization"
}

func toPascalCase(
	value string,
) string {
	parts := strings.FieldsFunc(
		value,
		func(character rune) bool {
			return character == '-' ||
				character == '_' ||
				unicode.IsSpace(character)
		},
	)

	var result strings.Builder

	for _, part := range parts {
		characters := []rune(part)

		if len(characters) == 0 {
			continue
		}

		characters[0] = unicode.ToUpper(
			characters[0],
		)

		result.WriteString(
			string(characters),
		)
	}

	return result.String()
}
