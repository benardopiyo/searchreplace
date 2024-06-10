# Search and Replace

This program takes three arguments: a string, a target letter to replace, and a replacement letter. It replaces all instances of the target letter in the string with the replacement letter. If the number of arguments is not exactly three, the program displays nothing. If the target letter is not found in the string, the original string is returned followed by a newline (\n).

## Instructions

* The program processes exactly three arguments:
    1. The input string.
    2. The target letter to be replaced.
    3. The replacement letter.

* If the number of arguments is not exactly three, the program will display nothing.

* If the target letter is not found in the input string, the program will return the original string followed by a newline (\n).

## Usage

The program is written in Go. To run it, use the go run command followed by the program and the three arguments. Below are some examples:

``` bash

$ go run . "hella there" "a" "o"
hello there

$ go run . "hallo thara" "a" "e"
hello there

$ go run . "abcd" "z" "l"
abcd

$ go run . "something" "a" "o" "b" "c"
$
```

* First example, the input string "hella there" has all instances of "a" replaced with "o", resulting in "hello there".

* Second example, the input string "hallo thara" has all instances of "a" replaced with "e", resulting in "hello there".

* Third example, the input string "abcd" does not contain the letter "z", so it is returned unchanged.

* Fourth example, the program is provided with more than three arguments, so it displays nothing.
