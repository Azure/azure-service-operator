/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package v4

import (
	"errors"
	"io"
	"net/http"
	"net/url"
	"strings"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/go-logr/logr"
	"gopkg.in/dnaeon/go-vcr.v4/pkg/cassette"
)

type closeTrackingBody struct {
	io.Reader
	closed   bool
	closeErr error
}

func (body *closeTrackingBody) Close() error {
	body.closed = true
	return body.closeErr
}

func TestDeleteAwareRoundTripper_SkipsStaleSuccessfulGetsAfterDelete(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		staleStatuses []int
		finalStatus   int
	}{
		"single OK before not found": {
			staleStatuses: []int{http.StatusOK},
			finalStatus:   http.StatusNotFound,
		},
		"multiple successful responses before gone": {
			staleStatuses: []int{http.StatusOK, http.StatusCreated, http.StatusOK},
			finalStatus:   http.StatusGone,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			fake := NewFakeRoundTripper()
			deleteRequest := testRequest(http.MethodDelete, "/parents/p1")
			getRequest := testRequest(http.MethodGet, "/parents/p1/children/c1")
			//nolint:bodyclose // The fake transport owns the response body.
			fake.AddResponse(deleteRequest, testResponse(http.StatusAccepted, http.NoBody))

			bodies := make([]*closeTrackingBody, 0, len(tc.staleStatuses))
			for _, status := range tc.staleStatuses {
				body := &closeTrackingBody{Reader: strings.NewReader("stale")}
				bodies = append(bodies, body)
				//nolint:bodyclose // The fake transport owns the response body.
				fake.AddResponse(getRequest, testResponse(status, body))
			}
			//nolint:bodyclose // The fake transport owns the response body.
			fake.AddResponse(getRequest, testResponse(tc.finalStatus, http.NoBody))

			transport := newDeleteAwareRoundTripper(fake, logr.Discard())
			response, err := transport.RoundTrip(deleteRequest)
			g.Expect(err).NotTo(HaveOccurred())
			g.Expect(response.Body.Close()).To(Succeed())

			response, err = transport.RoundTrip(getRequest)
			g.Expect(err).NotTo(HaveOccurred())
			g.Expect(response.StatusCode).To(Equal(tc.finalStatus))
			g.Expect(response.Body.Close()).To(Succeed())
			for _, body := range bodies {
				g.Expect(body.closed).To(BeTrue())
			}
		})
	}
}

func TestDeleteAwareRoundTripper_ReturnsDeleteLROPollResponses(t *testing.T) {
	t.Parallel()

	tests := map[string]string{
		"operation location":    "Operation-Location",
		"azure async operation": "Azure-AsyncOperation",
		"location":              "Location",
	}

	for name, header := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			fake := NewFakeRoundTripper()
			deleteRequest := &http.Request{
				Method: http.MethodDelete,
				URL: &url.URL{
					Path:     "/parents/p1",
					RawQuery: "api-version=2024-05-01",
				},
			}
			pollRequest := &http.Request{
				Method: http.MethodGet,
				URL: &url.URL{
					Path:     "/parents/p1",
					RawQuery: "api-version=2024-05-01&azure-asyncId=operation1",
				},
			}
			//nolint:bodyclose // The fake transport owns the response body.
			deleteResponse := testResponse(http.StatusAccepted, http.NoBody)
			deleteResponse.Header = make(http.Header)
			deleteResponse.Header.Set(header, pollRequest.URL.String())

			fake.AddResponse(deleteRequest, deleteResponse)
			//nolint:bodyclose // The fake transport owns the response body.
			fake.AddResponse(pollRequest, testResponse(http.StatusOK, io.NopCloser(strings.NewReader(`{"status":"Succeeded"}`))))

			transport := newDeleteAwareRoundTripper(fake, logr.Discard())
			response, err := transport.RoundTrip(deleteRequest)
			g.Expect(err).NotTo(HaveOccurred())
			g.Expect(response.Body.Close()).To(Succeed())

			response, err = transport.RoundTrip(pollRequest)
			g.Expect(err).NotTo(HaveOccurred())
			g.Expect(response.StatusCode).To(Equal(http.StatusOK))
			g.Expect(response.Body.Close()).To(Succeed())
		})
	}
}

func TestDeleteAwareRoundTripper_UsesLROHeaderPrecedence(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	fake := NewFakeRoundTripper()
	deleteRequest := testRequest(http.MethodDelete, "/parents/p1")
	pollRequest := testRequest(http.MethodGet, "/operations/1")
	resourceRequest := testRequest(http.MethodGet, "/parents/p1")
	//nolint:bodyclose // The fake transport owns the response body.
	deleteResponse := testResponse(http.StatusAccepted, http.NoBody)
	deleteResponse.Header = make(http.Header)
	deleteResponse.Header.Set("Azure-AsyncOperation", pollRequest.URL.String())
	deleteResponse.Header.Set("Location", resourceRequest.URL.String())

	fake.AddResponse(deleteRequest, deleteResponse)
	//nolint:bodyclose // The fake transport owns the response body.
	fake.AddResponse(resourceRequest, testResponse(http.StatusOK, http.NoBody))
	//nolint:bodyclose // The fake transport owns the response body.
	fake.AddResponse(resourceRequest, testResponse(http.StatusNotFound, http.NoBody))

	transport := newDeleteAwareRoundTripper(fake, logr.Discard())
	response, err := transport.RoundTrip(deleteRequest)
	g.Expect(err).NotTo(HaveOccurred())
	g.Expect(response.Body.Close()).To(Succeed())

	response, err = transport.RoundTrip(resourceRequest)
	g.Expect(err).NotTo(HaveOccurred())
	g.Expect(response.StatusCode).To(Equal(http.StatusNotFound))
	g.Expect(response.Body.Close()).To(Succeed())
}

func TestDeleteAwareRoundTripper_ReturnsResponsesThatAreNotStale(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		deleteStatus int
		getPath      string
		getStatus    int
	}{
		"failed delete does not mark path": {
			deleteStatus: http.StatusInternalServerError,
			getPath:      "/parents/p1/children/c1",
			getStatus:    http.StatusOK,
		},
		"sibling path is unaffected": {
			deleteStatus: http.StatusAccepted,
			getPath:      "/parents/p2/children/c1",
			getStatus:    http.StatusOK,
		},
		"similar prefix is not a descendant": {
			deleteStatus: http.StatusAccepted,
			getPath:      "/parents/p10",
			getStatus:    http.StatusCreated,
		},
		"transitional status is returned": {
			deleteStatus: http.StatusAccepted,
			getPath:      "/parents/p1/children/c1",
			getStatus:    http.StatusAccepted,
		},
		"server error is returned": {
			deleteStatus: http.StatusAccepted,
			getPath:      "/parents/p1/children/c1",
			getStatus:    http.StatusInternalServerError,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			fake := NewFakeRoundTripper()
			deleteRequest := testRequest(http.MethodDelete, "/parents/p1")
			getRequest := testRequest(http.MethodGet, tc.getPath)
			//nolint:bodyclose // The fake transport owns the response body.
			fake.AddResponse(deleteRequest, testResponse(tc.deleteStatus, http.NoBody))
			//nolint:bodyclose // The fake transport owns the response body.
			fake.AddResponse(getRequest, testResponse(tc.getStatus, http.NoBody))

			transport := newDeleteAwareRoundTripper(fake, logr.Discard())
			response, err := transport.RoundTrip(deleteRequest)
			g.Expect(err).NotTo(HaveOccurred())
			g.Expect(response.Body.Close()).To(Succeed())

			response, err = transport.RoundTrip(getRequest)
			g.Expect(err).NotTo(HaveOccurred())
			g.Expect(response.StatusCode).To(Equal(tc.getStatus))
			g.Expect(response.Body.Close()).To(Succeed())
		})
	}
}

func TestDeleteAwareRoundTripper_ReturnsLastSuccessfulGetWhenLookaheadFails(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	fake := NewFakeRoundTripper()
	deleteRequest := testRequest(http.MethodDelete, "/parents/p1")
	getRequest := testRequest(http.MethodGet, "/parents/p1/children/c1")
	staleBody := &closeTrackingBody{Reader: strings.NewReader("stale")}
	//nolint:bodyclose // The fake transport owns the response body.
	fake.AddResponse(deleteRequest, testResponse(http.StatusAccepted, http.NoBody))
	//nolint:bodyclose // The fake transport owns the response body.
	fake.AddResponse(getRequest, testResponse(http.StatusOK, staleBody))
	fake.AddError(getRequest, cassette.ErrInteractionNotFound)

	transport := newDeleteAwareRoundTripper(fake, logr.Discard())
	response, err := transport.RoundTrip(deleteRequest)
	g.Expect(err).NotTo(HaveOccurred())
	g.Expect(response.Body.Close()).To(Succeed())

	response, err = transport.RoundTrip(getRequest)
	g.Expect(err).NotTo(HaveOccurred())
	g.Expect(response.StatusCode).To(Equal(http.StatusOK))
	g.Expect(staleBody.closed).To(BeFalse())
	g.Expect(response.Body.Close()).To(Succeed())
}

func TestDeleteAwareRoundTripper_UpdatesDeletionStateAfterSuccessfulMutations(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		method         string
		mutationStatus int
		expectedStatus int
	}{
		"successful PUT clears deleted ancestor": {
			method:         http.MethodPut,
			mutationStatus: http.StatusOK,
			expectedStatus: http.StatusOK,
		},
		"successful POST clears deleted ancestor": {
			method:         http.MethodPost,
			mutationStatus: http.StatusCreated,
			expectedStatus: http.StatusOK,
		},
		"successful PATCH clears deleted ancestor": {
			method:         http.MethodPatch,
			mutationStatus: http.StatusOK,
			expectedStatus: http.StatusOK,
		},
		"failed PUT preserves deleted ancestor": {
			method:         http.MethodPut,
			mutationStatus: http.StatusBadRequest,
			expectedStatus: http.StatusNotFound,
		},
		"failed POST preserves deleted ancestor": {
			method:         http.MethodPost,
			mutationStatus: http.StatusBadRequest,
			expectedStatus: http.StatusNotFound,
		},
		"failed PATCH preserves deleted ancestor": {
			method:         http.MethodPatch,
			mutationStatus: http.StatusBadRequest,
			expectedStatus: http.StatusNotFound,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			fake := NewFakeRoundTripper()
			deleteRequest := testRequest(http.MethodDelete, "/parents/p1")
			mutationRequest := testRequest(tc.method, "/parents/p1/children/c1")
			getRequest := testRequest(http.MethodGet, "/parents/p1/children/c1")
			//nolint:bodyclose // The fake transport owns the response body.
			fake.AddResponse(deleteRequest, testResponse(http.StatusAccepted, http.NoBody))
			//nolint:bodyclose // The fake transport owns the response body.
			fake.AddResponse(mutationRequest, testResponse(tc.mutationStatus, http.NoBody))
			//nolint:bodyclose // The fake transport owns the response body.
			fake.AddResponse(getRequest, testResponse(http.StatusOK, http.NoBody))
			//nolint:bodyclose // The fake transport owns the response body.
			fake.AddResponse(getRequest, testResponse(http.StatusNotFound, http.NoBody))

			transport := newDeleteAwareRoundTripper(fake, logr.Discard())
			response, err := transport.RoundTrip(deleteRequest)
			g.Expect(err).NotTo(HaveOccurred())
			g.Expect(response.Body.Close()).To(Succeed())

			response, err = transport.RoundTrip(mutationRequest)
			g.Expect(err).NotTo(HaveOccurred())
			g.Expect(response.Body.Close()).To(Succeed())

			response, err = transport.RoundTrip(getRequest)
			g.Expect(err).NotTo(HaveOccurred())
			g.Expect(response.StatusCode).To(Equal(tc.expectedStatus))
			g.Expect(response.Body.Close()).To(Succeed())
		})
	}
}

func TestDeleteAwareRoundTripper_IsSameOrDescendantPath(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		path     string
		ancestor string
		expected bool
	}{
		"same path":        {path: "/foo", ancestor: "/foo", expected: true},
		"case insensitive": {path: "/FOO/child", ancestor: "/foo", expected: true},
		"descendant":       {path: "/foo/child", ancestor: "/foo", expected: true},
		"trailing slash":   {path: "/foo/child/", ancestor: "/foo/", expected: true},
		"similar prefix":   {path: "/foobar", ancestor: "/foo", expected: false},
		"ancestor longer":  {path: "/foo", ancestor: "/foo/child", expected: false},
		"empty ancestor":   {path: "/foo", ancestor: "", expected: false},
	}

	transport := &deleteAwareRoundTripper{}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)
			g.Expect(transport.isSameOrDescendantPath(tc.path, tc.ancestor)).To(Equal(tc.expected))
		})
	}
}

func TestDeleteAwareRoundTripper_ReturnsStaleBodyCloseError(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	closeErr := errors.New("close failed")
	fake := NewFakeRoundTripper()
	deleteRequest := testRequest(http.MethodDelete, "/parents/p1")
	getRequest := testRequest(http.MethodGet, "/parents/p1/children/c1")
	//nolint:bodyclose // The fake transport owns the response body.
	fake.AddResponse(deleteRequest, testResponse(http.StatusAccepted, http.NoBody))
	//nolint:bodyclose // The fake transport owns the response body.
	fake.AddResponse(getRequest, testResponse(http.StatusOK, &closeTrackingBody{
		Reader:   strings.NewReader("stale"),
		closeErr: closeErr,
	}))
	//nolint:bodyclose // The fake transport owns the response body.
	fake.AddResponse(getRequest, testResponse(http.StatusOK, http.NoBody))

	transport := newDeleteAwareRoundTripper(fake, logr.Discard())
	response, err := transport.RoundTrip(deleteRequest)
	g.Expect(err).NotTo(HaveOccurred())
	g.Expect(response.Body.Close()).To(Succeed())

	//nolint:bodyclose // The expected error has no response body.
	response, err = transport.RoundTrip(getRequest)
	g.Expect(response).To(BeNil())
	g.Expect(err).To(MatchError(ContainSubstring("closing stale GET response body")))
	g.Expect(errors.Is(err, closeErr)).To(BeTrue())
}

func testRequest(method string, path string) *http.Request {
	return &http.Request{
		Method: method,
		URL:    &url.URL{Path: path},
	}
}

func testResponse(status int, body io.ReadCloser) *http.Response {
	return &http.Response{
		StatusCode: status,
		Body:       body,
	}
}
