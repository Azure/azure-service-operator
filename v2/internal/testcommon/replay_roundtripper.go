/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package testcommon

import (
	"errors"
	"net/http"
	"sync"

	"github.com/go-logr/logr"
	"gopkg.in/dnaeon/go-vcr.v3/cassette"
)

// replayRoundTripper wraps an inner round tripper and replays requests in order to improve the resilience of ASO tests.
// PUT requests are cached by target URL, and may be replayed ONCE if an extra PUT occurs.
// GET requests are cached by target URL, and may be replayed multiple times.
// This combination should allow additional reconciles - an extra PUT gets returned the same long running operation as
// the original, which a GET then shows is complete.
type replayRoundTripper struct {
	inner   http.RoundTripper
	gets    map[string]*http.Response
	puts    map[string]*http.Response
	log     logr.Logger
	padlock sync.Mutex
}

var _ http.RoundTripper = &replayRoundTripper{}

// newReplayRoundTripper creates a new replayRoundTripper that will replay selected requests to improve test resilience.
func newReplayRoundTripper(
	inner http.RoundTripper,
	log logr.Logger,
) *replayRoundTripper {
	return &replayRoundTripper{
		inner: inner,
		gets:  make(map[string]*http.Response),
		puts:  make(map[string]*http.Response),
		log:   log,
	}
}

// RoundTrip implements http.RoundTripper.
func (replayer *replayRoundTripper) RoundTrip(request *http.Request) (*http.Response, error) {
	if request.Method == http.MethodGet {
		return replayer.roundTripGet(request)
	}

	if request.Method == http.MethodPut {
		return replayer.roundTripPut(request)
	}

	// For other kinds of request, just pass through to the inner round tripper.
	return replayer.inner.RoundTrip(request)
}

func (replayer *replayRoundTripper) roundTripGet(request *http.Request) (*http.Response, error) {
	// First use our inner round tripper to get the response.
	response, err := replayer.inner.RoundTrip(request)
	if err != nil {
		// We have an error - return it, unless it's from go-vcr
		if !errors.Is(err, cassette.ErrInteractionNotFound) {
			return response, err
		}

		replayer.padlock.Lock()
		defer replayer.padlock.Unlock()

		// We didn't find an interaction, see if we have a cached response to return
		if cachedResponse, ok := replayer.gets[request.URL.String()]; ok {
			replayer.log.Info("Replaying GET request", "url", request.URL.String())
			return cachedResponse, nil
		}

		// No cached response, return the original response and error
		return response, err
	}

	replayer.padlock.Lock()
	defer replayer.padlock.Unlock()

	// We have a response, cache it and return it
	replayer.gets[request.URL.String()] = response
	return response, nil
}

func (replayer *replayRoundTripper) roundTripPut(request *http.Request) (*http.Response, error) {
	response, err := replayer.inner.RoundTrip(request)
	if err != nil {
		// We have an error - return it, unless it's from go-vcr
		if !errors.Is(err, cassette.ErrInteractionNotFound) {
			return response, err
		}

		replayer.padlock.Lock()
		defer replayer.padlock.Unlock()

		// We didn't find an interaction, see if we have a cached response to return
		if cachedResponse, ok := replayer.puts[request.URL.String()]; ok {
			delete(replayer.puts, request.URL.String())
			replayer.log.Info("Replaying PUT request", "url", request.URL.String())
			return cachedResponse, nil
		}

		// No cached response, return the original response and error
		return response, err
	}

	replayer.padlock.Lock()
	defer replayer.padlock.Unlock()

	// We have a response, cache it and return it
	replayer.puts[request.URL.String()] = response
	return response, nil
}
