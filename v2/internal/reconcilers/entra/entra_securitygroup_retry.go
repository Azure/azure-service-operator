/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package entra

import (
	"net/http"
	"strconv"
	"time"

	"github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
	"github.com/rotisserie/eris"
	ctrl "sigs.k8s.io/controller-runtime"
)

const (
	relationshipRetryFast = 20 * time.Second
	relationshipRetrySlow = 5 * time.Minute
)

func classifyRelationshipError(err error) (ctrl.Result, error) {
	if result, ok := tryThrottleRequeue(err); ok {
		return result, nil
	}

	if isPermissionError(err) {
		return ctrl.Result{RequeueAfter: relationshipRetrySlow},
			eris.Wrap(err, "permission denied reconciling SecurityGroup owners/members")
	}

	return ctrl.Result{RequeueAfter: relationshipRetryFast},
		eris.Wrap(err, "error reconciling SecurityGroup owners/members")
}

func tryThrottleRequeue(err error) (ctrl.Result, bool) {
	odataError, ok := asODataError(err)
	if !ok || odataError.ResponseStatusCode != http.StatusTooManyRequests {
		return ctrl.Result{}, false
	}

	retryAfter := retryAfterFromODataError(odataError)
	if retryAfter <= 0 {
		return ctrl.Result{}, false
	}

	return ctrl.Result{RequeueAfter: retryAfter}, true
}

func isPermissionError(err error) bool {
	odataError, ok := asODataError(err)
	if !ok {
		return false
	}

	return odataError.ResponseStatusCode == http.StatusForbidden
}

func asODataError(err error) (*odataerrors.ODataError, bool) {
	var odataError *odataerrors.ODataError
	if eris.As(err, &odataError) {
		return odataError, true
	}

	return nil, false
}

func retryAfterFromODataError(odataError *odataerrors.ODataError) time.Duration {
	if odataError == nil || odataError.ResponseHeaders == nil {
		return 0
	}

	values := odataError.ResponseHeaders.Get("Retry-After")
	if len(values) == 0 {
		return 0
	}

	retryAfterStr := values[0]
	if retryAfterVal, parseErr := strconv.ParseInt(retryAfterStr, 10, 64); parseErr == nil {
		return time.Duration(retryAfterVal) * time.Second
	}

	if retryAfterTime, parseErr := parseHTTPDate(retryAfterStr); parseErr == nil {
		result := time.Until(retryAfterTime)
		if result > 0 {
			return result
		}
	}

	return 0
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
