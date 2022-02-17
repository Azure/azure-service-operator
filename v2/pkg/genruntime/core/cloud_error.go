/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package core

// CloudError - An error response for a resource management request.
type CloudError struct {
	error error

	// Common error response for all Azure Resource Manager APIs to return error details for failed operations. (This also follows the OData error response
	// format.)
	InnerError *ErrorResponse `json:"error,omitempty"`
}

// NewCloudError returns a new CloudError
func NewCloudError(err error) CloudError {
	return CloudError{
		error: err,
	}
}

// Error implements the error interface for type CloudError.
// The contents of the error text are not contractual and subject to change.
func (e CloudError) Error() string {
	return e.error.Error()
}

func (e CloudError) Unwrap() error { return e.error }
