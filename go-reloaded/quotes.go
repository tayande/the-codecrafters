package main

import (
	"strings"
)
func quotes(sentence string) string {
	words := strings.Split(sentence, "'")

	for i := 1; i < len(words); i++ {
		words[i] = strings.TrimSpace(words[i])
	}
	return strings.Join(words, "'")
}