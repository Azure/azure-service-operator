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
// surfaces any HTTP 429 Retry-After delay via ctrl.Result.RequeueAfter. The
// interval.Calculator combines these signals: throttling can only slow us down,
// never speed us up beyond the classification-based exponential backoff.
func classifyRelationshipError(err error) (ctrl.Result, error) {
	result, _ := tryThrottleRequeue(err)

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

// tryThrottleRequeue extracts an HTTP 429 Retry-After header from an OData error and
// returns it as a ctrl.Result.RequeueAfter. The bool indicates whether a usable
// Retry-After value was found. A zero result is returned when no throttle signal is
// present — callers can pass this straight to the calculator.
func tryThrottleRequeue(err error) (ctrl.Result, bool) {
	retryAfter, ok := maxRetryAfterFromError(err)
	if !ok {
		return ctrl.Result{}, false
	}

	return ctrl.Result{
		RequeueAfter: retryAfter,
	}, true
}

func maxRetryAfterFromError(err error) (time.Duration, bool) {
	if err == nil {
		// No error, no Retry-After
		return time.Duration(0), false
	}

	switch x := err.(type) {
	case *odataerrors.ODataError:
		// For OData Errors, look at the Retry-After header and return it if present
		retryAfter, ok := retryAfterFromODataError(x)
		if ok {
			return retryAfter, true
		}

		return time.Duration(0), false

	case interface{ Unwrap() []error }:
		// If we wrap a sequence of errors, return the largest Retry-After from any of the children.
		var max time.Duration
		for _, child := range x.Unwrap() {
			if child == nil {
				continue
			}

			if childMax, ok := maxRetryAfterFromError(child); ok && childMax > max {
				max = childMax
			}
		}

		if max > 0 {
			return max, true
		}

		return 0, false

	case interface{ Unwrap() error }:
		// If we wrap a single error, return the Retry-After from that child.
		child := x.Unwrap()
		if child == nil {
			return 0, false
		}

		return maxRetryAfterFromError(child)

	default:
		return time.Duration(0), false
	}
}

func isPermissionError(err error) bool {
	odataError, ok := asODataError(err)
	if !ok {
		return false
	}

	return odataError.ResponseStatusCode == http.StatusForbidden
}

func asODataError(err error) (*odataerrors.ODataError, bool) {
	// AsType() walks both `Unwrap() error`` and `Unwrap() []error` we we'll find anything nested
	if odataError, ok := errors.AsType[*odataerrors.ODataError](err); ok {
		return odataError, true
	}

	return nil, false
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

	if retryAfterTime, parseErr := parseHTTPDate(retryAfterStr); parseErr == nil {
		result := time.Until(retryAfterTime)
		if result > 0 {
			return result, true
		}
	}

	return 0, false
}

func parseHTTPDate(s string) (time.Time, error) {
	if t, err := time.Parse("Mon, 02 Jan 2006 15:04:05 MST", s); err == nil {
		return t, nil
	} else if t, err = time.Parse("Monday, 02-Jan-06 15:04:05 MST", s); err == nil {
		return t, nil
	} else if t, err = time.Parse("Mon Jan  2 15:04:05 2006", s); err == nil {
		return t, nil
	}

	return time.Time{}, eris.New("unable to parse date")
}
