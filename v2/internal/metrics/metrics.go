/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package metrics

import "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"

type Metrics interface {
	RegisterMetrics()
}

func RegisterMetrics(metrics ...Metrics) {
	for _, metric := range metrics {
		metric.RegisterMetrics()
	}
}

// GetTypeFromResourceID is a helper method for metrics to extract resourceType from resourceID
func GetTypeFromResourceID(resourceID string) (string, error) {
	id, err := arm.ParseResourceID(resourceID)
	if err != nil {
		return "", err
	}
	return id.ResourceType.String(), nil
}
