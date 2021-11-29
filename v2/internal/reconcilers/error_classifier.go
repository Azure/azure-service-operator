/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package reconcilers

import (
	"github.com/Azure/go-autorest/autorest/to"

	"github.com/Azure/azure-service-operator/v2/internal/genericarmclient"
)

type CloudErrorClassification string

const (
	CloudErrorRetryable = CloudErrorClassification("retryable")
	CloudErrorFatal     = CloudErrorClassification("fatal")
)

type CloudErrorDetails struct {
	Classification CloudErrorClassification
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

func ClassifyCloudError(err *genericarmclient.CloudError) CloudErrorDetails {
	if err == nil || err.InnerError == nil {
		// Default to retrying if we're asked to classify a nil error
		return CloudErrorDetails{
			Classification: CloudErrorRetryable,
			Code:           UnknownErrorCode,
			Message:        UnknownErrorMessage,
		}
	}

	classification := classifyInnerCloudError(err.InnerError)
	return CloudErrorDetails{
		Classification: classification,
		Code:           stringOrDefault(to.String(err.InnerError.Code), UnknownErrorCode),
		Message:        stringOrDefault(to.String(err.InnerError.Message), UnknownErrorMessage),
	}
}

func classifyInnerCloudError(err *genericarmclient.ErrorResponse) CloudErrorClassification {
	// See https://docs.microsoft.com/en-us/azure/azure-resource-manager/templates/common-deployment-errors
	// for a breakdown of common deployment error codes. Note that the error codes documented there are
	// the inner error codes we're parsing here.

	code := to.String(err.Code)
	if code == "" {
		// If there's no code, assume we can retry on it
		return CloudErrorRetryable
	}

	switch code {
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
		return CloudErrorRetryable
	case "BadRequestFormat",
		// TODO: See https://github.com/Azure/azure-service-operator/issues/1997 for why this is commented out
		// "BadRequest",
		"Conflict",                     // TODO: is conflict always not retryable?
		"PublicIpForGatewayIsRequired", // TODO: There's not a great way to look at an arbitrary error returned by this API and determine if it's a 4xx or 5xx level... ugh
		"InvalidParameter",
		"InvalidParameterValue",
		"InvalidResourceGroupLocation",
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
		return CloudErrorFatal
	default:
		// TODO: We could technically avoid listing the above Retryable errors since that's the default anyway
		// If we don't know what the error is, default to retrying on it
		return CloudErrorRetryable
	}
}
