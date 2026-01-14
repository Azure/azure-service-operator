/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package v3

import (
	"bytes"
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
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
	gets     map[string]*replayResponse
	puts     map[string]*replayResponse
	log      logr.Logger
	padlock  sync.Mutex
	redactor *vcr.Redactor
}

type replayResponse struct {
	response         *http.Response
	remainingReplays int
}

const (
	// Timing variations during test replay may result in a resource being reconciled an additional time - we don't
	// want to fail the test in this situation because we've already successfully achieved our goal state.
	// Thus we allow the last PUT for each resource to be replayed one extra time.
	maxPutReplays = 1 // Maximum number of times to replay a PUT request

	// GET requests may be replayed multiple times to allow multiple reconciles to observe the same stable final state.
	// We set this to accommodate timing variations during test replay, while avoiding unbounded replays as they might
	// result in a test getting stuck and continuing to run until the entire test suite times out.
	//
	// Surprisingly, the test Test_DBForPostgreSQL_FlexibleServer_20250801_CRUD_2 requires this value to be at least 62
	// before the test will pass locally for @theunrepentantgeek; be cautious about lowering this value.
	maxGetReplays = 80 // Maximum number of times to replay a GET request
)

var _ http.RoundTripper = &replayRoundTripper{}

// newReplayRoundTripper creates a new replayRoundTripper that will replay selected requests to improve test resilience.
func NewReplayRoundTripper(
	inner http.RoundTripper,
	log logr.Logger,
	redactor *vcr.Redactor,
) http.RoundTripper {
	return &replayRoundTripper{
		inner:    inner,
		gets:     make(map[string]*replayResponse),
		puts:     make(map[string]*replayResponse),
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
	requestURL := request.URL.RequestURI()

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
		if cachedResponse, ok := replayer.gets[requestURL]; ok {
			if cachedResponse.remainingReplays > 0 {
				cachedResponse.remainingReplays--
				replayer.log.Info("Replaying GET request", "url", requestURL)
				return cachedResponse.response, nil
			}

			// It's expired, remove it from the cache to ensure we don't replay it again
			delete(replayer.gets, requestURL)
		}

		// No cached response, return the original response and error
		return response, err
	}

	// We have a response; if it has a status, cache only if that represents a terminal state
	cacheable := true
	if state, ok := replayer.resourceStateFromBody(response); ok {
		cacheable = replayer.isTerminalProvisioningState(state)
	}
	if status, ok := replayer.operationStatusFromBody(response); ok {
		cacheable = replayer.isTerminalOperationStatus(status)
	}

	if cacheable {
		replayer.padlock.Lock()
		defer replayer.padlock.Unlock()

		replayer.gets[requestURL] = newReplayResponse(response, maxGetReplays)
	}

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
			if cachedResponse.remainingReplays > 0 {
				replayer.log.Info("Replaying PUT request", "url", request.URL.String(), "hash", hash)
				cachedResponse.remainingReplays--
				return cachedResponse.response, nil
			}

			// It's expired, remove it from the cache to ensure we don't replay it again
			delete(replayer.puts, hash)

		}

		// No cached response, return the original response and error
		return response, err
	}

	replayer.padlock.Lock()
	defer replayer.padlock.Unlock()

	// We have a response, cache it and return it
	replayer.puts[hash] = newReplayResponse(response, maxPutReplays)
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

type operation struct {
	Status string `json:"status"`
}

// operationStatusFromBody extracts the operation status from the response body, if present.
func (replayer *replayRoundTripper) operationStatusFromBody(response *http.Response) (string, bool) {
	body := replayer.bodyOfResponse(response)

	var op operation
	if err := json.Unmarshal([]byte(body), &op); err == nil {
		if op.Status != "" {
			return op.Status, true
		}
	}

	return "", false
}

// isTerminalOperationStatus returns true if the specified status represents a terminal state for a long-running operation
func (replayer *replayRoundTripper) isTerminalOperationStatus(status string) bool {
	return !strings.EqualFold(status, "InProgress")
}

type resource struct {
	Properties struct {
		ProvisioningState string `json:"provisioningState"`
	} `json:"properties"`
}

// resourceStateFromBody extracts the provisioning state from the response body, if present.
func (replayer *replayRoundTripper) resourceStateFromBody(response *http.Response) (string, bool) {
	body := replayer.bodyOfResponse(response)

	// Treat the body as an operation and deserialize it

	// Treat the body as a resource and deserialize it
	var res resource
	if err := json.Unmarshal([]byte(body), &res); err == nil {
		if res.Properties.ProvisioningState != "" {
			return res.Properties.ProvisioningState, true
		}
	}

	return "", false
}

// isTerminalStatus returns true if the specified status represents a terminal state a resource
func (*replayRoundTripper) isTerminalProvisioningState(status string) bool {
	return strings.EqualFold(status, "Succeeded") ||
		strings.EqualFold(status, "Failed") ||
		strings.EqualFold(status, "Canceled")
}

func (replayer *replayRoundTripper) bodyOfResponse(response *http.Response) string {
	// Read all the content of the response body
	var body bytes.Buffer
	_, err := body.ReadFrom(response.Body)
	if err != nil {
		// Should never fail
		panic(fmt.Sprintf("reading response.Body failed: %s", err))
	}

	// Reset the body so it can be read again
	response.Body = io.NopCloser(&body)
	return body.String()
}

func newReplayResponse(resp *http.Response, maxReplays int) *replayResponse {
	return &replayResponse{
		response:         resp,
		remainingReplays: maxReplays,
	}
}
