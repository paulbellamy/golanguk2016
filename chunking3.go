package main

import (
	"io"
	"io/ioutil"
	"net/http/httputil"
)

func main() {
	// START OMIT
	io.Copy(
		ioutil.Discard,
		httputil.NewChunkedReader(
			io.TeeReader(
				responseBody,
				clientWriter,
			),
		),
	)
	parseTrailers(responseBody)
	// END OMIT
}
