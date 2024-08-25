/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package v3

import (
	"bytes"
	"crypto/sha256"
	"errors"
	"fmt"
	"io"
	"net/http"
	"sync"

	"github.com/go-logr/logr"
	"gopkg.in/dnaeon/go-vcr.v3/cassette"

	"github.com/Azure/azure-service-operator/v2/internal/testcommon/vcr"
)

// replayRoundTripper wraps an inner round tripper and replays requests in order to improve the resilience of ASO tests.
//
// PUT requests are cached by hash of the (sanitised) PUT body and may be replayed ONCE if an extra PUT occurs.
//
// GET requests are cached by target URL, and may be replayed multiple times.
//
// This combination should allow additional reconciles - an extra PUT gets returned the same long running operation as
// the original, which a GET then shows is complete.
type replayRoundTripper struct {
	inner    http.RoundTripper
	gets     map[string]*http.Response
	puts     map[string]*http.Response
	log      logr.Logger
	padlock  sync.Mutex
	redactor *vcr.Redactor
}

var _ http.RoundTripper = &replayRoundTripper{}

// newReplayRoundTripper creates a new replayRoundTripper that will replay selected requests to improve test resilience.
func NewReplayRoundTripper(
	inner http.RoundTripper,
	log logr.Logger,
	redactor *vcr.Redactor,
) http.RoundTripper {
	return &replayRoundTripper{
		inner:    inner,
		gets:     make(map[string]*http.Response),
		puts:     make(map[string]*http.Response),
		log:      log,
		redactor: redactor,
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
	// Calculate a hash of the request body to use as a cache key
	// We need this whether we are updating our cache or replaying
	hash := replayer.hashOfBody(request)

	response, err := replayer.inner.RoundTrip(request)
	if err != nil {
		// We have an error - return it, unless it's from go-vcr
		if !errors.Is(err, cassette.ErrInteractionNotFound) {
			return response, err
		}

		replayer.padlock.Lock()
		defer replayer.padlock.Unlock()

		// We didn't find an interaction, see if we have a cached response to return
		if cachedResponse, ok := replayer.puts[hash]; ok {
			// Remove it from the cache to ensure we only replay it once
			delete(replayer.puts, hash)
			replayer.log.Info("Replaying PUT request", "url", request.URL.String(), "hash", hash)
			return cachedResponse, nil
		}

		// No cached response, return the original response and error
		return response, err
	}

	replayer.padlock.Lock()
	defer replayer.padlock.Unlock()

	// We have a response, cache it and return it
	replayer.puts[hash] = response
	return response, nil
}

// hashOfBody calculates a hash of the body of a request, for use as a cache key.
// The body is santised before calculating the hash to ensure that the same request body always results in the same hash.
func (replayer *replayRoundTripper) hashOfBody(request *http.Request) string {
	// Read all the content of the request body
	var body bytes.Buffer
	_, err := body.ReadFrom(request.Body)
	if err != nil {
		// Should never fail
		panic(fmt.Sprintf("reading request.Body failed: %s", err))
	}

	// Apply the same body filtering that we do in recordings so that the hash is consistent
	bodyString := replayer.redactor.HideRecordingData(body.String())

	// Calculate a hash based on body string
	hash := sha256.Sum256([]byte(bodyString))

	// Reset the body so it can be read again
	request.Body = io.NopCloser(&body)

	return fmt.Sprintf("%x", hash)
}
