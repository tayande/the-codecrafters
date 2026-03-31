### THE NUMBER BASE CONVERTER 
The Base Converter is a terminal-based number base conversion tool written in Go using the standard library only.
The program runs in a loop and keeps going until you decide to quit. From the main menu you type 1 to enter the converter, then select an operation by typing 1 for hexadecimal to decimal, 2 for binary to decimal, or 3 for decimal to any base of your choice. You are then prompted to enter the number you want to convert and the program displays the result immediately.
The program handles all bad inputs perfectly. Invalid characters for a given base, non-numeric decimal input, and out of range bases are all caught with clear error messages. Nothing you type will ever crash the program.
Type quit at any time to exit cleanly.
#### To run it:
- cd thecodecrafterthon-day2
- go run main.go
