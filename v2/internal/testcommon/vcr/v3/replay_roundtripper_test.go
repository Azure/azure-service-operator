/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package v3

import (
	"io"
	"net/http"
	"net/url"
	"strings"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/go-logr/logr"

	"github.com/Azure/azure-service-operator/v2/internal/testcommon/vcr"
)

func TestReplayRoundTripperRoundTrip_GivenSingleGET_ReturnsMultipleTimes(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	// Arrange
	req := &http.Request{
		URL:    &url.URL{Path: "/foo"},
		Method: http.MethodGet,
		Body:   io.NopCloser(strings.NewReader("GET body goes here")),
	}

	resp := &http.Response{
		StatusCode: 200,
	}

	fake := vcr.NewFakeRoundTripper()
	fake.AddResponse(req, resp)

	// Act
	replayer := NewReplayRoundTripper(fake, logr.Discard())

	// Assert - first request works
	//nolint:bodyclose // there's no actual body in this response to close
	resp, err := replayer.RoundTrip(req)
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(resp.StatusCode).To(Equal(200))

	// Assert - second request works by replaying the first
	//nolint:bodyclose // there's no actual body in this response to close
	resp, err = replayer.RoundTrip(req)
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(resp.StatusCode).To(Equal(200))

	// Assert - third request works by replaying the first
	//nolint:bodyclose // there's no actual body in this response to close
	resp, err = replayer.RoundTrip(req)
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(resp.StatusCode).To(Equal(200))
}

func TestReplayRoundTripperRoundTrip_GivenSinglePUT_ReturnsOnceExtra(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	// Arrange
	req := &http.Request{
		URL:    &url.URL{Path: "/foo"},
		Method: http.MethodPut,
		Body:   io.NopCloser(strings.NewReader("PUT body goes here")),
	}

	resp := &http.Response{
		StatusCode: 200,
	}

	fake := vcr.NewFakeRoundTripper()
	fake.AddResponse(req, resp)

	// Act
	replayer := NewReplayRoundTripper(fake, logr.Discard())

	// Assert - first request works
	//nolint:bodyclose // there's no actual body in this response to close
	resp, err := replayer.RoundTrip(req)
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(resp.StatusCode).To(Equal(200))

	// Assert - second request works by replaying the first
	//nolint:bodyclose // there's no actual body in this response to close
	resp, err = replayer.RoundTrip(req)
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(resp.StatusCode).To(Equal(200))

	// Assert - third request fails because we've had our one replay
	//nolint:bodyclose // there's no actual body in this response to close
	_, err = replayer.RoundTrip(req)
	g.Expect(err).To(HaveOccurred())
}
