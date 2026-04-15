/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package v4

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
	"gopkg.in/dnaeon/go-vcr.v4/pkg/cassette"

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
	posts    map[string]*replayResponse
	patches  map[string]*replayResponse
	deletes  map[string]*replayResponse
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

	// POST requests may be replayed once, like PUT, to handle timing variations.
	maxPostReplays = 1 // Maximum number of times to replay a POST request

	// PATCH requests may be replayed once, like PUT, to handle timing variations.
	maxPatchReplays = 1 // Maximum number of times to replay a PATCH request

	// DELETE requests may be replayed once to handle timing variations.
	maxDeleteReplays = 1 // Maximum number of times to replay a DELETE request

	// GET requests may be replayed multiple times to allow multiple reconciles to observe the same stable final state.
	// We set this to accommodate timing variations during test replay, while avoiding unbounded replays as they might
	// result in a test getting stuck and continuing to run until the entire test suite times out.
	//
	// Currently we need to set this to an apparently comical limit because some of our tests requiring many many
	// repetitions to work due to changes in timing used during test replay. Once we've addressed other issues causing
	// test instability, we should be able to reduce this limit significantly.
	maxGetReplays = 1000 // Maximum number of times to replay a GET request (effectively unlimited for now)
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
		posts:    make(map[string]*replayResponse),
		patches:  make(map[string]*replayResponse),
		deletes:  make(map[string]*replayResponse),
		log:      log,
		redactor: redactor,
	}
}

// RoundTrip implements http.RoundTripper.
func (replayer *replayRoundTripper) RoundTrip(request *http.Request) (*http.Response, error) {
	switch request.Method {
	case http.MethodGet:
		return replayer.roundTripGet(request)
	case http.MethodPut:
		return replayer.roundTripPut(request)
	case http.MethodPost:
		return replayer.roundTripPost(request)
	case http.MethodPatch:
		return replayer.roundTripPatch(request)
	case http.MethodDelete:
		return replayer.roundTripDelete(request)
	default:
		return replayer.inner.RoundTrip(request)
	}
}

func (replayer *replayRoundTripper) roundTripGet(request *http.Request) (*http.Response, error) {
	requestURL := request.URL.RequestURI()

	response, err := replayer.inner.RoundTrip(request)
	if err != nil {
		if !errors.Is(err, cassette.ErrInteractionNotFound) {
			return response, err
		}

		return replayer.replayFromCache(replayer.gets, requestURL, http.MethodGet)
	}

	// Cache only terminal-state responses
	cacheable := replayer.isTerminalHTTPStatus(response.StatusCode)
	if state, ok := replayer.resourceStateFromBody(response); ok {
		cacheable = replayer.isTerminalProvisioningState(state)
	}
	if status, ok := replayer.operationStatusFromBody(response); ok {
		cacheable = replayer.isTerminalOperationStatus(status)
	}

	if cacheable {
		replayer.cacheResponse(replayer.gets, requestURL, response, maxGetReplays)
	}

	return response, nil
}

func (replayer *replayRoundTripper) roundTripPut(request *http.Request) (*http.Response, error) {
	hash := replayer.hashOfCanonicalBody(request)

	response, err := replayer.inner.RoundTrip(request)
	if err != nil {
		if !errors.Is(err, cassette.ErrInteractionNotFound) {
			return response, err
		}

		return replayer.replayFromCache(replayer.puts, hash, http.MethodPut)
	}

	replayer.invalidateCachedGets(urlPath(request.URL.RequestURI()))
	replayer.cacheResponse(replayer.puts, hash, response, maxPutReplays)
	return response, nil
}

func (replayer *replayRoundTripper) roundTripPost(request *http.Request) (*http.Response, error) {
	hash := replayer.hashOfCanonicalBody(request)

	response, err := replayer.inner.RoundTrip(request)
	if err != nil {
		if !errors.Is(err, cassette.ErrInteractionNotFound) {
			return response, err
		}

		return replayer.replayFromCache(replayer.posts, hash, http.MethodPost)
	}

	replayer.invalidateCachedGets(urlPath(request.URL.RequestURI()))
	replayer.cacheResponse(replayer.posts, hash, response, maxPostReplays)
	return response, nil
}

func (replayer *replayRoundTripper) roundTripPatch(request *http.Request) (*http.Response, error) {
	hash := replayer.hashOfCanonicalBody(request)

	response, err := replayer.inner.RoundTrip(request)
	if err != nil {
		if !errors.Is(err, cassette.ErrInteractionNotFound) {
			return response, err
		}

		return replayer.replayFromCache(replayer.patches, hash, http.MethodPatch)
	}

	replayer.invalidateCachedGets(urlPath(request.URL.RequestURI()))
	replayer.cacheResponse(replayer.patches, hash, response, maxPatchReplays)
	return response, nil
}

func (replayer *replayRoundTripper) roundTripDelete(request *http.Request) (*http.Response, error) {
	requestURL := request.URL.RequestURI()

	response, err := replayer.inner.RoundTrip(request)
	if err != nil {
		if !errors.Is(err, cassette.ErrInteractionNotFound) {
			return response, err
		}

		return replayer.replayFromCache(replayer.deletes, requestURL, http.MethodDelete)
	}

	replayer.invalidateCachedGets(urlPath(requestURL))
	replayer.cacheResponse(replayer.deletes, requestURL, response, maxDeleteReplays)
	return response, nil
}

// replayFromCache attempts to return a cached response when go-vcr returns ErrInteractionNotFound.
// Returns the cached response if available, or the original error if not.
func (replayer *replayRoundTripper) replayFromCache(
	cache map[string]*replayResponse,
	key string,
	method string,
) (*http.Response, error) {
	replayer.padlock.Lock()
	defer replayer.padlock.Unlock()

	if cachedResponse, ok := cache[key]; ok {
		if cachedResponse.remainingReplays > 0 {
			cachedResponse.remainingReplays--
			replayer.log.Info("Replaying request", "method", method, "key", key)
			return cachedResponse.response, nil
		}

		// It's expired, remove it from the cache to ensure we don't replay it again
		delete(cache, key)
	}

	return nil, cassette.ErrInteractionNotFound
}

// cacheResponse stores a response in the given cache for future replay.
func (replayer *replayRoundTripper) cacheResponse(
	cache map[string]*replayResponse,
	key string,
	response *http.Response,
	maxReplays int,
) {
	replayer.padlock.Lock()
	defer replayer.padlock.Unlock()

	cache[key] = newReplayResponse(response, maxReplays)
}

// invalidateCachedGets removes all cached GET responses whose URL path matches the given path.
// Called when a mutating operation is observed, as cached GETs are now stale.
func (replayer *replayRoundTripper) invalidateCachedGets(mutationPath string) {
	replayer.padlock.Lock()
	defer replayer.padlock.Unlock()

	for key := range replayer.gets {
		if urlPath(key) == mutationPath {
			replayer.log.Info("Invalidating cached GET", "url", key, "reason", "mutation observed")
			delete(replayer.gets, key)
		}
	}
}

// hashOfCanonicalBody calculates a hash of the canonicalized body of a request, for use as a cache key.
// JSON bodies are canonicalized (keys sorted) before hashing to ensure deterministic results regardless
// of map key ordering in the serialized JSON.
func (replayer *replayRoundTripper) hashOfCanonicalBody(request *http.Request) string {
	var body bytes.Buffer
	if request.Body != nil {
		_, err := body.ReadFrom(request.Body)
		if err != nil {
			panic(fmt.Sprintf("reading request.Body failed: %s", err))
		}
	}

	bodyString := replayer.redactor.HideRecordingData(body.String())
	bodyString = canonicalizeJSON(bodyString)
	hash := sha256.Sum256([]byte(bodyString))

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

// isTerminalHTTPStatus returns true if the specified HTTP status code represents a terminal state for a resource.
// "Not existing" is pretty terminal.
func (*replayRoundTripper) isTerminalHTTPStatus(status int) bool {
	// Why these statuses?
	// - 404 Not Found and 410 Gone indicate that the resource doesn't exist, that's not going to change unless we do a PUT
	// - 201 Created and 200 OK indicate that the resource exists, and that no further processing is required by the server
	return status == http.StatusNotFound ||
		status == http.StatusGone ||
		status == http.StatusCreated ||
		status == http.StatusOK
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
