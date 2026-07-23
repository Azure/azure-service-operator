/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package entra

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
	"github.com/rotisserie/eris"
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/retry"
)

// Domain-specific Reasons for Entra SecurityGroup owner/member reconciliation.
// We pick retry classifications so that the standard interval.Calculator gives us the
// existing semantics: permission errors back off slowly, generic Graph errors back off
// quickly. The calculator combines these classifications with any caller-supplied
// RequeueAfter (e.g. parsed from a 429 Retry-After header) by taking the larger value.
var (
	reasonRelationshipPermissionDenied = conditions.Reason{
		Name:                "GraphPermissionDenied",
		RetryClassification: retry.Slow,
	}
	reasonRelationshipFailed = conditions.Reason{
		Name:                "GraphRelationshipReconcileFailed",
		RetryClassification: retry.Fast,
	}
)

// classifyRelationshipError wraps an error from owner/member reconciliation as a
// ReadyConditionImpactingError carrying the appropriate retry classification, and
// forwards the caller-supplied throttle result unchanged. Callers are expected to
// have already extracted any HTTP 429 Retry-After per-side (see
// reconcileOwnersAndMembers) so this function does not walk the error tree. The
// interval.Calculator combines these signals: throttling can only slow us down,
// never speed us up beyond the classification-based exponential backoff.
func classifyRelationshipError(result ctrl.Result, err error) (ctrl.Result, error) {
	if isPermissionError(err) {
		return result, conditions.NewReadyConditionImpactingError(
			eris.Wrap(err, "permission denied reconciling SecurityGroup owners/members"),
			conditions.ConditionSeverityWarning,
			reasonRelationshipPermissionDenied,
		)
	}

	return result, conditions.NewReadyConditionImpactingError(
		eris.Wrap(err, "error reconciling SecurityGroup owners/members"),
		conditions.ConditionSeverityWarning,
		reasonRelationshipFailed,
	)
}

// retryAfterResult extracts an HTTP 429 Retry-After header from an OData error in
// err (unwrapping through single-error wrappers via errors.AsType) and returns it
// as a ctrl.Result.RequeueAfter. A zero result is returned when no throttle signal
// is present. This helper is intended to be called per single-wrapped error at the
// point the error is produced; call sites that need to combine throttles from
// multiple independent sources should take the max of the resulting durations.
func retryAfterResult(err error) ctrl.Result {
	if err == nil {
		return ctrl.Result{}
	}

	odataError, ok := errors.AsType[*odataerrors.ODataError](err)
	if !ok {
		return ctrl.Result{}
	}

	retryAfter, ok := retryAfterFromODataError(odataError)
	if !ok {
		return ctrl.Result{}
	}

	return ctrl.Result{RequeueAfter: retryAfter}
}

// maxThrottleResult returns whichever of a and b has the larger RequeueAfter. Used
// to collapse per-side throttle signals into a single result while reconciling
// owners and members in parallel.
func maxThrottleResult(a, b ctrl.Result) ctrl.Result {
	if b.RequeueAfter > a.RequeueAfter {
		return b
	}
	return a
}

func isPermissionError(err error) bool {
	// errors.AsType walks both `Unwrap() error` and `Unwrap() []error`, so it
	// finds an ODataError nested inside any wrapping (including errors.Join).
	odataError, ok := errors.AsType[*odataerrors.ODataError](err)
	if !ok {
		return false
	}

	return odataError.ResponseStatusCode == http.StatusForbidden
}

func retryAfterFromODataError(
	odataError *odataerrors.ODataError,
) (time.Duration, bool) {
	if odataError == nil || odataError.ResponseHeaders == nil {
		return 0, false
	}

	values := odataError.ResponseHeaders.Get("Retry-After")
	if len(values) == 0 {
		return 0, false
	}

	retryAfterStr := values[0]
	if retryAfterVal, parseErr := strconv.ParseInt(retryAfterStr, 10, 64); parseErr == nil {
		return time.Duration(retryAfterVal) * time.Second, true
	}

	if retryAfterTime, parseErr := http.ParseTime(retryAfterStr); parseErr == nil {
		result := time.Until(retryAfterTime)
		if result > 0 {
			return result, true
		}
	}

	return 0, false
}
