// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package telemetry

import (
	log "github.com/go-logr/logr"
	"github.com/prometheus/client_golang/prometheus"
)

const namespace string = "Azure"
const subsystem string = "ServiceOperators"

var (
	infoCounter    *prometheus.CounterVec
	warningCounter *prometheus.CounterVec
	errorCounter   *prometheus.CounterVec
)

// LogTrace logs a trace message
func LogTrace(location string, typeTrace string, message string, logger log.Logger) {
	logger.Info(message, "Trace Type", typeTrace, "Location", location)
}

// LogInfo logs an informational message
func LogInfo(location string, typeInfo string, message string, logger log.Logger) {
	// init if needed
	if infoCounter == nil {
		infoCounter = prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Namespace: namespace,
				Subsystem: subsystem,
				Name:      "Info",
			},
			[]string{
				"location",
				"type",
				"message",
			},
		)
		prometheus.MustRegister(infoCounter)
	}

	logger.Info(message, "Info Type", typeInfo, "Location", location)
	infoCounter.WithLabelValues(location, typeInfo, message).Inc()
}

// LogWarning logs a warning
func LogWarning(location string, typeWarning string, message string, logger log.Logger) {

	// init if needed
	if warningCounter == nil {
		warningCounter = prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Namespace: namespace,
				Subsystem: subsystem,
				Name:      "Warning",
			},
			[]string{
				"location",
				"type",
				"message",
			},
		)
		prometheus.MustRegister(warningCounter)
	}

	// logs this as info as there's no go-logr warning level
	logger.Info(message, "Warning Type", typeWarning, "Location", location)
	warningCounter.WithLabelValues(location, typeWarning, message).Inc()
}

// LogError logs an error
func LogError(location string, message string, err error, logger log.Logger) {
	errorString := err.Error()

	// init if needed
	if errorCounter == nil {
		errorCounter = prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Namespace: namespace,
				Subsystem: subsystem,
				Name:      "Error",
			},
			[]string{
				"location",
				"error",
			},
		)
		prometheus.MustRegister(errorCounter)
	}

	// logs the error as info (eventhough there's an error) as this follows the previous pattern
	logger.Info(message, "Location", location, "Error", errorString)
	infoCounter.WithLabelValues(location, errorString).Inc()
}

// CreateHistogram creates a histogram, start = what value the historgram starts at, width = how wide are the buckets,
// numberOfBuckets = the number of buckets in the histogram
func CreateHistogram(name string, start float64, width float64, numberOfBuckets int) (histogram prometheus.Histogram) {
	histogram = prometheus.NewHistogram(
		prometheus.HistogramOpts{
			Namespace: namespace,
			Subsystem: subsystem,
			Name:      name,
			Buckets:   prometheus.LinearBuckets(start, width, numberOfBuckets),
		})

	return histogram
}
