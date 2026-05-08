package main

import (
	"strings"
)

func GenerateArt(input string, banner map[rune][]string) string {
	if input == "" {
		return ""
	}

	parts := SplitInput(input)

	var result strings.Builder

	for i, part := range parts {
		if part == "" {
			if i < len(parts)-1 {
				result.WriteString("\n")
			}
		} else {
			rows := RenderLine(part, banner)
			for _, row := range rows {
				result.WriteString(row + "\n")
			}
		}
	}
	return result.String()
}
