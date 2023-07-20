package core

import (
	"regexp"
	"strings"
)

var (
	spaceRegexp = regexp.MustCompile(`\s+`)
	punctRegexp = regexp.MustCompile(`([.,!?])(\S)`)
)

// CleanupDescription replaces multiple spaces with a single space
// adds space after a full stop if it's missing
func CleanupDescription(description string) string {
	// Replace NBSP with regular space
	description = strings.ReplaceAll(description, "\u00A0", " ")

	description = spaceRegexp.ReplaceAllString(description, " ")

	description = punctRegexp.ReplaceAllString(description, "$1 $2")

	return description
}
