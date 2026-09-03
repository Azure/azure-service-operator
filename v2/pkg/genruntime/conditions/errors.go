/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package conditions

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
	"github.com/rotisserie/eris"

	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/retry"
)

// ReadyConditionImpactingError is an error that requires notification in the Ready condition
type ReadyConditionImpactingError struct {
	Severity            ConditionSeverity
	Reason              string
	cause               error
	RetryClassification retry.Classification
	RetryAfter          time.Duration
}

// NewReadyConditionImpactingError creates a new ReadyConditionImpactingError
func NewReadyConditionImpactingError(cause error, severity ConditionSeverity, reason Reason) *ReadyConditionImpactingError {
	result := &ReadyConditionImpactingError{
		cause:               cause,
		Severity:            severity,
		Reason:              reason.Name,
		RetryClassification: reason.RetryClassification,
	}

	if retryAfter, ok := retryAfterFromError(cause); ok {
		result.RetryAfter = retryAfter
	}

	return result
}

var _ error = &ReadyConditionImpactingError{}

func AsReadyConditionImpactingError(err error) (*ReadyConditionImpactingError, bool) {
	var typedErr *ReadyConditionImpactingError
	if eris.As(err, &typedErr) {
		return typedErr, true
	}

	return nil, false
}

func (e *ReadyConditionImpactingError) WithRetryClassification(classification retry.Classification) *ReadyConditionImpactingError {
	e.RetryClassification = classification
	return e
}

func (e *ReadyConditionImpactingError) Error() string {
	return fmt.Sprintf("Reason: %s, Severity: %s, RetryClassification: %s, Cause: %s",
		e.Reason,
		e.Severity,
		e.RetryClassification,
		e.cause.Error())
}

func (e *ReadyConditionImpactingError) Is(err error) bool {
	var typedErr *ReadyConditionImpactingError
	if eris.As(err, &typedErr) {
		return e.Severity == typedErr.Severity && e.Reason == typedErr.Reason
	}
	return false
}

func (e *ReadyConditionImpactingError) Cause() error {
	return e.cause
}

// This was adapted from the function in errors
func (e *ReadyConditionImpactingError) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		if s.Flag('+') {
			n, _ := fmt.Fprintf(s, "%s", e.Cause())
			if n > 0 {
				_, _ = fmt.Fprintf(s, "\n")
			}
			_, _ = io.WriteString(s, e.Error())
			return
		}
		fallthrough
	case 's':
		_, _ = io.WriteString(s, e.Error())
	case 'q':
		_, _ = fmt.Fprintf(s, "%q", e.Error())
	}
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
