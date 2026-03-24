# VCR Mutation Barriers Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Prevent timing-jitter test failures by adding mutation barriers to VCR cassette matching and GET cache invalidation to the replay round tripper.

**Architecture:** A barrier matcher wraps go-vcr's matching to prevent GETs from matching past unconsumed mutations for the same URL path. The replay round tripper's GET cache is invalidated when mutations are observed. Count-based tracking headers are removed.

**Tech Stack:** Go, go-vcr v4, Gomega test assertions

---

## File Structure

All files are in `v2/internal/testcommon/vcr/v4/`:

| File                          | Action | Responsibility                                                             |
| ----------------------------- | ------ | -------------------------------------------------------------------------- |
| `barrier_matcher.go`          | Create | Stateful matcher that tracks mutation barriers during cassette scans       |
| `barrier_matcher_test.go`     | Create | Unit tests for barrier matching behaviour                                  |
| `replay_roundtripper.go`      | Modify | Add POST/PATCH/DELETE caches, GET invalidation on mutation, shared helpers |
| `replay_roundtripper_test.go` | Modify | New tests for expanded verbs and cache invalidation, update existing tests |
| `test_recorder.go`            | Modify | Wire barrier matcher, remove CountHeader check from matching               |
| `tracking_roundtripper.go`    | Modify | Remove count header logic, keep hash headers for PUT/POST/PATCH            |

---

### Task 1: Add `urlPath` helper

**Files:**
- Create: `v2/internal/testcommon/vcr/v4/barrier_matcher.go` (initial file with just the helper)
- Create: `v2/internal/testcommon/vcr/v4/barrier_matcher_test.go`

- [ ] **Step 1: Write the failing tests for `urlPath`**

Create `barrier_matcher_test.go`:

```go
package v4

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestUrlPath_GivenURLWithQueryParams_ReturnsPathOnly(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	g.Expect(urlPath("https://management.azure.com/subscriptions/sub1/resourceGroups/rg1?api-version=2021-04-01")).
		To(Equal("/subscriptions/sub1/resourceGroups/rg1"))
}

func TestUrlPath_GivenURLWithoutQueryParams_ReturnsPathOnly(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	g.Expect(urlPath("https://management.azure.com/subscriptions/sub1/resourceGroups/rg1")).
		To(Equal("/subscriptions/sub1/resourceGroups/rg1"))
}

func TestUrlPath_GivenRelativePath_ReturnsPath(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	g.Expect(urlPath("/subscriptions/sub1/resourceGroups/rg1?api-version=2021-04-01")).
		To(Equal("/subscriptions/sub1/resourceGroups/rg1"))
}
```

- [ ] **Step 2: Run tests to verify they fail**

Run: `cd v2 && go test ./internal/testcommon/vcr/v4/ -run TestUrlPath -v`
Expected: compilation error — `urlPath` undefined

- [ ] **Step 3: Write minimal implementation**

Create `barrier_matcher.go` with:

```go
/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package v4

import (
	"net/url"
)

// urlPath extracts the path portion of a URL string, stripping query parameters.
// Used by both the barrier matcher and the replay round tripper to match resources by path.
func urlPath(rawURL string) string {
	u, err := url.Parse(rawURL)
	if err != nil {
		// Should never happen with URLs from HTTP requests
		return rawURL
	}

	return u.Path
}
```

- [ ] **Step 4: Run tests to verify they pass**

Run: `cd v2 && go test ./internal/testcommon/vcr/v4/ -run TestUrlPath -v`
Expected: all 3 PASS

- [ ] **Step 5: Commit**

```bash
git add v2/internal/testcommon/vcr/v4/barrier_matcher.go v2/internal/testcommon/vcr/v4/barrier_matcher_test.go
git commit -m "Add urlPath helper for mutation barrier matching"
```

---

### Task 2: Implement `barrierMatcher`

**Files:**
- Modify: `v2/internal/testcommon/vcr/v4/barrier_matcher.go`
- Modify: `v2/internal/testcommon/vcr/v4/barrier_matcher_test.go`

- [ ] **Step 1: Write failing tests for barrier behaviour**

Add to `barrier_matcher_test.go`:

```go
import (
	"net/http"
	"net/url"

	"gopkg.in/dnaeon/go-vcr.v4/pkg/cassette"
	"github.com/go-logr/logr"
)

// fakeMatcher returns a MatcherFunc that always returns true for interactions
// whose method matches the request method, simulating default matching.
func fakeMatcher() cassette.MatcherFunc {
	return func(r *http.Request, i cassette.Request) bool {
		return r.Method == i.Method && r.URL.String() == i.URL
	}
}

func TestBarrierMatcher_GivenGETWithNoBarrier_DelegatesToInner(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	matcher := newBarrierMatcher(fakeMatcher(), logr.Discard())
	req := &http.Request{Method: http.MethodGet, URL: &url.URL{Path: "/foo"}}
	candidate := cassette.Request{Method: http.MethodGet, URL: "/foo"}

	g.Expect(matcher.Match(req, candidate)).To(BeTrue())
}

func TestBarrierMatcher_GivenGETPastMutation_RejectsGET(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	matcher := newBarrierMatcher(fakeMatcher(), logr.Discard())
	req := &http.Request{Method: http.MethodGet, URL: &url.URL{Path: "/foo"}}

	// Matcher sees a PUT candidate first (different method, won't match a GET request
	// via delegate, but the barrier is still recorded)
	putCandidate := cassette.Request{Method: http.MethodPut, URL: "/foo"}
	matcher.Match(req, putCandidate)

	// Now a GET candidate for the same path should be rejected
	getCandidate := cassette.Request{Method: http.MethodGet, URL: "/foo"}
	g.Expect(matcher.Match(req, getCandidate)).To(BeFalse())
}

func TestBarrierMatcher_GivenGETPastMutationForDifferentPath_AllowsGET(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	matcher := newBarrierMatcher(fakeMatcher(), logr.Discard())
	req := &http.Request{Method: http.MethodGet, URL: &url.URL{Path: "/bar"}}

	// PUT on /foo should not block GET on /bar
	putCandidate := cassette.Request{Method: http.MethodPut, URL: "/foo"}
	matcher.Match(req, putCandidate)

	getCandidate := cassette.Request{Method: http.MethodGet, URL: "/bar"}
	g.Expect(matcher.Match(req, getCandidate)).To(BeTrue())
}

func TestBarrierMatcher_GivenQueryParamDifference_BarrierStillApplies(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	matcher := newBarrierMatcher(fakeMatcher(), logr.Discard())
	req := &http.Request{Method: http.MethodGet, URL: &url.URL{Path: "/foo", RawQuery: "api-version=v2"}}

	putCandidate := cassette.Request{Method: http.MethodPut, URL: "/foo?api-version=v1"}
	matcher.Match(req, putCandidate)

	getCandidate := cassette.Request{Method: http.MethodGet, URL: "/foo?api-version=v2"}
	g.Expect(matcher.Match(req, getCandidate)).To(BeFalse())
}

func TestBarrierMatcher_GivenNewRequest_ResetsBarriers(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	matcher := newBarrierMatcher(fakeMatcher(), logr.Discard())

	// First scan: PUT creates a barrier
	req1 := &http.Request{Method: http.MethodGet, URL: &url.URL{Path: "/foo"}}
	putCandidate := cassette.Request{Method: http.MethodPut, URL: "/foo"}
	matcher.Match(req1, putCandidate)

	// Second scan: new request pointer, barriers should be cleared
	req2 := &http.Request{Method: http.MethodGet, URL: &url.URL{Path: "/foo"}}
	getCandidate := cassette.Request{Method: http.MethodGet, URL: "/foo"}
	g.Expect(matcher.Match(req2, getCandidate)).To(BeTrue())
}

func TestBarrierMatcher_GivenMultipleMutationVerbs_AllCreateBarriers(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	for _, method := range []string{http.MethodPut, http.MethodPost, http.MethodPatch, http.MethodDelete} {
		matcher := newBarrierMatcher(fakeMatcher(), logr.Discard())
		req := &http.Request{Method: http.MethodGet, URL: &url.URL{Path: "/foo"}}

		mutationCandidate := cassette.Request{Method: method, URL: "/foo"}
		matcher.Match(req, mutationCandidate)

		getCandidate := cassette.Request{Method: http.MethodGet, URL: "/foo"}
		g.Expect(matcher.Match(req, getCandidate)).To(BeFalse(), "expected %s to create barrier", method)
	}
}
```

- [ ] **Step 2: Run tests to verify they fail**

Run: `cd v2 && go test ./internal/testcommon/vcr/v4/ -run TestBarrierMatcher -v`
Expected: compilation error — `newBarrierMatcher` undefined

- [ ] **Step 3: Implement `barrierMatcher`**

Add to `barrier_matcher.go`:

```go
import (
	"net/http"
	"net/url"

	"github.com/go-logr/logr"
	"gopkg.in/dnaeon/go-vcr.v4/pkg/cassette"
)

// barrierMatcher wraps a delegate MatcherFunc to add mutation-barrier awareness.
//
// During cassette scanning, go-vcr calls the matcher for each candidate interaction in
// cassette order within a single GetInteraction() call. This matcher tracks when it
// encounters mutating operations (PUT/POST/PATCH/DELETE) and prevents subsequent GET
// candidates for the same URL path from matching. This stops timing jitter from causing
// GETs to match responses recorded after a mutation.
//
// Concurrency safety: go-vcr holds the cassette mutex (c.Lock()) for the entire
// interaction scan loop in cassette.getInteraction(). All matcher invocations within
// a scan are serialized, and concurrent goroutines calling GetInteraction will have
// their scans fully serialized. The barrier state maintained here is therefore safe
// from concurrent access without additional synchronization.
type barrierMatcher struct {
	delegate    cassette.MatcherFunc
	lastRequest *http.Request
	barriers    map[string]bool
	log         logr.Logger
}

func newBarrierMatcher(delegate cassette.MatcherFunc, log logr.Logger) *barrierMatcher {
	return &barrierMatcher{
		delegate: delegate,
		barriers: make(map[string]bool),
		log:      log,
	}
}

// Match implements the matching logic with mutation barriers.
func (m *barrierMatcher) Match(r *http.Request, i cassette.Request) bool {
	// Detect new scan by comparing request pointer identity.
	// go-vcr passes the same *http.Request for every candidate in a single scan.
	if r != m.lastRequest {
		m.lastRequest = r
		m.barriers = make(map[string]bool)
	}

	candidatePath := urlPath(i.URL)

	// If the candidate is a mutating operation, record a barrier for its path
	if isMutatingMethod(i.Method) {
		m.barriers[candidatePath] = true
	}

	// If the candidate is a GET and there's a barrier for its path, reject it
	if i.Method == http.MethodGet && m.barriers[candidatePath] {
		m.log.V(1).Info("Barrier rejected GET candidate", "url", i.URL)
		return false
	}

	return m.delegate(r, i)
}

// isMutatingMethod returns true if the HTTP method is one that modifies resources.
func isMutatingMethod(method string) bool {
	return method == http.MethodPut ||
		method == http.MethodPost ||
		method == http.MethodPatch ||
		method == http.MethodDelete
}
```

- [ ] **Step 4: Run tests to verify they pass**

Run: `cd v2 && go test ./internal/testcommon/vcr/v4/ -run "TestBarrierMatcher|TestUrlPath" -v`
Expected: all PASS

- [ ] **Step 5: Commit**

```bash
git add v2/internal/testcommon/vcr/v4/barrier_matcher.go v2/internal/testcommon/vcr/v4/barrier_matcher_test.go
git commit -m "Implement barrierMatcher with mutation barriers"
```

---

### Task 3: Wire barrier matcher into recorder

**Files:**
- Modify: `v2/internal/testcommon/vcr/v4/test_recorder.go:83-86` (recorder creation)
- Modify: `v2/internal/testcommon/vcr/v4/test_recorder.go:162-195` (matchOnHeadersAndBody)

- [ ] **Step 1: Run existing tests to confirm green baseline**

Run: `cd v2 && go test ./internal/testcommon/vcr/v4/ -v`
Expected: all PASS

- [ ] **Step 2: Update `NewTestRecorder` to create barrier matcher**

In `test_recorder.go`, replace the recorder creation block:

```go
	r, err := recorder.New(
		cassetteName,
		recorder.WithMode(mode),
		recorder.WithMatcher(matchOnHeadersAndBody(log)),
		recorder.WithHook(redactRecording(redactor), recorder.BeforeSaveHook),
	)
```

With:

```go
	barrier := newBarrierMatcher(matchOnHeadersAndBody(log), log)
	r, err := recorder.New(
		cassetteName,
		recorder.WithMode(mode),
		recorder.WithMatcher(barrier.Match),
		recorder.WithHook(redactRecording(redactor), recorder.BeforeSaveHook),
	)
```

- [ ] **Step 3: Remove `CountHeader` check from `matchOnHeadersAndBody`**

In `test_recorder.go`, remove the CountHeader check block from `matchOnHeadersAndBody`:

```go
	// verify custom request count header matches, if present
	if header := r.Header.Get(CountHeader); header != "" {
		interactionHeader := i.Headers.Get(CountHeader)
		if header != interactionHeader {
			log.Info("Request count header mismatch", CountHeader, header, "interaction", interactionHeader)
			return false
		}
	}
```

Remove only this block, leaving the `HashHeader` check and the body check intact.

- [ ] **Step 4: Run all v4 tests to verify nothing broke**

Run: `cd v2 && go test ./internal/testcommon/vcr/v4/ -v`
Expected: all PASS

- [ ] **Step 5: Commit**

```bash
git add v2/internal/testcommon/vcr/v4/test_recorder.go
git commit -m "Wire barrier matcher into recorder, remove CountHeader matching"
```

---

### Task 4: Simplify tracking round tripper

**Files:**
- Modify: `v2/internal/testcommon/vcr/v4/tracking_roundtripper.go`

- [ ] **Step 1: Run existing tests to confirm green baseline**

Run: `cd v2 && go test ./internal/testcommon/vcr/v4/ -v`
Expected: all PASS

- [ ] **Step 2: Simplify `RoundTrip` to only add hash headers**

Replace the `RoundTrip` method, `useHash` method, and `addCountHeader` method. The new `RoundTrip` should be:

```go
func (rt *requestCounter) RoundTrip(req *http.Request) (*http.Response, error) {
	if rt.hasBody(req) {
		rt.addHashHeader(req)
	}

	return rt.inner.RoundTrip(req)
}

// hasBody returns true if the request has a body that can be hashed
func (rt *requestCounter) hasBody(req *http.Request) bool {
	if req.Method != http.MethodPut && req.Method != http.MethodPost && req.Method != http.MethodPatch {
		return false
	}

	return req.Body != nil
}
```

Remove `useHash` and `addCountHeader` methods. Remove the `counts` and `countsMutex` fields from the `requestCounter` struct. Update `AddTrackingHeaders` to not initialise the removed fields.

- [ ] **Step 3: Write unit tests for simplified tracking roundtripper**

Add tests (in an appropriate test file, e.g. `tracking_roundtripper_test.go` if one exists, or add to `replay_roundtripper_test.go`):

```go
func TestTrackingRoundTripper_GivenGETRequest_AddsNoTrackingHeader(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	var captured *http.Request
	fake := vcr.NewFakeRoundTripper(cassette.ErrInteractionNotFound)
	redactor := vcr.NewRedactor(creds.DummyAzureIDs())
	tracker := AddTrackingHeaders(fake, redactor)

	req := &http.Request{
		URL:    &url.URL{Path: "/foo"},
		Method: http.MethodGet,
		Header: make(http.Header),
		Body:   io.NopCloser(strings.NewReader("")),
	}

	//nolint:bodyclose
	_, _ = tracker.RoundTrip(req)

	captured = req
	g.Expect(captured.Header.Get(CountHeader)).To(BeEmpty())
	g.Expect(captured.Header.Get(HashHeader)).To(BeEmpty())
}

func TestTrackingRoundTripper_GivenPUTRequest_AddsHashHeader(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	fake := vcr.NewFakeRoundTripper(cassette.ErrInteractionNotFound)
	redactor := vcr.NewRedactor(creds.DummyAzureIDs())
	tracker := AddTrackingHeaders(fake, redactor)

	req := &http.Request{
		URL:    &url.URL{Path: "/foo"},
		Method: http.MethodPut,
		Header: make(http.Header),
		Body:   io.NopCloser(strings.NewReader("PUT body")),
	}

	//nolint:bodyclose
	_, _ = tracker.RoundTrip(req)

	g.Expect(req.Header.Get(HashHeader)).ToNot(BeEmpty())
	g.Expect(req.Header.Get(CountHeader)).To(BeEmpty())
}

func TestTrackingRoundTripper_GivenDELETERequest_AddsNoTrackingHeader(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	fake := vcr.NewFakeRoundTripper(cassette.ErrInteractionNotFound)
	redactor := vcr.NewRedactor(creds.DummyAzureIDs())
	tracker := AddTrackingHeaders(fake, redactor)

	req := &http.Request{
		URL:    &url.URL{Path: "/foo"},
		Method: http.MethodDelete,
		Header: make(http.Header),
		Body:   io.NopCloser(strings.NewReader("")),
	}

	//nolint:bodyclose
	_, _ = tracker.RoundTrip(req)

	g.Expect(req.Header.Get(CountHeader)).To(BeEmpty())
	g.Expect(req.Header.Get(HashHeader)).To(BeEmpty())
}
```

- [ ] **Step 4: Run all v4 tests including new tracking tests**

Run: `cd v2 && go test ./internal/testcommon/vcr/v4/ -v`
Expected: all PASS

- [ ] **Step 5: Commit**

```bash
git add v2/internal/testcommon/vcr/v4/tracking_roundtripper.go v2/internal/testcommon/vcr/v4/tracking_roundtripper_test.go
git commit -m "Simplify tracking roundtripper: remove count headers, add PATCH hashing"
```

---

### Task 5: Refactor `replayRoundTripper` with shared helpers

**Files:**
- Modify: `v2/internal/testcommon/vcr/v4/replay_roundtripper.go`

This task refactors existing code without changing behaviour. The goal is to extract shared helpers before adding new verb caches.

- [ ] **Step 1: Run existing tests to confirm green baseline**

Run: `cd v2 && go test ./internal/testcommon/vcr/v4/ -run TestReplay -v`
Expected: all PASS

- [ ] **Step 2: Extract `replayFromCache` helper**

Add a shared method that handles the "check cache on ErrInteractionNotFound" pattern:

```go
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
```

- [ ] **Step 3: Extract `cacheResponse` helper**

```go
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
```

- [ ] **Step 4: Extract `invalidateCachedGets` helper**

```go
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
```

- [ ] **Step 5: Rewrite `roundTripGet` using helpers**

```go
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
```

- [ ] **Step 6: Rewrite `roundTripPut` using helpers**

```go
func (replayer *replayRoundTripper) roundTripPut(request *http.Request) (*http.Response, error) {
	hash := replayer.hashOfBody(request)

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
```

Note: `invalidateCachedGets` is now called on PUT success. This is the new cache invalidation behaviour.

- [ ] **Step 7: Run all existing tests to verify refactor is behaviour-preserving**

Run: `cd v2 && go test ./internal/testcommon/vcr/v4/ -run TestReplay -v`
Expected: all PASS

- [ ] **Step 8: Commit**

```bash
git add v2/internal/testcommon/vcr/v4/replay_roundtripper.go
git commit -m "Refactor replayRoundTripper: extract shared helpers, add GET invalidation on PUT"
```

---

### Task 6: Expand replay round tripper to cover POST, PATCH, DELETE

**Files:**
- Modify: `v2/internal/testcommon/vcr/v4/replay_roundtripper.go`
- Modify: `v2/internal/testcommon/vcr/v4/replay_roundtripper_test.go`

- [ ] **Step 1: Write failing tests for new verbs**

Add to `replay_roundtripper_test.go`:

```go
func TestReplayRoundTripper_GivenSinglePost_ReturnsOnceExtra(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	req := &http.Request{
		URL:    &url.URL{Path: "/foo"},
		Method: http.MethodPost,
		Body:   io.NopCloser(strings.NewReader("POST body goes here")),
	}
	resp := &http.Response{
		StatusCode: 200,
		Body:       io.NopCloser(strings.NewReader("POST response goes here")),
	}

	fake := vcr.NewFakeRoundTripper(cassette.ErrInteractionNotFound)
	fake.AddResponse(req, resp)
	redactor := vcr.NewRedactor(creds.DummyAzureIDs())
	replayer := NewReplayRoundTripper(fake, logr.Discard(), redactor)

	assertExpectedResponse(t, replayer, req, 200, "POST response goes here")
	assertExpectedResponse(t, replayer, req, 200, "POST response goes here")

	//nolint:bodyclose
	_, err := replayer.RoundTrip(req)
	g.Expect(err).To(HaveOccurred())
}

func TestReplayRoundTripper_GivenSinglePatch_ReturnsOnceExtra(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	req := &http.Request{
		URL:    &url.URL{Path: "/foo"},
		Method: http.MethodPatch,
		Body:   io.NopCloser(strings.NewReader("PATCH body goes here")),
	}
	resp := &http.Response{
		StatusCode: 200,
		Body:       io.NopCloser(strings.NewReader("PATCH response goes here")),
	}

	fake := vcr.NewFakeRoundTripper(cassette.ErrInteractionNotFound)
	fake.AddResponse(req, resp)
	redactor := vcr.NewRedactor(creds.DummyAzureIDs())
	replayer := NewReplayRoundTripper(fake, logr.Discard(), redactor)

	assertExpectedResponse(t, replayer, req, 200, "PATCH response goes here")
	assertExpectedResponse(t, replayer, req, 200, "PATCH response goes here")

	//nolint:bodyclose
	_, err := replayer.RoundTrip(req)
	g.Expect(err).To(HaveOccurred())
}

func TestReplayRoundTripper_GivenSingleDelete_ReturnsOnceExtra(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	req := &http.Request{
		URL:    &url.URL{Path: "/foo"},
		Method: http.MethodDelete,
		Body:   io.NopCloser(strings.NewReader("")),
	}
	resp := &http.Response{
		StatusCode: 200,
		Body:       io.NopCloser(strings.NewReader("DELETE response goes here")),
	}

	fake := vcr.NewFakeRoundTripper(cassette.ErrInteractionNotFound)
	fake.AddResponse(req, resp)
	redactor := vcr.NewRedactor(creds.DummyAzureIDs())
	replayer := NewReplayRoundTripper(fake, logr.Discard(), redactor)

	assertExpectedResponse(t, replayer, req, 200, "DELETE response goes here")
	assertExpectedResponse(t, replayer, req, 200, "DELETE response goes here")

	//nolint:bodyclose
	_, err := replayer.RoundTrip(req)
	g.Expect(err).To(HaveOccurred())
}
```

- [ ] **Step 2: Run tests to verify they fail**

Run: `cd v2 && go test ./internal/testcommon/vcr/v4/ -run "TestReplayRoundTripper_GivenSingle(Post|Patch|Delete)" -v`
Expected: FAIL — POST/PATCH/DELETE pass through without caching

- [ ] **Step 3: Add new replay constants and maps**

Add constants to `replay_roundtripper.go` alongside existing ones:

```go
	// POST requests may be replayed once, like PUT, to handle timing variations.
	maxPostReplays = 1 // Maximum number of times to replay a POST request

	// PATCH requests may be replayed once, like PUT, to handle timing variations.
	maxPatchReplays = 1 // Maximum number of times to replay a PATCH request

	// DELETE requests may be replayed once to handle timing variations.
	maxDeleteReplays = 1 // Maximum number of times to replay a DELETE request
```

Add new maps to the `replayRoundTripper` struct:

```go
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
```

Update `NewReplayRoundTripper` to initialise the new maps.

- [ ] **Step 4: Add handler methods for new verbs**

```go
func (replayer *replayRoundTripper) roundTripPost(request *http.Request) (*http.Response, error) {
	hash := replayer.hashOfBody(request)

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
	hash := replayer.hashOfBody(request)

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
```

- [ ] **Step 5: Update `RoundTrip` dispatch**

```go
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
```

- [ ] **Step 6: Run tests to verify new verb tests pass**

Run: `cd v2 && go test ./internal/testcommon/vcr/v4/ -run TestReplay -v`
Expected: all PASS

- [ ] **Step 7: Commit**

```bash
git add v2/internal/testcommon/vcr/v4/replay_roundtripper.go v2/internal/testcommon/vcr/v4/replay_roundtripper_test.go
git commit -m "Add POST/PATCH/DELETE replay caching with GET invalidation"
```

---

### Task 7: Add GET cache invalidation tests

**Files:**
- Modify: `v2/internal/testcommon/vcr/v4/replay_roundtripper_test.go`

- [ ] **Step 1: Write test for GET invalidation on PUT**

```go
func TestReplayRoundTripper_GivenPUTAfterCachedGET_InvalidatesGETCache(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	getReq := &http.Request{
		URL:    &url.URL{Path: "/foo"},
		Method: http.MethodGet,
		Body:   io.NopCloser(strings.NewReader("")),
	}
	getResp := &http.Response{
		StatusCode: 200,
		Body:       io.NopCloser(strings.NewReader(`{"properties":{"provisioningState": "Succeeded"}}`)),
	}

	putReq := &http.Request{
		URL:    &url.URL{Path: "/foo"},
		Method: http.MethodPut,
		Body:   io.NopCloser(strings.NewReader("PUT body")),
	}
	putResp := &http.Response{
		StatusCode: 200,
		Body:       io.NopCloser(strings.NewReader("PUT response")),
	}

	fake := vcr.NewFakeRoundTripper(cassette.ErrInteractionNotFound)
	fake.AddResponse(getReq, getResp)
	fake.AddResponse(putReq, putResp)
	redactor := vcr.NewRedactor(creds.DummyAzureIDs())
	replayer := NewReplayRoundTripper(fake, logr.Discard(), redactor)

	// GET caches a terminal response
	assertExpectedResponse(t, replayer, getReq, 200, `"provisioningState": "Succeeded"`)

	// PUT invalidates the cached GET
	assertExpectedResponse(t, replayer, putReq, 200, "PUT response")

	// GET replay should now fail (cache was invalidated)
	//nolint:bodyclose
	_, err := replayer.RoundTrip(getReq)
	g.Expect(err).To(HaveOccurred())
}

func TestReplayRoundTripper_GivenPUTForDifferentPath_DoesNotInvalidateGETCache(t *testing.T) {
	t.Parallel()

	getFooReq := &http.Request{
		URL:    &url.URL{Path: "/foo"},
		Method: http.MethodGet,
		Body:   io.NopCloser(strings.NewReader("")),
	}
	getFooResp := &http.Response{
		StatusCode: 200,
		Body:       io.NopCloser(strings.NewReader(`{"properties":{"provisioningState": "Succeeded"}}`)),
	}

	putBarReq := &http.Request{
		URL:    &url.URL{Path: "/bar"},
		Method: http.MethodPut,
		Body:   io.NopCloser(strings.NewReader("PUT body")),
	}
	putBarResp := &http.Response{
		StatusCode: 200,
		Body:       io.NopCloser(strings.NewReader("PUT response")),
	}

	fake := vcr.NewFakeRoundTripper(cassette.ErrInteractionNotFound)
	fake.AddResponse(getFooReq, getFooResp)
	fake.AddResponse(putBarReq, putBarResp)
	redactor := vcr.NewRedactor(creds.DummyAzureIDs())
	replayer := NewReplayRoundTripper(fake, logr.Discard(), redactor)

	// GET /foo caches
	assertExpectedResponse(t, replayer, getFooReq, 200, `"provisioningState": "Succeeded"`)

	// PUT /bar should NOT invalidate GET /foo
	assertExpectedResponse(t, replayer, putBarReq, 200, "PUT response")

	// GET /foo replay should still work
	assertExpectedResponse(t, replayer, getFooReq, 200, `"provisioningState": "Succeeded"`)
}
```

- [ ] **Step 2: Run tests**

Run: `cd v2 && go test ./internal/testcommon/vcr/v4/ -run "TestReplayRoundTripper_Given(PUT|POST|PATCH|DELETE)" -v`
Expected: all PASS

- [ ] **Step 3: Write test for DELETE invalidation**

Add to `replay_roundtripper_test.go` — this covers the URL-keyed invalidation path (different from hash-keyed PUT/POST/PATCH):

```go
func TestReplayRoundTripper_GivenDELETEAfterCachedGET_InvalidatesGETCache(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	getReq := &http.Request{
		URL:    &url.URL{Path: "/foo"},
		Method: http.MethodGet,
		Body:   io.NopCloser(strings.NewReader("")),
	}
	getResp := &http.Response{
		StatusCode: 200,
		Body:       io.NopCloser(strings.NewReader(`{"properties":{"provisioningState": "Succeeded"}}`)),
	}

	deleteReq := &http.Request{
		URL:    &url.URL{Path: "/foo"},
		Method: http.MethodDelete,
		Body:   io.NopCloser(strings.NewReader("")),
	}
	deleteResp := &http.Response{
		StatusCode: 200,
		Body:       io.NopCloser(strings.NewReader("DELETE response")),
	}

	fake := vcr.NewFakeRoundTripper(cassette.ErrInteractionNotFound)
	fake.AddResponse(getReq, getResp)
	fake.AddResponse(deleteReq, deleteResp)
	redactor := vcr.NewRedactor(creds.DummyAzureIDs())
	replayer := NewReplayRoundTripper(fake, logr.Discard(), redactor)

	// GET caches a terminal response
	assertExpectedResponse(t, replayer, getReq, 200, `"provisioningState": "Succeeded"`)

	// DELETE invalidates the cached GET
	assertExpectedResponse(t, replayer, deleteReq, 200, "DELETE response")

	// GET replay should now fail (cache was invalidated)
	//nolint:bodyclose
	_, err := replayer.RoundTrip(getReq)
	g.Expect(err).To(HaveOccurred())
}
```

- [ ] **Step 4: Run tests**

Run: `cd v2 && go test ./internal/testcommon/vcr/v4/ -run "TestReplayRoundTripper_Given" -v`
Expected: all PASS

- [ ] **Step 5: Commit**

```bash
git add v2/internal/testcommon/vcr/v4/replay_roundtripper_test.go
git commit -m "Add tests for GET cache invalidation on mutation"
```

---

### Task 8: Run full test suite and verify

**Files:** None (validation only)

- [ ] **Step 1: Run all v4 package tests**

Run: `cd v2 && go test ./internal/testcommon/vcr/v4/ -v -count=1`
Expected: all PASS

- [ ] **Step 2: Run linting**

Run: `cd /workspaces/azure-service-operator && ./hack/tools/task controller:lint`
Expected: no new lint errors in modified files

- [ ] **Step 3: Build controller**

Run: `cd /workspaces/azure-service-operator && ./hack/tools/task controller:build`
Expected: successful build

- [ ] **Step 4: Commit any fixups**

If lint or build exposed issues, fix them and commit:

```bash
git add -A
git commit -m "Fix lint and build issues from mutation barrier changes"
```
