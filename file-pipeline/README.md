# File Pipeline — Operation Gopher Protocol

## Overview

The File Pipeline is a command-line text processing program built in Go.
It reads a raw field report from an input file, runs every line through
a series of transformation rules, and writes a clean processed report
to an output file. A terminal summary is printed after every run.

This project was built as part of the CodeCrafters program
Operation Gopher Protocol, Mission: The File Pipeline.

---

## How to Run

```bash
go run . input.txt output.txt
```

- `input.txt` — the raw field report you want to process
- `output.txt` — where the cleaned report will be written

---

## Transformation Rules (Applied in This Order)

Every line in the input file passes through all rules in sequence.

| # | Rule | What It Does |
|---|------|--------------|
| 1 | Trim Whitespace | Removes all leading and trailing spaces from every line |
| 2 | Remove Blank Lines | Any line that is empty after trimming is dropped entirely |
| 3 | Replace TODO: | Replaces `TODO:` with `✦ ACTION:` |
| 4 | Replace CLASSIFIED: | Replaces `CLASSIFIED:` with `[REDACTED]:` |
| 5 | Flag Long Lines | Appends `[TRUNCATED]` to any line longer than 80 characters |
| 6 | ALL CAPS to Title Case | Converts fully capitalised lines to Title Case |
| 7 | Lowercase to UPPERCASE | Converts fully lowercase lines to UPPERCASE |

---

## Output Format

The output file is structured as follows:
SENTINEL FIELD REPORT — PROCESSED
═══════════════════════════════════
001. First processed line
002. Second processed line
...
═══════════════════════════════════
SUMMARY
Lines read    : 20
Lines written : 17
Lines removed : 3
Rules applied : Trim, RemoveBlanks, ReplaceTodo, ReplaceClassified, FlagLong, CapsToTitle, LowerToUpper

---

## Terminal Summary

After every run the program prints a summary to the terminal:
✔ Output written to: output.txt
───────────────────────────────────
✦ Lines read    : 20
✦ Lines written : 17
✦ Lines removed : 3
✦ Rules applied : Trim, RemoveBlanks, ReplaceTodo, ReplaceClassified, FlagLong, CapsToTitle, LowerToUpper

---

## What Happens When Rules Overlap

Because all rules fire in order on every line, a single line can
trigger multiple rules at once. Here is how each overlap is handled:

**TODO: + long line**
Rule 3 fires first and replaces `TODO:` with `✦ ACTION:`.
Rule 5 then checks the updated line length and appends `[TRUNCATED]`
if it exceeds 80 characters.

**ALL CAPS + CLASSIFIED:**
Rule 4 fires first and replaces `CLASSIFIED:` with `[REDACTED]:`.
Because the line is now mixed case, Rule 6 no longer treats it as
ALL CAPS and skips it. Order matters here.

**Blank line after trimming**
Rule 1 trims the line. If it becomes empty, Rule 2 removes it
immediately. No further rules run on that line.

**ALL CAPS + longer than 80 characters**
Rule 5 appends `[TRUNCATED]` first. Rule 6 then checks — because
the line now contains lowercase letters in `[TRUNCATED]`, it is no
longer ALL CAPS and is skipped by Rule 6.

---

## Edge Cases Handled

| Situation | Behaviour |
|---|---|
| Input file does not exist | Prints error and exits cleanly |
| No arguments provided | Prints usage instructions and exits |
| Same file used as input and output | Detected immediately, rejected with error |
| Input file is completely empty | Writes empty output, prints warning |
| All lines are blank or dashes | Output contains only header and summary |
| Output path is a directory | Detected, rejected with clear error |
| Line is only whitespace | Trimmed to empty, counted as removed |
| Windows line endings (`\r\n`) | Stripped during reading, no garbled output |

---

## Example Run

**input.txt**
ALERT LEVEL FIVE DETECTED IN NORTHERN SECTOR
todo: confirm coordinates with Agent Bulus
classified: target has been relocated
the perimeter was breached at 03:47 local time

**output.txt**
SENTINEL FIELD REPORT — PROCESSED
═══════════════════════════════════
001. Alert Level Five Detected In Northern Sector
002. ✦ ACTION: confirm coordinates with Agent Bulus
003. [REDACTED]: target has been relocated
004. THE PERIMETER WAS BREACHED AT 03:47 LOCAL TIME
═══════════════════════════════════
SUMMARY
Lines read    : 4
Lines written : 4
Lines removed : 0
Rules applied : Trim, RemoveBlanks, ReplaceTodo, ReplaceClassified, FlagLong, CapsToTitle, LowerToUpper

---

## Author

**Ayande**
CodeCrafters Program — Operation Gopher Protocol