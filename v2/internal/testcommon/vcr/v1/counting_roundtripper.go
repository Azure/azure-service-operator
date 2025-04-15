/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package v1

import (
	"fmt"
	"net/http"
	"sync"
)

// Wraps an inner HTTP roundtripper to add a
// counter for duplicated request URIs. This
// is then used to match up requests in the recorder
// - it is needed as we have multiple requests with
// the same Request URL and it will return the first
// one that matches.
type requestCounter struct {
	inner http.RoundTripper

	countsMutex sync.Mutex
	counts      map[string]uint32
}

func AddCountHeader(inner http.RoundTripper) *requestCounter {
	return &requestCounter{
		inner:       inner,
		counts:      make(map[string]uint32),
		countsMutex: sync.Mutex{},
	}
}

var CountHeader string = "TEST-REQUEST-ATTEMPT"

func (rt *requestCounter) RoundTrip(req *http.Request) (*http.Response, error) {
	key := req.Method + ":" + req.URL.String()

	// Read the current
	rt.countsMutex.Lock()
	count := rt.counts[key]
	rt.counts[key] = count + 1
	rt.countsMutex.Unlock()

	req.Header.Set(CountHeader, fmt.Sprintf("%d", count))
	return rt.inner.RoundTrip(req)
}

var _ http.RoundTripper = &requestCounter{}
