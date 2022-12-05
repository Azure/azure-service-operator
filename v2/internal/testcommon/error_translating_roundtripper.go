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
	"testing"

	"github.com/dnaeon/go-vcr/cassette"
	"github.com/dnaeon/go-vcr/recorder"
	"github.com/google/go-cmp/cmp"
	"github.com/pkg/errors"

	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
)

// translateErrors wraps the given Recorder to handle any "Requested interaction not found"
// and log better information about what the expected request was.
//
// By default the error will be returned to the controller which might ignore/retry it
// and not log any useful information. So instead here we find the recorded request with
// the body that most closely matches what was sent and report the "expected" body.
//
// Ideally we would panic on this error but we don't have a good way to deal with the following
// problem at the moment:
//   - during record the controller does GET (404), PUT, … GET (OK)
//   - during playback the controller does GET (which now returns OK), DELETE, PUT, …
//     and fails due to a missing DELETE recording
func translateErrors(r *recorder.Recorder, cassetteName string, t *testing.T) http.RoundTripper {
	return errorTranslation{r, cassetteName, nil, t}
}

type errorTranslation struct {
	recorder     *recorder.Recorder
	cassetteName string

	cassette *cassette.Cassette
	t        *testing.T
}

func (w errorTranslation) ensureCassette() *cassette.Cassette {
	if w.cassette == nil {
		cassette, err := cassette.Load(w.cassetteName)
		if err != nil {
			panic(fmt.Sprintf("unable to load cassette %q", w.cassetteName))
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

		// Apply the same body filtering that we do in recordings so that the diffs don't show things
		// that we've just removed
		sentBodyString = hideRecordingData(string(bodyBytes))
	}

	// find all request bodies for the specified method/URL combination
	matchingBodies := w.findMatchingBodies(req)

	if len(matchingBodies) == 0 {
		return nil, conditions.NewReadyConditionImpactingError(
			errors.Errorf("cannot find go-vcr recording for request from test %q (cassette: %q) (no responses recorded for this method/URL): %s %s (attempt: %s)\n\n",
				w.t.Name(),
				w.cassetteName,
				req.Method,
				req.URL.String(),
				req.Header.Get(COUNT_HEADER)),
			conditions.ConditionSeverityError,
			conditions.ReasonReconciliationFailedPermanently)
	}

	// locate the request body with the shortest diff from the sent body
	shortestDiff := ""
	for i, bodyString := range matchingBodies {
		diff := cmp.Diff(bodyString, sentBodyString)
		if i == 0 || len(diff) < len(shortestDiff) {
			shortestDiff = diff
		}
	}

	return nil, conditions.NewReadyConditionImpactingError(
		errors.Errorf("cannot find go-vcr recording for request from test %q (cassette: %q) (body mismatch): %s %s\nShortest body diff: %s\n\n",
			w.t.Name(),
			w.cassetteName,
			req.Method,
			req.URL.String(),
			shortestDiff),
		conditions.ConditionSeverityError,
		conditions.ReasonReconciliationFailedPermanently)
}

// finds bodies for interactions where request method, URL, and COUNT_HEADER match
func (w errorTranslation) findMatchingBodies(r *http.Request) []string {
	urlString := r.URL.String()
	var result []string
	for _, interaction := range w.ensureCassette().Interactions {
		if urlString == interaction.URL && r.Method == interaction.Request.Method &&
			r.Header.Get(COUNT_HEADER) == interaction.Request.Headers.Get(COUNT_HEADER) {
			result = append(result, interaction.Request.Body)
		}
	}

	return result
}
