/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package v4

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/go-logr/logr"
	"gopkg.in/dnaeon/go-vcr.v4/pkg/cassette"

	"github.com/Azure/azure-service-operator/v2/internal/testcommon/creds"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon/vcr"
)

func TestReplayRoundTripperRoundTrip_GivenSingleGETReturningTerminalState_ReturnsMultipleTimes(t *testing.T) {
	t.Parallel()

	// Arrange
	req := &http.Request{
		URL:    &url.URL{Path: "/foo"},
		Method: http.MethodGet,
		Body:   io.NopCloser(strings.NewReader("GET body goes here")),
	}

	resp := &http.Response{
		StatusCode: 200,
		Body:       io.NopCloser(strings.NewReader(`{"properties":{"provisioningState": "Succeeded"}}`)),
	}

	fake := vcr.NewFakeRoundTripper(cassette.ErrInteractionNotFound)
	fake.AddResponse(req, resp)

	redactor := vcr.NewRedactor(creds.DummyAzureIDs())

	// Act
	replayer := NewReplayRoundTripper(fake, logr.Discard(), redactor)

	// Assert - first request works
	assertExpectedResponse(t, replayer, req, 200, `{"provisioningState": "Succeeded"`)

	// Assert - second request works by replaying the first
	assertExpectedResponse(t, replayer, req, 200, `{"provisioningState": "Succeeded"`)

	// Assert - third request works by replaying the first
	assertExpectedResponse(t, replayer, req, 200, `{"provisioningState": "Succeeded"`)
}

func TestReplayRoundTripperRoundTrip_GivenSingleGETReturningNonterminalState_ReturnsOnce(t *testing.T) {
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
		Body:       io.NopCloser(strings.NewReader(`{"properties":{"provisioningState": "Deleting"}}`)),
	}

	fake := vcr.NewFakeRoundTripper(cassette.ErrInteractionNotFound)
	fake.AddResponse(req, resp)

	redactor := vcr.NewRedactor(creds.DummyAzureIDs())

	// Act
	replayer := NewReplayRoundTripper(fake, logr.Discard(), redactor)

	// Assert - first request works
	assertExpectedResponse(t, replayer, req, 200, `{"provisioningState": "Deleting"`)

	// Assert - second request fails
	//nolint:bodyclose // there's no actual body in this response to close
	_, err := fake.RoundTrip(req)
	g.Expect(err).To(MatchError(ContainSubstring("requested interaction not found")))
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
		Body:       io.NopCloser(strings.NewReader("PUT response goes here")),
	}

	fake := vcr.NewFakeRoundTripper(cassette.ErrInteractionNotFound)
	fake.AddResponse(req, resp)

	redactor := vcr.NewRedactor(creds.DummyAzureIDs())

	// Act
	replayer := NewReplayRoundTripper(fake, logr.Discard(), redactor)

	// Assert - first request works
	assertExpectedResponse(t, replayer, req, 200, "PUT response goes here")

	// Assert - second request works by replaying the first
	assertExpectedResponse(t, replayer, req, 200, "PUT response goes here")

	// Assert - third request fails because we've had our one replay
	//nolint:bodyclose // response body is a string, no need to close
	_, err := replayer.RoundTrip(req)
	g.Expect(err).To(HaveOccurred())
}

func TestReplayRoundTripperRoundTrip_GivenMultiplePUTsToSameURL_ReturnsExpectedBodies(t *testing.T) {
	t.Parallel()

	// Arrange
	//nolint:bodyclose // response body is a string, no need to close
	alphaRequest, alphaResponse := createPutRequestAndResponse(
		"/foo",
		"Alpha goes here",
		200)

	//nolint:bodyclose // response body is a string, no need to close
	betaRequest, betaResponse := createPutRequestAndResponse(
		"/foo",
		"Beta goes here",
		203)

	//nolint:bodyclose // response body is a string, no need to close
	gammaRequest, gammaResponse := createPutRequestAndResponse(
		"/foo",
		"Gamma goes here",
		200)

	fake := vcr.NewFakeRoundTripper(cassette.ErrInteractionNotFound)
	fake.AddResponse(alphaRequest, alphaResponse)
	fake.AddResponse(betaRequest, betaResponse)
	fake.AddResponse(gammaRequest, gammaResponse)

	redactor := vcr.NewRedactor(creds.DummyAzureIDs())

	// Act
	replayer := NewReplayRoundTripper(fake, logr.Discard(), redactor)

	// Assert - first alpha request works
	assertExpectedResponse(t, replayer, alphaRequest, 200, "Alpha goes here")

	// Assert - first beta request works
	assertExpectedResponse(t, replayer, betaRequest, 203, "Beta goes here")

	// Assert - first gamma request works
	assertExpectedResponse(t, replayer, gammaRequest, 200, "Gamma goes here")

	// Assert - second alpha request works by replaying the first
	assertExpectedResponse(t, replayer, alphaRequest, 200, "Alpha goes here")

	// Assert - second beta request works by replaying the first
	assertExpectedResponse(t, replayer, betaRequest, 203, "Beta goes here")

	// Assert - second gamma request works by replaying the first
	assertExpectedResponse(t, replayer, gammaRequest, 200, "Gamma goes here")
}

func Test_ReplayRoundTripper_WhenCombinedWithTrackingRoundTripper_GivesDesiredResult(t *testing.T) {
	t.Parallel()

	// Arrange - Request and response to create the resource
	//nolint:bodyclose // there's no actual body in this response to close
	creationRequest, creationResponse := createPutRequestAndResponse(
		"/sub/id/resource/A",
		"create resource A",
		200)

	// Arrange - Request and response to update the resource
	//nolint:bodyclose // there's no actual body in this response to close
	updateRequest, updateResponse := createPutRequestAndResponse(
		"/sub/id/resource/A",
		"update resource A",
		200)

	// Arrange - set up fake replayer
	fake := vcr.NewFakeRoundTripper(cassette.ErrInteractionNotFound)
	fake.AddResponse(creationRequest, creationResponse)
	fake.AddError(creationRequest, cassette.ErrInteractionNotFound)
	fake.AddResponse(updateRequest, updateResponse)
	fake.AddError(updateRequest, cassette.ErrInteractionNotFound)

	redactor := vcr.NewRedactor(creds.DummyAzureIDs())

	// Act
	replayRoundTripper := NewReplayRoundTripper(fake, logr.Discard(), redactor)
	replayer := AddTrackingHeaders(replayRoundTripper, redactor)

	// Assert - first PUT to create the resource works
	assertExpectedResponse(t, replayer, creationRequest, 200, "create resource A")

	// Assert - second PUT to create the resource works because of replay
	assertExpectedResponse(t, replayer, creationRequest, 200, "create resource A")

	// Assert - first PUT to update the resource works
	assertExpectedResponse(t, replayer, updateRequest, 200, "update resource A")

	// Assert - second PUT to update the resource works due to replay
	assertExpectedResponse(t, replayer, updateRequest, 200, "update resource A")
}

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

func createPutRequestAndResponse(
	urlpath string,
	body string,
	statusCode int,
) (*http.Request, *http.Response) {
	req := &http.Request{
		URL:    &url.URL{Path: urlpath},
		Method: http.MethodPut,
		Body: io.NopCloser(
			bytes.NewBufferString(
				fmt.Sprintf("PUT for %s", body),
			)),
	}

	res := &http.Response{
		StatusCode: statusCode,
		Body: io.NopCloser(
			bytes.NewBufferString(
				fmt.Sprintf("PUT response for %s", body),
			)),
	}

	return req, res
}

func assertExpectedResponse(
	t *testing.T,
	client http.RoundTripper,
	request *http.Request,
	expectedStatus int,
	expectedBodyContent string,
) {
	t.Helper()
	g := NewGomegaWithT(t)

	// Get the response
	//nolint:bodyclose // there's no actual body in this response to close
	response, err := client.RoundTrip(request)
	g.Expect(err).ToNot(HaveOccurred())

	g.Expect(response.StatusCode).To(Equal(expectedStatus))

	var body bytes.Buffer
	if response.Body != nil {
		_, err = body.ReadFrom(response.Body)
		g.Expect(err).ToNot(HaveOccurred())
		g.Expect(string(body.String())).To(ContainSubstring(expectedBodyContent))
	}

	// Reset the body so it can be read again
	response.Body = io.NopCloser(&body)
}
