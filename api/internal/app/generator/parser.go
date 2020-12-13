package generator

import (
	"regexp"
	"strings"
)

// Clean converts to lowercase, removes spaces and invalid characters of a string.
func Clean(text string) string {
	return strings.ToLower(regexp.MustCompile(`\W+`).ReplaceAllString(text, ""))
}

// Prune removes extra spaces and invalid characters of a string
func Prune(text string) string {
	return strings.TrimSpace(regexp.MustCompile(`\W+`).ReplaceAllString(text, " "))
}

// PruneAndLowercase converts to lowercase, removes extra spaces and invalid characters of a string
func PruneAndLowercase(text string) string {
	return strings.ToLower(Prune(text))
}

// Hyphenate converts to lowercase, replaces spaces with hyphens and removes invalid characters of a string
func Hyphenate(text string) string {
	return strings.Join(Words(text), "-")
}

// Words finds spaces to split string into words
func Words(text string) []string {
	text = PruneAndLowercase(text)
	return regexp.MustCompile(`\W+`).Split(text, -1)
}
