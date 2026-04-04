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
		return 0, fmt.Errorf("invalid hex number: %s", hexstr)
	}
	return value, nil
}

func binToDec(binstr string) (int64, error) {
	value, err := strconv.ParseInt(binstr, 2, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid binary number: %s", binstr)
	}
	return value, nil
}

func decToAny(num int64, base int) (string, error) {
	if base < 2 || base > 36 {
		return "", fmt.Errorf("base must be between 2 and 36")
	}
	return strconv.FormatInt(num, base), nil
}

func isValidHex(s string) bool {
	for _, char := range strings.ToUpper(s) {
		if !strings.ContainsRune("0123456789ABCDEF", char) {
			return false
		}
	}
	return true
}

func isValidBin(s string) bool {
	for _, char := range s {
		if char != '0' && char != '1' {
			return false
		}
	}
	return true
}

func RunBaseConversion() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("select an operation with their respective values")
		fmt.Println("base dec <number> <targetBase>  → decimal to any base")
		fmt.Println("base hex <hexNumber>             → hex to decimal")
		fmt.Println("base bin <binaryNumber>          → binary to decimal")
		fmt.Println("exit")
		fmt.Print("> ")

		scanner.Scan()
		broken := strings.TrimSpace(scanner.Text())

		if broken == "" {
			continue
		}

		if broken == "exit" {
			fmt.Println("Returning to main menu...")
			return 
		}

		command := strings.Fields(broken)

		if len(command) < 3 {
			fmt.Println("Invalid input — usage: base <op> <number>")
			continue
		}

		base := command[0]
		op := command[1]
		num := command[2]

		if base != "base" {
			fmt.Println("Command must start with 'base'")
			continue
		}

		switch op {
		
		case "dec":
			if len(command) < 4 {
				fmt.Println("Usage: base dec <number> <targetBase>")
				continue
			}
			decNum, err := strconv.ParseInt(num, 10, 64)
			if err != nil {
				fmt.Println("Invalid decimal number")
				continue
			}
			targetBase, err := strconv.Atoi(command[3])
			if err != nil {
				fmt.Println("Invalid base — must be a number")
				continue
			}
			if targetBase < 2 || targetBase > 36 {
				fmt.Println("Base must be between 2 and 36")
				continue
			}
			value, err := decToAny(decNum, targetBase)
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Printf("%d converted to base %d is: %s\n", decNum, targetBase, value)

		
		case "hex":
			if !isValidHex(num) {
				fmt.Println("Invalid hex number — use only 0-9 and A-F")
				continue
			}
			result, err := hexToDec(strings.ToUpper(num))
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Printf("%s in hex = %d in decimal\n", num, result)

		case "bin":
			if !isValidBin(num) {
				fmt.Println("Invalid binary number — use only 0 and 1")
				continue
			}
			result, err := binToDec(num)
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Printf("%s in binary = %d in decimal\n", num, result)

		default:
			fmt.Println("Invalid operation — use dec, hex or bin")
		}
	}
}
