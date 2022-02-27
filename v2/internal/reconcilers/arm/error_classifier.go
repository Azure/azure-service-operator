/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package arm

import (
	"github.com/Azure/azure-service-operator/v2/internal/genericarmclient"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/core"
)

func ClassifyCloudError(err *genericarmclient.CloudError) (core.CloudErrorDetails, error) {
	if err == nil {
		// Default to retrying if we're asked to classify a nil error
		result := core.CloudErrorDetails{
			Classification: core.ErrorRetryable,
			Code:           core.UnknownErrorCode,
			Message:        core.UnknownErrorMessage,
		}
		return result, nil
	}

	classification := classifyCloudErrorCode(err.Code())
	result := core.CloudErrorDetails{
		Classification: classification,
		Code:           err.Code(),
		Message:        err.Message(),
	}
	return result, nil
}

func classifyCloudErrorCode(code string) core.ErrorClassification {
	// See https://docs.microsoft.com/en-us/azure/azure-resource-manager/templates/common-deployment-errors
	// for a breakdown of common deployment error codes. Note that the error codes documented there are
	// the inner error codes we're parsing here.

	if code == "" {
		// If there's no code, assume we can retry on it
		return core.ErrorRetryable
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
		return core.ErrorRetryable
	case "BadRequestFormat",
		"Conflict",
		"BadRequest",
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
		return core.ErrorFatal
	default:
		// TODO: We could technically avoid listing the above Retryable errors since that's the default anyway
		// If we don't know what the error is, default to retrying on it
		return core.ErrorRetryable
	}
}
