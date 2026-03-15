/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package podmetrics

import (
	"context"
	"fmt"
	"slices"
	"strings"
	"sync"
	"time"

	"github.com/go-logr/logr"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
	metricsv1beta1 "k8s.io/metrics/pkg/apis/metrics/v1beta1"
	metricsclient "k8s.io/metrics/pkg/client/clientset/versioned"
)

const (
	// DefaultNamespace is the namespace where ASO controller pods run.
	DefaultNamespace = "azureserviceoperator-system"

	// DefaultPodPrefix is the prefix of the ASO controller pod name.
	DefaultPodPrefix = "azureserviceoperator-controller-manager"

	// DefaultInterval is how frequently to poll for pod metrics.
	DefaultInterval = 5 * time.Second
)

// CollectorConfig holds configuration for a MetricsCollector.
type CollectorConfig struct {
	// Namespace is the Kubernetes namespace to monitor.
	Namespace string

	// PodPrefix filters pods whose names start with this prefix.
	PodPrefix string

	// Interval is how often to poll the metrics API.
	Interval time.Duration
}

// MetricsCollector polls the Kubernetes metrics API at a regular interval and stores samples.
type MetricsCollector struct {
	client    metricsclient.Interface
	namespace string
	podPrefix string
	interval  time.Duration

	mu      sync.Mutex
	samples []Sample
	start   time.Time

	cancel context.CancelFunc
	done   chan struct{}
}

// NewMetricsCollector creates a MetricsCollector that polls pod metrics from the given cluster.
// It targets pods in the specified namespace whose names start with podPrefix.
func NewMetricsCollector(cfg *rest.Config, collectorCfg CollectorConfig) (*MetricsCollector, error) {
	mc, err := metricsclient.NewForConfig(cfg)
	if err != nil {
		return nil, fmt.Errorf("creating metrics client: %w", err)
	}

	namespace := collectorCfg.Namespace
	if namespace == "" {
		namespace = DefaultNamespace
	}

	podPrefix := collectorCfg.PodPrefix
	if podPrefix == "" {
		podPrefix = DefaultPodPrefix
	}

	interval := collectorCfg.Interval
	if interval <= 0 {
		interval = DefaultInterval
	}

	return &MetricsCollector{
		client:    mc,
		namespace: namespace,
		podPrefix: podPrefix,
		interval:  interval,
	}, nil
}

// CheckAvailable probes the metrics API to verify that metrics-server is reachable.
// Returns an error if the metrics API is not available.
func (mc *MetricsCollector) CheckAvailable(ctx context.Context) error {
	_, err := mc.client.MetricsV1beta1().PodMetricses(mc.namespace).List(ctx, metav1.ListOptions{Limit: 1})
	if err != nil {
		return fmt.Errorf("metrics-server not available in namespace %q: %w", mc.namespace, err)
	}
	return nil
}

// Start begins collecting metrics in a background goroutine.
// Call Stop to end collection.
func (mc *MetricsCollector) Start(log logr.Logger) {
	mc.mu.Lock()
	mc.start = time.Now()
	mc.samples = nil
	mc.mu.Unlock()

	ctx, cancel := context.WithCancel(context.Background())
	mc.cancel = cancel
	mc.done = make(chan struct{})

	go func() {
		defer close(mc.done)
		ticker := time.NewTicker(mc.interval)
		defer ticker.Stop()

		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				if err := mc.collect(ctx); err != nil {
					log.Error(err, "metrics collection error")
				}
			}
		}
	}()

	log.Info("Metrics collector started",
		"namespace", mc.namespace,
		"podPrefix", mc.podPrefix,
		"interval", mc.interval)
}

// Stop ends metrics collection and waits for the collector goroutine to exit.
func (mc *MetricsCollector) Stop() {
	if mc.cancel != nil {
		mc.cancel()
		<-mc.done
	}
}

// Samples returns a copy of all collected samples.
func (mc *MetricsCollector) Samples() []Sample {
	mc.mu.Lock()
	defer mc.mu.Unlock()
	result := slices.Clone(mc.samples)
	return result
}

// collect performs a single metrics API poll and stores the results.
func (mc *MetricsCollector) collect(ctx context.Context) error {
	podMetricsList, err := mc.client.MetricsV1beta1().PodMetricses(mc.namespace).List(ctx, metav1.ListOptions{})
	if err != nil {
		return fmt.Errorf("listing pod metrics: %w", err)
	}

	now := time.Now()
	mc.mu.Lock()
	defer mc.mu.Unlock()

	for i := range podMetricsList.Items {
		pod := &podMetricsList.Items[i]
		if !strings.HasPrefix(pod.Name, mc.podPrefix) {
			continue
		}

		mc.collectPodContainers(pod, now)
	}

	return nil
}

// collectPodContainers extracts metrics from each container in a pod.
func (mc *MetricsCollector) collectPodContainers(pod *metricsv1beta1.PodMetrics, now time.Time) {
	for i := range pod.Containers {
		container := &pod.Containers[i]

		cpuMillis := container.Usage.Cpu().MilliValue()
		memBytes := container.Usage.Memory().Value()

		sample := Sample{
			Timestamp:     now,
			Elapsed:       now.Sub(mc.start),
			PodName:       pod.Name,
			ContainerName: container.Name,
			CPUMillicores: float64(cpuMillis),
			MemoryBytes:   memBytes,
		}
		mc.samples = append(mc.samples, sample)
	}
}
