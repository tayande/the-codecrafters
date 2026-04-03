# String Transformer — Operation Gopher Protocol

## What the Program Does

The String Transformer is a terminal-based intelligence text processing tool
built as part of the CodeCrafters Operation Gopher Protocol mission.
It takes corrupted text input and produces clean, correctly formatted output
through six core transformation commands. The program runs interactively —
it stays alive, takes commands, processes them, and waits for the next one
until the user decides to exit.

---

## How to Run

Make sure you have Go installed on your machine. Then:

- cd string-transformer
- go run main.go


## Commands

| Option | Command    | What it does                                              |
|--------|------------|-----------------------------------------------------------|
| 1      | upper      | Converts every letter to UPPERCASE                        |
| 2      | lower      | Converts every letter to lowercase                        |
| 3      | cap        | Capitalises the first letter of every word                |
| 4      | title      | Title case — small connector words stay lowercase         |
| 5      | snake      | Converts to snake_case — removes symbols, lowercases all  |
| 6      | reverse    | Reverses each word individually, word order stays the same|
| 7      | palindrome | Checks if the text reads the same forwards and backwards  |
| 8      | history    | Shows the last 5 transformations you ran                  |
| 9      | exit       | Shuts down the program                                    |

---

## Example Interactions

**upper**

Enter the sentence
sentinel is watching
SENTINEL IS WATCHING


**lower**

Enter the sentence
ALERT LEVEL FIVE DETECTED
alert level five detected


**cap**

Enter the sentence
director adaeze okonkwo
Director Adaeze Okonkwo


**title**

Enter the sentence
the fall of the western power grid
The Fall of the Western Power Grid


**snake**

Enter the sentence
Operation Gopher Protocol!
operation_gopher_protocol


**reverse**

Enter the sentence
Lagos Nigeria
sogaL airegiN


**palindrome**

Enter the sentence
never odd or even
✦ "never odd or even" is a palindrome!


**history**
─── Last operations ───

[upper] "sentinel is watching" → "SENTINEL IS WATCHING"
[snake] "Operation Gopher Protocol!" → "operation_gopher_protocol"
[palindrome] "never odd or even" → ✦ "never odd or even" is a palindrome!
───────────────────────


---

## Edge Cases Handled

- Unknown commands are rejected with a clear error message
- Empty input is caught before any transformation runs
- Extra spaces around input are trimmed automatically
- Numbers and symbols pass through transformations without crashing
- Single word input works on all commands
- All-whitespace input is treated as empty
- Commands are case-insensitive
- Empty lines do not crash the program
- Title case always capitalises the first word even if it is a small word
- History only stores the last 5 operations and drops the oldest automatically

---

## Author

- **Name:** [Ayande David Terngu]
- **Squad:** [solo]
- **Program:** The string converter
- **Mission:** Operation Gopher Protocol
