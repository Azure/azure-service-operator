/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package testcommon

import (
	"fmt"
	"net/http"
)

// Wraps an inner HTTP roundtripper to add a
// counter for duplicated request URIs. This
// is then used to match up requests in the recorder
// - it is needed as we have multiple requests with
// the same Request URL and it will return the first
// one that matches.
type requestCounter struct {
	inner  http.RoundTripper
	counts map[string]uint32
}

func addCountHeader(inner http.RoundTripper) *requestCounter {
	return &requestCounter{
		inner:  inner,
		counts: make(map[string]uint32),
	}
}

var COUNT_HEADER string = "TEST-REQUEST-ATTEMPT"

func (rt *requestCounter) RoundTrip(req *http.Request) (*http.Response, error) {
	key := req.Method + ":" + req.URL.String()
	count := rt.counts[key]
	req.Header.Add(COUNT_HEADER, fmt.Sprintf("%v", count))
	rt.counts[key] = count + 1
	return rt.inner.RoundTrip(req)
}

var _ http.RoundTripper = &requestCounter{}
