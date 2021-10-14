/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package reconcilers

import (
	"github.com/Azure/azure-service-operator/v2/internal/armclient"
)

type DeploymentErrorClassification string

const (
	DeploymentErrorRetryable = DeploymentErrorClassification("retryable")
	DeploymentErrorFatal     = DeploymentErrorClassification("fatal")
)

type DeploymentErrorDetails struct {
	Classification DeploymentErrorClassification
	Code           string
	Message        string
}

const (
	UnknownErrorCode    = "UnknownError"
	UnknownErrorMessage = "There was an unknown deployment error"
)

func stringOrDefault(str string, def string) string {
	if str == "" {
		return def
	}

	return str
}

func stringPtrOrDefault(str *string, def string) string {
	if str == nil {
		return def
	}

	if *str == "" {
		return def
	}

	return *str
}

func ClassifyDeploymentError(deploymentError *armclient.DeploymentError) DeploymentErrorDetails {
	if deploymentError == nil {
		// Default to retrying if we're asked to classify a nil error
		return DeploymentErrorDetails{
			Classification: DeploymentErrorRetryable,
			Code:           UnknownErrorCode,
			Message:        UnknownErrorMessage,
		}
	}

	if len(deploymentError.Details) == 0 {
		// Default to retrying if we're asked to classify an error with no details
		return DeploymentErrorDetails{
			Classification: DeploymentErrorRetryable,
			Code:           stringOrDefault(deploymentError.Code, UnknownErrorCode),
			Message:        stringOrDefault(deploymentError.Message, UnknownErrorMessage),
		}
	}

	// Classify all of the details -- there may ALWAYS be only one but
	// since the API technically allows a list just deal with it in case
	// it actually happens in some rare case.

	// First check if any errors are fatal
	for _, detail := range deploymentError.Details {
		classification := classifyInnerDeploymentError(detail)
		// A single fatal sub-error means the error as a whole is fatal
		if classification == DeploymentErrorFatal {
			return DeploymentErrorDetails{
				Classification: classification,
				Code:           stringOrDefault(detail.Code, UnknownErrorCode),
				Message:        stringOrDefault(detail.Message, UnknownErrorMessage),
			}
		}
	}

	// Otherwise return the first error (which must have been retryable since we didn't return above)
	return DeploymentErrorDetails{
		Classification: DeploymentErrorRetryable,
		Code:           stringOrDefault(deploymentError.Details[0].Code, UnknownErrorCode),
		Message:        stringOrDefault(deploymentError.Details[0].Message, UnknownErrorMessage),
	}
}

func classifyInnerDeploymentError(deploymentError armclient.DeploymentError) DeploymentErrorClassification {
	// See https://docs.microsoft.com/en-us/azure/azure-resource-manager/templates/common-deployment-errors
	// for a breakdown of common deployment error codes. Note that the error codes documented there are
	// the inner error codes we're parsing here.
	if deploymentError.Code == "" {
		// If there's no code, assume we can retry on it
		return DeploymentErrorRetryable
	}

	switch deploymentError.Code {
	case "AnotherOperationInProgress",
		"AuthorizationFailed",
		"AllocationFailed",
		"InvalidResourceReference",
		"InvalidSubscriptionRegistrationState",
		"LinkedAuthorizationFailed",
		"MissingRegistrationForLocation",
		"MissingSubscriptionRegistration",
		"NoRegisteredProviderFound",
		"NotFound",
		// It sounds weird to retry on "OperationNotAllowed" but according to the docs
		// it's a quota issue, so we can in theory retry through it
		"OperationNotAllowed",
		"ParentResourceNotFound",
		"ResourceGroupNotFound",
		"ResourceNotFound",
		"ResourceQuotaExceeded",
		"SubscriptionNotRegistered":
		return DeploymentErrorRetryable
	case "BadRequestFormat",
		"Conflict",                     // TODO: is conflict always not retryable?
		"PublicIpForGatewayIsRequired", // TODO: There's not a great way to look at an arbitrary error returned by this API and determine if it's a 4xx or 5xx level... ugh
		"InvalidParameter",
		"InvalidParameterValue",
		"InvalidResourceType",
		"InvalidRequestContent",
		"InvalidTemplate",
		"InvalidValuesForRequestParameters",
		"InvalidGatewaySkuProvidedForGatewayVpnType",
		"InvalidGatewaySize",
		"LocationRequired",
		"MissingRequiredParameter",
		"PasswordTooLong",
		"PrivateIPAddressInReservedRange",
		"PrivateIPAddressNotInSubnet",
		"PropertyChangeNotAllowed",
		"RequestDisallowedByPolicy", // TODO: Technically could probably retry through this?
		"ReservedResourceName",
		"SkuNotAvailable",
		"SubscriptionNotFound":
		return DeploymentErrorFatal
	default:
		// TODO: We could technically avoid listing the above Retryable errors since that's the default anyway
		// If we don't know what the error is, default to retrying on it
		return DeploymentErrorRetryable
	}
}
