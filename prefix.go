package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
)

func main() {
	done := make(chan error)
	// START USAGE OMIT
	commands := [][]string{
		{"echo", "Hello,\nWorld!"},
		{"make", "all"},
		{"git", "status"},
	}

	for _, cmd := range commands {
		c := exec.Command(cmd[0], cmd[1])

		// Wire up our prefix writer
		prefixedOutput := prefixingWriter(cmd[0], os.Stdout)
		c.Stdout = prefixedOutput
		c.Stderr = prefixedOutput

		// Start each in a goroutine
		go func() { done <- c.Run() }()
	}
	// END USAGE OMIT
	for i := 0; i < len(commands); i++ {
		<-done
	}
}

// START DECL OMIT
func prefixingWriter(prefix string, output io.Writer) io.Writer

// END DECL OMIT

// START IMPL OMIT
func prefixingWriter(prefix string, output io.Writer) io.Writer {
	pipeReader, pipeWriter := io.Pipe()

	scanner := bufio.NewScanner(pipeReader)

	go func() {
		for scanner.Scan() {
			fmt.Fprintf(output, "[%s] ", prefix) // Write the prefix into the output
			output.Write(scanner.Bytes())        // Copy the line
			fmt.Fprint(output, "\n")             // Re-add a newline (scanner removes it)
		}
	}()

	return pipeWriter
}

// END IMPL OMIT
