/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package entra

import (
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

	classifiedErr := classifyRelationshipError(err)

	g.Expect(classifiedErr).To(HaveOccurred())
	g.Expect(classifiedErr.Error()).To(ContainSubstring("permission denied reconciling SecurityGroup owners/members"))

	readyErr, ok := conditions.AsReadyConditionImpactingError(classifiedErr)
	g.Expect(ok).To(BeTrue())
	g.Expect(readyErr.Reason).To(Equal(reasonRelationshipPermissionDenied.Name))
	g.Expect(readyErr.RetryClassification).To(Equal(reasonRelationshipPermissionDenied.RetryClassification))
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
