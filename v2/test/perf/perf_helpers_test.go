/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package perf_test

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/samber/lo"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon/podmetrics"
)

// ResourceFactory is a function that, when called, returns a collection of resources to be created.
// The resources are returned in dependency order: owner resources first, dependent resources after.
// The tc parameter provides access to the test context for name generation, region, etc.
// The index parameter is a unique index for each invocation, useful for generating unique names or
// varying resource properties across invocations.
type ResourceFactory func(tc *testcommon.KubePerTestContext, index int) []client.Object

// Pattern defines the type of performance test pattern to run.
type Pattern string

const (
	// PatternStatic creates a fixed number of resources, waits for a configurable duration,
	// then deletes them all. Useful for measuring steady-state resource overhead.
	PatternStatic Pattern = "static"

	// PatternDynamic creates and deletes resources at a steady rate over a configurable duration.
	// Useful for measuring throughput and operator responsiveness under continuous load.
	PatternDynamic Pattern = "dynamic"
)

// Config holds all configuration for a performance test run.
type Config struct {
	// Pattern is the test pattern to execute (static or dynamic).
	Pattern Pattern

	// ResourceFactory is the function used to generate resources for each iteration.
	ResourceFactory ResourceFactory

	// StaticConfig holds configuration specific to the static pattern.
	// Only used when Pattern is PatternStatic.
	StaticConfig *StaticConfig

	// DynamicConfig holds configuration specific to the dynamic pattern.
	// Only used when Pattern is PatternDynamic.
	DynamicConfig *DynamicConfig
}

// StaticConfig holds configuration for the static test pattern.
type StaticConfig struct {
	// ResourceSetCount is the total number of resource sets to create (each set is one call to ResourceFactory).
	ResourceSetCount int

	// Duration is how long to wait after all resources are created before deleting them.
	Duration time.Duration
}

// DynamicConfig holds configuration for the dynamic test pattern.
type DynamicConfig struct {
	// Duration is the total duration over which to create and delete resources.
	Duration time.Duration

	// CreationRatePerMin is the rate at which resource sets are created (sets per minute).
	// Each "set" is one call to ResourceFactory.
	CreationRatePerMin float64

	// MaxConcurrentSets is the maximum number of resource sets that may exist at any time.
	// When this limit is reached, the oldest set is deleted before a new one is created.
	// If 0, there is no limit aside from what rate and duration imply.
	MaxConcurrentSets int
}

// Result captures the outcome of a performance test run.
type Result struct {
	// TotalSetsCreated is the total number of resource sets created during the test.
	TotalSetsCreated int

	// TotalSetsDeleted is the total number of resource sets deleted during the test.
	TotalSetsDeleted int

	// TotalResourcesCreated is the total number of individual resources created.
	TotalResourcesCreated int

	// TotalResourcesDeleted is the total number of individual resources deleted.
	TotalResourcesDeleted int

	// Elapsed is the total wall-clock time for the test.
	Elapsed time.Duration
}

// RunPerfTest executes a performance test based on the given configuration.
// It automatically collects CPU and memory metrics from the ASO controller pod
// and writes a CSV report to the output directory. The test fails immediately
// if metrics-server is not available on the cluster.
func RunPerfTest(t *testing.T, tc *testcommon.KubePerTestContext, cfg Config) (result Result, err error) {
	t.Helper()

	// Set up metrics collection
	collectorCfg := podmetrics.CollectorConfig{
		Namespace: podmetrics.DefaultNamespace,
		PodPrefix: podmetrics.DefaultPodPrefix,
		Interval:  podmetrics.DefaultInterval,
	}

	collector, err := podmetrics.NewMetricsCollector(tc.KubeConfig, collectorCfg)
	if err != nil {
		return Result{}, fmt.Errorf("creating metrics collector: %w", err)
	}

	// Probe metrics-server availability — fail fast if not present
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := collector.CheckAvailable(ctx); err != nil {
		return Result{}, fmt.Errorf("metrics-server is required for perf tests but is not available: %w", err)
	}

	log := ctrl.Log.WithName("perf-metrics")
	collector.Start(log)
	defer func() {
		collector.Stop()
		if metricsErr := writeMetricsCSV(t, collector); metricsErr != nil && err == nil {
			err = metricsErr
		}
	}()

	switch cfg.Pattern {
	case PatternStatic:
		return runStaticPattern(t, tc, cfg.ResourceFactory, cfg.StaticConfig)
	case PatternDynamic:
		return runDynamicPattern(t, tc, cfg.ResourceFactory, cfg.DynamicConfig)
	default:
		return Result{}, fmt.Errorf("unknown perf test pattern: %s", cfg.Pattern)
	}
}

// writeMetricsCSV writes the collected metrics samples to a CSV file.
func writeMetricsCSV(t *testing.T, collector *podmetrics.MetricsCollector) error {
	t.Helper()

	samples := collector.Samples()
	if len(samples) == 0 {
		t.Log("No metrics samples were collected")
		return nil
	}

	outDir := filepath.Join(".", "reports")

	if err := os.MkdirAll(outDir, 0o755); err != nil {
		return fmt.Errorf("failed to create output directory %s: %w", outDir, err)
	}

	// Sanitize test name for filename
	testName := strings.ReplaceAll(t.Name(), "/", "_")
	fileName := fmt.Sprintf("metrics_%s_%s.csv", testName, time.Now().Format("20060102_1504"))
	csvPath := filepath.Join(outDir, fileName)

	f, err := os.Create(csvPath)
	if err != nil {
		return fmt.Errorf("failed to create metrics CSV file: %w", err)
	}
	defer f.Close()

	if err := podmetrics.WriteCSV(f, samples); err != nil {
		return fmt.Errorf("failed to write metrics CSV: %w", err)
	}

	t.Logf("Metrics CSV written to %s (%d samples)", csvPath, len(samples))
	return nil
}

// resourceSet tracks a group of resources created from a single ResourceFactory invocation.
type resourceSet struct {
	index     int
	resources []client.Object
}

// runStaticPattern implements the static performance test pattern:
// 1. Create all resource sets
// 2. Wait for them to be provisioned
// 3. Hold for the configured duration
// 4. Delete all resource sets
func runStaticPattern(
	t *testing.T,
	tc *testcommon.KubePerTestContext,
	factory ResourceFactory,
	cfg *StaticConfig,
) (Result, error) { //nolint:unparam
	result := Result{}
	start := time.Now()

	t.Logf("Static pattern: creating %d resource sets", cfg.ResourceSetCount)

	// Phase 1: Create all resource sets
	sets := make([]resourceSet, 0, cfg.ResourceSetCount)
	for i := 0; i < cfg.ResourceSetCount; i++ {
		resources := factory(tc, i)
		set := resourceSet{
			index:     i,
			resources: resources,
		}

		t.Logf("Creating resource set %d/%d (%d resources)", i+1, cfg.ResourceSetCount, len(resources))
		sets = append(sets, set)
	}
	allResources := lo.FlatMap(sets, func(s resourceSet, _ int) []client.Object {
		return s.resources
	})
	tc.CreateResourcesAndWait(allResources...)
	result.TotalResourcesCreated += len(allResources)
	result.TotalSetsCreated += len(sets)

	t.Logf("All %d resource sets created (%d total resources). Holding for %s",
		len(sets), result.TotalResourcesCreated, cfg.Duration)

	// Phase 2: Hold
	time.Sleep(cfg.Duration)

	// Phase 3: Delete all resource sets
	t.Logf("Hold complete. Deleting %d resource sets", len(sets))
	tc.DeleteResourcesAndWait(allResources...)
	result.TotalResourcesDeleted += len(allResources)
	result.TotalSetsDeleted += len(sets)

	result.Elapsed = time.Since(start)
	t.Logf("Static pattern complete. Created %d sets (%d resources), deleted %d sets (%d resources) in %s",
		result.TotalSetsCreated,
		result.TotalResourcesCreated,
		result.TotalSetsDeleted,
		result.TotalResourcesDeleted,
		result.Elapsed)

	return result, nil
}

// runDynamicPattern implements the dynamic performance test pattern:
// Resources are created at a steady rate. When MaxConcurrentSets is reached, the oldest set is
// deleted before creating a new one. After Duration, all remaining sets are cleaned up.
//
// Each tick spawns a goroutine so that the ticker loop remains non-blocking even when resource
// creation or deletion takes longer than the tick interval. A mutex protects shared state (sets,
// result) and a WaitGroup ensures all in-flight goroutines complete before final cleanup.
func runDynamicPattern(
	t *testing.T,
	tc *testcommon.KubePerTestContext,
	factory ResourceFactory,
	cfg *DynamicConfig,
) (Result, error) {
	result := Result{}
	start := time.Now()

	if cfg.CreationRatePerMin <= 0 {
		return Result{}, fmt.Errorf("DynamicConfig.CreationRate must be > 0")
	}

	interval := time.Duration(float64(time.Minute) / cfg.CreationRatePerMin)

	t.Logf("Dynamic pattern: creating resources at %.2f sets/min for %s (interval %s, max concurrent %d)",
		cfg.CreationRatePerMin, cfg.Duration, interval, cfg.MaxConcurrentSets)

	var mu sync.Mutex
	var wg sync.WaitGroup
	var sets []resourceSet
	setIndex := 0

	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	deadline := time.After(cfg.Duration)

	done := false
	for !done {
		select {
		case <-deadline:
			done = true
			continue

		case <-ticker.C:
			mu.Lock()
			idx := setIndex
			setIndex++

			// If we've hit the max concurrent sets, delete the oldest one first
			var toDelete *resourceSet
			if cfg.MaxConcurrentSets > 0 && len(sets) >= cfg.MaxConcurrentSets {
				oldest := sets[0]
				sets = sets[1:]
				toDelete = &oldest
			}
			mu.Unlock()

			wg.Add(1)
			go func() {
				defer wg.Done()

				// Delete the evicted set if needed
				if toDelete != nil {
					t.Logf("Max concurrent sets reached (%d). Deleting oldest set (index %d)",
						cfg.MaxConcurrentSets, toDelete.index)
					tc.DeleteResourcesAndWait(toDelete.resources...)

					mu.Lock()
					result.TotalResourcesDeleted += len(toDelete.resources)
					result.TotalSetsDeleted++
					mu.Unlock()
				}

				// Create a new resource set
				resources := factory(tc, idx)
				set := resourceSet{
					index:     idx,
					resources: resources,
				}

				t.Logf("Creating resource set %d (%d resources)",
					idx, len(resources))
				tc.CreateResourcesAndWait(resources...)

				mu.Lock()
				result.TotalResourcesCreated += len(resources)
				result.TotalSetsCreated++
				sets = append(sets, set)
				mu.Unlock()
			}()
		}
	}

	// Wait for all in-flight goroutines to finish
	t.Logf("Duration complete. Waiting for in-flight operations to finish...")
	wg.Wait()

	// Cleanup: delete all remaining active sets
	t.Logf("Cleaning up %d remaining resource sets", len(sets))
	allResources := lo.FlatMap(sets, func(s resourceSet, _ int) []client.Object {
		return s.resources
	})
	tc.DeleteResourcesAndWait(allResources...)

	result.Elapsed = time.Since(start)
	t.Logf("Dynamic pattern complete. Created %d sets (%d resources), deleted %d sets (%d resources) in %s",
		result.TotalSetsCreated,
		result.TotalResourcesCreated,
		result.TotalSetsDeleted,
		result.TotalResourcesDeleted,
		result.Elapsed)

	return result, nil
}
