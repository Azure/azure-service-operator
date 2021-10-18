/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package reconcilers

import (
	"github.com/Azure/go-autorest/autorest/azure"

	"github.com/Azure/azure-service-operator/v2/internal/armclient"
)

// stringOrEmptyString takes the provided input value and transforms it into a string.
// If the supplied value is nil, an empty string is returned.
// If the supplied value is not a string, an empty string is returned.
// If the supplied value is a non-nil string, the value (cast as a string) is returned
func stringOrEmptyString(value interface{}) string {
	if value == nil {
		return ""
	}

	result := ""
	temp, ok := value.(string)
	if ok {
		result = temp
	}

	return result
}

func translateAzureErrorDetailsToDeploymentErrorDetails(details []map[string]interface{}) []armclient.DeploymentError {
	if len(details) == 0 {
		return nil
	}

	result := make([]armclient.DeploymentError, len(details))

	for i, detail := range details {
		code := detail["code"]
		message := detail["message"]
		target := detail["target"]

		result[i] = armclient.DeploymentError{
			Code:    stringOrEmptyString(code),
			Message: stringOrEmptyString(message),
			Target:  stringOrEmptyString(target),
		}
	}

	return result
}

func TranslateAzureErrorToDeploymentError(err *azure.RequestError) *armclient.DeploymentError {
	if err == nil {
		return nil
	}

	if err.ServiceError == nil {
		return nil
	}

	result := &armclient.DeploymentError{
		Code:    err.ServiceError.Code,
		Message: err.ServiceError.Message,
		Target:  stringPtrOrDefault(err.ServiceError.Target, ""),
		Details: translateAzureErrorDetailsToDeploymentErrorDetails(err.ServiceError.Details),
	}

	return result
}
