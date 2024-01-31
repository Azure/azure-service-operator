/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package testcommon

import (
	"net/http"

	"gopkg.in/dnaeon/go-vcr.v3/cassette"
)

// fakeRoundTripper is a fake implementation of http.RoundTripper used in testing.
type fakeRoundTripper struct {
	responses map[string][]fakeRoundTripResponse
}

type fakeRoundTripResponse struct {
	response *http.Response
	err      error
}

var _ http.RoundTripper = &fakeRoundTripper{}

func newFakeRoundTripper() *fakeRoundTripper {
	return &fakeRoundTripper{
		responses: make(map[string][]fakeRoundTripResponse),
	}
}

// RoundTrip implements http.RoundTripper.
func (fake *fakeRoundTripper) RoundTrip(request *http.Request) (*http.Response, error) {
	key := request.URL.String()
	if available, ok := fake.responses[key]; ok {
		if len(available) > 0 {
			response := available[0]
			fake.responses[key] = available[1:]
			return response.response, response.err
		}

		return nil, cassette.ErrInteractionNotFound
	}

	return nil, cassette.ErrInteractionNotFound
}

// AddResponse adds a response to the fake round tripper.
func (fake *fakeRoundTripper) AddResponse(request *http.Request, response *http.Response) {
	key := request.URL.String()
	fake.responses[key] = append(fake.responses[key], fakeRoundTripResponse{response: response})
}

// AddError adds an error to the fake round tripper.
func (fake *fakeRoundTripper) AddError(request *http.Request, err error) {
	key := request.URL.String()
	fake.responses[key] = append(fake.responses[key], fakeRoundTripResponse{err: err})
}
