// CodeCrafters — Operation Gopher Protocol
// Module: String Transformer
// Author: [Ayande David Terngu]
// Squad:  [One man squad]


// CodeCrafters — Operation Gopher Protocol
// Module: String Transformer
// Author: [Your Name]
// Squad:  [Your Squad Name]

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

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("========================================")
	fmt.Println("WELCOME TO SENTINEL TRANSFORMER — ONLINE")
	fmt.Println()

	for {
		fmt.Println("Enter '1' to continue and 'quit' to exit")
		fmt.Print("> ")

		var firstCommand string
		fmt.Scanln(&firstCommand)
		firstCommand = strings.ToLower(firstCommand)

		if firstCommand == "quit" {
			fmt.Println("Shutting down String Transformer. Goodbye...")
			fmt.Println("============================================")
			return
		}

		if firstCommand != "1" {
			fmt.Println("Invalid command")
			continue
		}

		for {
			fmt.Println()
			fmt.Println("Select operation:")
			fmt.Println("1. upper")
			fmt.Println("2. lower")
			fmt.Println("3. cap")
			fmt.Println("4. title")
			fmt.Println("5. snake")
			fmt.Println("6. reverse")
			fmt.Println("7. palindrome check")
			fmt.Println("8. history")
			fmt.Println("9. exit")
			fmt.Println("Enter 'back' to return")
			fmt.Print("> ")

			var operation string
			fmt.Scanln(&operation)

			if operation == "back" {
				break
			}

			if operation != "1" && operation != "2" && operation != "3" &&
				operation != "4" && operation != "5" && operation != "6" &&
				operation != "7" && operation != "8" && operation != "9" {
				fmt.Println("Invalid operation")
				continue
			}

			switch operation {
			case "1":
				fmt.Println("Enter the sentence")
				fmt.Print("> ")
				scanner.Scan()
				input := strings.TrimSpace(scanner.Text())
				if len(input) == 0 {
					fmt.Println("You need to enter a sentence")
					continue
				}
				output := upperCase(input)
				fmt.Println(output)
				saveHistory("upper", input, output)

			case "2":
				fmt.Println("Enter the sentence")
				fmt.Print("> ")
				scanner.Scan()
				input := strings.TrimSpace(scanner.Text())
				if len(input) == 0 {
					fmt.Println("You need to enter a sentence")
					continue
				}
				output := lowerCase(input)
				fmt.Println(output)
				saveHistory("lower", input, output)

			case "3":
				fmt.Println("Enter the sentence")
				fmt.Print("> ")
				scanner.Scan()
				input := strings.TrimSpace(scanner.Text())
				if len(input) == 0 {
					fmt.Println("You need to enter a sentence")
					continue
				}
				output := capitalize(input)
				fmt.Println(output)
				saveHistory("cap", input, output)

			case "4":
				fmt.Println("Enter the sentence")
				fmt.Print("> ")
				scanner.Scan()
				input := strings.TrimSpace(scanner.Text())
				if len(input) == 0 {
					fmt.Println("You need to enter a sentence")
					continue
				}
				output := title(input)
				fmt.Println(output)
				saveHistory("title", input, output)

			case "5":
				fmt.Println("Enter the sentence")
				fmt.Print("> ")
				scanner.Scan()
				input := strings.TrimSpace(scanner.Text())
				if len(input) == 0 {
					fmt.Println("You need to enter a sentence")
					continue
				}
				output := snake(input)
				fmt.Println(output)
				saveHistory("snake", input, output)

			case "6":
				fmt.Println("Enter the sentence")
				fmt.Print("> ")
				scanner.Scan()
				input := strings.TrimSpace(scanner.Text())
				if len(input) == 0 {
					fmt.Println("You need to enter a sentence")
					continue
				}
				output := reverse(input)
				fmt.Println(output)
				saveHistory("reverse", input, output)

			case "7":
				fmt.Println("Enter the sentence")
				fmt.Print("> ")
				scanner.Scan()
				input := strings.TrimSpace(scanner.Text())
				if len(input) == 0 {
					fmt.Println("You need to enter a sentence")
					continue
				}
				output := isPalindrome(input)
				fmt.Println(output)
				saveHistory("palindrome", input, output)

			case "8":
				showHistory()

			case "9":
				fmt.Println("Shutting down String Transformer. Goodbye...")
				fmt.Println("============================================")
				return
			}
		}
	}
}