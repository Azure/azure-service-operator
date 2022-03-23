/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package metrics

import (
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"sigs.k8s.io/controller-runtime/pkg/metrics"
)

const (
	putLabel    = "PUT"
	deleteLabel = "DELETE"
)

// TODO: Add more status codes here if there are more expected codes.
// Basic http API request status codes
var statusCodes = []float64{
	http.StatusContinue, http.StatusSwitchingProtocols, http.StatusProcessing,
	http.StatusEarlyHints, http.StatusOK, http.StatusAccepted,
	http.StatusMovedPermanently, http.StatusNotModified, http.StatusTemporaryRedirect,
	http.StatusBadRequest, http.StatusUnauthorized, http.StatusNotFound,
	http.StatusInternalServerError, http.StatusNotImplemented, http.StatusBadGateway,
	http.StatusServiceUnavailable, http.StatusGatewayTimeout,
}

var (
	azureRequestsTotal = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "azure_requests_total",
		Help: "Total number of requests to azure",
	}, []string{"name", "requestType"})

	azureFailedRequestsTotal = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "azure_failed_requests_total",
		Help: "Total number of failed requests to azure",
	}, []string{"name", "requestType"})

	azureRequestsTime = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name: "azure_requests_time_seconds",
		Help: "Length of time per ARM request",
	}, []string{"name", "requestType"})

	azureResponseCode = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "azure_response_code",
		Help:    "Status codes returned by ARM request",
		Buckets: statusCodes,
	}, []string{"name", "requestType"})

	requeueTotal = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "requests_requeue_total",
		Help: "Total number of request requeue(s) for resource",
	}, []string{"name", "requestType"})
)

func init() {
	metrics.Registry.MustRegister(azureRequestsTotal, azureFailedRequestsTotal, azureRequestsTime, azureResponseCode)
}

func RecordRequeueTotalPUT(resourceName string) {
	requeueTotal.WithLabelValues(resourceName, putLabel)
}

func RecordRequeueTotalDELETE(resourceName string) {
	requeueTotal.WithLabelValues(resourceName, deleteLabel)
}

func RecordAzureRequestsTotalPUT(resourceName string) {
	azureRequestsTotal.WithLabelValues(resourceName, putLabel).Inc()
}

func RecordAzureRequestsTotalDELETE(resourceName string) {
	azureRequestsTotal.WithLabelValues(resourceName, deleteLabel).Inc()
}

func RecordAzureFailedRequestsTotalPUT(resourceName string) {
	azureFailedRequestsTotal.WithLabelValues(resourceName, putLabel).Inc()
}

func RecordAzureFailedRequestsTotalDELETE(resourceName string) {
	azureFailedRequestsTotal.WithLabelValues(resourceName, deleteLabel).Inc()
}

func RecordAzureRequestsTimePUT(resourceName string, requestTime time.Duration) {
	azureRequestsTime.WithLabelValues(resourceName, putLabel).Observe(requestTime.Seconds())
}

func RecordAzureRequestsTimeDELETE(resourceName string, requestTime time.Duration) {
	azureRequestsTime.WithLabelValues(resourceName, deleteLabel).Observe(requestTime.Seconds())
}

func RecordAzureResponseCodePUT(resourceName string, statusCode int) {
	azureResponseCode.WithLabelValues(resourceName, putLabel).Observe(float64(statusCode))
}

func RecordAzureResponseCodeDELETE(resourceName string, statusCode int) {
	azureResponseCode.WithLabelValues(resourceName, deleteLabel).Observe(float64(statusCode))
}
