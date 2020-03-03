// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package telemetry

// BaseTelemetry contains the helper functions for basic telemetry
type BaseTelemetry interface {
	LogTrace(typeTrace string, message string)
	LogInfo(typeInfo string, message string)
	LogWarning(typeWarning string, message string)
	LogError(message string, err error)
	LogStart()
	LogSuccess()
	LogFailure()
}
