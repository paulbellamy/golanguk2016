package main

import "io"

func main() {
	// START OMIT
	io.Copy(clientWriter, responseBody)
	// END OMIT
}
