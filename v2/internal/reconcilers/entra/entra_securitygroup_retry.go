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

	"github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
	"github.com/rotisserie/eris"

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
func classifyRelationshipError(err error) error {
	var result *conditions.ReadyConditionImpactingError
	if isPermissionError(err) {
		result = conditions.NewReadyConditionImpactingError(
			eris.Wrap(err, "permission denied reconciling SecurityGroup owners/members"),
			conditions.ConditionSeverityWarning,
			reasonRelationshipPermissionDenied,
		)
	} else {
		result = conditions.NewReadyConditionImpactingError(
			eris.Wrap(err, "error reconciling SecurityGroup owners/members"),
			conditions.ConditionSeverityWarning,
			reasonRelationshipFailed,
		)
	}

	if retryAfter, ok := retryAfterFromError(err); ok {
		result = result.WithRetryAfter(retryAfter)
	}

	return result
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

// retryAfterFromError extracts the Retry-After header from an ODataError, if present, and returns it as a time.Duration.
func retryAfterFromError(
	err error,
) (time.Duration, bool) {
	odataError, ok := errors.AsType[*odataerrors.ODataError](err)
	if !ok {
		return 0, false
	}

	if odataError == nil || odataError.ResponseHeaders == nil {
		return 0, false
	}

	values := odataError.ResponseHeaders.Get("Retry-After")
	if len(values) == 0 {
		return 0, false
	}

	retryAfterStr := values[0]
	if retryAfterVal, parseErr := strconv.ParseInt(retryAfterStr, 10, 64); parseErr == nil {
		// Clamp to a reasonable range just in case we get a crazy value from the service.
		retryAfterVal = max(0, min(retryAfterVal, 3600)) // 1 hour
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
