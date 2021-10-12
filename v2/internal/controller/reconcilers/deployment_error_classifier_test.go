/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package reconcilers_test

import (
	"testing"

	. "github.com/onsi/gomega"

	"github.com/Azure/azure-service-operator/v2/internal/controller/armclient"
	"github.com/Azure/azure-service-operator/v2/internal/controller/reconcilers"
)

var badRequestError = armclient.DeploymentError{
	Code:    "BadRequest",
	Message: "That was not a good request",
}

var badRequestError2 = armclient.DeploymentError{
	Code:    "BadRequest",
	Message: "There was something wrong with that request",
}

var conflictError = armclient.DeploymentError{
	Code:    "Conflict",
	Message: "That doesn't match what I have",
}

var badRequestFormatError = armclient.DeploymentError{
	Code:    "BadRequestFormat",
	Message: "I couldn't understand you",
}

var resourceGroupNotFoundError = armclient.DeploymentError{
	Code:    "ResourceGroupNotFound",
	Message: "The resource group was not found",
}

var resourceNotFound = armclient.DeploymentError{
	Code:    "ResourceNotFound",
	Message: "The resource was not found",
}

var unknownError = armclient.DeploymentError{
	Code:    "ThisCodeIsNotACodeUnderstoodByTheClassifier",
	Message: "No idea what went wrong",
}

func newError(inner ...armclient.DeploymentError) *armclient.DeploymentError {
	return &armclient.DeploymentError{
		Code:    "DeploymentFailed",
		Message: "The deployment failed",
		Details: inner,
	}
}

func Test_NilError_IsRetryable(t *testing.T) {
	g := NewGomegaWithT(t)
	expected := reconcilers.DeploymentErrorDetails{
		Classification: reconcilers.DeploymentErrorRetryable,
		Code:           reconcilers.UnknownErrorCode,
		Message:        reconcilers.UnknownErrorMessage,
	}
	g.Expect(reconcilers.ClassifyDeploymentError(nil)).To(Equal(expected))
}

func Test_EmptyErrorDetails_IsRetryable(t *testing.T) {
	g := NewGomegaWithT(t)

	err := newError()
	expected := reconcilers.DeploymentErrorDetails{
		Classification: reconcilers.DeploymentErrorRetryable,
		Code:           err.Code,
		Message:        err.Message,
	}
	g.Expect(reconcilers.ClassifyDeploymentError(err)).To(Equal(expected))
}

func Test_BadRequest_IsNotRetryable(t *testing.T) {
	g := NewGomegaWithT(t)

	err := newError(badRequestError)
	expected := reconcilers.DeploymentErrorDetails{
		Classification: reconcilers.DeploymentErrorRetryable,
		Code:           badRequestError.Code,
		Message:        badRequestError.Message,
	}
	g.Expect(reconcilers.ClassifyDeploymentError(err)).To(Equal(expected))
}

func Test_Conflict_IsNotRetryable(t *testing.T) {
	g := NewGomegaWithT(t)

	err := newError(conflictError)
	expected := reconcilers.DeploymentErrorDetails{
		Classification: reconcilers.DeploymentErrorFatal,
		Code:           conflictError.Code,
		Message:        conflictError.Message,
	}
	g.Expect(reconcilers.ClassifyDeploymentError(err)).To(Equal(expected))
}

func Test_ResourceGroupNotFound_IsRetryable(t *testing.T) {
	g := NewGomegaWithT(t)

	err := newError(resourceGroupNotFoundError)
	expected := reconcilers.DeploymentErrorDetails{
		Classification: reconcilers.DeploymentErrorRetryable,
		Code:           resourceGroupNotFoundError.Code,
		Message:        resourceGroupNotFoundError.Message,
	}
	g.Expect(reconcilers.ClassifyDeploymentError(err)).To(Equal(expected))
}

func Test_UnknownError_IsRetryable(t *testing.T) {
	g := NewGomegaWithT(t)

	err := newError(unknownError)
	expected := reconcilers.DeploymentErrorDetails{
		Classification: reconcilers.DeploymentErrorRetryable,
		Code:           unknownError.Code,
		Message:        unknownError.Message,
	}
	g.Expect(reconcilers.ClassifyDeploymentError(err)).To(Equal(expected))
}

func Test_MultipleRetryableErrors_IsRetryable(t *testing.T) {
	g := NewGomegaWithT(t)

	err := newError(resourceGroupNotFoundError, resourceNotFound)
	// expected to match the first error, if all errors are retryable
	expected := reconcilers.DeploymentErrorDetails{
		Classification: reconcilers.DeploymentErrorRetryable,
		Code:           resourceGroupNotFoundError.Code,
		Message:        resourceGroupNotFoundError.Message,
	}
	g.Expect(reconcilers.ClassifyDeploymentError(err)).To(Equal(expected))
}

func Test_MultipleRetryableErrorsSingleFatalError_IsFatal(t *testing.T) {
	g := NewGomegaWithT(t)

	err := newError(resourceGroupNotFoundError, resourceNotFound, conflictError)
	// Exported to match the first fatal error, if at least one error is fatal
	expected := reconcilers.DeploymentErrorDetails{
		Classification: reconcilers.DeploymentErrorFatal,
		Code:           conflictError.Code,
		Message:        conflictError.Message,
	}
	g.Expect(reconcilers.ClassifyDeploymentError(err)).To(Equal(expected))
}

func Test_MultipleFatalErrors_IsFatal(t *testing.T) {
	g := NewGomegaWithT(t)

	err := newError(conflictError, badRequestFormatError)
	// Exported to match the first fatal error, if at least one error is fatal
	expected := reconcilers.DeploymentErrorDetails{
		Classification: reconcilers.DeploymentErrorFatal,
		Code:           conflictError.Code,
		Message:        conflictError.Message,
	}
	g.Expect(reconcilers.ClassifyDeploymentError(err)).To(Equal(expected))
}
