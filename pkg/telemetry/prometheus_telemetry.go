// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package telemetry

import (
	"github.com/prometheus/client_golang/prometheus"
)

// PrometheusTelemetry contains the helper functions for Prometheus-based telemetry
type PrometheusTelemetry interface {
	BaseTelemetry
	CreateHistogram(name string, start float64, width float64, numberOfBuckets int) (histogram prometheus.Histogram)
}
