/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package vcr

import (
	"net/http"
	"regexp"
)

var RequestHeadersToRemove = []string{
	// remove all Authorization headers from stored requests
	"Authorization",

	// Not needed, adds to diff churn:
	"User-Agent",
}

func RedactRequestHeaders(headers http.Header) {
	for _, header := range RequestHeadersToRemove {
		delete(headers, header)
	}
}

var ResponseHeadersToRemove = []string{
	// Request IDs
	"X-Ms-Arm-Service-Request-Id",
	"X-Ms-Correlation-Request-Id",
	"X-Ms-Request-Id",
	"X-Ms-Routing-Request-Id",
	"X-Ms-Client-Request-Id",
	"Client-Request-Id",

	// Quota limits
	"X-Ms-Ratelimit-Remaining-Subscription-Deletes",
	"X-Ms-Ratelimit-Remaining-Subscription-Reads",
	"X-Ms-Ratelimit-Remaining-Subscription-Writes",

	// Not needed, adds to diff churn
	"Date",
}

func RedactResponseHeaders(headers http.Header) {
	for _, header := range ResponseHeadersToRemove {
		delete(headers, header)
	}
}

func HideRecordingData(s string) string {
	result := hideDates(s)
	result = hideSSHKeys(result)
	result = hidePasswords(result)
	result = hideKubeConfigs(result)
	result = hideKeys(result)
	result = hideCustomKeys(result)

	return result
}

var dateMatcher = regexp.MustCompile(`\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(.\d+)?Z`)

// hideDates replaces all ISO8601 datetimes with a fixed value
// this lets us match requests that may contain time-sensitive information (timestamps, etc)
func hideDates(s string) string {
	return dateMatcher.ReplaceAllLiteralString(s, "2001-02-03T04:05:06Z") // this should be recognizable/parseable as a fake date
}

var sshKeyMatcher = regexp.MustCompile("ssh-rsa [0-9a-zA-Z+/=]+")

// hideSSHKeys hides anything that looks like SSH keys
func hideSSHKeys(s string) string {
	return sshKeyMatcher.ReplaceAllLiteralString(s, "ssh-rsa {KEY}")
}

var passwordMatcher = regexp.MustCompile("\"pass[^\"]*?pass\"")

// hidePasswords hides anything that looks like a generated password
func hidePasswords(s string) string {
	return passwordMatcher.ReplaceAllLiteralString(s, "\"{PASSWORD}\"")
}

// kubeConfigMatcher specifically matches base64 data returned by the AKS get keys API
var kubeConfigMatcher = regexp.MustCompile(`"value": "[a-zA-Z0-9+/]+={0,2}"`)

func hideKubeConfigs(s string) string {
	return kubeConfigMatcher.ReplaceAllLiteralString(s, `"value": "IA=="`) // Have to replace with valid base64 data, so replace with " "
}

// keyMatcher matches any valid base64 value with at least 10 sets of 4 bytes of data that ends in = or ==.
// Both storage account keys and Redis account keys are longer than that and end in = or ==. Note that technically
// base64 values need not end in == or =, but allowing for that in the match will flag tons of false positives as
// any text (including long URLs) have strings of characters that meet this requirement. There are other base64 values
// in the payloads (such as operationResults URLs for polling async operations for some services) that seem to use
// very long base64 strings as well.
var keyMatcher = regexp.MustCompile("(?:[A-Za-z0-9+/]{4}){10,}(?:[A-Za-z0-9+/]{2}==|[A-Za-z0-9+/]{3}=)")

func hideKeys(s string) string {
	return keyMatcher.ReplaceAllLiteralString(s, "{KEY}")
}

var (
	// customKeyMatcher is used to match 'key' or 'Key' followed by the base64 patterns without '=' padding.
	customKeyMatcher  = regexp.MustCompile(`"([a-z]+)?[K-k]ey":"[a-zA-Z0-9+/]+"`)
	customKeyReplacer = regexp.MustCompile(`"(?:[A-Za-z0-9+/]{4}){10,}(?:[A-Za-z0-9+/]{4}|[A-Za-z0-9+/])"`)
)

func hideCustomKeys(s string) string {
	return customKeyMatcher.ReplaceAllStringFunc(s, func(matched string) string {
		return customKeyReplacer.ReplaceAllString(matched, `"{KEY}"`)
	})
}

func HideURLData(s string) string {
	return HideBaseRequestURL(s)
}

// baseURLMatcher matches the base part of a URL
var baseURLMatcher = regexp.MustCompile(`^https://[^/]+/`)

func HideBaseRequestURL(s string) string {
	return baseURLMatcher.ReplaceAllLiteralString(s, `https://management.azure.com/`)
}
