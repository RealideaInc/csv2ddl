package main

import "strings"

func ParseBool(value string) bool {
	lower := strings.ToLower(value)
	return lower == "true" || lower == "1" || lower == "yes" || lower == "y"
}
