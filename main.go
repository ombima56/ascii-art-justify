package main

import (
	"fmt"
	"os"
	"strings"

	ascii "asci-art/banner"
)

func main() {
	// Check for the correct number of arguments.
	if len(os.Args) < 2 || len(os.Args) > 3 {
		fmt.Println("Usage: go run . [STRING] [BANNER]")
		fmt.Println("EX: go run . something standard")
		os.Exit(1)
	}

	input := os.Args[1]

	input = strings.Replace(input, "\\n", "\n", -1)
	//  Removing the non-printable characters in the input string.
	input = ascii.HandleSpecialCase(input)

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
			ascii.PrintBanner(word)
		}
	}
}
