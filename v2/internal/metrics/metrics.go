/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package metrics

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-service-operator/v2/internal/logging"
	ctrl "sigs.k8s.io/controller-runtime"
)

type Metrics interface {
	RegisterMetrics()
}

func RegisterMetrics(metrics ...Metrics) {
	for _, metric := range metrics {
		metric.RegisterMetrics()
	}
}

// GetTypeFromResourceID is a helper method for metrics to extract resourceType from resourceID
func GetTypeFromResourceID(resourceID string) string {
	id, err := arm.ParseResourceID(resourceID)
	if err != nil {
		ctrl.Log.V(logging.Debug).Error(err, "Error while parsing", "resourceID", resourceID)
		// if error while parsing the resourceID, resourceType would be empty. Rather we send
		// the resourceID back as it is.
		return resourceID
	}

	return id.ResourceType.String()
}
