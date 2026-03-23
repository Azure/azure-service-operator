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

	g.Expect(req.Header.Get(CountHeader)).To(BeEmpty())
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
