// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package telemetry

import (
	"time"

	"github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/prometheus/client_golang/prometheus"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// LogTrace logs a trace message, it does not send telemetry to Prometheus
func (client *PrometheusClient) LogTrace(typeTrace string, message string) {
	client.Logger.Info(message, "Trace Type", typeTrace, "Component", client.Component)
}

// LogInfo logs an informational message
func (client *PrometheusClient) LogInfo(typeInfo string, message string) {
	client.Logger.Info(message, "Info Type", typeInfo, "Component", client.Component)
	statusCounter.WithLabelValues(client.Component, "Info", typeInfo, message).Inc()
}

// LogWarning logs a warning
func (client *PrometheusClient) LogWarning(typeWarning string, message string) {
	// logs this as info as there's no go-logr warning level
	client.Logger.Info(message, "Warning Type", typeWarning, "Component", client.Component)
	statusCounter.WithLabelValues(client.Component, "Warning", typeWarning, message).Inc()
}

// LogError logs an error
func (client *PrometheusClient) LogError(message string, err error) {
	errorString := err.Error()

	// logs the error as info (eventhough there's an error) as this follows the previous pattern
	client.Logger.Info(message, "Component", client.Component, "Error", errorString)
	statusCounter.WithLabelValues(client.Component, "Error", "Error", errorString).Inc()
}

// LogStart logs the start of a component
func (client *PrometheusClient) LogStart(s *v1alpha1.ASOStatus) {

	// mark now as the start time
	//client.StartTime = time.Now()

	if !s.Provisioning {
		s.Provisioning = true
		s.RequestedAt = metav1.NewTime(time.Now())
	}

	// log that operator is running
	//client.Logger.Info("Component has started", "Component", client.Component)
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
func (client *PrometheusClient) LogSuccess(s *v1alpha1.ASOStatus) {

	// log the completion
	//logCompleted(client)

	if s.Provisioning {
		s.Provisioning = false
		s.Provisioned = true
		s.CompletedAt = metav1.NewTime(time.Now())
		durationInSecs := s.CompletedAt.Sub(time.Now()).Seconds()
		executionTime.WithLabelValues(client.Component).Observe(durationInSecs)
	}

	// log success
	client.Logger.Info("Component completed successfully", "Component", client.Component)
	successCounter.WithLabelValues(client.Component).Inc()
}

// LogFailureWithError logs the unsuccessful completion of a component along with an error
func (client *PrometheusClient) LogFailureWithError(message string, err error) {

	// log error
	client.LogError(message, err)

	// log completion
	client.LogFailure()
}

// LogFailure logs the unsuccessful completion of a component
func (client *PrometheusClient) LogFailure() {

	// log the completion
	logCompleted(client)

	// log failure
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
