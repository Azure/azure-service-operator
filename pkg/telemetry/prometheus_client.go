// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package telemetry

import (
	"time"

	"github.com/Azure/go-autorest/autorest/to"

	log "github.com/go-logr/logr"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	namespace      *string
	subsystem      *string
	infoCounter    *prometheus.CounterVec
	warningCounter *prometheus.CounterVec
	errorCounter   *prometheus.CounterVec
	activeGuage    *prometheus.GaugeVec
	successCounter *prometheus.CounterVec
	failureCounter *prometheus.CounterVec
	executionTime  *prometheus.HistogramVec
)

// defaultNamespace is the default namespace for this project
const defaultNamespace = "Azure"

// defaultSubsystem is the default subsystem for this project
const defaultSubsystem = "Operators"

// exeuctionTimeStart base time == 0
const exeuctionTimeStart = 0

// exeuctionTimeWidth is the width of a bucket in the histogram, here it is 10s
const exeuctionTimeWidth = 10

// executionTimeBuckets is the number of buckets, here it 10 minutes worth of 10s buckets
const executionTimeBuckets = 6 * 10

// PrometheusClient stores information for the Prometheus implementation of telemetry
type PrometheusClient struct {
	Logger    log.Logger
	Component string
	StartTime time.Time
}

// InitializePrometheusDefault initializes a Prometheus client
func InitializePrometheusDefault(logger log.Logger, component string) (client *PrometheusClient) {
	return InitializePrometheus(logger, component, defaultNamespace, defaultSubsystem)
}

// InitializePrometheus initializes a Prometheus client
func InitializePrometheus(logger log.Logger, component string, namespaceLog string, subsystemLog string) (client *PrometheusClient) {

	// initialize globals if neccessary
	if namespace == nil {
		namespace = to.StringPtr(namespaceLog)
		subsystem = to.StringPtr(subsystemLog)
		initializeAllCounters()
	}

	// create the client
	return &PrometheusClient{
		Logger:    logger,
		Component: component,
		StartTime: time.Now(),
	}
}

// initializeAllCounters inits all counts
func initializeAllCounters() {

	infoCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: *namespace,
			Subsystem: *subsystem,
			Name:      "Info",
		},
		[]string{
			"component",
			"type",
			"message",
		},
	)
	prometheus.MustRegister(infoCounter)

	warningCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: *namespace,
			Subsystem: *subsystem,
			Name:      "Warning",
		},
		[]string{
			"component",
			"type",
			"message",
		},
	)
	prometheus.MustRegister(warningCounter)

	errorCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: *namespace,
			Subsystem: *subsystem,
			Name:      "Error",
		},
		[]string{
			"component",
			"error",
		},
	)
	prometheus.MustRegister(errorCounter)

	executionTime = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: *namespace,
			Subsystem: *subsystem,
			Name:      "ExecutionTime",
			Buckets: prometheus.LinearBuckets(
				exeuctionTimeStart,
				exeuctionTimeWidth,
				executionTimeBuckets),
		},
		[]string{
			"component",
		},
	)
	prometheus.MustRegister(executionTime)

	activeGuage = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: *namespace,
			Subsystem: *subsystem,
			Name:      "ActiveGuage",
		},
		[]string{
			"component",
		},
	)
	prometheus.MustRegister(activeGuage)

	successCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: *namespace,
			Subsystem: *subsystem,
			Name:      "Success",
		},
		[]string{
			"component",
		},
	)
	prometheus.MustRegister(successCounter)

	failureCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: *namespace,
			Subsystem: *subsystem,
			Name:      "Failure",
		},
		[]string{
			"component",
		},
	)
	prometheus.MustRegister(failureCounter)
}
