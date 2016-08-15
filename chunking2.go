package main

import (
	"io"
	"net/http/httputil"
)

func main() {
	// START OMIT
	io.Copy(
		clientWriter,
		httputil.NewChunkedReader(responseBody),
	)
	parseTrailers(responseBody)
	// END OMIT
}
