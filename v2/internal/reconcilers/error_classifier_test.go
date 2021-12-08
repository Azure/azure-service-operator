/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package reconcilers_test

import (
	"testing"

	"github.com/Azure/go-autorest/autorest/to"
	. "github.com/onsi/gomega"

	"github.com/Azure/azure-service-operator/v2/internal/genericarmclient"
	"github.com/Azure/azure-service-operator/v2/internal/reconcilers"
)

var badRequestError = &genericarmclient.CloudError{
	InnerError: &genericarmclient.ErrorResponse{
		Code:    to.StringPtr("BadRequest"),
		Message: to.StringPtr("That was not a good request"),
	},
}

var conflictError = &genericarmclient.CloudError{
	InnerError: &genericarmclient.ErrorResponse{
		Code:    to.StringPtr("Conflict"),
		Message: to.StringPtr("That doesn't match what I have"),
	},
}

var retryableConflictError = &genericarmclient.CloudError{
	InnerError: &genericarmclient.ErrorResponse{
		Code:    to.StringPtr("Conflict"),
		Message: to.StringPtr("Umm, other stuff is going on. Try again later?"),
	},
}

var resourceGroupNotFoundError = &genericarmclient.CloudError{
	InnerError: &genericarmclient.ErrorResponse{
		Code:    to.StringPtr("ResourceGroupNotFound"),
		Message: to.StringPtr("The resource group was not found"),
	},
}

var unknownError = &genericarmclient.CloudError{
	InnerError: &genericarmclient.ErrorResponse{
		Code:    to.StringPtr("ThisCodeIsNotACodeUnderstoodByTheClassifier"),
		Message: to.StringPtr("No idea what went wrong"),
	},
}

func Test_NilError_IsRetryable(t *testing.T) {
	g := NewGomegaWithT(t)
	expected := reconcilers.CloudErrorDetails{
		Classification: reconcilers.CloudErrorRetryable,
		Code:           reconcilers.UnknownErrorCode,
		Message:        reconcilers.UnknownErrorMessage,
	}
	g.Expect(reconcilers.ClassifyCloudError(nil)).To(Equal(expected))
}

func Test_Conflict_IsNotRetryable(t *testing.T) {
	g := NewGomegaWithT(t)

	expected := reconcilers.CloudErrorDetails{
		Classification: reconcilers.CloudErrorFatal,
		Code:           to.String(conflictError.InnerError.Code),
		Message:        to.String(conflictError.InnerError.Message),
	}
	g.Expect(reconcilers.ClassifyCloudError(conflictError)).To(Equal(expected))
}

func Test_RetryableConflict_IsRetryable(t *testing.T) {
	g := NewGomegaWithT(t)

	expected := reconcilers.CloudErrorDetails{
		Classification: reconcilers.CloudErrorRetryable,
		Code:           to.String(retryableConflictError.InnerError.Code),
		Message:        to.String(retryableConflictError.InnerError.Message),
	}
	g.Expect(reconcilers.ClassifyCloudError(retryableConflictError)).To(Equal(expected))
}

func Test_BadRequest_IsRetryable(t *testing.T) {
	g := NewGomegaWithT(t)

	expected := reconcilers.CloudErrorDetails{
		Classification: reconcilers.CloudErrorRetryable,
		Code:           to.String(badRequestError.InnerError.Code),
		Message:        to.String(badRequestError.InnerError.Message),
	}
	g.Expect(reconcilers.ClassifyCloudError(badRequestError)).To(Equal(expected))
}

func Test_ResourceGroupNotFound_IsRetryable(t *testing.T) {
	g := NewGomegaWithT(t)

	expected := reconcilers.CloudErrorDetails{
		Classification: reconcilers.CloudErrorRetryable,
		Code:           to.String(resourceGroupNotFoundError.InnerError.Code),
		Message:        to.String(resourceGroupNotFoundError.InnerError.Message),
	}
	g.Expect(reconcilers.ClassifyCloudError(resourceGroupNotFoundError)).To(Equal(expected))
}

func Test_UnknownError_IsRetryable(t *testing.T) {
	g := NewGomegaWithT(t)

	expected := reconcilers.CloudErrorDetails{
		Classification: reconcilers.CloudErrorRetryable,
		Code:           to.String(unknownError.InnerError.Code),
		Message:        to.String(unknownError.InnerError.Message),
	}
	g.Expect(reconcilers.ClassifyCloudError(unknownError)).To(Equal(expected))
}
