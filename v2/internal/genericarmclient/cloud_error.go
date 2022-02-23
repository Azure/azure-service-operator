/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package genericarmclient

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/core"
)

// CloudError - An error response for a resource management request
// We have to support two different formats for the error as some services do things differently.
//
// The ARM spec says that error details should be nested inside an `error` element.
// See https://github.com/Azure/azure-resource-manager-rpc/blob/master/v1.0/common-api-details.md#error-response-content
//
// However, some services put the code & message at the top level instead.
// This is common enough that the Azure Python SDK has specific handling to promote a nested error to the top level.
// See https://github.com/Azure/azure-sdk-for-python/blob/9791fb5bc4cb6001768e6e1fb986b8d8f8326c43/sdk/core/azure-core/azure/core/exceptions.py#L153
//
type CloudError struct {
	error error

	// READ-ONLY; The error code.
	Code *string `json:"code,omitempty" azure:"ro"`

	// READ-ONLY; The error message.
	Message *string `json:"message,omitempty" azure:"ro"`

	// READ-ONLY; The error target.
	Target *string `json:"target,omitempty" azure:"ro"`

	// Common error response for all Azure Resource Manager APIs to return error details for failed operations.
	// (This also follows the OData error response format.)
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

// ErrorCode returns the error code from the message.
// If InnerError is present, we have an ARM compliant error (our preferred kind), so we use that code if present.
// If InnerError is not present, we return the outer error code if present.
// If neither is present, we return UnknownErrorCode
func (e CloudError) ErrorCode() string {
	if e.InnerError != nil && e.InnerError.Code != nil && *e.InnerError.Code != "" {
		return *e.InnerError.Code
	}

	if e.Code != nil && *e.Code != "" {
		return *e.Code
	}

	return core.UnknownErrorCode
}

// ErrorMessage returns the message from the error.
// If InnerError is present, we have an ARM compliant error (our preferred kind), so we use that message if present.
// If InnerError is not present, we return the outer message if present.
// If neither is present, we return UnknownErrorCode
func (e CloudError) ErrorMessage() string {
	if e.InnerError != nil && e.InnerError.Message != nil && *e.InnerError.Message != "" {
		return *e.InnerError.Message
	}

	if e.Message != nil && *e.Message != "" {
		return *e.Message
	}

	return core.UnknownErrorMessage
}

// ErrorTarget returns the target of the error.
// If InnerError is present, we have an ARM compliant error (our preferred kind), so we use that target if present.
// If InnerError is not present, we return the outer target if present.
// If neither is present, we return an empty string
func (e CloudError) ErrorTarget() string {
	if e.InnerError != nil && e.InnerError.Target != nil {
		return *e.InnerError.Target
	}

	return ""
}

func (e CloudError) Unwrap() error {
	return e.error
}
