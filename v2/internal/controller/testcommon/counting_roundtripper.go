/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package testcommon

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

func addCountHeader(inner http.RoundTripper) *requestCounter {
	return &requestCounter{
		inner:       inner,
		counts:      make(map[string]uint32),
		countsMutex: sync.Mutex{},
	}
}

var COUNT_HEADER string = "TEST-REQUEST-ATTEMPT"

func (rt *requestCounter) RoundTrip(req *http.Request) (*http.Response, error) {
	key := req.Method + ":" + req.URL.String()
	rt.countsMutex.Lock()
	count := rt.counts[key]
	rt.counts[key] = count + 1
	rt.countsMutex.Unlock()
	req.Header.Set(COUNT_HEADER, fmt.Sprintf("%d", count))
	return rt.inner.RoundTrip(req)
}

var _ http.RoundTripper = &requestCounter{}
