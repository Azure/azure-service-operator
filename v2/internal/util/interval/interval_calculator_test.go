/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package interval

import (
	"math/rand"
	"testing"
	"time"

	. "github.com/onsi/gomega"
	"github.com/pkg/errors"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/Azure/azure-service-operator/v2/internal/util/lockedrand"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
)

func newCalculator(params CalculatorParameters) Calculator {
	seed := time.Now().UnixNano()
	//nolint:gosec // do not want cryptographic randomness here
	rng := rand.New(lockedrand.NewSource(seed))
	params.Rand = rng

	calc := NewCalculator(params)

	return calc
}

func Test_Success_WithoutSyncPeriod_ReturnsEmptyResult(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	calc := newCalculator(
		CalculatorParameters{
			ErrorBaseDelay:    1 * time.Second,
			ErrorMaxFastDelay: 5 * time.Second,
			ErrorMaxSlowDelay: 10 * time.Second,
		})

	req := ctrl.Request{NamespacedName: types.NamespacedName{Namespace: "foo", Name: "bar"}}

	result, err := calc.NextInterval(req, ctrl.Result{}, nil)
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(result).To(Equal(ctrl.Result{}))

	g.Expect(calc.(*calculator).failures).To(HaveLen(0))
}

func Test_Success_WithSyncPeriod_ReturnsSyncPeriod(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	syncPeriod := 12 * time.Second
	calc := newCalculator(
		CalculatorParameters{
			ErrorBaseDelay:    1 * time.Second,
			ErrorMaxFastDelay: 5 * time.Second,
			ErrorMaxSlowDelay: 10 * time.Second,
			SyncPeriod:        &syncPeriod,
		})

	req := ctrl.Request{NamespacedName: types.NamespacedName{Namespace: "foo", Name: "bar"}}

	result, err := calc.NextInterval(req, ctrl.Result{}, nil)
	g.Expect(err).ToNot(HaveOccurred())

	// Threshold here needs to be at or below the lower-bound of the RequeuePeriod, taking Jitter into account
	// Currently jitter is 0.25, so the 12s above becomes 9s to 15s
	g.Expect(result.RequeueAfter >= 9*time.Second).To(BeTrue())

	g.Expect(calc.(*calculator).failures).To(HaveLen(0))
}

func Test_Requeue_ReturnsResultUnmodified(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	syncPeriod := 10 * time.Second
	calc := newCalculator(
		CalculatorParameters{
			ErrorBaseDelay:    1 * time.Second,
			ErrorMaxFastDelay: 5 * time.Second,
			ErrorMaxSlowDelay: 10 * time.Second,
			SyncPeriod:        &syncPeriod,
		})

	req := ctrl.Request{NamespacedName: types.NamespacedName{Namespace: "foo", Name: "bar"}}

	result, err := calc.NextInterval(req, ctrl.Result{Requeue: true}, nil)
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(result).To(Equal(ctrl.Result{Requeue: true}))

	g.Expect(calc.(*calculator).failures).To(HaveLen(0))
}

func Test_RequeueWithDelayOverride_ReturnsResultWithDelay(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	syncPeriod := 10 * time.Second
	calc := newCalculator(
		CalculatorParameters{
			ErrorBaseDelay:       1 * time.Second,
			ErrorMaxFastDelay:    5 * time.Second,
			ErrorMaxSlowDelay:    10 * time.Second,
			SyncPeriod:           &syncPeriod,
			RequeueDelayOverride: 77 * time.Second,
		})

	req := ctrl.Request{NamespacedName: types.NamespacedName{Namespace: "foo", Name: "bar"}}

	result, err := calc.NextInterval(req, ctrl.Result{Requeue: true}, nil)
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(result).To(Equal(ctrl.Result{Requeue: true, RequeueAfter: 77 * time.Second}))

	g.Expect(calc.(*calculator).failures).To(HaveLen(0))
}

func Test_Error_ReturnedAsIs(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	calc := newCalculator(
		CalculatorParameters{
			ErrorBaseDelay:    1 * time.Second,
			ErrorMaxFastDelay: 5 * time.Second,
			ErrorMaxSlowDelay: 10 * time.Second,
		})

	req := ctrl.Request{NamespacedName: types.NamespacedName{Namespace: "foo", Name: "bar"}}

	inputErr := errors.New("An error")
	result, err := calc.NextInterval(req, ctrl.Result{}, inputErr)
	g.Expect(err).To(Equal(inputErr))
	g.Expect(result).To(Equal(ctrl.Result{}))

	g.Expect(calc.(*calculator).failures).To(HaveLen(1))
}

func Test_ErrorFollowedBySuccess_ClearsFailureTracking(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	calc := newCalculator(
		CalculatorParameters{
			ErrorBaseDelay:    1 * time.Second,
			ErrorMaxFastDelay: 5 * time.Second,
			ErrorMaxSlowDelay: 10 * time.Second,
		})

	req := ctrl.Request{NamespacedName: types.NamespacedName{Namespace: "foo", Name: "bar"}}

	inputErr := errors.New("An error")
	result, err := calc.NextInterval(req, ctrl.Result{}, inputErr)
	g.Expect(err).To(Equal(inputErr))
	g.Expect(result).To(Equal(ctrl.Result{}))

	result, err = calc.NextInterval(req, ctrl.Result{}, nil)
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(result).To(Equal(ctrl.Result{}))

	g.Expect(calc.(*calculator).failures).To(HaveLen(0))
}

func Test_ReadyConditionErrorWithErrorSeverity_ReturnsSuccess(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	calc := newCalculator(
		CalculatorParameters{
			ErrorBaseDelay:    1 * time.Second,
			ErrorMaxFastDelay: 5 * time.Second,
			ErrorMaxSlowDelay: 10 * time.Second,
		})

	req := ctrl.Request{NamespacedName: types.NamespacedName{Namespace: "foo", Name: "bar"}}

	inputErr := conditions.NewReadyConditionImpactingError(errors.New("problem"), conditions.ConditionSeverityError, conditions.ReasonFailed)
	result, err := calc.NextInterval(req, ctrl.Result{}, inputErr)
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(result).To(Equal(ctrl.Result{}))

	g.Expect(calc.(*calculator).failures).To(HaveLen(0))
}

func Test_ReadyConditionErrorWithSlowBackoff_UsesSlowBackoff(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	calc := newCalculator(
		CalculatorParameters{
			ErrorBaseDelay:    1 * time.Second,
			ErrorMaxFastDelay: 5 * time.Second,
			ErrorMaxSlowDelay: 10 * time.Second,
		})

	req := ctrl.Request{NamespacedName: types.NamespacedName{Namespace: "foo", Name: "bar"}}

	inputErr := conditions.NewReadyConditionImpactingError(
		errors.New("problem"),
		conditions.ConditionSeverityWarning,
		conditions.Reason{Name: "Abc", RetryClassification: conditions.RetrySlow})

	expectedDelaySec := []int64{1, 2, 4, 8, 10, 10, 10}

	for _, delay := range expectedDelaySec {
		result, err := calc.NextInterval(req, ctrl.Result{}, inputErr)
		g.Expect(err).ToNot(HaveOccurred()) // No error because we're manually controlling backoff here
		g.Expect(result).To(Equal(ctrl.Result{RequeueAfter: time.Duration(delay) * time.Second}))
	}

	g.Expect(calc.(*calculator).failures).To(HaveLen(1))

	// Success should then clear failure tracking
	result, err := calc.NextInterval(req, ctrl.Result{}, nil)
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(result).To(Equal(ctrl.Result{}))

	g.Expect(calc.(*calculator).failures).To(HaveLen(0))
}

func Test_ReadyConditionErrorWithFastBackoff_UsesFastBackoff(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	calc := newCalculator(
		CalculatorParameters{
			ErrorBaseDelay:    1 * time.Second,
			ErrorMaxFastDelay: 5 * time.Second,
			ErrorMaxSlowDelay: 10 * time.Second,
		})

	req := ctrl.Request{NamespacedName: types.NamespacedName{Namespace: "foo", Name: "bar"}}

	inputErr := conditions.NewReadyConditionImpactingError(
		errors.New("problem"),
		conditions.ConditionSeverityWarning,
		conditions.Reason{Name: "Abc", RetryClassification: conditions.RetryFast})

	expectedDelaySec := []int64{1, 2, 4, 5, 5, 5, 5}

	for _, delay := range expectedDelaySec {
		result, err := calc.NextInterval(req, ctrl.Result{}, inputErr)
		g.Expect(err).ToNot(HaveOccurred()) // No error because we're manually controlling backoff here
		g.Expect(result).To(Equal(ctrl.Result{RequeueAfter: time.Duration(delay) * time.Second}))
	}

	g.Expect(calc.(*calculator).failures).To(HaveLen(1))

	// Success should then clear failure tracking
	result, err := calc.NextInterval(req, ctrl.Result{}, nil)
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(result).To(Equal(ctrl.Result{}))

	g.Expect(calc.(*calculator).failures).To(HaveLen(0))
}

func Test_KubeClientNotFoundError_ReturnsSuccess(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	calc := newCalculator(
		CalculatorParameters{
			ErrorBaseDelay:    1 * time.Second,
			ErrorMaxFastDelay: 5 * time.Second,
			ErrorMaxSlowDelay: 10 * time.Second,
		})

	req := ctrl.Request{NamespacedName: types.NamespacedName{Namespace: "foo", Name: "bar"}}

	inputErr := &apierrors.StatusError{
		ErrStatus: metav1.Status{
			Reason: metav1.StatusReasonNotFound,
		},
	}

	result, err := calc.NextInterval(req, ctrl.Result{}, inputErr)
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(result).To(Equal(ctrl.Result{}))

	g.Expect(calc.(*calculator).failures).To(HaveLen(0))
}

func Test_KubeClientConflict_ReturnsBackoff(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	calc := newCalculator(
		CalculatorParameters{
			ErrorBaseDelay:    1 * time.Second,
			ErrorMaxFastDelay: 5 * time.Second,
			ErrorMaxSlowDelay: 10 * time.Second,
		})

	req := ctrl.Request{NamespacedName: types.NamespacedName{Namespace: "foo", Name: "bar"}}

	inputErr := &apierrors.StatusError{
		ErrStatus: metav1.Status{
			Reason: metav1.StatusReasonConflict,
		},
	}

	result, err := calc.NextInterval(req, ctrl.Result{}, inputErr)
	g.Expect(err).To(HaveOccurred())
	g.Expect(result).To(Equal(ctrl.Result{}))

	g.Expect(calc.(*calculator).failures).To(HaveLen(1))
}
