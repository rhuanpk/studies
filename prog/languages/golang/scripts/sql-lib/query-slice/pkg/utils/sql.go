package utils

import (
	"fmt"
	"strings"
)

// BuildQuerySlice build the query with the placeholders.
func BuildQuerySlice(query string, args ...any) string {
	if strings.Contains(query, "???") {
		return fmt.Sprintf(
			strings.ReplaceAll(query, "???", "%v"),
			strings.TrimSuffix(
				strings.Repeat(
					"?,",
					len(args),
				),
				",",
			),
		)
	}
	return ""
}
