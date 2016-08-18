package main

import (
	"compress/gzip"
	"io"
	"os"
	"strings"
)

func main() {
	// START SETUP OMIT
	var r io.Reader

	r = strings.NewReader("1234567890")

	r = io.LimitReader(r, 5)
	// END SETUP OMIT

	// START OUTPUT OMIT
	io.Copy(os.Stdout, r)
	// END OUTPUT OMIT

	// START GZIP OMIT
	io.Copy(gzip.NewReader(os.Stdout), r)
	// END GZIP OMIT
}
