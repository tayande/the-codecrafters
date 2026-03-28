## Introduction
**David's Calculator** is a simple terminal-based calculator written in Go using the standard library only.
The calculator runs in a loop, meaning it keeps going until you decide to quit. From the main menu you type 1 to enter the calculator, then select an operation by typing 1 for addition, 2 for multiplication, 3 for subtraction, or 4 for division. You are then prompted to enter two whole numbers and the result is displayed immediately.
The calculator handles all bad inputs like dividing by zero, typing letters where numbers are expected, and selecting invalid options are all caught with clear error messages. Nothing you type will ever crash the program.
Type help at any time to see usage guidance, and type quit to exit cleanly from anywhere in the program.

### To run it:
- cd thecodecrafterthon-day1
- go run main.go
