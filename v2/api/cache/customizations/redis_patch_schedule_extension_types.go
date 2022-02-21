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

var _ extensions.ErrorClassifier = &RedisPatchScheduleExtension{}

// ClassifyError evaluates the provided error, returning including whether it is fatal or can be retried.
// A conflict error (409) is normally fatal, but RedisPatchSchedule resources may return 409 whilst a dependency is
// being created, so we override for that case.
// cloudError is the error returned from ARM.
// apiVersion is the ARM API version used for the request.
// log is a logger than can be used for telemetry.
// next is the next implementation to call.
func (e *RedisPatchScheduleExtension) ClassifyError(
	cloudError *genericarmclient.CloudError,
	apiVersion string,
	log logr.Logger,
	next extensions.ErrorClassifierFunc) (genericarmclient.CloudErrorDetails, error) {
	details, err := next(cloudError)
	if err != nil {
		return genericarmclient.CloudErrorDetails{}, err
	}

	// Override is to treat Conflict as retryable for RedisPatchSchedules, if the message contains "try again later"
	if cloudError.InnerError != nil &&
		cloudError.InnerError.Message != nil {
		inner := cloudError.InnerError
		if to.String(inner.Code) == "Conflict" &&
			strings.Contains(strings.ToLower(*inner.Message), "try again later") {
			details.Classification = genericarmclient.ErrorRetryable
		}
	}

	return details, nil
}
