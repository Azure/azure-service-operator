/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package v4

import (
	"io"
	"net/http"
	"net/url"
	"strings"
	"testing"

	. "github.com/onsi/gomega"

	"gopkg.in/dnaeon/go-vcr.v4/pkg/cassette"

	"github.com/Azure/azure-service-operator/v2/internal/testcommon/creds"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon/vcr"
)

func TestTrackingRoundTripper_GivenGETRequest_AddsNoTrackingHeader(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

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

	g.Expect(req.Header.Get(HashHeader)).To(BeEmpty())
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

	g.Expect(req.Header.Get(HashHeader)).To(BeEmpty())
}

func TestTrackingRoundTripper_GivenPATCHRequestsWithReorderedJSON_ProducesSameCanonicalHash(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	fake := vcr.NewFakeRoundTripper(cassette.ErrInteractionNotFound)
	redactor := vcr.NewRedactor(creds.DummyAzureIDs())
	tracker := AddTrackingHeaders(fake, redactor)

	// Two JSON bodies with the same content but different key order
	bodyA := `{"@odata.type":"#microsoft.graph.group","infoCatalogs":[],"creationOptions":[],"@odata.context":"https://graph.microsoft.com/v1.0/$metadata#groups/$entity"}`
	bodyB := `{"@odata.type":"#microsoft.graph.group","@odata.context":"https://graph.microsoft.com/v1.0/$metadata#groups/$entity","creationOptions":[],"infoCatalogs":[]}`

	reqA := &http.Request{
		URL:    &url.URL{Path: "/v1.0/groups/some-id"},
		Method: http.MethodPatch,
		Header: make(http.Header),
		Body:   io.NopCloser(strings.NewReader(bodyA)),
	}

	reqB := &http.Request{
		URL:    &url.URL{Path: "/v1.0/groups/some-id"},
		Method: http.MethodPatch,
		Header: make(http.Header),
		Body:   io.NopCloser(strings.NewReader(bodyB)),
	}

	//nolint:bodyclose
	_, _ = tracker.RoundTrip(reqA)
	//nolint:bodyclose
	_, _ = tracker.RoundTrip(reqB)

	// Raw hashes differ because key order differs
	rawA := reqA.Header.Get(HashHeader)
	rawB := reqB.Header.Get(HashHeader)
	g.Expect(rawA).ToNot(BeEmpty())
	g.Expect(rawA).ToNot(Equal(rawB), "Raw hashes should differ for reordered JSON")

	// Canonical hashes match because canonicalization sorts keys
	canonA := reqA.Header.Get(CanonicalHashHeader)
	canonB := reqB.Header.Get(CanonicalHashHeader)
	g.Expect(canonA).ToNot(BeEmpty())
	g.Expect(canonA).To(Equal(canonB), "Canonical hashes should match for reordered JSON")
}

func TestCanonicalizeJSON_GivenValidJSON_SortsKeys(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	input := `{"z":"last","a":"first","m":"middle"}`
	expected := `{"a":"first","m":"middle","z":"last"}`

	g.Expect(canonicalizeJSON(input)).To(Equal(expected))
}

func TestCanonicalizeJSON_GivenInvalidJSON_ReturnsUnchanged(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	input := `not valid json`
	g.Expect(canonicalizeJSON(input)).To(Equal(input))
}

func TestCanonicalizeJSON_GivenNestedJSON_SortsAllKeys(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	input := `{"b":{"z":1,"a":2},"a":1}`
	expected := `{"a":1,"b":{"a":2,"z":1}}`

	g.Expect(canonicalizeJSON(input)).To(Equal(expected))
}

func TestCanonicalizeJSON_GivenHTMLChars_DoesNotEscape(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	input := `{"url":"https://example.com?a=1&b=2","tag":"<div>"}`
	expected := `{"tag":"<div>","url":"https://example.com?a=1&b=2"}`

	g.Expect(canonicalizeJSON(input)).To(Equal(expected))
}
