package Ascii

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// LoadBanner loads the banner characters from a file into a map
func LoadBanner(name string) map[rune]string {
	var height int                  // Tracks the current height of the character being read
	Banner := make(map[rune]string) // Map to store the banner characters
	currentChar := rune(32)
	charLine := []string{}                       // Slice to store lines of the current character
	filePath := "./bannerFiles/" + name + ".txt" // Construct the file path

	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	defer file.Close() // Ensure the file is closed after function returns

	scanner := bufio.NewScanner(file)
	scanner.Scan() // Skip the first line (assuming it's a header or empty)
	for scanner.Scan() {
		line := scanner.Text()

		if height == 8 {
			// If the character height reaches 8, store the character and reset for the next
			Banner[currentChar] = strings.Join(charLine, "\n")

			currentChar++
			height = 0
			charLine = []string{}
		} else {
			// Otherwise, append the line to the current character
			charLine = append(charLine, line)
			height++

		}
	}
	// After the loop, check if there's a partially read character and store it
	if height > 0 {
		Banner[currentChar] = strings.Join(charLine, "\n")
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading:", err)
	}
	return Banner
}
