/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package metrics

import (
	"net/http"
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"sigs.k8s.io/controller-runtime/pkg/metrics"
)

type HttpMethod string

const (
	PutLabel    HttpMethod = http.MethodPut
	DeleteLabel HttpMethod = http.MethodGet
	GetLabel    HttpMethod = http.MethodGet
)

type ArmClientMetrics struct {
	azureRequestsTotal       *prometheus.CounterVec
	azureFailedRequestsTotal *prometheus.CounterVec
	azureRequestsTime        *prometheus.HistogramVec
}

func NewArmClientMetrics() ArmClientMetrics {

	azureRequestsTotal := prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "azure_requests_total",
		Help: "Total number of requests to azure",
	}, []string{"resource", "requestType", "responseCode"})

	azureFailedRequestsTotal := prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "azure_failed_requests_total",
		Help: "Total number of failed requests to azure",
	}, []string{"resource", "requestType"})

	azureRequestsTime := prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name: "azure_requests_time_seconds",
		Help: "Length of time per ARM request",
	}, []string{"resource", "requestType"})

	return ArmClientMetrics{
		azureRequestsTotal:       azureRequestsTotal,
		azureFailedRequestsTotal: azureFailedRequestsTotal,
		azureRequestsTime:        azureRequestsTime,
	}
}

func (a ArmClientMetrics) RegisterMetrics() {
	metrics.Registry.MustRegister(a.azureRequestsTime, a.azureRequestsTotal, a.azureFailedRequestsTotal)
}

func (a ArmClientMetrics) RecordAzureRequestsTotal(resourceName string, statusCode int, method HttpMethod) {
	a.azureRequestsTotal.WithLabelValues(resourceName, string(method), strconv.Itoa(statusCode)).Inc()
}

func (a ArmClientMetrics) RecordAzureFailedRequestsTotal(resourceName string, method HttpMethod) {
	a.azureFailedRequestsTotal.WithLabelValues(resourceName, string(method)).Inc()
}

func (a ArmClientMetrics) RecordAzureRequestsTime(resourceName string, requestTime time.Duration, method HttpMethod) {
	a.azureRequestsTime.WithLabelValues(resourceName, string(method)).Observe(requestTime.Seconds())
}
