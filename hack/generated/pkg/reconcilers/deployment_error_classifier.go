package reconcilers

import "github.com/Azure/azure-service-operator/hack/generated/pkg/armclient"

type DeploymentErrorClassification string

const (
	DeploymentErrorRetryable = DeploymentErrorClassification("retryable")
	DeploymentErrorFatal     = DeploymentErrorClassification("fatal")
)

// TODO: need Some basic unit tests for this

func ClassifyDeploymentError(deploymentError *armclient.DeploymentError) DeploymentErrorClassification {
	if deploymentError == nil {
		// Default to retrying if we're asked to classify a nil error
		return DeploymentErrorRetryable
	}

	if len(deploymentError.Details) == 0 {
		// Default to retrying if we're asked to classify an error with no details
		return DeploymentErrorRetryable
	}

	// Classify all of the details -- there may ALWAYS be only one but
	// since the API technically allows a list just deal with it in case
	// it actually happens in some rare case.
	var classification DeploymentErrorClassification
	for _, detail := range deploymentError.Details {
		classification = classifyInnerDeploymentError(detail)
		// A single fatal sub-error means the error as a whole is fatal
		if classification == DeploymentErrorFatal {
			return classification
		}
	}

	return classification
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
	case "BadRequest",
		"Conflict", // TODO: is conflict always not retryable?
		"InvalidParameter",
		"InvalidRequestContent",
		"InvalidTemplate",
		"InvalidValuesForRequestParameters",
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
