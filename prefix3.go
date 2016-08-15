package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	imr "github.com/paulbellamy/infinite_multi_reader"
)

const inputStr = `The servers are on fire
My pants are on fire
The fire is on fire`

func main() {
	input := strings.NewReader(inputStr)
	// START OMIT
	scanner := bufio.NewScanner(input)
	output := imr.InfiniteMultiReader(func() (io.Reader, error) {
		if !scanner.Scan() {
			return nil, io.EOF // If we have no more input
		}

		prefix := fmt.Sprintf("[%v] ", time.Now()) // Generate timestamp
		line := scanner.Text()

		// Create a new reader with timestamp prepended
		lineReader := io.MultiReader(
			strings.NewReader(prefix),
			strings.NewReader(line),
			strings.NewReader("\n"),
		)

		return lineReader, nil // Begin reading from the next reader
	})
	// END OMIT
	io.Copy(os.Stdout, output)
}
