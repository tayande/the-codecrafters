// package main

// import (
// 	"fmt"
// 	"strconv"
// 	"strings"
// )
// func hexToDec(hexstr string) (int64, error) {
// 	value, err := strconv.ParseInt(hexstr, 16, 64)
// 	if err != nil {
// 		fmt.Println("The input you used is not valid")
// 	}
// 	return value, nil
// }
// func binToDec(binstr string) (int64, error) {
// 	value, err := strconv.ParseInt(binstr, 2, 64)
// 	if err != nil {
// 		fmt.Println("The input you used is not valid")
// 	}
// 	return value, nil
// }

// func decToAny(num int64, base int) (string, error) {

// 	if base < 2 || base > 36 {
// 		return "", fmt.Errorf("The input you used is not valid\n")
// 	}
// 	 return strconv.FormatInt(num, base), nil
// }
// func main() {

// 	fmt.Println("========================================")
// 	fmt.Println("WELCOME TO THE BEST BASE CONVERTER EVER")
// 	fmt.Println()

// 	for {
// 		fmt.Println("Select '1' to continue and 'quit' to exit the program")
// 		fmt.Print("> ")

// 		var reply string
// 		fmt.Scan(&reply)

// 		if reply == "quit" {
// 			fmt.Println("Goodbye!")
// 			return
// 		}
// 		if reply != "1" {
// 			fmt.Println("Invalid input. please enter '1' to continue or 'quit' to exit")
// 			fmt.Println()
// 			continue

// 		}
// 		for {

// 			fmt.Println()
// 			fmt.Println("select the operation you want to perform")
// 			fmt.Println("1. Hexadecimal to decimal")
// 			fmt.Println("2. Binary to decimal")
// 			fmt.Println("3. Decimal to both Hexadecimal and Binary")
// 			fmt.Println("Enter 'quit' to exit and 'back' to return")
// 			fmt.Print("> ")

// 			var operation string
// 			fmt.Scan(&operation)

// 			if operation == "quit" {
// 				fmt.Println("Goodbye!")
// 				return
// 			}
// 			if operation == "back" {
// 				break
// 			}

// 			if operation != "1" && operation != "2" && operation != "3" {
// 				fmt.Println("Invalid operation, please enter '1' or '2' or '3'")
// 				continue
// 			}

// 			switch operation {
// 			case "1":
// 				var num string
// 				fmt.Println("Enter the number you want to convert")
// 				fmt.Scan(&num)

// 				if num == "" {
// 					fmt.Println("You must input a number")
// 					continue
// 				}

// 				num = strings.ToUpper(num)
// 				if !strings.ContainsAny(num, "0123456789ABCDEF") {
// 					fmt.Println("Invalid hexadecimal input")
// 					continue
// 				}

// 				var base int
// 				fmt.Println("Enter the base you want to convert from")
// 				fmt.Print("> ")
// 				fmt.Scan(&base)

// 				if base != 16 {
// 					fmt.Println("The base you entered is invalid for hexadecimal conversion")
// 					continue
// 				}

// 				if base == 0 {
// 					fmt.Println("You must enter the base from which you want to convert")
// 				}

// 				value, err := hexToDec(num)
// 				if err != nil {
// 					fmt.Println(err)
// 					continue
// 				}
// 				fmt.Printf("%q converted to decimal is:\n", num)
// 				fmt.Println(value)
// 			case "2":
// 				var num string
// 				fmt.Println("Enter the number you want to convert")
// 				fmt.Scan(&num)

// 				if num == "" {
// 					fmt.Println("You must input a number")
// 					continue
// 				}
// 				if !strings.ContainsAny(num, "01")  {
// 					fmt.Println("The input you used is not valid for binary conversion")
// 					continue
// 				}

// 				var base int
// 				fmt.Println("Enter the base you want to convert from")
// 				fmt.Print("> ")
// 				fmt.Scan(&base)

// 				if base != 2 {
// 					fmt.Println("The base you entered is invalid for hexadecimal conversion")
// 					continue
// 				}

// 				if base == 0 {
// 					fmt.Println("You must enter the base from which you want to convert")
// 				}

// 				value, err := binToDec(num)
// 				if err != nil {
// 					fmt.Println(err)
// 					continue
// 				}
// 				fmt.Printf("%q converted to decimal is:\n", num)
// 				fmt.Println(value)
// 			case "3":
// 				var a string
// 				fmt.Println("Enter the decimal number you want to convert")
// 				fmt.Scan(&a)

// 				if a == "" {
// 					fmt.Println("You must enter a number")
// 					continue
// 				}

// 				num, err := strconv.Atoi(a)
// 				if err != nil {
// 					fmt.Println("The input you used is invalid. You must use numbers only")
// 					continue
// 				}

// 				var base int
// 				fmt.Println("Enter the base you want to convert from")
// 				fmt.Print("> ")
// 				fmt.Scan(&base)

// 				if base < 2 || base > 36 {
// 					fmt.Println("The base you entered is not valid")
// 					continue

// 				}

// 				value, err := decToAny(int64(num), base)
// 				if err != nil {
// 					fmt.Println(err)
// 					continue
// 				}
// 				fmt.Printf("%s in base %d is:\n", a, base)
// 				fmt.Println(value)
// 			}
// 		}
// 	}
// }

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func hexToDec(hexstr string) (int64, error) {
	value, err := strconv.ParseInt(hexstr, 16, 64)
	if err != nil {
		fmt.Println("The input you used is not valid")
	}
	return value, nil
}
func binToDec(binstr string) (int64, error) {
	value, err := strconv.ParseInt(binstr, 2, 64)
	if err != nil {
		fmt.Println("The input you used is not valid")
	}
	return value, nil
}

func decToAny(num int64, base int) (string, error) {

	if base < 2 || base > 36 {
		return "", fmt.Errorf("The input you used is not valid\n")
	}
	 return strconv.FormatInt(num, base), nil
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)



	for {

		fmt.Println("select an operation with their respective values")
		fmt.Println("base dec <number> ")
		fmt.Println("base hex <number>")
		fmt.Println("base bin <number>")
		fmt.Println("exit")	
		fmt.Print("> ")	
		
		scanner.Scan()

		broken := strings.TrimSpace(scanner.Text())
		if broken == "exit" {
			fmt.Println("Goodbye...")
			return
		}
		
		command := strings.SplitN(broken, " ", 3)

		if len(command) < 3 {
    		fmt.Printf("invalid input\n")
   			 continue
		}

		base := command[0]
		op := command[1]
		num := command[2]

		if base != "base" {
			fmt.Println("Invalid input")
			continue
		}


		switch op {

		case "dec":
			value1, err1 := strconv.Atoi(base)
			value2, err2 := strconv.Atoi(num)
			if err1 != nil || err2 != nil {
				fmt.Println("Invalid input")
				continue
			}
			if value1 < 2 && value1 > 36 {
				fmt.Println("Invalid base")
				continue
			}
			value, err := decToAny(int64(value2), value1)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Printf("%s in base %d is:\n", value1, value2)
			fmt.Println(value)

		case "hex":
			if base != "16" {
				fmt.Println("Invalid base for hexadecimal conversion")
				continue
			}
			if !strings.ContainsAny(num, "0123456789ABCDEF") {
				fmt.Println("Invalid  input")
				continue
			}
			result, err := hexToDec(num)
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Println(result)

		case "bin":
			if base != "2" {
				fmt.Println("Invalid base for binary conversion")
				continue
			}
			if !strings.ContainsAny(num, "01")  {
				fmt.Println("The input you used is invalid for binary conversion")
				continue
			}
			result, err := binToDec(num)
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Println(result)

		// case "exit":
		// 	fmt.Println("Goodbye...")
		// 	return

		default:
			fmt.Println("Invalid operation")
			continue
		}
	}
}