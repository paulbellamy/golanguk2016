package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
	"time"
)

const inputStr = `The servers are on fire
My pants are on fire
The fire is on fire`

func main() {
	input := strings.NewReader(inputStr)
	// START OMIT
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		prefix := fmt.Sprintf("[%v] ", time.Now()) // Generate timestamp
		line := scanner.Text()

		// Create a new reader with timestamp prepended
		outputLine := io.MultiReader(
			strings.NewReader(prefix),
			strings.NewReader(line),
			strings.NewReader("\n"),
		)
	}
	// END OMIT
}
