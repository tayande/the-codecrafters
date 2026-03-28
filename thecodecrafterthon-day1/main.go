package main

import (
	"fmt"
	"strconv"
	"strings"
)

func addition(firstNumber, secondNumber int) int {
	return firstNumber + secondNumber
}

func multiplication(firstNumber, secondNumber int) int {
	return firstNumber * secondNumber
}

func subtraction(firstNumber, secondNumber int) int {
	return firstNumber - secondNumber
}

func division(firstNumber, secondNumber int) float64 {
	return float64(firstNumber) / float64(secondNumber)
}

func readInt(prompt string) (int, error) {
	var input string
	fmt.Print(prompt)
	fmt.Scan(&input)
	num, err := strconv.Atoi(input)
	if err != nil {
		return 0, fmt.Errorf("'%s' is not a valid number. Please enter a whole number", input)
	}
	return num, nil
}

func printHelp() {
	fmt.Println("Here is how to use David's Calculator:")
	fmt.Println("  - From the main menu, type 1 to open the calculator")
	fmt.Println("  - Then select an operation by typing 1, 2, 3, or 4")
	fmt.Println("  - Enter the two numbers when prompted")
	fmt.Println("  - Type 'back' to return to the main menu")
	fmt.Println("  - Type 'quit' at any time to exit")
}

func main() {
	fmt.Println("Welcome to David's Calculator!")

	for {
		fmt.Println()
		fmt.Println("Type 1 to continue, 'help' for guidance, or 'quit' to exit:")
		fmt.Print("> ")

		var choice string
		fmt.Scanln(&choice)
		choice = strings.TrimSpace(strings.ToLower(choice))

		if choice == "quit" {
			fmt.Println("Goodbye!")
			return
		}

		if choice == "help" {
			printHelp()
			continue
		}

		if choice != "1" {
			fmt.Println("Invalid input. Type 1 to continue, 'help' for guidance, or 'quit' to exit.")
			continue
		}

		for {
			fmt.Println()
			fmt.Println("Select an operation:")
			fmt.Println("  1. Addition")
			fmt.Println("  2. Multiplication")
			fmt.Println("  3. Subtraction")
			fmt.Println("  4. Division")
			fmt.Println("Type 'back' to return to the main menu or 'quit' to exit.")
			fmt.Print("> ")

			var operation string
			fmt.Scanln(&operation)
			operation = strings.TrimSpace(strings.ToLower(operation))

			if operation == "quit" {
				fmt.Println("Goodbye!")
				return
			}

			if operation == "back" {
				break
			}

			if operation != "1" && operation != "2" && operation != "3" && operation != "4" {
				fmt.Println("Invalid option. Please type 1, 2, 3, or 4.")
				continue
			}

			num1, err := readInt("Enter first number: ")
			if err != nil {
				fmt.Println(err)
				continue
			}

			num2, err := readInt("Enter second number: ")
			if err != nil {
				fmt.Println(err)
				continue
			}

			switch operation {
			case "1":
				fmt.Printf("Result: %d + %d = %d\n", num1, num2, addition(num1, num2))
			case "2":
				fmt.Printf("Result: %d * %d = %d\n", num1, num2, multiplication(num1, num2))
			case "3":
				fmt.Printf("Result: %d - %d = %d\n", num1, num2, subtraction(num1, num2))
			case "4":
				if num2 == 0 {
					fmt.Println("Error: you cannot divide by zero.")
					continue
				}
				fmt.Printf("Result: %d / %d = %.2f\n", num1, num2, division(num1, num2))
			}
		}
	}
}