/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package perf_test

import (
	"fmt"
	"os"
	"testing"
	"time"

	. "github.com/onsi/gomega"
)

func Test_Perf_Static_VirtualNetworks(t *testing.T) {
	t.Parallel()
	tc := globalTestContext.ForTest(t)

	cfg := Config{
		Pattern:         PatternStatic,
		ResourceFactory: VirtualNetworkFactory(),
		StaticConfig: &StaticConfig{
			Duration:         5 * time.Minute,
			ResourceSetCount: 500, // This is actually 2 (small) resources, so total resource count = 1000
		},
	}
	staticConfigFromEnv(cfg.StaticConfig)

	result, err := RunPerfTest(t, tc, cfg)
	tc.G.Expect(err).ToNot(HaveOccurred())
	t.Logf("Result: %d sets created, %d sets deleted, elapsed %s",
		result.TotalSetsCreated,
		result.TotalSetsDeleted,
		result.Elapsed)
}

func Test_Perf_Dynamic_ResourceGroups(t *testing.T) {
	t.Parallel()
	tc := globalTestContext.ForTest(t)

	cfg := Config{
		Pattern:         PatternDynamic,
		ResourceFactory: ResourceGroupFactory(),
		DynamicConfig: &DynamicConfig{
			Duration:           5 * time.Minute,
			CreationRatePerMin: 60,
			MaxConcurrentSets:  200,
		},
	}
	dynamicConfigFromEnv(cfg.DynamicConfig)

	result, err := RunPerfTest(t, tc, cfg)
	tc.G.Expect(err).ToNot(HaveOccurred())
	t.Logf("Result: %d sets created, %d sets deleted, elapsed %s",
		result.TotalSetsCreated,
		result.TotalSetsDeleted,
		result.Elapsed)
}

func Test_Perf_Dynamic_VirtualNetworks(t *testing.T) {
	t.Parallel()
	tc := globalTestContext.ForTest(t)

	cfg := Config{
		Pattern:         PatternDynamic,
		ResourceFactory: VirtualNetworkFactory(),
		DynamicConfig: &DynamicConfig{
			Duration:           5 * time.Minute,
			CreationRatePerMin: 60,
			MaxConcurrentSets:  200,
		},
	}
	dynamicConfigFromEnv(cfg.DynamicConfig)

	result, err := RunPerfTest(t, tc, cfg)
	tc.G.Expect(err).ToNot(HaveOccurred())
	t.Logf("Result: %d sets created, %d sets deleted, elapsed %s",
		result.TotalSetsCreated,
		result.TotalSetsDeleted,
		result.Elapsed)
}

// staticConfigFromEnv overrides StaticConfig fields from environment variables, if set.
// Supported env vars:
//
//	PERF_STATIC_RESOURCE_COUNT - number of resource sets for static pattern
//	PERF_DURATION              - hold duration for static pattern (e.g. "1m", "30s")
func staticConfigFromEnv(cfg *StaticConfig) {
	if v := getEnvInt("PERF_STATIC_RESOURCE_COUNT"); v > 0 {
		cfg.ResourceSetCount = v
	}
	if v := getEnvDuration("PERF_DURATION"); v > 0 {
		cfg.Duration = v
	}
}

// dynamicConfigFromEnv overrides DynamicConfig fields from environment variables, if set.
// Supported env vars:
//
//	PERF_DURATION               - total duration for dynamic pattern
//	PERF_DYNAMIC_RATE           - creation rate (sets/min) for dynamic pattern
//	PERF_DYNAMIC_MAX_CONCURRENT - max concurrent sets for dynamic pattern
func dynamicConfigFromEnv(cfg *DynamicConfig) {
	if v := getEnvDuration("PERF_DURATION"); v > 0 {
		cfg.Duration = v
	}
	if v := getEnvFloat("PERF_DYNAMIC_RATE"); v > 0 {
		cfg.CreationRatePerMin = v
	}
	if v := getEnvInt("PERF_DYNAMIC_MAX_CONCURRENT"); v > 0 {
		cfg.MaxConcurrentSets = v
	}
}

// TODO: Use helpers from v2/internal/config/vars.go ?
func getEnvInt(key string) int {
	v, ok := os.LookupEnv(key)
	if !ok {
		return 0
	}
	var result int
	_, err := fmt.Sscanf(v, "%d", &result)
	if err != nil {
		return 0
	}
	return result
}

func getEnvFloat(key string) float64 {
	v, ok := os.LookupEnv(key)
	if !ok {
		return 0
	}
	var result float64
	_, err := fmt.Sscanf(v, "%f", &result)
	if err != nil {
		return 0
	}
	return result
}

func getEnvDuration(key string) time.Duration {
	v, ok := os.LookupEnv(key)
	if !ok {
		return 0
	}
	d, err := time.ParseDuration(v)
	if err != nil {
		return 0
	}
	return d
}
