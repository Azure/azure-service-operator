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
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/extensions"
)

var _ extensions.ErrorClassifier = &RedisExtension{}

// ClassifyError evaluates the provided error, returning including whether it is fatal or can be retried.
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
	next extensions.ErrorClassifierFunc) (genericarmclient.CloudErrorDetails, error) {
	details, err := next(cloudError)
	if err != nil {
		return details, err
	}

	// Override is to treat Conflict as retryable for Redis, if the message contains "try again later"
	// TODO: Do we need to check the message, or was that just a discriminator for when the code was generic?
	if details.Classification == genericarmclient.ErrorFatal &&
		cloudError.InnerError != nil &&
		cloudError.InnerError.Message != nil {
		inner := cloudError.InnerError
		if to.String(cloudError.InnerError.Code) == "Conflict" &&
			strings.Contains(strings.ToLower(*inner.Message), "try again later") {
			details.Classification = genericarmclient.ErrorRetryable
		}
	}

	return details, nil
}
