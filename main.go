package main

import (
	Ascii "ascii/banner"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Define the alignment flag
	alignment := flag.String("align", "left", "Alignment of the text: left, right, center, or justify")

	// Parse the flags
	flag.Parse()

	// Validate the alignment flag
	if *alignment != "left" && *alignment != "right" && *alignment != "center" && *alignment != "justify" {
		printUsage()
		return
	}

	// Check for the correct number of arguments.
	if len(flag.Args()) < 1 || len(flag.Args()) > 2 {
		printUsage()
		return
	}

	input := flag.Args()[0]

	if len(flag.Args()) == 2 {
		banner := flag.Args()[1]
		os.Args = append(os.Args[:2], banner) // Pass banner file if provided
	} else {
		os.Args = append(os.Args[:2], "standard")
	}

	input = strings.Replace(input, "\\n", "\n", -1)
	// Removing the non-printable characters in the input string.
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
			// Print the banner for non-empty strings with the specified alignment.
			Ascii.PrintBanner(word, *alignment)
		}
	}
}

// printUsage prints the usage message
func printUsage() {
	fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]")
	fmt.Println("Example: go run . --align=right something standard")
}
