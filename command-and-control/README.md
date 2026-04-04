# CodeCrafters Suite

## What the Program Does

The CodeCrafters Suite is a terminal-based program that combines three 
powerful tools into one interactive experience. Built entirely in Go using 
only the standard library, the suite gives you access to a calculator, a 
base converter, and a string transformer, all from a single entry point. 
The program runs interactively, stays alive between operations, and returns 
you to the main menu after every tool session.

---

## How to Run

Make sure you have Go installed on your machine. Then:
```bash
cd combined-program
go run .
```

---

## Main Menu

When the program starts you will see:
=============================
CODECRAFTERS SUITE

Calculator
Base Converter
String Transformer
Exit


Pick a number to enter that tool. Type `exit` inside any tool to return 
to the main menu at any time.

---

## Tool 1 — Calculator

Performs arithmetic operations directly from the command line.

### Commands

| Command | What it does | Example |
|---|---|---|
| `add <a> <b>` | Adds two numbers | `add 5 3` |
| `sub <a> <b>` | Subtracts two numbers | `sub 10 4` |
| `mul <a> <b>` | Multiplies two numbers | `mul 6 7` |
| `div <a> <b>` | Divides two numbers | `div 20 4` |
| `pow <a> <b>` | Raises a to the power of b | `pow 2 8` |
| `rem <a> <b>` | Remainder of a divided by b | `rem 10 3` |
| `history` | Shows last 5 calculations | `history` |
| `last` | Shows the most recent result | `last` |
| `exit` | Returns to main menu | `exit` |

### Example Session

add 5 3
'5' + '3' = '8'


pow 2 8
'2' raised to the power of '8' = '256'


div 20 4
'20' / '4' = '5.00'


history
─── Last 5 Calculations ───


'5' + '3' = '8'
'2' raised to the power of '8' = '256'
'20' / '4' = '5.00'


---

## Tool 2 — Base Converter

Converts numbers between different number bases.

### Commands

| Command | What it does | Example |
|---|---|---|
| `base dec <number> <targetBase>` | Decimal to any base | `base dec 255 16` |
| `base hex <hexNumber>` | Hex to decimal | `base hex ff` |
| `base bin <binaryNumber>` | Binary to decimal | `base bin 1010` |
| `exit` | Returns to main menu | `exit` |

### Example Session

base dec 255 16
255 converted to base 16 is: ff


base dec 10 2
10 converted to base 2 is: 1010


base hex ff
ff in hex = 255 in decimal


base bin 1010
1010 in binary = 10 in decimal


---

## Tool 3 — String Transformer

Transforms text input in various ways.

### Commands

| Command | What it does |
|---|---|
| `upper` | Converts text to UPPERCASE |
| `lower` | Converts text to lowercase |
| `cap` | Capitalises the first letter of every word |
| `title` | Title case — small connector words stay lowercase |
| `snake` | Converts to snake_case |
| `reverse` | Reverses each word individually |
| `palindrome` | Checks if the text is a palindrome |
| `history` | Shows last 5 operations |
| `last` | Shows the most recent operation |
| `exit` | Returns to main menu |

### Example Session

upper
Enter text
sentinel is watching
'sentinel is watching' <-> uppercase = SENTINEL IS WATCHING


snake
Enter text
Operation Gopher Protocol!
'Operation Gopher Protocol!' <-> snake = operation_gopher_protocol


palindrome
Enter text
never odd or even
'never odd or even' <-> isPalindrome = "never odd or even" is a palindrome!


history
─── Last 5 operations ───


'sentinel is watching' <-> uppercase = SENTINEL IS WATCHING
'Operation Gopher Protocol!' <-> snake = operation_gopher_protocol
'never odd or even' <-> isPalindrome = "never odd or even" is a palindrome!


---

## Edge Cases Handled

- Empty input is caught before any operation runs
- Invalid commands are rejected with a clear error message
- Division and remainder by zero are blocked
- Non-numeric input to the calculator is caught and rejected
- Invalid hex and binary characters are caught and rejected
- Extra spaces around input are trimmed automatically
- History shows whatever is available even if less than 5 entries
- Typing exit in any tool returns cleanly to the main menu

---

## Project Structure
combined-program/
├── main.go                 ← entry point and main menu
├── calculator.go           ← calculator logic
├── base_converter.go       ← base conversion logic
└── string_converter.go     ← string transformation logic

---

## Author

- **Name:** [Your Name]
- **Squad:** [Your Squad Name]
- **Program:** Learn2Earn — CodeCrafters