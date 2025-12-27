package main

import(
	"strings"
)

func cleanInput(text string) []string {
	var result []string
	lowerText := strings.ToLower(text)
	trimmedText := strings.TrimSpace(lowerText)
	textSlice := strings.Split(trimmedText, " ")
	for _, s := range textSlice {
		if s != "" && s != " " {
			result = append(result, s)
		}
	}
	return result
}