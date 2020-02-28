// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package telemetry

import (
	"github.com/go-logr/logr"
	"github.com/prometheus/client_golang/prometheus"
	ctrlmetrics "sigs.k8s.io/controller-runtime/pkg/metrics"
)

// namespaceNameDefault is the default namespace for this project
const namespaceNameDefault = "Azure"

// subsystemNameDefault is the default subsystem for this project
const subsystemNameDefault = "Operators"

// exeuctionTimeStart base time == 0
const exeuctionTimeStart = 0

// exeuctionTimeWidth is the width of a bucket in the histogram, here it is 1m
const exeuctionTimeWidth = 60

// executionTimeBuckets is the number of buckets, here it 20 minutes worth of 1m buckets
const executionTimeBuckets = 20

var (
	statusCounter     *prometheus.CounterVec
	durationHistogram *prometheus.HistogramVec
)

// Telemetry contains data for the TelemetryClient
type Telemetry struct {
	NamespaceName string
	SubsystemName string
	ComponentName string
	Logger        logr.Logger
	Instance      string
}

// TelemetryClient contains the functions for Telemetry
type TelemetryClient interface {
	SetInstance(instance string)
	LogTrace(typeTrace string, message string)
	LogInfo(typeInfo string, message string)
	LogWarning(typeWarning string, message string)
	LogError(message string, err error)
	LogDuration(duration int)
	CreateHistogram(name string, start float64, width float64, numberOfBuckets int) (histogram prometheus.Histogram)
}

// InitializeTelemetryDefault initializes a TelemetryFactory client
func InitializeTelemetryDefault(componentName string, logger logr.Logger) *Telemetry {

	// initialize global metrics if neccessary
	if statusCounter == nil {
		initializeGlobalPrometheusMetrics()
	}

	return &Telemetry{
		NamespaceName: namespaceNameDefault,
		SubsystemName: subsystemNameDefault,
		ComponentName: componentName,
		Logger:        logger,
		Instance:      "",
	}
}

// initializeGlobalPrometheusMetrics inits all counts
func initializeGlobalPrometheusMetrics() {

	statusCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: namespaceNameDefault,
			Subsystem: subsystemNameDefault,
			Name:      "Status",
			Help:      "Status messages",
		},
		[]string{
			"component",
			"instance",
			"level",
			"type",
			"message",
		},
	)
	ctrlmetrics.Registry.MustRegister(statusCounter)

	durationHistogram = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: namespaceNameDefault,
			Subsystem: subsystemNameDefault,
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
	ctrlmetrics.Registry.MustRegister(durationHistogram)
}

// SetInstance must be called first if you want to set an instance for the subsequent telemetry calls
func (t *Telemetry) SetInstance(instance string) {
	t.Instance = instance
}

// LogTrace logs a trace message, it does not send telemetry to Prometheus
func (t *Telemetry) LogTrace(typeTrace string, message string) {
	t.Logger.Info(message, "Trace Type", typeTrace, "Component", t.ComponentName, "Instance", t.Instance)
}

// LogInfo logs an informational message
func (t *Telemetry) LogInfo(typeInfo string, message string) {
	t.Logger.Info(message, "Info Type", typeInfo, "Component", t.ComponentName, "Instance", t.Instance)
	statusCounter.WithLabelValues(t.ComponentName, t.Instance, "Info", typeInfo, message).Inc()
}

// LogWarning logs a warning
func (t *Telemetry) LogWarning(typeWarning string, message string) {

	// logs this as info as there's no go-logr warning level
	t.Logger.Info(message, "Warning Type", typeWarning, "Component", t.ComponentName, "Instance", t.Instance)
	statusCounter.WithLabelValues(t.ComponentName, t.Instance, "Warning", typeWarning, message).Inc()
}

// LogError logs an error
func (t *Telemetry) LogError(message string, err error) {
	errorString := err.Error()

	// logs the error as info (eventhough there's an error) as this follows the previous pattern
	t.Logger.Info(message, "Component", t.ComponentName, "Error", errorString, "Instance", t.Instance)
	statusCounter.WithLabelValues(t.ComponentName, t.Instance, "Error", "Error", errorString).Inc()
}

// LogDuration logs the duration of an operation in seconds
func (t *Telemetry) LogDuration(durationInSecs float64) {
	durationHistogram.WithLabelValues(t.ComponentName).Observe(durationInSecs)
}

// CreateHistogram creates a histogram, start = what value the historgram starts at, width = how wide are the buckets,
// numberOfBuckets = the number of buckets in the histogram
func (t *Telemetry) CreateHistogram(name string, start float64, width float64, numberOfBuckets int) (histogram prometheus.Histogram) {
	histogram = prometheus.NewHistogram(
		prometheus.HistogramOpts{
			Namespace: t.NamespaceName,
			Subsystem: t.SubsystemName,
			Name:      name,
			Buckets:   prometheus.LinearBuckets(start, width, numberOfBuckets),
		})

	return histogram
}
