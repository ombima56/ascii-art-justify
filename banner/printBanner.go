package Ascii

import (
	"fmt"
	"os"
	"strings"
	"syscall"
	"unsafe"
)

// PrintBanner prints the input string using the loaded banner characters and specified alignment.
func PrintBanner(line string, alignment string) {
	outPut := make([][]string, 8) // Output slice to store the banner lines.
	fileName := ""
	if len(os.Args) == 3 {
		fileName = os.Args[2]
	} else {
		fileName = "standard" // Default to "standard" if no file name is provided.
	}
	// Checks if the banner file size is not altered.
	filePath, err := FileCheck(fileName)
	if err != nil {
		fmt.Println(err, filePath)
		os.Exit(1)
	}

	banner := LoadBanner(filePath) // Load the banner characters

	for _, char := range line {
		if char < 32 || char > 126 {
			fmt.Printf("Character out of range:%q\n", char)
			os.Exit(1)
		}
		if ascii, Ok := banner[char]; Ok {

			// If the character is found, split it into lines and append to the output
			asciiLines := strings.Split(ascii, "\n")
			for i := 0; i < len(asciiLines); i++ {
				outPut[i] = append(outPut[i], asciiLines[i])
			}
		} else {
			// If the character is not found, print an error message and continue
			fmt.Printf("Character not found: %q\n", char)
			continue
		}
	}

	// Print the assembled output lines with alignment
	for _, line := range outPut {
		outputLine := strings.Join(line, "")
		switch alignment {
		case "left":
			fmt.Println(outputLine)
		case "right":
			fmt.Println(rightAlign(outputLine, getTerminalWidth()))
		case "center":
			fmt.Println(centerAlign(outputLine, getTerminalWidth()))
		case "justify":
			fmt.Println(justifyAlign(outputLine, getTerminalWidth()))
		default:
			fmt.Println(outputLine)
		}
	}
}

// Get the terminal width manually
func getTerminalWidth() int {
	var ws struct {
		rows uint16
		cols uint16
		xpix uint16
		ypix uint16
	}
	retCode, _, errno := syscall.Syscall(
		syscall.SYS_IOCTL,
		uintptr(syscall.Stdin),
		uintptr(syscall.TIOCGWINSZ),
		uintptr(unsafe.Pointer(&ws)),
	)
	if int(retCode) == -1 {
		fmt.Printf("Error getting terminal size: %v\n", errno)
		return 80 // Default width if an error occurs
	}
	return int(ws.cols)
}

// Right-align the text
func rightAlign(text string, width int) string {
	if len(text) >= width {
		return text
	}
	padding := width - len(text)
	return strings.Repeat(" ", padding) + text
}

// Center-align the text
func centerAlign(text string, width int) string {
	if len(text) >= width {
		return text
	}
	padding := (width - len(text)) / 2
	return strings.Repeat(" ", padding) + text
}

// Justify-align the text
func justifyAlign(text string, width int) string {
	words := strings.Fields(text)
	if len(words) == 1 {
		return words[0] + strings.Repeat(" ", width-len(words[0]))
	}

	// Calculate the total length of all words without spaces
	totalWordsLength := 0
	for _, word := range words {
		totalWordsLength += len(word)
	}

	// Calculate total spaces needed to fill the width
	spacesNeeded := width - totalWordsLength
	spacesBetweenWords := spacesNeeded / (len(words) - 1)
	extraSpaces := spacesNeeded % (len(words) - 1)

	var result strings.Builder
	for i, word := range words {
		result.WriteString(word)
		if i < len(words)-1 {
			spaces := spacesBetweenWords
			if extraSpaces > 0 {
				spaces++
				extraSpaces--
			}
			result.WriteString(strings.Repeat(" ", spaces))
		}
	}

	return result.String()
}
