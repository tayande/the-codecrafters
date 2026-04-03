package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)
func addition(a, b int) int  {
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
	
func runCalculator() {

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

		
		if broken[0] == "history" {
			if len(history) == 0 {
    			fmt.Println("No calculations yet.")
    			continue
			}
			fmt.Println("─── Last 5 Calculations ───")
			for i, entry := range history[max(0, len(history)-5):] {
 			   fmt.Printf("%d. %s", i+1, entry)
			}
 		}
		if broken[0] == "last" {
			if len(history) < 1 {
				fmt.Println("Do one more calculation before you can see the last one")
				continue
			}
			fmt.Println(history[len(history)-1:])
			continue
		}
		
		if len(broken) < 3 {
    		fmt.Printf("Usage:\ncommand <a> <b>\n")
   			 continue
		}
		
		switch newCommand {

		case "add":
			if len(broken[1]) == 0 || len(broken[2]) == 0 {
				fmt.Println("You need to provide all inputs")
				continue
			}
			value1, err1 := strconv.Atoi(broken[1])
			value2, err2 := strconv.Atoi(broken[2])
			if err1 != nil || err2 != nil {
				fmt.Println("Invalid input, use only unmbers")	
				continue
			}
			
			recentResult := addition(value1, value2)
			history = append(history, fmt.Sprintf("'%d' + '%d' = '%d'\n", value1, value2, recentResult))
			fmt.Printf("'%d' + '%d' = '%d'\n", value1, value2, recentResult)
			continue
		case "sub":

			if len(broken[1]) == 0 || len(broken[2]) == 0 {
				fmt.Println("You need to provide all inputs")
				continue
			}
			value1, err1 := strconv.Atoi(broken[1])
			value2, err2 := strconv.Atoi(broken[2])
			if err1 != nil || err2 != nil {
				fmt.Println("Invalid input, use only unmbers")	
				continue
			}
			recentResult := subtraction(value1, value2)
			history = append(history, fmt.Sprintf("'%d' - '%d' = '%d'\n", value1, value2, recentResult))
			fmt.Printf("'%d' - '%d' = '%d'\n", value1, value2, recentResult)
			continue
		case "mul":

			if len(broken[1]) == 0 || len(broken[2]) == 0 {
				fmt.Println("You need to provide all inputs")
				continue
			}
			value1, err1 := strconv.Atoi(broken[1])
			value2, err2 := strconv.Atoi(broken[2])
			if err1 != nil || err2 != nil {
				fmt.Println("Invalid input, use only unmbers")	
				continue
			}
			recentResult := multiplication(value1, value2)
			history = append(history, fmt.Sprintf("'%d' * '%d' = '%d'\n", value1, value2, recentResult))
			fmt.Printf("'%d' * '%d' = '%d'\n", value1, value2, recentResult)
			continue
		case "div":

			if len(broken[1]) == 0 || len(broken[2]) == 0 {
				fmt.Println("You need to provide all inputs")
				continue
			}
			value1, err1 := strconv.Atoi(broken[1])
			value2, err2 := strconv.Atoi(broken[2])
			if err1 != nil || err2 != nil {
				fmt.Println("Invalid input, use only unmbers")	
				continue
			}
			recentResult := division(value1, value2)
			history = append(history, fmt.Sprintf("'%d' / '%d' = '%.2f'\n", value1, value2, recentResult))
			fmt.Printf("'%d' / '%d' = '%.2f'\n", value1, value2, recentResult)
			continue

		case "pow":
			if len(broken[1]) == 0 || len(broken[2]) == 0 {
				fmt.Println("You need to provide all inputs")
				continue
			}
			value1, err1 := strconv.Atoi(broken[1])
			value2, err2 := strconv.Atoi(broken[2])
			if err1 != nil || err2 != nil {
				fmt.Println("Invalid input, use only unmbers")	
				continue
			}
			recentResult := power(value1, value2)
			history = append(history, fmt.Sprintf("'%d' raised to the power of '%d' = '%d'\n", value1, value2, recentResult))
			fmt.Printf("'%d' raised to the power of '%d' = '%d'\n", value1, value2, recentResult)
			continue

		case "rem":
			if len(broken[1]) == 0 || len(broken[2]) == 0 {
				fmt.Println("You need to provide all inputs")
				continue
			}
			if broken[1] == "0" || broken[2] == "0" {
				fmt.Println("Can't perform remainder operation with '0'")
				continue
			}
			value1, err1 := strconv.Atoi(broken[1])
			value2, err2 := strconv.Atoi(broken[2])
			if err1 != nil || err2 != nil {
				fmt.Println("Invalid input, use only unmbers")	
				continue
			}
			recentResult := remainder(value1, value2)
			history = append(history, fmt.Sprintf("'%d' modulus '%d' = '%d'\n", value1, value2, recentResult))
			fmt.Printf("'%d' modulus '%d' = '%d'\n", value1, value2, recentResult)
			continue

		case "exit":
    		fmt.Println("Returning to main menu...")
    		return
		default:
			fmt.Println("Invalid command")
			continue

		}
	}	
}