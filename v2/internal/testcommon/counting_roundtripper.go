/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package testcommon

import (
	"crypto/sha256"
	"fmt"
	"io"
	"net/http"
	"sync"

	"k8s.io/klog/v2"
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
var HASH_HEADER string = "TEST-REQUEST-HASH"

func (rt *requestCounter) RoundTrip(req *http.Request) (*http.Response, error) {
	if req.Method == "PUT" || req.Method == "POST" {
		rt.addHashHeader(req)
	} else {
		rt.addCountHeader(req)
	}

	return rt.inner.RoundTrip(req)
}

func (rt *requestCounter) addCountHeader(req *http.Request) {
	key := req.Method + ":" + req.URL.String()
	rt.countsMutex.Lock()
	count := rt.counts[key]
	rt.counts[key] = count + 1
	rt.countsMutex.Unlock()
	req.Header.Set(COUNT_HEADER, fmt.Sprintf("%d", count))
}

func (rt *requestCounter) addHashHeader(req *http.Request) {
	klog.Warning(fmt.Sprintf("Adding hash header to request: %s %s %T", req.Method, req.URL.String(), req.Body))

	rsc, ok := req.Body.(io.ReadSeekCloser)
	if !ok {
		panic(fmt.Sprintf("req.Body is not an io.ReadSeekCloser, it is a %T", req.Body))
	}

	// Calculate a hash based on the request body, but sanitise it first so it's the same for each test run
	bodyBytes, bodyErr := io.ReadAll(rsc)
	if bodyErr != nil {
		// see invocation of SetMatcher in the createRecorder, which does this
		panic("io.ReadAll(req.Body) failed, this should always succeed because req.Body has been replaced by a buffer")
	}

	// Apply the same body filtering that we do in recordings so that the hash is consistent
	bodyString := hideRecordingData(string(bodyBytes))

	// Calculate a hash based on body string
	hash := sha256.Sum256([]byte(bodyString))

	// Set the header
	req.Header.Set(HASH_HEADER, fmt.Sprintf("%x", hash))

	// Reset the body so it can be read again
	_, _ = rsc.Seek(0, io.SeekStart)
}

var _ http.RoundTripper = &requestCounter{}
