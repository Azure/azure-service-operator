// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package telemetry

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

// LogTrace logs a trace message, it does not send telemetry to Prometheus
func (client *PrometheusClient) LogTrace(typeTrace string, message string) {
	client.Logger.Info(message, "Trace Type", typeTrace, "Component", client.Component)
}

// LogInfo logs an informational message
func (client *PrometheusClient) LogInfo(typeInfo string, message string) {
	client.Logger.Info(message, "Info Type", typeInfo, "Component", client.Component)
	infoCounter.WithLabelValues(client.Component, typeInfo, message).Inc()
}

// LogWarning logs a warning
func (client *PrometheusClient) LogWarning(typeWarning string, message string) {
	// logs this as info as there's no go-logr warning level
	client.Logger.Info(message, "Warning Type", typeWarning, "Component", client.Component)
	warningCounter.WithLabelValues(client.Component, typeWarning, message).Inc()
}

// LogError logs an error
func (client *PrometheusClient) LogError(message string, err error) {
	errorString := err.Error()

	// logs the error as info (eventhough there's an error) as this follows the previous pattern
	client.Logger.Info(message, "Component", client.Component, "Error", errorString)
	infoCounter.WithLabelValues(client.Component, errorString).Inc()
}

// LogStart logs the start of a component
func (client *PrometheusClient) LogStart() {

	// mark now as the start time
	client.StartTime = time.Now()

	// log that operator is running
	client.Logger.Info("Component has started", "Component", client.Component)
	activeGuage.WithLabelValues(client.Component).Inc()
}

// logCompletedOperation logs a component as completed
func logCompleted(client *PrometheusClient) {

	// log that operator has completed
	activeGuage.WithLabelValues(client.Component).Dec()

	// record the time for the histogram
	timeNow := time.Now()
	durationInSecs := timeNow.Sub(client.StartTime).Seconds()
	client.Logger.Info("Component has completed", "Component", client.Component, "Duration", durationInSecs)
	executionTime.WithLabelValues(client.Component).Observe(durationInSecs)
}

// LogSuccess logs the successful completion of a component
func (client *PrometheusClient) LogSuccess() {

	// log the completion
	logCompleted(client)

	// log success
	client.Logger.Info("Component completed successfully", "Component", client.Component)
	successCounter.WithLabelValues(client.Component).Inc()
}

// LogFailure logs the successful completion of a component
func (client *PrometheusClient) LogFailure() {

	// log the completion
	logCompleted(client)

	// log success
	client.Logger.Info("Component completed unsuccessfully", "Component", client.Component)
	failureCounter.WithLabelValues(client.Component).Inc()
}

// CreateHistogram creates a histogram, start = what value the historgram starts at, width = how wide are the buckets,
// numberOfBuckets = the number of buckets in the histogram
func (client *PrometheusClient) CreateHistogram(name string, start float64, width float64, numberOfBuckets int) (histogram prometheus.Histogram) {
	histogram = prometheus.NewHistogram(
		prometheus.HistogramOpts{
			Namespace: *namespace,
			Subsystem: *subsystem,
			Name:      name,
			Buckets:   prometheus.LinearBuckets(start, width, numberOfBuckets),
		})

	return histogram
}
