/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package testcommon

import (
	"context"
	"errors"
	"reflect"
	"sync"
	"sync/atomic"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/go-logr/logr"
	"golang.org/x/sync/semaphore"

	"github.com/Azure/azure-service-operator/v2/internal/config"
)

func Test_CfgToKey_HasAllConfigDotValuesKeys(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	// This makes sure that if config.Values or testConfig changes
	// then cfgToKey is updated to match.
	key := cfgToKey(testConfig{})

	testConfigType := reflect.TypeOf(testConfig{})

	for i, field := range reflect.VisibleFields(testConfigType) {
		// Skip the embedded struct and TerminateWhenDone field
		if (i == 0 && field.Name == "Values") || field.Name == "CountsTowardsLimit" {
			continue
		}
		g.Expect(key).To(ContainSubstring(field.Name + ":"))
	}
}

func Test_GetEnvTestForConfig_ConcurrentCallersShareEnvironment(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	const callers = 8
	ready := make(chan struct{}, callers)
	release := make(chan struct{})
	var creations atomic.Int32
	var stops atomic.Int32
	set := sharedEnvTests{
		concurrencyLimitSemaphore: semaphore.NewWeighted(1),
		envtests:                  make(map[string]*runningEnvTest),
		envtestsBeingCreated:      make(map[string]*envTestCreation),
		afterInitialLookup: func() {
			ready <- struct{}{}
			<-release
		},
		createEnvTest: func(_ context.Context, cfg testConfig, _ *namespaceResources) (*runningEnvTest, error) {
			creations.Add(1)
			return &runningEnvTest{
				Cfg:     cfg,
				Callers: 1,
				Stop: func() {
					stops.Add(1)
				},
			}, nil
		},
	}
	cfg := testConfig{CountsTowardsLimit: true}

	results := make([]*runningEnvTest, callers)
	errs := make([]error, callers)
	var wg sync.WaitGroup
	wg.Add(callers)
	for i := range callers {
		go func() {
			defer wg.Done()
			results[i], errs[i] = set.getEnvTestForConfig(t.Context(), cfg, logr.Discard())
		}()
	}

	for range callers {
		<-ready
	}
	close(release)
	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()
	g.Eventually(done).Should(BeClosed())

	g.Expect(errs).To(ConsistOf(make([]error, callers)))
	g.Expect(creations.Load()).To(Equal(int32(1)))
	for _, result := range results {
		g.Expect(result).To(BeIdenticalTo(results[0]))
	}
	g.Expect(results[0].Callers).To(Equal(callers))
	g.Expect(set.concurrencyLimitSemaphore.TryAcquire(1)).To(BeFalse())

	set.stopAll()
	g.Expect(stops.Load()).To(Equal(int32(1)))
	g.Expect(set.concurrencyLimitSemaphore.TryAcquire(1)).To(BeTrue())
}

func Test_GetEnvTestForConfig_CreationFailureReleasesPermit(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	limit := semaphore.NewWeighted(1)
	set := sharedEnvTests{
		concurrencyLimitSemaphore: limit,
		envtests:                  make(map[string]*runningEnvTest),
		envtestsBeingCreated:      make(map[string]*envTestCreation),
		createEnvTest: func(_ context.Context, _ testConfig, _ *namespaceResources) (*runningEnvTest, error) {
			return nil, errors.New("expected error")
		},
	}

	_, err := set.getEnvTestForConfig(
		t.Context(),
		testConfig{CountsTowardsLimit: true},
		logr.Discard(),
	)

	g.Expect(err).To(MatchError(ContainSubstring("expected error")))
	g.Expect(limit.TryAcquire(1)).To(BeTrue())
}

func Test_GetEnvTestForConfig_LimitsDistinctEnvironments(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	const limit = 2
	var live atomic.Int32
	var maximum atomic.Int32
	set := sharedEnvTests{
		concurrencyLimitSemaphore: semaphore.NewWeighted(limit),
		envtests:                  make(map[string]*runningEnvTest),
		envtestsBeingCreated:      make(map[string]*envTestCreation),
		createEnvTest: func(_ context.Context, cfg testConfig, _ *namespaceResources) (*runningEnvTest, error) {
			current := live.Add(1)
			for current > maximum.Load() && !maximum.CompareAndSwap(maximum.Load(), current) {
			}
			return &runningEnvTest{
				Cfg:     cfg,
				Callers: 1,
				Stop: func() {
					live.Add(-1)
				},
			}, nil
		},
	}

	defaultCfg := testConfig{}
	_, err := set.getEnvTestForConfig(t.Context(), defaultCfg, logr.Discard())
	g.Expect(err).NotTo(HaveOccurred())

	type result struct {
		cfg testConfig
		err error
	}
	results := make(chan result, 3)
	for i := range 3 {
		cfg := testConfig{
			Values: config.Values{
				PodNamespace: string(rune('a' + i)),
			},
			CountsTowardsLimit: true,
		}
		go func() {
			_, err := set.getEnvTestForConfig(t.Context(), cfg, logr.Discard())
			results <- result{cfg: cfg, err: err}
		}()
	}

	first := <-results
	second := <-results
	g.Expect(first.err).NotTo(HaveOccurred())
	g.Expect(second.err).NotTo(HaveOccurred())
	g.Expect(live.Load()).To(Equal(int32(1 + limit)))
	g.Expect(maximum.Load()).To(Equal(int32(1 + limit)))
	g.Consistently(results).ShouldNot(Receive())

	set.garbageCollect(first.cfg, logr.Discard())
	third := <-results
	g.Expect(third.err).NotTo(HaveOccurred())
	g.Expect(maximum.Load()).To(Equal(int32(1 + limit)))

	set.stopAll()
	g.Expect(live.Load()).To(Equal(int32(0)))
}
