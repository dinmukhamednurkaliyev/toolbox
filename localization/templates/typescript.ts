export const languages = {{.Languages}} as const

export type Language = keyof typeof languages

export function createLocalization(language: Language) {
  return languages[language]
}
