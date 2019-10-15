// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package telemetry

import (
	"time"

	log "github.com/go-logr/logr"
	"github.com/prometheus/client_golang/prometheus"
)

// AzureOperatorNamespace is the default namespace for this project
const AzureOperatorNamespace = "Azure"

// AzureOperatorSubsystem is the default subsystem for this project
const AzureOperatorSubsystem = "Operators"

// exeuctionTimeStart base time == 0
const exeuctionTimeStart = 0

// exeuctionTimeWidth is the width of a bucket in the histogram, here it is 10s
const exeuctionTimeWidth = 10

// executionTimeBuckets is the number of buckets, here it 10 minutes worth of 10s buckets
const executionTimeBuckets = 6 * 10

// PrometheusClient stores information for the Prometheus implementation of telemetry
type PrometheusClient struct {
	Logger         log.Logger
	Namespace      string
	Subsystem      string
	Component      string
	StartTime      time.Time
	InfoCounter    *prometheus.CounterVec
	WarningCounter *prometheus.CounterVec
	ErrorCounter   *prometheus.CounterVec
	ActiveGuage    *prometheus.GaugeVec
	SuccessCounter *prometheus.CounterVec
	FailureCounter *prometheus.CounterVec
	ExecutionTime  *prometheus.HistogramVec
}

// InitializePrometheusClient easily initializes a PrometheusClient for an operator and logs the start of an operation
func InitializePrometheusClient(logger log.Logger, component string) (client *PrometheusClient) {
	client = &PrometheusClient{
		Logger:    logger,
		Namespace: AzureOperatorNamespace,
		Subsystem: AzureOperatorSubsystem,
		Component: component,
	}
	client.LogStart()
	return client
}
