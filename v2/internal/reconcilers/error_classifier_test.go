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
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/core"
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
	t.Parallel()
	g := NewGomegaWithT(t)
	expected := core.CloudErrorDetails{
		Classification: core.ErrorRetryable,
		Code:           core.UnknownErrorCode,
		Message:        core.UnknownErrorMessage,
	}
	
	g.Expect(reconcilers.ClassifyCloudError(nil)).To(Equal(expected))
}

func Test_Conflict_IsNotRetryable(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	expected := core.CloudErrorDetails{
		Classification: core.ErrorFatal,
		Code:           conflictError.ErrorCode(),
		Message:        conflictError.ErrorMessage(),
	}

	g.Expect(reconcilers.ClassifyCloudError(conflictError)).To(Equal(expected))
}

func Test_BadRequest_IsRetryable(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	expected := core.CloudErrorDetails{
		Classification: core.ErrorRetryable,
		Code:           badRequestError.ErrorCode(),
		Message:        badRequestError.ErrorMessage(),
	}

	g.Expect(reconcilers.ClassifyCloudError(badRequestError)).To(Equal(expected))
}

func Test_ResourceGroupNotFound_IsRetryable(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	expected := core.CloudErrorDetails{
		Classification: core.ErrorRetryable,
		Code:           resourceGroupNotFoundError.ErrorCode(),
		Message:        resourceGroupNotFoundError.ErrorMessage(),
	}

	g.Expect(reconcilers.ClassifyCloudError(resourceGroupNotFoundError)).To(Equal(expected))
}

func Test_UnknownError_IsRetryable(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	expected := core.CloudErrorDetails{
		Classification: core.ErrorRetryable,
		Code:           unknownError.ErrorCode(),
		Message:        unknownError.ErrorMessage(),
	}

	g.Expect(reconcilers.ClassifyCloudError(unknownError)).To(Equal(expected))
}
