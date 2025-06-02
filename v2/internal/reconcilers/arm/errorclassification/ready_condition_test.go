/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package errorclassification_test

import (
	"errors"
	"fmt"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/rotisserie/eris"

	"github.com/Azure/azure-service-operator/v2/internal/genericarmclient"
	"github.com/Azure/azure-service-operator/v2/internal/reconcilers/arm/errorclassification"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/core"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/retry"
)

// Mock classifier function for the unknown classification test
func mockUnknownClassifier(err *genericarmclient.CloudError) (core.CloudErrorDetails, error) {
	return core.CloudErrorDetails{
		Classification: "Unknown", // Not a valid classification
		Code:           err.Code(),
		Message:        err.Message(),
	}, nil
}

func Test_MakeReadyConditionImpactingErrorFromError_AlreadyReadyConditionImpactingError(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	// Create a ReadyConditionImpactingError
	originalErr := errors.New("original error")
	reason := conditions.MakeReason("SomeCode", retry.Fast)
	readyErr := conditions.NewReadyConditionImpactingError(originalErr, conditions.ConditionSeverityWarning, reason)
	result := errorclassification.MakeReadyConditionImpactingErrorFromError(readyErr, errorclassification.ClassifyCloudError)

	// Check that the original error is returned
	g.Expect(result).To(MatchError(readyErr))
}

func Test_MakeReadyConditionImpactingErrorFromError_NotCloudError(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	// Create a non-CloudError
	regularErr := errors.New("regular error")
	result := errorclassification.MakeReadyConditionImpactingErrorFromError(regularErr, errorclassification.ClassifyCloudError)

	// Check that a new ReadyConditionImpactingError is created with appropriate defaults
	readyErr, ok := conditions.AsReadyConditionImpactingError(result)
	g.Expect(ok).To(BeTrue())
	g.Expect(readyErr.Severity).To(Equal(conditions.ConditionSeverityWarning))
	g.Expect(readyErr.Reason).To(ContainSubstring(core.UnknownErrorCode))
}

func Test_MakeReadyConditionImpactingErrorFromError_RetryableCloudError(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	// Create a CloudError using "Conflict" code which is classified as retryable
	innerErr := fmt.Errorf("This is an inner error")
	retryableErr := genericarmclient.NewTestCloudError("Conflict", "This is a retryable error", genericarmclient.WithTestInnerError(innerErr))

	result := errorclassification.MakeReadyConditionImpactingErrorFromError(retryableErr, errorclassification.ClassifyCloudError)

	// Check that a new ReadyConditionImpactingError is created with appropriate severity
	readyErr, ok := conditions.AsReadyConditionImpactingError(result)
	g.Expect(ok).To(BeTrue())
	g.Expect(readyErr.Severity).To(Equal(conditions.ConditionSeverityWarning))
	g.Expect(readyErr.Reason).To(ContainSubstring("Conflict"))

	// Verify the cause contains our original error
	g.Expect(eris.Unwrap(readyErr.Cause())).To(Equal(retryableErr))
	// Verify that the inner error is not included in the message
	g.Expect(readyErr).ToNot(MatchError(ContainSubstring("This is an inner error")))
}

func Test_MakeReadyConditionImpactingErrorFromError_FatalCloudError(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	// Create a fatal CloudError using "BadRequest" code which is classified as fatal
	innerErr := fmt.Errorf("This is an inner error")
	fatalErr := genericarmclient.NewTestCloudError("BadRequest", "This is a fatal error", genericarmclient.WithTestInnerError(innerErr))

	// Pass it to MakeReadyConditionImpactingErrorFromError with the real classifier
	result := errorclassification.MakeReadyConditionImpactingErrorFromError(fatalErr, errorclassification.ClassifyCloudError)

	// Check that a new ReadyConditionImpactingError is created with appropriate severity
	readyErr, ok := conditions.AsReadyConditionImpactingError(result)
	g.Expect(ok).To(BeTrue())
	g.Expect(readyErr.Severity).To(Equal(conditions.ConditionSeverityError))
	g.Expect(readyErr.Reason).To(ContainSubstring("BadRequest"))

	// Verify the cause contains our original error
	g.Expect(eris.Unwrap(readyErr.Cause())).To(Equal(fatalErr))
	// Verify that the inner error is not included in the message
	g.Expect(readyErr).ToNot(MatchError(ContainSubstring("This is an inner error")))
}

func Test_MakeReadyConditionImpactingErrorFromError_UnknownClassification(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	// Create a CloudError
	unknownErr := genericarmclient.NewTestCloudError("SomeCode", "This has an unknown classification")

	// Pass it to MakeReadyConditionImpactingErrorFromError with a mock classifier that returns an invalid classification
	result := errorclassification.MakeReadyConditionImpactingErrorFromError(unknownErr, mockUnknownClassifier)

	// Check that this returns an error about the unknown classification
	g.Expect(result).To(MatchError("unknown error classification \"Unknown\" while making Ready condition"))
}
