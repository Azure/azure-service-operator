/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package telemetry

import (
	"fmt"

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
}

// TelemetryClient contains the functions for Telemetry
type TelemetryClient interface {
	LogTrace(typeTrace string, message string)
	LogInfo(typeInfo string, message string)
	LogWarning(typeWarning string, message string)
	LogError(message string, err error)
	LogDuration(duration float64)
	LogTraceByInstance(typeTrace string, message string, instance string)
	LogInfoByInstance(typeInfo string, message string, instance string)
	LogWarningByInstance(typeWarning string, message string, instance string)
	LogErrorByInstance(message string, err error, instance string)
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

// LogTrace logs a trace message, it does not send telemetry to Prometheus
func (t *Telemetry) LogTrace(typeTrace string, message string) {
	t.LogTraceByInstance(typeTrace, message, "")
}

// LogInfo logs an informational message
func (t *Telemetry) LogInfo(typeInfo string, message string) {
	t.LogInfoByInstance(typeInfo, message, "")
}

// LogWarning logs a warning
func (t *Telemetry) LogWarning(typeWarning string, message string) {
	t.LogWarningByInstance(typeWarning, message, "")
}

// LogError logs an error
func (t *Telemetry) LogError(message string, err error) {
	t.LogErrorByInstance(message, err, "")
}

// LogDuration logs the duration of an operation in seconds
func (t *Telemetry) LogDuration(durationInSecs float64) {
	t.Logger.Info("duration", "Info Type", "duration", "Component", t.ComponentName, "Time", fmt.Sprintf("%f seconds", durationInSecs))
	durationHistogram.WithLabelValues(t.ComponentName).Observe(durationInSecs)
}

// LogTraceByInstance logs a trace
func (t *Telemetry) LogTraceByInstance(typeTrace string, message string, instance string) {
	t.Logger.Info(message, "Trace Type", typeTrace, "Component", t.ComponentName, "Instance", instance)
}

// LogInfoByInstance logs an informational message
func (t *Telemetry) LogInfoByInstance(typeInfo string, message string, instance string) {
	t.Logger.Info(message, "Info Type", typeInfo, "Component", t.ComponentName, "Instance", instance)
	statusCounter.WithLabelValues(t.ComponentName, instance, "Info", typeInfo, message).Inc()
}

// LogWarningByInstance logs a warning
func (t *Telemetry) LogWarningByInstance(typeWarning string, message string, instance string) {

	// logs this as info as there's no go-logr warning level
	t.Logger.Info(message, "Warning Type", typeWarning, "Component", t.ComponentName, "Instance", instance)
	statusCounter.WithLabelValues(t.ComponentName, instance, "Warning", typeWarning, message).Inc()
}

// LogErrorByInstance logs an error
func (t *Telemetry) LogErrorByInstance(message string, err error, instance string) {
	errorString := err.Error()

	// logs the error as info (eventhough there's an error) as this follows the previous pattern
	t.Logger.Info(message, "Component", t.ComponentName, "Error", errorString, "Instance", instance)
	statusCounter.WithLabelValues(t.ComponentName, instance, "Error", "Error", errorString).Inc()
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
