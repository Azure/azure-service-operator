/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package v3

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"io"
	"net/http"
	"sync"

	"github.com/Azure/azure-service-operator/v2/internal/testcommon/vcr"
)

const (
	// COUNT_HEADER is the name of the header used to record the sequence number of a request
	COUNT_HEADER = "TEST-REQUEST-ATTEMPT"

	// HASH_HEADER is the name of the header used to record the hash of a request body
	HASH_HEADER = "TEST-REQUEST-HASH"
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

func AddTrackingHeaders(inner http.RoundTripper) *requestCounter {
	return &requestCounter{
		inner:       inner,
		counts:      make(map[string]uint32),
		countsMutex: sync.Mutex{},
	}
}

var _ http.RoundTripper = &requestCounter{}

func (rt *requestCounter) RoundTrip(req *http.Request) (*http.Response, error) {
	if rt.useHash(req) {
		rt.addContentHeaders(req)
	} else {
		rt.addCountHeader(req)
	}

	return rt.inner.RoundTrip(req)
}

// useHash returns true if we should use a hash to match this request
func (rt *requestCounter) useHash(req *http.Request) bool {
	if req.Method != "PUT" && req.Method != "POST" {
		// Only use a hash for PUT and POST methods
		return false
	}

	// Can only use a hash if there is a body
	return req.Body != nil
}

// addCountHeader adds a header to the request based on the URL requested
func (rt *requestCounter) addCountHeader(req *http.Request) {
	// Count keys are based on method and URL
	key := req.Method + ":" + req.URL.String()

	// Allocate a number
	rt.countsMutex.Lock()
	count := rt.counts[key]
	rt.counts[key] = count + 1
	rt.countsMutex.Unlock()

	// Apply the header
	req.Header.Set(COUNT_HEADER, fmt.Sprintf("%d", count))
}

// addContentHeaders adds headers to the request based on the content of the request body
func (rt *requestCounter) addContentHeaders(request *http.Request) {
	// Read all the content of the request body
	var body bytes.Buffer
	_, err := body.ReadFrom(request.Body)
	if err != nil {
		// Should never fail
		panic(fmt.Sprintf("reading request.Body failed: %s", err))
	}

	// Apply the same body filtering that we do in recordings so that the hash is consistent
	bodyString := vcr.HideRecordingData(string(body.Bytes()))

	// Calculate a hash based on body string
	hash := sha256.Sum256([]byte(bodyString))

	// Format hash as a hex string
	hashString := fmt.Sprintf("%x", hash)

	// Allocate a number based on the hash (not the URL)
	rt.countsMutex.Lock()
	count := rt.counts[hashString]
	rt.counts[hashString] = count + 1
	rt.countsMutex.Unlock()

	// Set the headers
	if request.Header == nil {
		request.Header = make(http.Header)
	}

	request.Header.Set(HASH_HEADER, fmt.Sprintf("%x", hash))
	request.Header.Set(COUNT_HEADER, fmt.Sprintf("%d", count))

	// Reset the body so it can be read again
	request.Body = io.NopCloser(&body)
}
