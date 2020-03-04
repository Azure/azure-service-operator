// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package telemetry

import (
	"time"

	"github.com/Azure/go-autorest/autorest/to"

	log "github.com/go-logr/logr"
	"github.com/prometheus/client_golang/prometheus"

	ctrlmetrics "sigs.k8s.io/controller-runtime/pkg/metrics"
)

var (
	namespace      *string
	subsystem      *string
	statusCounter  *prometheus.CounterVec
	activeGuage    *prometheus.GaugeVec
	successCounter *prometheus.CounterVec
	failureCounter *prometheus.CounterVec
	executionTime  *prometheus.HistogramVec
)

// defaultNamespace is the default namespace for this project
const defaultNamespace = "Azure"

// defaultSubsystem is the default subsystem for this project
const defaultSubsystem = "Operators"

// exeuctionTimeStart base time == 0
const exeuctionTimeStart = 0

// exeuctionTimeWidth is the width of a bucket in the histogram, here it is 10s
const exeuctionTimeWidth = 10

// executionTimeBuckets is the number of buckets, here it 10 minutes worth of 10s buckets
const executionTimeBuckets = 6 * 10

// PrometheusClient stores information for the Prometheus implementation of telemetry
type PrometheusClient struct {
	Logger    log.Logger
	Component string
	StartTime time.Time
}

// InitializePrometheusDefault initializes a Prometheus client
func InitializePrometheusDefault(logger log.Logger, component string) (client *PrometheusClient) {
	return InitializePrometheus(logger, component, defaultNamespace, defaultSubsystem)
}

// InitializePrometheus initializes a Prometheus client
func InitializePrometheus(logger log.Logger, component string, namespaceLog string, subsystemLog string) (client *PrometheusClient) {

	// initialize globals if neccessary
	if namespace == nil {
		namespace = to.StringPtr(namespaceLog)
		subsystem = to.StringPtr(subsystemLog)
		initializeGlobalPrometheusMetricsEtc()
	}

	// create the client
	return &PrometheusClient{
		Logger:    logger,
		Component: component,
		StartTime: time.Now(),
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
