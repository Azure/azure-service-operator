// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package telemetry

import "github.com/Azure/azure-service-operator/api/v1alpha1"
import "github.com/prometheus/client_golang/prometheus"

var (
	statusCounter  *prometheus.CounterVec
	activeGuage    *prometheus.GaugeVec
	successCounter *prometheus.CounterVec
	failureCounter *prometheus.CounterVec
	executionTime  *prometheus.HistogramVec
)

// namespaceName is the default namespace for this project
const namespaceName = "Azure"

// subsystemName is the default subsystem for this project
const subsystemName = "Operators"

type TelemetryFactory interface {
	Logger    log.Logger
	CreateTelemetry(componentName string, instanceName string) (*Telemtry, error)
}

// InitializeTelemetryFactoryDefault initializes a TelemetryFactory client
func InitializeTelemetryFactoryDefault(logger log.Logger) (factory *TelemetryFactory) {

	// initialize globals if neccessary
	if statusCounter == nil {
		initializeGlobalPrometheusMetricsEtc()
	}

	// create the client
	return &TelemetryFactory{
		Logger:    logger,
	}
}

// initializeAllPrometheusMetricsEtc inits all counts
func initializeGlobalPrometheusMetricsEtc() {

	statusCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: *namespace,
			Subsystem: *subsystem,
			Name:      "Status",
			Help:      "Status messages",
		},
		[]string{
			"component",
			"level",
			"type",
			"message",
		},
	)
	ctrlmetrics.Registry.MustRegister(statusCounter)

	executionTime = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: *namespace,
			Subsystem: *subsystem,
			Name:      "ExecutionTime",
			Buckets: prometheus.LinearBuckets(
				exeuctionTimeStart,
				exeuctionTimeWidth,
				executionTimeBuckets),
			Help: "Time to execute",
		},
		[]string{
			"component",
		},
	)
	ctrlmetrics.Registry.MustRegister(executionTime)

	activeGuage = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: *namespace,
			Subsystem: *subsystem,
			Name:      "ActiveGuage",
			Help:      "How many active segments of code",
		},
		[]string{
			"component",
		},
	)
	ctrlmetrics.Registry.MustRegister(activeGuage)

	successCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: *namespace,
			Subsystem: *subsystem,
			Name:      "Success",
			Help:      "Count of Successes",
		},
		[]string{
			"component",
		},
	)
	ctrlmetrics.Registry.MustRegister(successCounter)

	failureCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: *namespace,
			Subsystem: *subsystem,
			Name:      "Failure",
			Help:      "Count of Failures",
		},
		[]string{
			"component",
		},
	)
	ctrlmetrics.Registry.MustRegister(failureCounter)
}
