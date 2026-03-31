/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package v4

import (
	"bytes"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/Azure/azure-service-operator/v2/internal/testcommon/vcr"
)

const (
	// HashHeader is the name of the header used to record the hash of a request body
	HashHeader = "TEST-REQUEST-HASH"

	// CanonicalHashHeader is the name of the header used to record the hash of a canonicalized request body.
	// JSON canonicalization sorts keys alphabetically, ensuring deterministic hashing regardless of Go map
	// iteration order. Both headers are recorded; during matching, a request matches if either hash matches.
	CanonicalHashHeader = "TEST-REQUEST-CANONICAL-HASH"
)

// Wraps an inner HTTP roundtripper to add a
// counter for duplicated request URIs. This
// is then used to match up requests in the recorder
// - it is needed as we have multiple requests with
// the same Request URL and it will return the first
// one that matches.
type requestCounter struct {
	inner    http.RoundTripper
	redactor *vcr.Redactor
}

func AddTrackingHeaders(inner http.RoundTripper, redactor *vcr.Redactor) *requestCounter {
	return &requestCounter{
		inner:    inner,
		redactor: redactor,
	}
}

var _ http.RoundTripper = &requestCounter{}

func (rt *requestCounter) RoundTrip(req *http.Request) (*http.Response, error) {
	if rt.hasBody(req) {
		rt.addHashHeader(req)
	}

	return rt.inner.RoundTrip(req)
}

// hasBody returns true if the request has a body that can be hashed
func (rt *requestCounter) hasBody(req *http.Request) bool {
	if req.Method != http.MethodPut && req.Method != http.MethodPost && req.Method != http.MethodPatch {
		return false
	}

	return req.Body != nil
}

// addHashHeader adds headers to the request based on hashes of the request body content.
// Two hash headers are written: one for the raw body and one for the canonicalized (key-sorted) body.
// During matching, a request matches a recorded interaction if either hash matches. This provides
// backward compatibility with existing recordings (raw hash matches) while also handling
// non-deterministic JSON key ordering from Go map serialization (canonical hash matches).
func (rt *requestCounter) addHashHeader(request *http.Request) {
	// Read all the content of the request body
	var body bytes.Buffer
	_, err := body.ReadFrom(request.Body)
	if err != nil {
		// Should never fail
		panic(fmt.Sprintf("reading request.Body failed: %s", err))
	}

	// Apply the same body filtering that we do in recordings so that the hash is consistent
	bodyString := rt.redactor.HideRecordingData(body.String())

	// Set the headers
	if request.Header == nil {
		request.Header = make(http.Header)
	}

	// Raw hash — matches existing recordings and deterministic bodies
	rawHash := sha256.Sum256([]byte(bodyString))
	request.Header.Set(HashHeader, fmt.Sprintf("%x", rawHash))

	// Canonical hash — matches even when JSON key ordering is non-deterministic
	canonicalBody := canonicalizeJSON(bodyString)
	canonicalHash := sha256.Sum256([]byte(canonicalBody))
	request.Header.Set(CanonicalHashHeader, fmt.Sprintf("%x", canonicalHash))

	// Reset the body so it can be read again
	request.Body = io.NopCloser(&body)
}

// canonicalizeJSON returns a canonical form of the given JSON string with keys sorted.
// If the input is not valid JSON, it is returned unchanged.
// Go's encoding/json sorts map keys alphabetically, so unmarshaling and re-marshaling
// produces a deterministic byte sequence regardless of the original key order.
// We use SetEscapeHTML(false) to avoid altering characters like &, <, > that json.Marshal
// would otherwise escape to \u0026, \u003c, \u003e.
func canonicalizeJSON(s string) string {
	var v any
	if err := json.Unmarshal([]byte(s), &v); err != nil {
		return s
	}

	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	enc.SetEscapeHTML(false)
	if err := enc.Encode(v); err != nil {
		return s
	}

	// json.Encoder.Encode appends a trailing newline; remove it
	return strings.TrimRight(buf.String(), "\n")
}
