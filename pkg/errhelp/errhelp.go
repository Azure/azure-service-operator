// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package errhelp

import (
	"strings"
)

// IsParentNotFound checks if the error is about a parent resrouce not existing
func IsParentNotFound(err error) bool {
	return strings.Contains(err.Error(), "ParentResourceNotFound")
}

// IsGroupNotFound checks if error is about resource group not existing
func IsGroupNotFound(err error) bool {
	return strings.Contains(err.Error(), "ResourceGroupNotFound")
}

// IsNotActive checks if error is mentioning a non active resource
func IsNotActive(err error) bool {
	return strings.Contains(err.Error(), "not active")
}

// IsAsynchronousOperationNotComplete checks if error reports an asynchronous operation not completed
func IsAsynchronousOperationNotComplete(err error) bool {
	return strings.Contains(err.Error(), "asynchronous operation has not completed")
}

// IsStatusCode204 checks if the error reports a status code 204 failure to respond to request
func IsStatusCode204(err error) bool {
	return strings.Contains(err.Error(), "StatusCode=204")
}

// IsStatusCode404 checks if the error reports a status code 404 resource not found
func IsStatusCode404(err error) bool {
	return strings.Contains(err.Error(), "StatusCode=404")
}

// IsResourceNotFound checks if error reports that a referenced resource is not found
func IsResourceNotFound(err error) bool {
	return strings.Contains(err.Error(), "ResourceNotFound")
}
