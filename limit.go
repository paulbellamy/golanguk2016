package main

import (
	"io"
	"os"
	"strings"
)

func main() {
	// START OMIT
	var r io.Reader

	r = strings.NewReader("1234567890")

	r = io.LimitReader(r, 5)

	io.Copy(os.Stdout, r)
	// END OMIT
}
