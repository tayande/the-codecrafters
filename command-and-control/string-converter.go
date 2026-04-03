package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

type HistoryEntry struct {
	command string
	input   string
	output  string
}

var history []HistoryEntry

func saveHistory(command, input, output string) {
	entry := HistoryEntry{command: command, input: input, output: output}
	history = append(history, entry)
	if len(history) > 5 {
		history = history[len(history)-5:]
	}
}

func showHistory() {
	if len(history) == 0 {
		fmt.Println("No operations yet.")
		return
	}
	fmt.Println("─── Last operations ───")
	for i, entry := range history {
		fmt.Printf("%d. [%s] \"%s\" → \"%s\"\n", i+1, entry.command, entry.input, entry.output)
	}
	fmt.Println("───────────────────────")
}

func upperCase(sentence string) string {
	if len(sentence) == 0 {
		fmt.Println("Please enter a word or sentence")
	}
	return strings.ToUpper(sentence)
}

func lowerCase(sentence string) string {
	if len(sentence) == 0 {
		fmt.Println("Please enter a word or sentence")
	}
	return strings.ToLower(sentence)
}

func capitalize(sentence string) string {
	if len(sentence) == 0 {
		fmt.Println("Please enter a word or sentence")
	}
	return strings.Title(strings.ToLower(sentence))
}

func title(sentence string) string {
	smallWords := map[string]bool{
		"a": true, "an": true, "the": true, "and": true,
		"but": true, "or": true, "for": true, "nor": true,
		"on": true, "at": true, "to": true, "by": true,
		"in": true, "of": true, "up": true, "is": true, "it": true,
	}
	if len(sentence) == 0 {
		fmt.Println("Please enter a word or sentence")
		return ""
	}
	words := strings.Fields(sentence)
	result := []string{}

	for i, word := range words {
		lowered := strings.ToLower(word)
		if i == 0 || !smallWords[lowered] {
			result = append(result, strings.Title(lowered))
		} else {
			result = append(result, lowered)
		}
	}
	return strings.Join(result, " ")
}

func snake(sentence string) string {
	if len(sentence) == 0 {
		fmt.Println("Please enter a word or sentence")
		return ""
	}
	words := strings.Fields(sentence)
	cleaned := []string{}
	for _, word := range words {
		word = strings.ToLower(word)
		var kept []rune
		for _, char := range word {
			if unicode.IsLetter(char) || unicode.IsDigit(char) {
				kept = append(kept, char)
			}
		}
		if len(kept) > 0 {
			cleaned = append(cleaned, string(kept))
		}
	}
	return strings.Join(cleaned, "_")
}

func reverse(sentence string) string {
	if len(sentence) == 0 {
		fmt.Println("Please enter a word or sentence")
		return ""
	}
	words := strings.Fields(sentence)
	for i, word := range words {
		runes := []rune(word)
		for first, last := 0, len(runes)-1; first < last; first, last = first+1, last-1 {
			runes[first], runes[last] = runes[last], runes[first]
		}
		words[i] = string(runes)
	}
	return strings.Join(words, " ")
}

func isPalindrome(sentence string) string {
	cleaned := strings.ReplaceAll(strings.ToLower(sentence), " ", "")
	runes := []rune(cleaned)
	for first, last := 0, len(runes)-1; first < last; first, last = first+1, last-1 {
		if runes[first] != runes[last] {
			return fmt.Sprintf("✗ \"%s\" is not a palindrome.", sentence)
		}
	}
	return fmt.Sprintf("✦ \"%s\" is a palindrome!", sentence)
}

func RunStringConversion() {

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("select an operation with their respective values")
		fmt.Println("upper <text>")
		fmt.Println("lower <text>")
		fmt.Println("cap <text>")
		fmt.Println("title <text>")
		fmt.Println("snake <text>")
		fmt.Println("reverse <text>")
		fmt.Println("'history' | 'last' | 'exit'")
		fmt.Print("> ")	

		
	}

	

}