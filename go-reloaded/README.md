# Sentinel Reborn

A Go text processing program that reads an input file, applies a series
of intelligent transformations, and writes the result to an output file.

## How to run
. go run . input.txt output.txt
## How to test
. go test -v

## Transformations

**Case conversion**
Words followed by `(up)`, `(low)`, or `(cap)` are converted to uppercase,
lowercase, or title case respectively. A number can be added to affect
multiple words — for example `(cap, 3)` capitalizes the three words before it.

**Number base conversion**
Numbers followed by `(hex)` are converted from hexadecimal to decimal.
Numbers followed by `(bin)` are converted from binary to decimal.

**Article correction**
`a` is changed to `an` before words starting with a vowel sound, and
`an` is changed to `a` before words starting with a consonant.

**Punctuation spacing**
Spaces before punctuation marks are removed. A single space is ensured
after every punctuation mark.

**Quote formatting**
Single quotes are trimmed of any inner spacing so `' hello '` becomes
`'hello'`.

## My contribution

Since I am built this without a squad and i am just alone on this project, I literally did everything by myself.

## What I found hardest

The hardest part of the project was trying to use the regexp package for the punctuations function.

## What I understand now that I did not understand this before

Alot can change in just a small time. It may not seem as if its yielding the right results we expect but if you keep showing up and doing your best, all will just click and look simple with time.

## Appreciation to mr. X
What can I say really, but thank you so much for being the santa of my coding journey. I appreciate you in no small measure
