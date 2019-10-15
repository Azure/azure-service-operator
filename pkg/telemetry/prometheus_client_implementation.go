// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package telemetry

import (
	"math"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

// LogTrace logs a trace message, it does not send telemetry to Prometheus
func (client PrometheusClient) LogTrace(typeTrace string, message string) {
	client.Logger.Info(message, "Trace Type", typeTrace, "Component", client.Component)
}

// LogInfo logs an informational message
func (client PrometheusClient) LogInfo(typeInfo string, message string) {
	// init if needed
	if client.InfoCounter == nil {
		client.InfoCounter = prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Namespace: client.Namespace,
				Subsystem: client.Subsystem,
				Name:      "Info",
			},
			[]string{
				"component",
				"type",
				"message",
			},
		)
		prometheus.MustRegister(client.InfoCounter)
	}

	client.Logger.Info(message, "Info Type", typeInfo, "Component", client.Component)
	client.InfoCounter.WithLabelValues(client.Component, typeInfo, message).Inc()
}

// LogWarning logs a warning
func (client PrometheusClient) LogWarning(typeWarning string, message string) {

	// init if needed
	if client.WarningCounter == nil {
		client.WarningCounter = prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Namespace: client.Namespace,
				Subsystem: client.Subsystem,
				Name:      "Warning",
			},
			[]string{
				"component",
				"type",
				"message",
			},
		)
		prometheus.MustRegister(client.WarningCounter)
	}

	// logs this as info as there's no go-logr warning level
	client.Logger.Info(message, "Warning Type", typeWarning, "Component", client.Component)
	client.WarningCounter.WithLabelValues(client.Component, typeWarning, message).Inc()
}

// LogError logs an error
func (client PrometheusClient) LogError(message string, err error) {
	errorString := err.Error()

	// init if needed
	if client.ErrorCounter == nil {
		client.ErrorCounter = prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Namespace: client.Namespace,
				Subsystem: client.Subsystem,
				Name:      "Error",
			},
			[]string{
				"component",
				"error",
			},
		)
		prometheus.MustRegister(client.ErrorCounter)
	}

	// logs the error as info (eventhough there's an error) as this follows the previous pattern
	client.Logger.Info(message, "Component", client.Component, "Error", errorString)
	client.InfoCounter.WithLabelValues(client.Component, errorString).Inc()
}

// LogStart logs the start of a component
func (client PrometheusClient) LogStart() {

	// mark now as the start time
	client.StartTime = time.Now()

	// init if needed
	if client.ExecutionTime == nil {
		client.ExecutionTime = prometheus.NewHistogramVec(
			prometheus.HistogramOpts{
				Namespace: client.Namespace,
				Subsystem: client.Subsystem,
				Name:      "ExecutionTime",
				Buckets: prometheus.LinearBuckets(exeuctionTimeStart,
					exeuctionTimeWidth,
					executionTimeBuckets),
			},
			[]string{
				"component",
			},
		)
		prometheus.MustRegister(client.ExecutionTime)
	}

	// init if needed
	if client.ActiveGuage == nil {
		client.ActiveGuage = prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Namespace: client.Namespace,
				Subsystem: client.Subsystem,
				Name:      "ActiveGuage",
			},
			[]string{
				"component",
			},
		)
		prometheus.MustRegister(client.ActiveGuage)
	}

	// log that operator is running
	client.Logger.Info("Component has started", "Component", client.Component)
	client.ActiveGuage.WithLabelValues(client.Component).Inc()
}

// logCompletedOperation logs a component as completed
func logCompleted(client PrometheusClient) {

	// init if needed
	if client.ExecutionTime == nil {
		client.StartTime = time.Now()
		client.ExecutionTime = prometheus.NewHistogramVec(
			prometheus.HistogramOpts{
				Namespace: client.Namespace,
				Subsystem: client.Subsystem,
				Name:      "ExecutionTime",
				Buckets: prometheus.LinearBuckets(exeuctionTimeStart,
					exeuctionTimeWidth,
					executionTimeBuckets),
			},
			[]string{
				"component",
			},
		)
		prometheus.MustRegister(client.ExecutionTime)
	}

	// init if needed
	if client.ActiveGuage == nil {
		client.ActiveGuage = prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Namespace: client.Namespace,
				Subsystem: client.Subsystem,
				Name:      "ActiveGuage",
			},
			[]string{
				"component",
			},
		)
		prometheus.MustRegister(client.ActiveGuage)
	}

	// log that operator has completed
	client.ActiveGuage.WithLabelValues(client.Component).Dec()

	// record the time for the histogram
	durationInSecs := float64(time.Since(client.StartTime)) * math.Pow10(9)
	client.Logger.Info("Component has completed", "Component", client.Component, "Duration", durationInSecs)
	client.ExecutionTime.WithLabelValues(client.Component).Observe(durationInSecs)
}

// LogSuccess logs the successful completion of a component
func (client PrometheusClient) LogSuccess() {

	// log the completion
	logCompleted(client)

	// init if needed
	if client.SuccessCounter == nil {
		client.SuccessCounter = prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Namespace: client.Namespace,
				Subsystem: client.Subsystem,
				Name:      "Success",
			},
			[]string{
				"component",
			},
		)
		prometheus.MustRegister(client.SuccessCounter)
	}

	// log success
	client.Logger.Info("Component completed successfully", "Component", client.Component)
	client.SuccessCounter.WithLabelValues(client.Component).Inc()
}

// LogFailure logs the successful completion of a component
func (client PrometheusClient) LogFailure() {

	// log the completion
	logCompleted(client)

	// init if needed
	if client.FailureCounter == nil {
		client.FailureCounter = prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Namespace: client.Namespace,
				Subsystem: client.Subsystem,
				Name:      "Failure",
			},
			[]string{
				"component",
			},
		)
		prometheus.MustRegister(client.FailureCounter)
	}

	// log success
	client.Logger.Info("Component completed unsuccessfully", "Component", client.Component)
	client.FailureCounter.WithLabelValues(client.Component).Inc()
}

// CreateHistogram creates a histogram, start = what value the historgram starts at, width = how wide are the buckets,
// numberOfBuckets = the number of buckets in the histogram
func (client PrometheusClient) CreateHistogram(name string, start float64, width float64, numberOfBuckets int) (histogram prometheus.Histogram) {
	histogram = prometheus.NewHistogram(
		prometheus.HistogramOpts{
			Namespace: client.Namespace,
			Subsystem: client.Subsystem,
			Name:      name,
			Buckets:   prometheus.LinearBuckets(start, width, numberOfBuckets),
		})

	return histogram
}
