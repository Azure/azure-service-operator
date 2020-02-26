// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package telemetry

import "github.com/Azure/azure-service-operator/api/v1alpha1"

// Telemetry contains the helper functions for basic telemetry
type Telemetry interface {
	ComponentName string
	InstanceName string
	Factory TelemetryFactory
	LogTrace(typeTrace string, message string)
	LogInfo(typeInfo string, message string)
	LogWarning(typeWarning string, message string)
	LogError(message string, err error)
	CreateHistogram(name string, start float64, width float64, numberOfBuckets int) (histogram prometheus.Histogram)
}
