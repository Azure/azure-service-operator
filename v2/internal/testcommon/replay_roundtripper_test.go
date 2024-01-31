/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package testcommon

import (
	"net/http"
	"net/url"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/go-logr/logr"
)

func TestReplayRoundTripperRoundTrip_GivenSingleGET_ReturnsMultipleTimes(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	// Arrange
	req := &http.Request{
		URL:    &url.URL{Path: "/foo"},
		Method: http.MethodGet,
	}

	resp := &http.Response{
		StatusCode: 200,
	}

	fake := newFakeRoundTripper()
	fake.AddResponse(req, resp)

	// Act
	replayer := newReplayRoundTripper(fake, logr.Discard())

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
	}

	resp := &http.Response{
		StatusCode: 200,
	}

	fake := newFakeRoundTripper()
	fake.AddResponse(req, resp)

	// Act
	replayer := newReplayRoundTripper(fake, logr.Discard())

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
	resp, err = replayer.RoundTrip(req)
	g.Expect(err).To(HaveOccurred())
}
