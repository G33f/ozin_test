package utils

import "strings"

// FormatQuery removes all tabs and line breaks that can be used in
// the query for more readable code
func FormatQuery(q string) string {
	return strings.ReplaceAll(strings.ReplaceAll(q, "\t", ""), "\n", " ")
}
