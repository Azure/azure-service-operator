/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package vcr

import (
	"net/http"
	"regexp"
	"strings"

	"github.com/google/uuid"

	"github.com/Azure/azure-service-operator/v2/internal/genericarmclient"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon/creds"
)

type Redactor struct {
	azureIDs   creds.AzureIDs
	redactions map[string]string
}

func NewRedactor(azureIDs creds.AzureIDs) *Redactor {
	return &Redactor{
		azureIDs:   azureIDs,
		redactions: make(map[string]string),
	}
}

func (r *Redactor) AddRedaction(redactionValue string, replacementValue string) {
	r.redactions[redactionValue] = replacementValue
}

var nilGuid = uuid.Nil.String()

// requestHeadersToRemove is the list of request headers to remove when recording or replaying.
var requestHeadersToRemove = []string{
	// remove all Authorization headers from stored requests
	"Authorization",

	// Not needed, adds to diff churn:
	"User-Agent",
}

func (r *Redactor) RedactRequestHeaders(headers http.Header) {
	for _, header := range requestHeadersToRemove {
		delete(headers, header)
	}

	azureIDs := r.azureIDs
	// Hide sensitive request headers
	for _, values := range headers {
		for i := range values {
			values[i] = strings.ReplaceAll(values[i], azureIDs.SubscriptionID, nilGuid)
			values[i] = strings.ReplaceAll(values[i], azureIDs.TenantID, nilGuid)
			if azureIDs.BillingInvoiceID != "" {
				values[i] = strings.ReplaceAll(values[i], azureIDs.BillingInvoiceID, creds.DummyBillingId)
			}
		}
	}
}

// responseHeadersToRemove is the list of response headers to remove when recording or replaying.
var responseHeadersToRemove = []string{
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

func (r *Redactor) RedactResponseHeaders(headers http.Header) {
	for _, header := range responseHeadersToRemove {
		delete(headers, header)
	}

	azureIDs := r.azureIDs
	// Hide sensitive response headers
	for key, values := range headers {
		for i := range values {
			values[i] = strings.ReplaceAll(values[i], azureIDs.SubscriptionID, nilGuid)
			values[i] = strings.ReplaceAll(values[i], azureIDs.TenantID, nilGuid)
			if azureIDs.BillingInvoiceID != "" {
				values[i] = strings.ReplaceAll(values[i], azureIDs.BillingInvoiceID, creds.DummyBillingId)
			}
		}

		// Hide the base request URL in the AzureOperation and Location headers
		if key == genericarmclient.AsyncOperationHeader || key == genericarmclient.LocationHeader {
			for i := range values {
				values[i] = r.HideURLData(values[i])
			}
		}
	}
}

func (r *Redactor) hideRecordingData(s string) string {
	// Hide the subscription ID
	azureIDs := r.azureIDs
	s = strings.ReplaceAll(s, azureIDs.SubscriptionID, nilGuid)

	// Hide the tenant ID
	s = strings.ReplaceAll(s, azureIDs.TenantID, nilGuid)

	// Hide the billing ID
	if azureIDs.BillingInvoiceID != "" {
		s = strings.ReplaceAll(s, azureIDs.BillingInvoiceID, creds.DummyBillingId)
	}

	s = hideDates(s)
	s = hideSSHKeys(s)
	s = hidePasswords(s)
	s = hideKubeConfigs(s)
	s = hideKeys(s)
	s = hideCustomKeys(s)

	return s
}

func (r *Redactor) HideRecordingDataWithCustomRedaction(s string) string {
	// Replace and hide all the custom data
	for k, v := range r.redactions {
		s = strings.ReplaceAll(s, k, v)
	}

	return r.hideRecordingData(s)
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

func (r *Redactor) HideURLData(s string) string {
	s = hideBaseRequestURL(s)

	azureIDs := r.azureIDs
	s = strings.ReplaceAll(s, azureIDs.SubscriptionID, nilGuid)

	// Hide the tenant ID
	s = strings.ReplaceAll(s, azureIDs.TenantID, nilGuid)

	// Hide the billing ID
	if azureIDs.BillingInvoiceID != "" {
		s = strings.ReplaceAll(s, azureIDs.BillingInvoiceID, creds.DummyBillingId)
	}

	return s
}

// baseURLMatcher matches the base part of a URL
var baseURLMatcher = regexp.MustCompile(`^https://[^/]+/`)

func hideBaseRequestURL(s string) string {
	return baseURLMatcher.ReplaceAllLiteralString(s, `https://management.azure.com/`)
}
