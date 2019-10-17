package adlsgen2s

// Copyright 2017 Microsoft Corporation
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

import (
	. "github.com/Azure/go-autorest/autorest" 
	"net/http"
	"strings"
)

// SharedKeyAuthorizer implements an authorization for Shared Key
// this can be used for interaction with Blob, File and Queue Storage Endpoints
type SharedKeyAuthorizer struct {
	storageAccountName string
	storageAccountKey  string
}

// NewSharedKeyAuthorizer creates a SharedKeyAuthorizer using the given credentials
func NewSharedKeyAuthorizer(accountName, accountKey string) *SharedKeyAuthorizer {
	return &SharedKeyAuthorizer{
		storageAccountName: accountName,
		storageAccountKey:  accountKey,
	}
}

// WithAuthorization returns a PrepareDecorator that adds an HTTP Authorization header whose
// value is "SharedKey " followed by the computed key.
// This can be used for the Blob, Queue, and File Services
//
// from: https://docs.microsoft.com/en-us/rest/api/storageservices/authorize-with-shared-key
// You may use Shared Key Lite authorization to authorize a request made against the
// 2009-09-19 version and later of the Blob and Queue services,
// and version 2014-02-14 and later of the File services.
func (skl *SharedKeyAuthorizer) WithAuthorization() PrepareDecorator {
	return func(p Preparer) Preparer {
		return PreparerFunc(func(r *http.Request) (*http.Request, error) {
			r, err := p.Prepare(r)
			if err != nil {
				return r, err
			}

			key, err := buildSharedKey(skl.storageAccountName, skl.storageAccountKey, r)
			if err != nil {
				return r, err
			}

			sharedKeyHeader := formatSharedKeyAuthorizationHeader(skl.storageAccountName, *key)
			return Prepare(r, WithHeader(HeaderAuthorization, sharedKeyHeader))
		})
	}
}

func buildSharedKey(accountName, storageAccountKey string, r *http.Request) (*string, error) {
	// first ensure the relevant headers are configured
	prepareHeadersForRequest(r)

	sharedKey, err := computeSharedKey(r.Method, r.URL.String(), accountName, r.Header)
	if err != nil {
		return nil, err
	}

	// we then need to HMAC that value
	hmacdValue := hmacValue(storageAccountKey, *sharedKey)
	return &hmacdValue, nil
}

// computeSharedKey computes the Shared Key Lite required for Storage Authentication
// NOTE: this function assumes that the `x-ms-date` field is set
func computeSharedKey(verb, url string, accountName string, headers http.Header) (*string, error) {
	canonicalizedResource, err := buildCanonicalizedResource(url, accountName, false)
	if err != nil {
		return nil, err
	}

	canonicalizedHeaders := buildCanonicalizedHeader(headers)
	canonicalizedString := buildCanonicalizedStringForSharedKey(verb, headers, canonicalizedHeaders, *canonicalizedResource)
	return &canonicalizedString, nil
}

func buildCanonicalizedStringForSharedKey(verb string, headers http.Header, canonicalizedHeaders, canonicalizedResource string) string {
	lengthString := headers.Get(HeaderContentLength)

	// empty string when zero
	if lengthString == "0" {
		lengthString = ""
	}

	return strings.Join([]string{
		verb,                                 // HTTP Verb
		headers.Get(HeaderContentEncoding),   // Content-Encoding
		headers.Get(HeaderContentLanguage),   // Content-Language
		lengthString,                         // Content-Length (empty string when zero)
		headers.Get(HeaderContentMD5),        // Content-MD5
		headers.Get(HeaderContentType),       // Content-Type
		"",                                   // date should be nil, apparently :shrug:
		headers.Get(HeaderIfModifiedSince),   // If-Modified-Since
		headers.Get(HeaderIfMatch),           // If-Match
		headers.Get(HeaderIfNoneMatch),       // If-None-Match
		headers.Get(HeaderIfUnmodifiedSince), // If-Unmodified-Since
		headers.Get(HeaderRange),             // Range
		canonicalizedHeaders,
		canonicalizedResource,
	}, "\n")
}

// SharedKeyLiteAuthorizer implements an authorization for Shared Key Lite
// this can be used for interaction with Blob, File and Queue Storage Endpoints
type SharedKeyLiteAuthorizer struct {
	storageAccountName string
	storageAccountKey  string
}

// NewSharedKeyLiteAuthorizer crates a SharedKeyLiteAuthorizer using the given credentials
func NewSharedKeyLiteAuthorizer(accountName, accountKey string) *SharedKeyLiteAuthorizer {
	return &SharedKeyLiteAuthorizer{
		storageAccountName: accountName,
		storageAccountKey:  accountKey,
	}
}

// WithAuthorization returns a PrepareDecorator that adds an HTTP Authorization header whose
// value is "SharedKeyLite " followed by the computed key.
// This can be used for the Blob, Queue, and File Services
//
// from: https://docs.microsoft.com/en-us/rest/api/storageservices/authorize-with-shared-key
// You may use Shared Key Lite authorization to authorize a request made against the
// 2009-09-19 version and later of the Blob and Queue services,
// and version 2014-02-14 and later of the File services.
func (skl *SharedKeyLiteAuthorizer) WithAuthorization() PrepareDecorator {
	return func(p Preparer) Preparer {
		return PreparerFunc(func(r *http.Request) (*http.Request, error) {
			r, err := p.Prepare(r)
			if err != nil {
				return r, err
			}

			key, err := buildSharedKeyLite(skl.storageAccountName, skl.storageAccountKey, r)
			if err != nil {
				return r, err
			}

			sharedKeyHeader := formatSharedKeyLiteAuthorizationHeader(skl.storageAccountName, *key)
			return Prepare(r, WithHeader(HeaderAuthorization, sharedKeyHeader))
		})
	}
}

func buildSharedKeyLite(accountName, storageAccountKey string, r *http.Request) (*string, error) {
	// first ensure the relevant headers are configured
	prepareHeadersForRequest(r)

	sharedKey, err := computeSharedKeyLite(r.Method, r.URL.String(), accountName, r.Header)
	if err != nil {
		return nil, err
	}

	// we then need to HMAC that value
	hmacdValue := hmacValue(storageAccountKey, *sharedKey)
	return &hmacdValue, nil
}

// computeSharedKeyLite computes the Shared Key Lite required for Storage Authentication
// NOTE: this function assumes that the `x-ms-date` field is set
func computeSharedKeyLite(verb, url string, accountName string, headers http.Header) (*string, error) {
	canonicalizedResource, err := buildCanonicalizedResource(url, accountName, false)
	if err != nil {
		return nil, err
	}

	canonicalizedHeaders := buildCanonicalizedHeader(headers)
	canonicalizedString := buildCanonicalizedStringForSharedKeyLite(verb, headers, canonicalizedHeaders, *canonicalizedResource)
	return &canonicalizedString, nil
}

func buildCanonicalizedStringForSharedKeyLite(verb string, headers http.Header, canonicalizedHeaders, canonicalizedResource string) string {
	return strings.Join([]string{
		verb,
		headers.Get(HeaderContentMD5),
		headers.Get(HeaderContentType),
		"", // date should be nil, apparently :shrug:
		canonicalizedHeaders,
		canonicalizedResource,
	}, "\n")
}

// SharedKeyLiteTableAuthorizer implements an authorization for Shared Key Lite
// this can be used for interaction with Table Storage Endpoints
type SharedKeyLiteTableAuthorizer struct {
	storageAccountName string
	storageAccountKey  string
}

// NewSharedKeyLiteTableAuthorizer creates a SharedKeyLiteAuthorizer using the given credentials
func NewSharedKeyLiteTableAuthorizer(accountName, accountKey string) *SharedKeyLiteTableAuthorizer {
	return &SharedKeyLiteTableAuthorizer{
		storageAccountName: accountName,
		storageAccountKey:  accountKey,
	}
}

// WithAuthorization returns a PrepareDecorator that adds an HTTP Authorization header whose
// value is "SharedKeyLite " followed by the computed key.
// This can be used for the Blob, Queue, and File Services
//
// from: https://docs.microsoft.com/en-us/rest/api/storageservices/authorize-with-shared-key
// The Blob, Queue, Table, and File services support the following Shared Key authorization
// schemes for version 2009-09-19 and later
func (skl *SharedKeyLiteTableAuthorizer) WithAuthorization() PrepareDecorator {
	return func(p Preparer) Preparer {
		return PreparerFunc(func(r *http.Request) (*http.Request, error) {
			r, err := p.Prepare(r)
			if err != nil {
				return r, err
			}

			key, err := buildSharedKeyLiteTable(skl.storageAccountName, skl.storageAccountKey, r)
			if err != nil {
				return r, err
			}

			sharedKeyHeader := formatSharedKeyLiteAuthorizationHeader(skl.storageAccountName, *key)
			return Prepare(r, WithHeader(HeaderAuthorization, sharedKeyHeader))
		})
	}
}

func buildSharedKeyLiteTable(accountName, storageAccountKey string, r *http.Request) (*string, error) {
	// first ensure the relevant headers are configured
	prepareHeadersForRequest(r)

	sharedKey, err := computeSharedKeyLiteTable(r.URL.String(), accountName, r.Header)
	if err != nil {
		return nil, err
	}

	// we then need to HMAC that value
	hmacdValue := hmacValue(storageAccountKey, *sharedKey)
	return &hmacdValue, nil
}

// computeSharedKeyLite computes the Shared Key Lite required for Storage Authentication
// NOTE: this function assumes that the `x-ms-date` field is set
func computeSharedKeyLiteTable(url string, accountName string, headers http.Header) (*string, error) {
	dateHeader := headers.Get("x-ms-date")
	canonicalizedResource, err := buildCanonicalizedResource(url, accountName, false)
	if err != nil {
		return nil, err
	}

	canonicalizedString := buildCanonicalizedStringForSharedKeyLiteTable(*canonicalizedResource, dateHeader)
	return &canonicalizedString, nil
}

func buildCanonicalizedStringForSharedKeyLiteTable(canonicalizedResource, dateHeader string) string {
	return strings.Join([]string{
		dateHeader,
		canonicalizedResource,
	}, "\n")
}
