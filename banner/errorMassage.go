package Ascii

import (
	"fmt"
	"os"
)

func PrintErr() {
	fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nExample: go run . --align=right something standard")
	os.Exit(1)
}
