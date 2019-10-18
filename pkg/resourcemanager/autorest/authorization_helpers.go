package autorest

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
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"
)

// HeaderAuthorization is the name of the HTTP Authorization Header
const HeaderAuthorization = "Authorization"

// HeaderContentEncoding is the name of the HTTP Content-Encoding Header
const HeaderContentEncoding = "Content-Encoding"

// HeaderContentLength is the name of the HTTP Content-Length Header
const HeaderContentLength = "Content-Length"

// HeaderContentMD5 is the name of the HTTP Content-MD5 Header
const HeaderContentMD5 = "Content-MD5"

// HeaderContentLanguage is the name of the HTTP Content-Language Header
const HeaderContentLanguage = "Content-Language"

// HeaderContentType is the name of the HTTP Content-Type Header
const HeaderContentType = "Content-Type"

// HeaderIfModifiedSince is the name of the HTTP If-Modified-Since Header
const HeaderIfModifiedSince = "If-Modified-Since"

// HeaderIfMatch is the name of the HTTP If-Match Header
const HeaderIfMatch = "If-Match"

// HeaderIfNoneMatch is the name of the HTTP If-None-Match Header
const HeaderIfNoneMatch = "If-None-Match"

// HeaderIfUnmodifiedSince is the name of the HTTP If-Unmodified-Since Header
const HeaderIfUnmodifiedSince = "If-Unmodified-Since"

// HeaderMSDate is the name of the HTTP X-Ms-Date Header
const HeaderMSDate = "X-Ms-Date"

// HeaderRange is the name of the HTTP Range Header
const HeaderRange = "Range"

// StorageEmulatorAccountName is the Storage Account Name for the Azure Storage Emulator
const StorageEmulatorAccountName = "devstoreaccount1"

// StorageEmulatorAccountKey is the Storage Account Key for the Azure Storage Emulator
const StorageEmulatorAccountKey = "Eby8vdM02xNOcqFlqUwJPLlmEtlCDXJ1OUzFT50uSRZ6IFsuFq2UVErCz4I6tq/K1SZFPTOtr/KBHBeksoGMGw=="

// buildCanonicalizedHeader builds the Canonicalized Header required to sign Storage Requests
func buildCanonicalizedHeader(headers http.Header) string {
	cm := make(map[string]string)

	for k, v := range headers {
		headerName := strings.TrimSpace(strings.ToLower(k))
		if strings.HasPrefix(headerName, "x-ms-") {
			cm[headerName] = v[0]
		}
	}

	if len(cm) == 0 {
		return ""
	}

	var keys []string
	for key := range cm {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	ch := bytes.NewBufferString("")

	for _, key := range keys {
		ch.WriteString(key)
		ch.WriteRune(':')
		ch.WriteString(cm[key])
		ch.WriteRune('\n')
	}

	return strings.TrimSuffix(string(ch.Bytes()), "\n")
}

// buildCanonicalizedResource builds the Canonical Resource required for to sign Storage Account requests
func buildCanonicalizedResource(uri, accountName string, sharedKeyLite bool) (*string, error) {
	u, err := url.Parse(uri)
	if err != nil {
		return nil, err
	}

	cr := bytes.NewBufferString("")
	if accountName != StorageEmulatorAccountName {
		cr.WriteString("/")
		cr.WriteString(primaryStorageAccountName(accountName))
	}

	if len(u.Path) > 0 {
		// Any portion of the CanonicalizedResource string that is derived from
		// the resource's URI should be encoded exactly as it is in the URI.
		// -- https://msdn.microsoft.com/en-gb/library/azure/dd179428.aspx
		cr.WriteString(u.EscapedPath())
	}

	if sharedKeyLite {
		if comp := u.Query().Get("comp"); comp != "" {
			cr.WriteString(fmt.Sprintf("?comp=%s", comp))
		}
	} else {
		// Convert all parameter names to lowercase.
		// URL-decode each query parameter name and value.
		// Sort the query parameters lexicographically by parameter name, in ascending order.
		keys := make([]string, 0)
		keysWithValues := make(map[string]string, len(u.Query()))
		for k, unsortedValues := range u.Query() {
			key := strings.ToLower(k)
			keys = append(keys, key)

			// If a query parameter has more than one value, sort all values lexicographically, then include them in a comma-separated list:
			sortedValues := make([]string, 0)
			sortedValues = append(sortedValues, unsortedValues...)
			sort.Strings(sortedValues)
			keysWithValues[key] = strings.Join(sortedValues, ",")
		}
		sort.Strings(keys)

		for _, key := range keys {
			values := keysWithValues[key]
			// Include a new-line character (\n) before each name-value pair.
			// Append each query parameter name and value to the string in the following format, making sure to include the colon (:) between the name and the value
			cr.WriteString(fmt.Sprintf("\n%s:%s", key, values))
		}
	}

	out := cr.String()
	return &out, nil
}

func formatSharedKeyAuthorizationHeader(accountName, key string) string {
	canonicalizedAccountName := primaryStorageAccountName(accountName)
	return fmt.Sprintf("SharedKey %s:%s", canonicalizedAccountName, key)
}

func formatSharedKeyLiteAuthorizationHeader(accountName, key string) string {
	canonicalizedAccountName := primaryStorageAccountName(accountName)
	return fmt.Sprintf("SharedKeyLite %s:%s", canonicalizedAccountName, key)
}

// hmacValue base-64 decodes the storageAccountKey, then signs the string with it
// as outlined here: https://docs.microsoft.com/en-us/rest/api/storageservices/authorize-with-shared-key
func hmacValue(storageAccountKey, canonicalizedString string) string {
	key, err := base64.StdEncoding.DecodeString(storageAccountKey)
	if err != nil {
		return ""
	}

	encr := hmac.New(sha256.New, []byte(key))
	encr.Write([]byte(canonicalizedString))
	return base64.StdEncoding.EncodeToString(encr.Sum(nil))
}

// prepareHeadersForRequest prepares a request so that it can be signed
// by ensuring the `date` and `x-ms-date` headers are set
func prepareHeadersForRequest(r *http.Request) {
	if r.Header == nil {
		r.Header = http.Header{}
	}

	date := time.Now().UTC().Format(http.TimeFormat)

	// a date must be set, X-Ms-Date should be used when both are set; but let's set both for completeness
	r.Header.Set("date", date)
	r.Header.Set("x-ms-date", date)
}

// primaryStorageAccountName returns the name of the primary for a given Storage Account
func primaryStorageAccountName(input string) string {
	// from https://docs.microsoft.com/en-us/rest/api/storageservices/authorize-with-shared-key
	// If you are accessing the secondary location in a storage account for which
	// read-access geo-replication (RA-GRS) is enabled, do not include the
	// -secondary designation in the authorization header.
	// For authorization purposes, the account name is always the name of the primary location,
	// even for secondary access.
	return strings.TrimSuffix(input, "-secondary")
}
