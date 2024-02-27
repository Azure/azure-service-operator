/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package v3

import (
	"bytes"
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

func TestReplayRoundTripperRoundTrip_GivenSinglePut_ReturnsOnceExtra(t *testing.T) {
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

func TestReplayRoundTripperRoundTrip_GivenMultiplePUTsToSameURL_ReturnsExpectedBodies(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	// Arrange
	alphaRequest := &http.Request{
		URL:    &url.URL{Path: "/foo"},
		Method: http.MethodPut,
		Body:   io.NopCloser(strings.NewReader("PUT body Alpha goes here")),
	}

	alphaResponse := &http.Response{
		StatusCode: 200,
		Body:       io.NopCloser(strings.NewReader("PUT response Alpha goes here")),
	}

	betaRequest := &http.Request{
		URL:    &url.URL{Path: "/foo"},
		Method: http.MethodPut,
		Body:   io.NopCloser(strings.NewReader("PUT body Beta goes here")),
	}

	betaResponse := &http.Response{
		StatusCode: 200,
		Body:       io.NopCloser(strings.NewReader("PUT response Beta goes here")),
	}

	gammaRequest := &http.Request{
		URL:    &url.URL{Path: "/foo"},
		Method: http.MethodPut,
		Body:   io.NopCloser(strings.NewReader("PUT body Gamma goes here")),
	}

	gammaResponse := &http.Response{
		StatusCode: 200,
		Body:       io.NopCloser(strings.NewReader("PUT response Gamma goes here")),
	}

	fake := vcr.NewFakeRoundTripper()
	fake.AddResponse(alphaRequest, alphaResponse)
	fake.AddResponse(betaRequest, betaResponse)
	fake.AddResponse(gammaRequest, gammaResponse)

	// Act
	replayer := NewReplayRoundTripper(fake, logr.Discard())

	// Assert - first alpha request works
	//nolint:bodyclose // there's no actual body in this response to close
	actual, err := replayer.RoundTrip(alphaRequest)
	g.Expect(err).ToNot(HaveOccurred())
	assertResponse(t, actual, 200, "PUT response Alpha goes here")

	// Assert - first beta request works
	//nolint:bodyclose // there's no actual body in this response to close
	actual, err = replayer.RoundTrip(betaRequest)
	g.Expect(err).ToNot(HaveOccurred())
	assertResponse(t, actual, 200, "PUT response Beta goes here")

	// Assert - first gamma request works
	//nolint:bodyclose // there's no actual body in this response to close
	actual, err = replayer.RoundTrip(gammaRequest)
	g.Expect(err).ToNot(HaveOccurred())
	assertResponse(t, actual, 200, "PUT response Gamma goes here")

	// Assert - second alpha request works by replaying the first
	//nolint:bodyclose // there's no actual body in this response to close
	actual, err = replayer.RoundTrip(alphaRequest)
	g.Expect(err).ToNot(HaveOccurred())
	assertResponse(t, actual, 200, "PUT response Alpha goes here")

	// Assert - second beta request works by replaying the first
	//nolint:bodyclose // there's no actual body in this response to close
	actual, err = replayer.RoundTrip(betaRequest)
	g.Expect(err).ToNot(HaveOccurred())
	assertResponse(t, actual, 200, "PUT response Beta goes here")

	// Assert - second gamma request works by replaying the first
	actual, err = replayer.RoundTrip(gammaRequest)
	g.Expect(err).ToNot(HaveOccurred())
	assertResponse(t, actual, 200, "PUT response Gamma goes here")
}

func assertResponse(
	t *testing.T,
	response *http.Response,
	expectedStatus int,
	expectedBodyContent string,
) {
	t.Helper()
	g := NewGomegaWithT(t)

	g.Expect(response.StatusCode).To(Equal(expectedStatus))

	var body bytes.Buffer
	_, err := body.ReadFrom(response.Body)
	g.Expect(err).ToNot(HaveOccurred())

	g.Expect(string(body.String())).To(ContainSubstring(expectedBodyContent))

	// Reset the body so it can be read again
	response.Body = io.NopCloser(&body)
}
