package main

import (
	"os"
	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 4 {
		return
	}

	input := os.Args[1]
	oldChar := os.Args[2]
	newChar := os.Args[3]

	// Check if the old character exists in the input string
	found := false
	for _, char := range input {
		if string(char) == oldChar {
			found = true
			break
		}
	}

	// If the old character is not found, print the input string followed by a newline
	if !found {
		println(input)
		return
	}

	// Replace the old character with the new character in the input string
	var output string
	for _, char := range input {
		if string(char) == oldChar {
			output += newChar
		} else {
			output += string(char)
		}
	}

	// Print the modified string followed by a newline
	println(output)
}

// println prints the given string to standard output
func println(s string) {
	for _, char := range s {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}
