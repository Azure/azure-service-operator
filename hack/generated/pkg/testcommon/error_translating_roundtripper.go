/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package testcommon

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/dnaeon/go-vcr/cassette"
	"github.com/dnaeon/go-vcr/recorder"
	"github.com/google/go-cmp/cmp"
)

// translateErrors wraps the given Recorder to handle any "Requested interaction not found"
// and log better information about what the expected request was.
//
// By default the error will be returned to the controller which might ignore/retry it
// and not log any useful information. So instead here we find  the recorded request with
// the body that most closely matches what was sent and report the "expected" body.
//
// Ideally we would panic on this error but we don't have a good way to deal with the following
// problem at the moment:
// - during record the controller does GET (404), PUT, … GET (OK)
// - during playback the controller does GET (which now returns OK), DELETE, PUT, …
//   and fails due to a missing DELETE recording
func translateErrors(r *recorder.Recorder, cassetteName string) http.RoundTripper {
	return errorTranslation{r, cassetteName, nil}
}

type errorTranslation struct {
	recorder     *recorder.Recorder
	cassetteName string

	cassette *cassette.Cassette
}

func (w errorTranslation) ensure_cassette() *cassette.Cassette {
	if w.cassette == nil {
		cassette, err := cassette.Load(w.cassetteName)
		if err != nil {
			panic(fmt.Sprintf("unable to load casette %q", w.cassetteName))
		}

		w.cassette = cassette
	}

	return w.cassette
}

func (w errorTranslation) RoundTrip(req *http.Request) (*http.Response, error) {
	resp, originalErr := w.recorder.RoundTrip(req)
	// sorry, go-vcr doesn't expose the error type or message
	if originalErr == nil || !strings.Contains(originalErr.Error(), "interaction not found") {
		return resp, originalErr
	}

	sentBodyString := "<nil>"
	if req.Body != nil {
		bodyBytes, bodyErr := io.ReadAll(req.Body)
		if bodyErr != nil {
			// see invocation of SetMatcher in the createRecorder, which does this
			panic("io.ReadAll(req.Body) failed, this should always succeed because req.Body has been replaced by a buffer")
		}

		sentBodyString = string(bodyBytes)
	}

	// find all request bodies for the specified method/URL combination
	urlString := req.URL.String()
	var bodiesForMethodAndURL []string
	for _, interaction := range w.ensure_cassette().Interactions {
		if urlString == interaction.URL && req.Method == interaction.Request.Method &&
			req.Header.Get(COUNT_HEADER) == interaction.Request.Headers.Get(COUNT_HEADER) {
			bodiesForMethodAndURL = append(bodiesForMethodAndURL, interaction.Request.Body)
			break
		}
	}

	if len(bodiesForMethodAndURL) == 0 {
		fmt.Printf("\n*** Cannot find go-vcr recording for request (no responses recorded for this method/URL): %s %s (attempt: %s)\n\n", req.Method, req.URL.String(), req.Header.Get(COUNT_HEADER))
		return nil, originalErr
	}

	// locate the request body with the shortest diff from the sent body
	shortestDiff := ""
	for i, bodyString := range bodiesForMethodAndURL {
		diff := cmp.Diff(sentBodyString, bodyString)
		if i == 0 || len(diff) < len(shortestDiff) {
			shortestDiff = diff
		}
	}

	fmt.Printf("\n*** Cannot find go-vcr recording for request (body mismatch): %s %s\nShortest body diff: %s\n\n", req.Method, req.URL.String(), shortestDiff)
	return nil, originalErr
}
