// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package errhelp

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/Azure/azure-service-operator/pkg/helpers"
)

// ErrIdsRegex is used to find and remove uuids from errors
var ErrIdsRegex *regexp.Regexp

// ErrTimesRegex allows timestamp seconds to be removed from error strings
var ErrTimesRegex *regexp.Regexp

// StripErrorIDs takes an error and returns its string representation after filtering some common ID patterns
func StripErrorIDs(err error) string {
	patterns := []string{
		"RequestID=",
		"CorrelationId:\\s",
		"Tracking ID: ",
		"requestId",
	}

	if ErrIdsRegex == nil {
		ErrIdsRegex = regexp.MustCompile(fmt.Sprintf(`(%s)\S+`, strings.Join(patterns, "|")))
	}

	return ErrIdsRegex.ReplaceAllString(err.Error(), "")

}

// StripErrorTimes removes the hours:minutes:seconds from a date to prevent updates to Status.Message from changing unnecessarily
func StripErrorTimes(err string) string {
	if ErrTimesRegex == nil {
		ErrTimesRegex = regexp.MustCompile(`(T\d\d:\d\d:\d\d)\"`)
	}
	return ErrTimesRegex.ReplaceAllString(err, "")

}

// IsErrorFatal checks the given error against the provided list of allowed and unrecoverable error types.
//   - Allowed errors are NOT fatal and no error is returned. When returned to the async_reconciler this means that
//     reconciliation is reattempted.
//   - Unrecoverable errors are fatal. When returned to the async_reconciler reconciliation is stopped until
//     a new modification is made to the resource in question. This is useful for things like client errors that
//     no amount of reconciliation will fix.
//
// If an error is not in the allowed list and also not in the unrecoverable list, it is classified as nonfatal,
// but an error is returned. When returned to the async_reconciler, reconciliation will continue but an error will
// be logged.
func IsErrorFatal(err error, allowedErrorTypes []string, unrecoverableErrorTypes []string) (bool, error) {
	azerr := NewAzureError(err)
	if helpers.ContainsString(allowedErrorTypes, azerr.Type) {
		return false, nil // false means the resource is not in a terminal state yet, keep trying to reconcile.
	}
	if helpers.ContainsString(unrecoverableErrorTypes, azerr.Type) {
		// Unrecoverable error, so stop reconcilation
		return true, nil
	}

	// We don't know how to classify this error, so bubble it up in the operator logs but don't assume it's
	// unrecoverable/terminal
	return false, err
}
