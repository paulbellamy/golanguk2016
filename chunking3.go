package main

import (
	"crypto/md5"
	"io"
	"net/http/httputil"
)

func main() {
	// START OMIT
	digest := md5.New()
	io.Copy(
		digest,
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
