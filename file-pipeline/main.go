// ═══════════════════════════════════════════
// PIPELINE CONTRACT
// ───────────────────────────────────────────
// Rules (in order):
//   1. Trim leading/trailing whitespace
//   2. Remove blank lines
//   3. Replace TODO: with ✦ ACTION:
//   4. Replace CLASSIFIED: with [REDACTED]:
//   5. Flag lines longer than 80 chars with [TRUNCATED]
//   6. Convert ALL CAPS lines to Title Case
//   7. Convert all lowercase lines to UPPERCASE
// ═══════════════════════════════════════════


package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func trim(sentence string) string {
	return strings.TrimSpace(sentence)
}

func replaceTodo(sentence string) string {
	words := strings.Fields(sentence)
	for i := 0; i < len(words); i++ {
		if words[i] == "TODO:" {
			words[i] = "✦ ACTION:"
		}
	}
	return strings.Join(words, " ")
}

func replaceClassified(sentence string) string {
	words := strings.Fields(sentence)
	for i := 0; i < len(words); i++ {
		if words[i] == "CLASSIFIED:" {
			words[i] = "[REDACTED]:"
		}
	}
	return strings.Join(words, " ")
}

func flag(sentence string) string {
	if len(sentence) > 80 {
		return sentence + " [TRUNCATED]"
	}
	return sentence
}

func isAllCaps(sentence string) bool {
	if sentence == strings.ToUpper(sentence) && strings.TrimSpace(sentence) != "" {
		return true
	}
	return false
}

func convertCapsToTitle(sentence string) string {
	if isAllCaps(sentence) {
		return strings.Title(strings.ToLower(sentence))
	}
	return sentence
}

func convertLowerToUpper(sentence string) string {
	if sentence == strings.ToLower(sentence) && sentence != "" {
		return strings.ToUpper(sentence)
	}
	return sentence
}

func main() {

	if len(os.Args) != 3 {
		fmt.Println("Usage: go run . <input.txt> <output.txt>")
		return
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	if inputFile == outputFile {
		fmt.Println("Input and output cannot be the same file.")
		return
	}

	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Printf("File not found: '%v'\n", inputFile)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	data := []string{}

	for scanner.Scan() {
		line := strings.TrimRight(scanner.Text(), "\r")
		data = append(data, line)
	}

	linesRead := len(data)

	// ── THE PIPELINE LOOP ──────────────────────────
	processed := []string{}
	linesRemoved := 0

	for _, line := range data {

		line = trim(line)
		if line == "" {
			linesRemoved++
			continue
		}
		line = replaceTodo(line)
		line = replaceClassified(line)
		line = flag(line)
		line = convertCapsToTitle(line)
		line = convertLowerToUpper(line)

		processed = append(processed, line)
	}

	linesWritten := len(processed)

	// ── CHECK IF OUTPUT PATH IS A DIRECTORY ────────
	info, err := os.Stat(outputFile)
	if err == nil && info.IsDir() {
		fmt.Println("Cannot write to output: path is a directory, not a file.")
		return
	}

	// ── CREATE AND WRITE THE OUTPUT FILE ───────────
	outFile, err := os.Create(outputFile)
	if err != nil {
		fmt.Printf("Could not create output file: '%v'\n", err)
		return
	}
	defer outFile.Close()

	writer := bufio.NewWriter(outFile)

	fmt.Fprintln(writer, "SENTINEL FIELD REPORT — PROCESSED")
	fmt.Fprintln(writer, "═══════════════════════════════════")

	// Write each processed line with a number prefix
	for i, line := range processed {
		fmt.Fprintf(writer, "%03d. %s\n", i+1, line)
	}

	// Write summary block at the end of the file
	fmt.Fprintln(writer, "═══════════════════════════════════")
	fmt.Fprintln(writer, "SUMMARY")
	fmt.Fprintf(writer, "Lines read    : %d\n", linesRead)
	fmt.Fprintf(writer, "Lines written : %d\n", linesWritten)
	fmt.Fprintf(writer, "Lines removed : %d\n", linesRemoved)
	fmt.Fprintln(writer, "Rules applied : Trim, RemoveBlanks, ReplaceTodo, ReplaceClassified, FlagLong, CapsToTitle, LowerToUpper")

	err = writer.Flush()
	if err != nil {
		fmt.Printf("Failed to flush output: %v\n", err)
		return
	}

	// ── PRINT TERMINAL SUMMARY ─────────────────────
	fmt.Println()
	fmt.Println("Output written to:", outputFile)
	fmt.Println("───────────────────────────────────")
	fmt.Printf("Lines read    : %d\n", linesRead)
	fmt.Printf("Lines written : %d\n", linesWritten)
	fmt.Printf("Lines removed : %d\n", linesRemoved)
	fmt.Println("Rules applied : Trim, RemoveBlanks, ReplaceTodo, ReplaceClassified, FlagLong, CapsToTitle, LowerToUpper")
}
