package main

import (
	Ascii "ascii/banner"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Check for the correct number of arguments.
	if len(os.Args) < 2 || len(os.Args) > 3 {
		Ascii.PrintErr()
	}

	input := os.Args[1]

	input = strings.Replace(input, "\\n", "\n", -1)
	//  Removing the non-printable characters in the input string.
	input = Ascii.HandleSpecialCase(input)

	if input == "\n" {
		fmt.Println()
		return
	} else if input == "" {
		return
	}
	// Split the input into lines based on newline characters.

	Input := strings.Split(input, "\n")

	spaceCount := 0
	// Iterate over each line of the input.
	for _, word := range Input {
		if word == "" {
			spaceCount++
			if spaceCount < len(Input) {
				fmt.Println()
			}
		} else {
			// Print the banner for non-empty strings.
			Ascii.PrintBanner(word)
		}
	}
}
