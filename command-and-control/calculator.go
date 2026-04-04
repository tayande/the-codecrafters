package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func addition(a, b int) int {
	return a + b
}

func multiplication(a, b int) int {
	return a * b
}

func subtraction(a, b int) int {
	return a - b
}

func division(a, b int) float64 {
	return float64(a) / float64(b)
}

func remainder(a, b int) int {
	return a % b
}

func power(a, b int) int {
	result := math.Pow(float64(a), float64(b))
	return int(result)
}

func RunCalculator() {
	scanner := bufio.NewScanner(os.Stdin)
	history := []string{}

	for {
		fmt.Println("select an operation with their respective values")
		fmt.Println("add <a> <b>")
		fmt.Println("sub <a> <b>")
		fmt.Println("mul <a> <b>")
		fmt.Println("div <a> <b>")
		fmt.Println("pow <a> <b>")
		fmt.Println("rem <a> <b>")
		fmt.Println("'history' | 'last' | 'exit'")
		fmt.Print("> ")

		scanner.Scan()
		command := strings.TrimSpace(scanner.Text())
		broken := strings.SplitN(command, " ", 3)
		newCommand := broken[0]

		if newCommand == "history" {
			if len(history) == 0 {
				fmt.Println("No calculations yet.")
				continue
			}
			fmt.Println("─── Last 5 Calculations ───")
			start := len(history) - 5
			if start < 0 {
				start = 0
			}
			for i, entry := range history[start:] {
				fmt.Printf("%d. %s", i+1, entry)
			}
			continue 
		}

		if newCommand == "last" {
			if len(history) == 0 {
				fmt.Println("You have not done any calculation yet")
				continue
			}
			fmt.Print("Last: ", history[len(history)-1])
			continue
		}

		if newCommand == "exit" {
			fmt.Println("Returning to main menu...")
			return 
		}

		if len(broken) < 3 {
			fmt.Printf("Usage: command <a> <b>\n")
			continue
		}

		switch newCommand {
		case "add":
			value1, err1 := strconv.Atoi(broken[1])
			value2, err2 := strconv.Atoi(broken[2])
			if err1 != nil || err2 != nil {
				fmt.Println("Invalid input, use only numbers")
				continue
			}
			recentResult := addition(value1, value2)
			entry := fmt.Sprintf("'%d' + '%d' = '%d'\n", value1, value2, recentResult)
			history = append(history, entry)
			fmt.Print(entry)

		case "sub":
			value1, err1 := strconv.Atoi(broken[1])
			value2, err2 := strconv.Atoi(broken[2])
			if err1 != nil || err2 != nil {
				fmt.Println("Invalid input, use only numbers")
				continue
			}
			recentResult := subtraction(value1, value2)
			entry := fmt.Sprintf("'%d' - '%d' = '%d'\n", value1, value2, recentResult)
			history = append(history, entry)
			fmt.Print(entry)

		case "mul":
			value1, err1 := strconv.Atoi(broken[1])
			value2, err2 := strconv.Atoi(broken[2])
			if err1 != nil || err2 != nil {
				fmt.Println("Invalid input, use only numbers")
				continue
			}
			recentResult := multiplication(value1, value2)
			entry := fmt.Sprintf("'%d' * '%d' = '%d'\n", value1, value2, recentResult)
			history = append(history, entry)
			fmt.Print(entry)

		case "div":
			value1, err1 := strconv.Atoi(broken[1])
			value2, err2 := strconv.Atoi(broken[2])
			if err1 != nil || err2 != nil {
				fmt.Println("Invalid input, use only numbers")
				continue
			}
			if value2 == 0 {
				fmt.Println("Cannot divide by zero")
				continue
			}
			recentResult := division(value1, value2)
			entry := fmt.Sprintf("'%d' / '%d' = '%.2f'\n", value1, value2, recentResult)
			history = append(history, entry)
			fmt.Print(entry)

		case "pow":
			value1, err1 := strconv.Atoi(broken[1])
			value2, err2 := strconv.Atoi(broken[2])
			if err1 != nil || err2 != nil {
				fmt.Println("Invalid input, use only numbers")
				continue
			}
			recentResult := power(value1, value2)
			entry := fmt.Sprintf("'%d' raised to the power of '%d' = '%d'\n", value1, value2, recentResult)
			history = append(history, entry)
			fmt.Print(entry)

		case "rem":
			value1, err1 := strconv.Atoi(broken[1])
			value2, err2 := strconv.Atoi(broken[2])
			if err1 != nil || err2 != nil {
				fmt.Println("Invalid input, use only numbers")
				continue
			}
			if value2 == 0 {
				fmt.Println("Can't perform remainder operation with '0'")
				continue
			}
			recentResult := remainder(value1, value2)
			entry := fmt.Sprintf("'%d' modulus '%d' = '%d'\n", value1, value2, recentResult)
			history = append(history, entry)
			fmt.Print(entry)

		default:
			fmt.Println("Invalid command")
		}
	}
}