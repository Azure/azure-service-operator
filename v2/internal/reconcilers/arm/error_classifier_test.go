/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package arm_test

import (
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	. "github.com/onsi/gomega"

	"github.com/Azure/azure-service-operator/v2/internal/genericarmclient"
	"github.com/Azure/azure-service-operator/v2/internal/reconcilers/arm"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/core"
)

var badRequestError = genericarmclient.NewTestCloudError("BadRequest", "That was not a good request")

var conflictError = genericarmclient.NewTestCloudError("Conflict", "That doesn't match what I have")

var resourceGroupNotFoundError = genericarmclient.NewTestCloudError("ResourceGroupNotFound", "The resource group was not found")

var http400Error = genericarmclient.NewCloudError(&azcore.ResponseError{StatusCode: 400})

var unknownError = genericarmclient.NewTestCloudError("ThisCodeIsNotACodeUnderstoodByTheClassifier", "No idea what went wrong")

func Test_NilError_IsRetryable(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	expected := core.CloudErrorDetails{
		Classification: core.ErrorRetryable,
		Code:           core.UnknownErrorCode,
		Message:        core.UnknownErrorMessage,
	}

	g.Expect(arm.ClassifyCloudError(nil)).To(Equal(expected))
}

func Test_Conflict_IsNotRetryable(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	expected := core.CloudErrorDetails{
		Classification: core.ErrorFatal,
		Code:           conflictError.Code(),
		Message:        conflictError.Message(),
	}
	g.Expect(arm.ClassifyCloudError(conflictError)).To(Equal(expected))
}

func Test_BadRequest_IsFatal(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	expected := core.CloudErrorDetails{
		Classification: core.ErrorFatal,
		Code:           badRequestError.Code(),
		Message:        badRequestError.Message(),
	}
	g.Expect(arm.ClassifyCloudError(badRequestError)).To(Equal(expected))
}

func Test_ResourceGroupNotFound_IsRetryable(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	expected := core.CloudErrorDetails{
		Classification: core.ErrorRetryable,
		Code:           resourceGroupNotFoundError.Code(),
		Message:        resourceGroupNotFoundError.Message(),
	}
	g.Expect(arm.ClassifyCloudError(resourceGroupNotFoundError)).To(Equal(expected))
}

func Test_UnknownError_IsRetryable(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	expected := core.CloudErrorDetails{
		Classification: core.ErrorRetryable,
		Code:           unknownError.Code(),
		Message:        unknownError.Message(),
	}
	g.Expect(arm.ClassifyCloudError(unknownError)).To(Equal(expected))
}

func Test_HTTP400_IsFatal(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	expected := core.CloudErrorDetails{
		Classification: core.ErrorFatal,
		Code:           core.UnknownErrorCode,
		Message:        core.UnknownErrorMessage,
	}

	g.Expect(arm.ClassifyCloudError(http400Error)).To(Equal(expected))
}
