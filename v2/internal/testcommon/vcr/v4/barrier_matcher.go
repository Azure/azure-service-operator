/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package v4

import (
	"net/http"
	"net/url"

	"github.com/go-logr/logr"
	"gopkg.in/dnaeon/go-vcr.v4/pkg/cassette"
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

// barrierMatcher wraps a delegate MatcherFunc to add mutation-barrier awareness.
//
// During cassette scanning, go-vcr calls the matcher for each unused candidate interaction in cassette order within a
// single GetInteraction() call. This matcher tracks when it encounters mutating operations (PUT/POST/PATCH/DELETE) and
// prevents subsequent GET candidates for the same URL path from matching. This stops timing jitter from causing GETs to
// match responses recorded after a mutation.
//
// Concurrency safety: go-vcr holds the cassette mutex (c.Lock()) for the entire interaction scan loop in
// cassette.getInteraction(). All matcher invocations within a scan are serialized, and concurrent goroutines calling
// GetInteraction will have their scans fully serialized. The barrier state maintained here is therefore safe from
// concurrent access without additional synchronization.
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
// Only unused candidates are passed to the matcher, they're passed strictly in cassette order, and a lock within
// cassette.getInteraction() ensures that concurrent scans are fully serialized.
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

	// If this is a GET request and the candidate is a GET past a mutation barrier, reject it.
	// We check r.Method (the incoming request) rather than i.Method (the candidate) because:
	// - For GET requests: this prevents the scan from jumping past mutations for the same path
	// - For non-GET requests: GET candidates would never match anyway (DefaultMatcher checks
	//   method equality), so blocking them would only produce misleading log messages
	if r.Method == http.MethodGet && i.Method == http.MethodGet && m.barriers[candidatePath] {
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
