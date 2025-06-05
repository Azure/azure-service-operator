/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package errorclassification_test

import (
	"testing"

	. "github.com/onsi/gomega"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"

	"github.com/Azure/azure-service-operator/v2/internal/genericarmclient"
	"github.com/Azure/azure-service-operator/v2/internal/reconcilers/arm/errorclassification"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/core"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/retry"
)

var (
	badRequestError            = genericarmclient.NewTestCloudError("BadRequest", "That was not a good request")
	conflictError              = genericarmclient.NewTestCloudError("Conflict", "That doesn't match what I have")
	resourceGroupNotFoundError = genericarmclient.NewTestCloudError("ResourceGroupNotFound", "The resource group was not found")
	http400Error               = genericarmclient.NewCloudError(&azcore.ResponseError{StatusCode: 400})
	unknownError               = genericarmclient.NewTestCloudError("ThisCodeIsNotACodeUnderstoodByTheClassifier", "No idea what went wrong")
	lockedError                = genericarmclient.NewTestCloudError("ScopeLocked", "The scope 'scope1' cannot perform write operation because following scope(s) are locked: 'scope2'. Please remove the lock and try again.")
)

func Test_NilError_IsRetryable(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	expected := core.CloudErrorDetails{
		Classification: core.ErrorRetryable,
		Code:           core.UnknownErrorCode,
		Message:        core.UnknownErrorMessage,
	}

	g.Expect(errorclassification.ClassifyCloudError(nil)).To(Equal(expected))
}

func Test_Conflict_IsRetryable(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	expected := core.CloudErrorDetails{
		Classification: core.ErrorRetryable,
		Code:           conflictError.Code(),
		Message:        conflictError.Message(),
	}
	g.Expect(errorclassification.ClassifyCloudError(conflictError)).To(Equal(expected))
}

func Test_BadRequest_IsFatal(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	expected := core.CloudErrorDetails{
		Classification: core.ErrorFatal,
		Code:           badRequestError.Code(),
		Message:        badRequestError.Message(),
	}
	g.Expect(errorclassification.ClassifyCloudError(badRequestError)).To(Equal(expected))
}

func Test_Locked_IsRetryableVerySlowly(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	expected := core.CloudErrorDetails{
		Classification: core.ErrorRetryable,
		Code:           lockedError.Code(),
		Message:        lockedError.Message(),
		Retry:          retry.VerySlow,
	}
	g.Expect(errorclassification.ClassifyCloudError(lockedError)).To(Equal(expected))
}

func Test_ResourceGroupNotFound_IsRetryable(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	expected := core.CloudErrorDetails{
		Classification: core.ErrorRetryable,
		Code:           resourceGroupNotFoundError.Code(),
		Message:        resourceGroupNotFoundError.Message(),
	}
	g.Expect(errorclassification.ClassifyCloudError(resourceGroupNotFoundError)).To(Equal(expected))
}

func Test_UnknownError_IsRetryable(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	expected := core.CloudErrorDetails{
		Classification: core.ErrorRetryable,
		Code:           unknownError.Code(),
		Message:        unknownError.Message(),
	}
	g.Expect(errorclassification.ClassifyCloudError(unknownError)).To(Equal(expected))
}

func Test_HTTP400_IsFatal(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	expected := core.CloudErrorDetails{
		Classification: core.ErrorFatal,
		Code:           core.UnknownErrorCode,
		Message:        core.UnknownErrorMessage,
	}

	g.Expect(errorclassification.ClassifyCloudError(http400Error)).To(Equal(expected))
}
