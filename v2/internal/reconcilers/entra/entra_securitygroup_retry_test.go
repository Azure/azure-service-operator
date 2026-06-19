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

	abstractions "github.com/microsoft/kiota-abstractions-go"
	"github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
	. "github.com/onsi/gomega"
	"github.com/rotisserie/eris"
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
)

func TestClassifyRelationshipError_PermissionDenied_ReturnsSlowReadyConditionError(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	err := makeODataError(http.StatusForbidden, nil)

	result, classifiedErr := classifyRelationshipError(err)

	g.Expect(result).To(Equal(ctrl.Result{}))
	g.Expect(classifiedErr).To(HaveOccurred())
	g.Expect(classifiedErr.Error()).To(ContainSubstring("permission denied reconciling SecurityGroup owners/members"))

	readyErr, ok := conditions.AsReadyConditionImpactingError(classifiedErr)
	g.Expect(ok).To(BeTrue())
	g.Expect(readyErr.Reason).To(Equal(reasonRelationshipPermissionDenied.Name))
	g.Expect(readyErr.RetryClassification).To(Equal(reasonRelationshipPermissionDenied.RetryClassification))
}

func TestClassifyRelationshipError_Throttle_PropagatesRetryAfterInResult(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	err := makeODataError(http.StatusTooManyRequests, map[string]string{"Retry-After": "42"})

	result, classifiedErr := classifyRelationshipError(err)

	// Throttle Retry-After is surfaced via result.RequeueAfter so the interval.Calculator
	// can take the max of it and the classification-based backoff.
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

	result, classifiedErr := classifyRelationshipError(err)

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

	result, classifiedErr := classifyRelationshipError(errors.New("boom"))

	g.Expect(result).To(Equal(ctrl.Result{}))
	g.Expect(classifiedErr).To(HaveOccurred())
	g.Expect(classifiedErr.Error()).To(ContainSubstring("error reconciling SecurityGroup owners/members"))

	readyErr, ok := conditions.AsReadyConditionImpactingError(classifiedErr)
	g.Expect(ok).To(BeTrue())
	g.Expect(readyErr.Reason).To(Equal(reasonRelationshipFailed.Name))
	g.Expect(readyErr.RetryClassification).To(Equal(reasonRelationshipFailed.RetryClassification))
}

func TestTryThrottleRequeue_UsesRetryAfterHTTPDate(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	retryAt := time.Now().Add(55 * time.Second).UTC().Format(http.TimeFormat)
	err := makeODataError(http.StatusTooManyRequests, map[string]string{"Retry-After": retryAt})

	result, ok := tryThrottleRequeue(err)

	g.Expect(ok).To(BeTrue())
	g.Expect(result.RequeueAfter).To(BeNumerically(">=", 50*time.Second))
	g.Expect(result.RequeueAfter).To(BeNumerically("<=", 56*time.Second))
}

func TestTryThrottleRequeue_NonThrottleReturnsFalse(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	result, ok := tryThrottleRequeue(makeODataError(http.StatusForbidden, nil))

	g.Expect(ok).To(BeFalse())
	g.Expect(result).To(Equal(ctrl.Result{}))
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
