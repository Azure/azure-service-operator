/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package v4

import (
	"net/http"
	"net/url"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/go-logr/logr"
	"gopkg.in/dnaeon/go-vcr.v4/pkg/cassette"
)

// fakeMatcher returns a MatcherFunc that always returns true for interactions
// whose method matches the request method, simulating default matching.
func fakeMatcher() cassette.MatcherFunc {
	return func(r *http.Request, i cassette.Request) bool {
		return r.Method == i.Method && r.URL.String() == i.URL
	}
}

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

func TestBarrierMatcher_GivenGETWithNoBarrier_DelegatesToInner(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	matcher := newBarrierMatcher(fakeMatcher(), logr.Discard())
	req := &http.Request{Method: http.MethodGet, URL: &url.URL{Path: "/foo"}}
	candidate := cassette.Request{Method: http.MethodGet, URL: "/foo"}

	g.Expect(matcher.Match(req, candidate)).To(BeTrue())
}

func TestBarrierMatcher_GivenGETPastMutation_DoesNotMatchGET(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	matcher := newBarrierMatcher(fakeMatcher(), logr.Discard())
	req := &http.Request{Method: http.MethodGet, URL: &url.URL{Path: "/foo"}}

	putCandidate := cassette.Request{Method: http.MethodPut, URL: "/foo"}
	matcher.Match(req, putCandidate)

	getCandidate := cassette.Request{Method: http.MethodGet, URL: "/foo"}
	g.Expect(matcher.Match(req, getCandidate)).To(BeFalse())
}

func TestBarrierMatcher_GivenGETPastMutationForDifferentPath_MatchesGET(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	matcher := newBarrierMatcher(fakeMatcher(), logr.Discard())
	req := &http.Request{Method: http.MethodGet, URL: &url.URL{Path: "/bar"}}

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

	req1 := &http.Request{Method: http.MethodGet, URL: &url.URL{Path: "/foo"}}
	putCandidate := cassette.Request{Method: http.MethodPut, URL: "/foo"}
	matcher.Match(req1, putCandidate)

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

func TestBarrierMatcher_GivenNonGETRequest_BarrierDoesNotBlockGETCandidates(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	// During a DELETE scan, GET candidates past a mutation barrier should NOT be blocked
	// (they wouldn't match on method anyway, and blocking them produces misleading logs)
	matcher := newBarrierMatcher(fakeMatcher(), logr.Discard())
	req := &http.Request{Method: http.MethodDelete, URL: &url.URL{Path: "/bar"}}

	// PUT for /foo creates a mutation barrier
	putCandidate := cassette.Request{Method: http.MethodPut, URL: "/foo"}
	matcher.Match(req, putCandidate)

	// GET candidate for /foo past the barrier — should NOT be blocked during DELETE scan
	// (delegate will reject it on method mismatch instead)
	getCandidate := cassette.Request{Method: http.MethodGet, URL: "/foo"}
	g.Expect(matcher.Match(req, getCandidate)).To(BeFalse(), "should reach delegate, not barrier")
}
