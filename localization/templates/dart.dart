{{.Classes}}

enum Language {
{{.Languages}}
}

final localizations = <Language, Localization>{
{{.Localizations}}
};

Localization createLocalization(Language language) {
  return localizations[language]!;
}
