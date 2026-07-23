/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package entra

import (
	"errors"
	"net/http"
	"testing"
	"time"

	. "github.com/onsi/gomega"

	abstractions "github.com/microsoft/kiota-abstractions-go"
	"github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
	"github.com/rotisserie/eris"
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
)

func TestClassifyRelationshipError_PermissionDenied_ReturnsSlowReadyConditionError(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	err := makeODataError(http.StatusForbidden, nil)

	result, classifiedErr := classifyRelationshipError(ctrl.Result{}, err)

	g.Expect(result).To(Equal(ctrl.Result{}))
	g.Expect(classifiedErr).To(HaveOccurred())
	g.Expect(classifiedErr.Error()).To(ContainSubstring("permission denied reconciling SecurityGroup owners/members"))

	readyErr, ok := conditions.AsReadyConditionImpactingError(classifiedErr)
	g.Expect(ok).To(BeTrue())
	g.Expect(readyErr.Reason).To(Equal(reasonRelationshipPermissionDenied.Name))
	g.Expect(readyErr.RetryClassification).To(Equal(reasonRelationshipPermissionDenied.RetryClassification))
}

func TestClassifyRelationshipError_Throttle_PropagatesCallerSuppliedResult(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	err := makeODataError(http.StatusTooManyRequests, map[string]string{"Retry-After": "42"})

	// Caller (reconcileOwnersAndMembers) is expected to extract the throttle result
	// per side and pass it in; classifyRelationshipError forwards it unchanged so
	// the interval.Calculator can take the max of it and the classification-based
	// backoff.
	throttle := retryAfterResult(err)
	result, classifiedErr := classifyRelationshipError(throttle, err)

	g.Expect(result).To(Equal(ctrl.Result{RequeueAfter: 42 * time.Second}))
	g.Expect(classifiedErr).To(HaveOccurred())

	readyErr, ok := conditions.AsReadyConditionImpactingError(classifiedErr)
	g.Expect(ok).To(BeTrue())
	g.Expect(readyErr.Reason).To(Equal(reasonRelationshipFailed.Name))
	g.Expect(readyErr.RetryClassification).To(Equal(reasonRelationshipFailed.RetryClassification))
}

func TestClassifyRelationshipError_ThrottleWithoutRetryAfter_ReturnsFastReadyConditionError(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	err := makeODataError(http.StatusTooManyRequests, nil)

	result, classifiedErr := classifyRelationshipError(retryAfterResult(err), err)

	g.Expect(result).To(Equal(ctrl.Result{}))
	g.Expect(classifiedErr).To(HaveOccurred())
	g.Expect(classifiedErr.Error()).To(ContainSubstring("error reconciling SecurityGroup owners/members"))

	readyErr, ok := conditions.AsReadyConditionImpactingError(classifiedErr)
	g.Expect(ok).To(BeTrue())
	g.Expect(readyErr.Reason).To(Equal(reasonRelationshipFailed.Name))
	g.Expect(readyErr.RetryClassification).To(Equal(reasonRelationshipFailed.RetryClassification))
}

func TestClassifyRelationshipError_GenericError_ReturnsFastReadyConditionError(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	err := errors.New("boom")
	result, classifiedErr := classifyRelationshipError(retryAfterResult(err), err)

	g.Expect(result).To(Equal(ctrl.Result{}))
	g.Expect(classifiedErr).To(HaveOccurred())
	g.Expect(classifiedErr.Error()).To(ContainSubstring("error reconciling SecurityGroup owners/members"))

	readyErr, ok := conditions.AsReadyConditionImpactingError(classifiedErr)
	g.Expect(ok).To(BeTrue())
	g.Expect(readyErr.Reason).To(Equal(reasonRelationshipFailed.Name))
	g.Expect(readyErr.RetryClassification).To(Equal(reasonRelationshipFailed.RetryClassification))
}

func TestRetryAfterResult_UsesRetryAfterHTTPDate(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	retryAt := time.Now().Add(55 * time.Second).UTC().Format(http.TimeFormat)
	err := makeODataError(http.StatusTooManyRequests, map[string]string{"Retry-After": retryAt})

	result := retryAfterResult(err)

	g.Expect(result.RequeueAfter).To(BeNumerically(">=", 50*time.Second))
	g.Expect(result.RequeueAfter).To(BeNumerically("<=", 56*time.Second))
}

func TestRetryAfterResult_NonThrottleReturnsZero(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	result := retryAfterResult(makeODataError(http.StatusForbidden, nil))

	g.Expect(result).To(Equal(ctrl.Result{}))
}

func TestRetryAfterResult_NilReturnsZero(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	g.Expect(retryAfterResult(nil)).To(Equal(ctrl.Result{}))
}

func TestMaxThrottleResult_PicksLarger(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	a := ctrl.Result{RequeueAfter: 5 * time.Second}
	b := ctrl.Result{RequeueAfter: 30 * time.Second}

	g.Expect(maxThrottleResult(a, b)).To(Equal(b))
	g.Expect(maxThrottleResult(b, a)).To(Equal(b))
	g.Expect(maxThrottleResult(ctrl.Result{}, ctrl.Result{})).To(Equal(ctrl.Result{}))
}

func makeODataError(statusCode int, headers map[string]string) error {
	oDataErr := odataerrors.NewODataError()
	oDataErr.SetStatusCode(statusCode)

	responseHeaders := abstractions.NewResponseHeaders()
	for key, value := range headers {
		responseHeaders.Add(key, value)
	}
	oDataErr.SetResponseHeaders(responseHeaders)

	return eris.Wrap(oDataErr, "wrapped")
}
