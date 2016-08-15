package main

import (
	"bufio"
	"strings"
)

const inputStr = `The servers are on fire
My pants are on fire
The fire is on fire`

func main() {
	input := strings.NewReader(inputStr)
	// START OMIT
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
	}
	// END OMIT
}
