// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package telemetry

import (
	"github.com/prometheus/client_golang/prometheus"
)

// PrometheusTelemetry contains the helper functions for Prometheus-based telemetry
type PrometheusTelemetry interface {
	BaseTelemetry
	CreateHistogram(name string, start float64, width float64, numberOfBuckets int) (histogram prometheus.Histogram)
}
