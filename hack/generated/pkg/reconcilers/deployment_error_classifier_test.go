/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package reconcilers_test

import (
	"testing"

	"github.com/Azure/azure-service-operator/hack/generated/pkg/armclient"
	"github.com/Azure/azure-service-operator/hack/generated/pkg/reconcilers"
	. "github.com/onsi/gomega"
)

var badRequestError = armclient.DeploymentError{
	Code:    "BadRequest",
	Message: "That was not a good request",
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
	g.Expect(reconcilers.ClassifyDeploymentError(nil)).To(Equal(reconcilers.DeploymentErrorRetryable))
}

func Test_EmptyErrorDetails_IsRetryable(t *testing.T) {
	g := NewGomegaWithT(t)

	err := newError()
	g.Expect(reconcilers.ClassifyDeploymentError(err)).To(Equal(reconcilers.DeploymentErrorRetryable))
}

func Test_BadRequest_IsNotRetryable(t *testing.T) {
	g := NewGomegaWithT(t)

	err := newError(badRequestError)
	g.Expect(reconcilers.ClassifyDeploymentError(err)).To(Equal(reconcilers.DeploymentErrorFatal))
}

func Test_ResourceGroupNotFound_IsRetryable(t *testing.T) {
	g := NewGomegaWithT(t)

	err := newError(resourceGroupNotFoundError)
	g.Expect(reconcilers.ClassifyDeploymentError(err)).To(Equal(reconcilers.DeploymentErrorRetryable))
}

func Test_UnknownError_IsRetryable(t *testing.T) {
	g := NewGomegaWithT(t)

	err := newError(unknownError)
	g.Expect(reconcilers.ClassifyDeploymentError(err)).To(Equal(reconcilers.DeploymentErrorRetryable))
}

func Test_MultipleRetryableErrors_IsRetryable(t *testing.T) {
	g := NewGomegaWithT(t)

	err := newError(resourceGroupNotFoundError, resourceNotFound)
	g.Expect(reconcilers.ClassifyDeploymentError(err)).To(Equal(reconcilers.DeploymentErrorRetryable))
}

func Test_MultipleRetryableErrorsSingleFatalError_IsFatal(t *testing.T) {
	g := NewGomegaWithT(t)

	err := newError(resourceGroupNotFoundError, resourceNotFound, badRequestError)
	g.Expect(reconcilers.ClassifyDeploymentError(err)).To(Equal(reconcilers.DeploymentErrorFatal))
}
