/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package v3

import (
	"crypto/sha256"
	"fmt"
	"io"
	"net/http"
	"sync"

	"github.com/Azure/azure-service-operator/v2/internal/testcommon/vcr"
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

const (
	COUNT_HEADER string = "TEST-REQUEST-ATTEMPT"
	HASH_HEADER  string = "TEST-REQUEST-HASH"
)

var _ http.RoundTripper = &requestCounter{}

func (rt *requestCounter) RoundTrip(req *http.Request) (*http.Response, error) {
	if req.Method == "PUT" || req.Method == "POST" {
		rt.addHashHeader(req)
	} else {
		rt.addCountHeader(req)
	}

	return rt.inner.RoundTrip(req)
}

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

func (rt *requestCounter) addHashHeader(req *http.Request) {
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
	bodyString := vcr.HideRecordingData(string(bodyBytes))

	// Calculate a hash based on body string
	hash := sha256.Sum256([]byte(bodyString))

	// Format hash as a hex string
	hashString := fmt.Sprintf("%x", hash)

	// Allocate a number
	rt.countsMutex.Lock()
	count := rt.counts[hashString]
	rt.counts[hashString] = count + 1
	rt.countsMutex.Unlock()

	// Set the headers
	req.Header.Set(HASH_HEADER, fmt.Sprintf("%x", hash))
	req.Header.Set(COUNT_HEADER, fmt.Sprintf("%d", count))

	// Reset the body so it can be read again
	_, _ = rsc.Seek(0, io.SeekStart)
}
