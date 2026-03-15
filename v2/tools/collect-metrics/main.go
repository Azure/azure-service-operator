/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

// collect-metrics is a standalone tool that polls the Kubernetes metrics API
// for ASO controller pod CPU and memory usage, writing results to a CSV file.
//
// Usage (from the v2/ directory):
//
//	go run ./tools/collect-metrics/ [flags]
//
// Examples:
//
//	# Collect until interrupted, writing to a timestamped file
//	go run ./tools/collect-metrics/
//
//	# Collect until interrupted, writing to the default reports folder
//	go run ./tools/collect-metrics/ --output-folder ./test/perf/reports
//
//	# Collect for 10 minutes with defaults, writing to metrics.csv
//	go run ./tools/collect-metrics/ --duration 10m --output metrics.csv
//
//	# Custom namespace, prefix, and interval
//	go run ./tools/collect-metrics/ --namespace my-ns --pod-prefix my-pod --interval 2s --duration 5m
//
//	# Run in background alongside tests
//	go run ./tools/collect-metrics/ --duration 30m --output metrics.csv &
//	METRICS_PID=$!
//	go test -v -run Test_Perf_Dynamic_ResourceGroups ./test/perf/...
//	kill $METRICS_PID
package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

	"k8s.io/klog/v2/textlogger"
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/Azure/azure-service-operator/v2/internal/testcommon/podmetrics"
)

func main() {
	namespace := flag.String("namespace", podmetrics.DefaultNamespace, "Kubernetes namespace to monitor")
	podPrefix := flag.String("pod-prefix", podmetrics.DefaultPodPrefix, "Pod name prefix to filter on")
	interval := flag.Duration("interval", podmetrics.DefaultInterval, "Polling interval (e.g. 5s, 10s) - default is 5s")
	duration := flag.Duration("duration", 0, "Collection duration (e.g. 5m, 1h). 0 means run until interrupted.")
	outputFolder := flag.String("output-folder", ".", "Output folder for the CSV file (used when --output is not set)")
	output := flag.String("output", "", "Output CSV file path (default: <output-folder>/metrics_<timestamp>.csv)")
	flag.Parse()

	log := textlogger.NewLogger(textlogger.NewConfig())
	ctrl.SetLogger(log)

	if *output == "" {
		if err := os.MkdirAll(*outputFolder, 0o755); err != nil {
			log.Error(err, "Failed to create output directory", "dir", *outputFolder)
			os.Exit(1)
		}
		*output = filepath.Join(*outputFolder, fmt.Sprintf("metrics_%s.csv", time.Now().Format("20060102_1504")))
	}

	cfg, err := ctrl.GetConfig()
	if err != nil {
		log.Error(err, "Failed to get kubeconfig")
		os.Exit(1)
	}

	collector, err := podmetrics.NewMetricsCollector(
		cfg,
		podmetrics.CollectorConfig{
			Namespace: *namespace,
			PodPrefix: *podPrefix,
			Interval:  *interval,
		})
	if err != nil {
		log.Error(err, "Failed to create metrics collector")
		os.Exit(1)
	}

	// Check that metrics-server is available
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := collector.CheckAvailable(ctx); err != nil {
		log.Error(err, "metrics-server is not available")
		os.Exit(1)
	}

	collector.Start(log)

	// Wait for duration or signal
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	if *duration > 0 {
		log.Info("Collecting metrics", "duration", *duration)
		select {
		case <-time.After(*duration):
			log.Info("Duration elapsed", "duration", *duration)
		case sig := <-sigCh:
			log.Info("Signal received", "signal", sig)
		}
	} else {
		log.Info("Collecting metrics until interrupted (Ctrl+C to stop)")
		sig := <-sigCh
		log.Info("Signal received", "signal", sig)
	}

	collector.Stop()

	// Write output
	samples := collector.Samples()
	if len(samples) == 0 {
		log.Info("No samples collected")
		return
	}

	f, err := os.Create(*output)
	if err != nil {
		log.Error(err, "Failed to create output file")
		os.Exit(1)
	}
	defer f.Close()

	if err := podmetrics.WriteCSV(f, samples); err != nil {
		log.Error(err, "Failed to write CSV")
		os.Exit(1)
	}

	log.Info("Wrote metrics", "samples", len(samples), "output", *output)
}
