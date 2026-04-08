/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package vcr

import (
	"net/http"
)

// FakeRoundTripper is a fake implementation of http.RoundTripper used in testing.
type FakeRoundTripper struct {
	responses map[string][]fakeRoundTripResponse
	notFound  error // Error to return when no response is found
}

type fakeRoundTripResponse struct {
	response *http.Response
	err      error
}

var _ http.RoundTripper = &FakeRoundTripper{}

func NewFakeRoundTripper(notFound error) *FakeRoundTripper {
	return &FakeRoundTripper{
		responses: make(map[string][]fakeRoundTripResponse),
		notFound:  notFound,
	}
}

// RoundTrip implements http.RoundTripper.
func (fake *FakeRoundTripper) RoundTrip(request *http.Request) (*http.Response, error) {
	key := request.URL.String()
	if available, ok := fake.responses[key]; ok {
		if len(available) > 0 {
			response := available[0]
			fake.responses[key] = available[1:]
			return response.response, response.err
		}

		return nil, fake.notFound
	}

	return nil, fake.notFound
}

// AddResponse adds a response to the fake round tripper.
func (fake *FakeRoundTripper) AddResponse(request *http.Request, response *http.Response) {
	key := request.URL.String()
	fake.responses[key] = append(fake.responses[key], fakeRoundTripResponse{response: response})
}

// AddError adds an error to the fake round tripper.
func (fake *FakeRoundTripper) AddError(request *http.Request, err error) {
	key := request.URL.String()
	fake.responses[key] = append(fake.responses[key], fakeRoundTripResponse{err: err})
}
