/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package customizations

import (
	"strings"

	"github.com/Azure/go-autorest/autorest/to"
	"github.com/go-logr/logr"

	"github.com/Azure/azure-service-operator/v2/internal/genericarmclient"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/core"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/extensions"
)

var _ extensions.ErrorClassifier = &RedisExtension{}

// ClassifyError evaluates the provided error, returning whether it is fatal or can be retried.
// A conflict error (409) is normally fatal, but Redis resources may return 409 whilst a dependency is being created,
// so we override for that case.
// cloudError is the error returned from ARM.
// apiVersion is the ARM API version used for the request.
// log is a logger than can be used for telemetry.
// next is the next implementation to call.
func (e *RedisExtension) ClassifyError(
	cloudError *genericarmclient.CloudError,
	apiVersion string,
	log logr.Logger,
	next extensions.ErrorClassifierFunc) (core.CloudErrorDetails, error) {
	details, err := next(cloudError)
	if err != nil {
		return core.CloudErrorDetails{}, err
	}

	// Override is to treat Conflict as retryable for Redis, if the message contains "try again later"
	if isRetryableConflict(cloudError.InnerError) {
		details.Classification = core.ErrorRetryable
	}

	return details, nil
}

// isRetryableConflict checks the passed error to see if it is a retryable conflict, returning true if it is.
func isRetryableConflict(err *genericarmclient.ErrorResponse) bool {
	if err == nil || err.Message == nil {
		return false
	}

	return to.String(err.Code) == "Conflict" &&
		strings.Contains(strings.ToLower(*err.Message), "try again later")
}
