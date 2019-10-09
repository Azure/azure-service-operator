// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package telemetry

import (
	"github.com/prometheus/client_golang/prometheus"
)

// PrometheusClient stores information for the Prometheus implementation of telemetry
type PrometheusClient struct {
	BaseClient
	Namespace      string
	Subsystem      string
	Operator       string
	InfoCounter    *prometheus.CounterVec
	WarningCounter *prometheus.CounterVec
	ErrorCounter   *prometheus.CounterVec
	ActiveGuage    *prometheus.GaugeVec
	SuccessCounter *prometheus.CounterVec
	FailureCounter *prometheus.CounterVec
}
