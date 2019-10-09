// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package telemetry

import (
	"github.com/prometheus/client_golang/prometheus"
)

// LogTrace logs a trace message
func (client PrometheusClient) LogTrace(typeTrace string, message string) {
	client.Logger.Info(message, "Trace Type", typeTrace, "Operator", client.Operator)
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
				"operator",
				"type",
				"message",
			},
		)
		prometheus.MustRegister(client.InfoCounter)
	}

	client.Logger.Info(message, "Info Type", typeInfo, "Operator", client.Operator)
	client.InfoCounter.WithLabelValues(client.Operator, typeInfo, message).Inc()
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
				"operator",
				"type",
				"message",
			},
		)
		prometheus.MustRegister(client.WarningCounter)
	}

	// logs this as info as there's no go-logr warning level
	client.Logger.Info(message, "Warning Type", typeWarning, "Operator", client.Operator)
	client.WarningCounter.WithLabelValues(client.Operator, typeWarning, message).Inc()
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
				"operator",
				"error",
			},
		)
		prometheus.MustRegister(client.ErrorCounter)
	}

	// logs the error as info (eventhough there's an error) as this follows the previous pattern
	client.Logger.Info(message, "Operator", client.Operator, "Error", errorString)
	client.InfoCounter.WithLabelValues(client.Operator, errorString).Inc()
}

// LogStart los the start of an operator reconciling
func (client PrometheusClient) LogStart() {

	// init if needed
	if client.ActiveGuage == nil {
		client.ActiveGuage = prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Namespace: client.Namespace,
				Subsystem: client.Subsystem,
				Name:      "CurrrentOperatorCount",
			},
			[]string{
				"operator",
			},
		)
		prometheus.MustRegister(client.ActiveGuage)
	}

	// log that operator is running
	client.ActiveGuage.With(prometheus.Labels{"operator": client.Operator}).Inc()
}

func logCompleted(operator string, client PrometheusClient) {

	// init if needed
	if client.ActiveGuage == nil {
		client.ActiveGuage = prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Namespace: client.Namespace,
				Subsystem: client.Subsystem,
				Name:      "CurrrentOperatorCount",
			},
			[]string{
				"operator",
			},
		)
		prometheus.MustRegister(client.ActiveGuage)
	}

	// log that operator has completed
	client.ActiveGuage.With(prometheus.Labels{"operator": client.Operator}).Dec()
}

// LogSuccess logs the successful completion of an operator
func (client PrometheusClient) LogSuccess(operator string) {

	// log the completion
	logCompleted(operator, client)

	// init if needed
	if client.SuccessCounter == nil {
		client.SuccessCounter = prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Namespace: client.Namespace,
				Subsystem: client.Subsystem,
				Name:      "OperatorSuccess",
			},
			[]string{
				"operator",
			},
		)
		prometheus.MustRegister(client.SuccessCounter)
	}

	// log success
	client.Logger.Info("Operator completed successfully", "Operator", client.Operator)
	client.SuccessCounter.WithLabelValues(client.Operator).Inc()
}

// LogFailure logs the successful completion of an operator
func (client PrometheusClient) LogFailure(operator string) {

	// log the completion
	logCompleted(operator, client)

	// init if needed
	if client.FailureCounter == nil {
		client.FailureCounter = prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Namespace: client.Namespace,
				Subsystem: client.Subsystem,
				Name:      "OperatorSuccess",
			},
			[]string{
				"operator",
			},
		)
		prometheus.MustRegister(client.FailureCounter)
	}

	// log success
	client.Logger.Info("Operator completed unsuccessfully", "Operator", client.Operator)
	client.FailureCounter.WithLabelValues(client.Operator).Inc()
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
