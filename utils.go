package pronoundb

import "strings"

func caps(inp string) string {
	if len(inp) == 0 {
		return inp
	}

	return strings.ToUpper(inp[:1]) + inp[1:]
}