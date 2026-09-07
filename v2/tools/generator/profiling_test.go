/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package main

import (
	"bytes"
	"os"
	"path/filepath"
	"runtime"
	"runtime/debug"
	"strings"
	"testing"
	"time"

	. "github.com/onsi/gomega"

	pprofprofile "github.com/google/pprof/profile"
)

//nolint:paralleltest // Go permits only one process-wide CPU profile at a time.
func TestProfiler_CPUOnlyLifecycle(t *testing.T) {
	g := NewGomegaWithT(t)

	tempDir := t.TempDir()
	cpuPath := filepath.Join(tempDir, "cpu.pprof")
	profiler := newProfiler()
	profiler.cpuProfilePath = cpuPath

	g.Expect(profiler.start()).To(Succeed())

	burnCPU(200 * time.Millisecond)

	g.Expect(profiler.stop()).To(Succeed())
	g.Expect(readProfile(g, cpuPath).Sample).NotTo(BeEmpty())
}

//nolint:paralleltest // Keep profiler lifecycle tests serialized for consistency.
func TestProfiler_MemoryOnlyLifecycle(t *testing.T) {
	g := NewGomegaWithT(t)

	tempDir := t.TempDir()
	memoryPath := filepath.Join(tempDir, "memory.pprof")
	profiler := newProfiler()
	profiler.memoryProfilePath = memoryPath

	g.Expect(profiler.start()).To(Succeed())

	var allocations [][]byte
	deadline := time.Now().Add(200 * time.Millisecond)
	for time.Now().Before(deadline) {
		allocations = append(allocations, allocateProfileBytes(1024))
		if len(allocations) > 1_000 {
			allocations = allocations[:0]
		}
	}

	g.Expect(profiler.stop()).To(Succeed())

	memoryProfile := readProfile(g, memoryPath)
	g.Expect(memoryProfile.SampleType).To(ContainElements(
		And(HaveField("Type", "alloc_objects"), HaveField("Unit", "count")),
		And(HaveField("Type", "alloc_space"), HaveField("Unit", "bytes")),
	))
	g.Expect(memoryProfile.Sample).NotTo(BeEmpty())
}

//nolint:paralleltest // Go permits only one process-wide CPU profile at a time.
func TestProfiler_CreatesProfilesConcurrently(t *testing.T) {
	g := NewGomegaWithT(t)

	tempDir := t.TempDir()
	cpuPath := filepath.Join(tempDir, "cpu.pprof")
	memoryPath := filepath.Join(tempDir, "memory.pprof")
	profiler := newProfiler()
	profiler.cpuProfilePath = cpuPath
	profiler.memoryProfilePath = memoryPath

	g.Expect(profiler.start()).To(Succeed())

	var allocations [][]byte
	deadline := time.Now().Add(200 * time.Millisecond)
	for time.Now().Before(deadline) {
		allocations = append(allocations, allocateProfileBytes(1024))
		burnCPU(1 * time.Millisecond)
		if len(allocations) > 1_000 {
			allocations = allocations[:0]
		}
	}

	g.Expect(profiler.stop()).To(Succeed())
	g.Expect(readProfile(g, cpuPath).Sample).NotTo(BeEmpty())

	memoryProfile := readProfile(g, memoryPath)
	g.Expect(memoryProfile.SampleType).To(ContainElements(
		And(HaveField("Type", "alloc_objects"), HaveField("Unit", "count")),
		And(HaveField("Type", "alloc_space"), HaveField("Unit", "bytes")),
	))
	g.Expect(memoryProfile.Sample).NotTo(BeEmpty())
}

//nolint:paralleltest // Keep profiler lifecycle tests serialized for consistency.
func TestProfiler_AllocationsAtEndOfWorkAreRepresented(t *testing.T) {
	g := NewGomegaWithT(t)

	originalRate := runtime.MemProfileRate
	runtime.MemProfileRate = 1
	defer func() {
		runtime.MemProfileRate = originalRate
	}()

	originalGCPercent := debug.SetGCPercent(-1)
	defer debug.SetGCPercent(originalGCPercent)

	runtime.GC()

	tempDir := t.TempDir()
	memoryPath := filepath.Join(tempDir, "memory.pprof")
	profiler := newProfiler()
	profiler.memoryProfilePath = memoryPath

	g.Expect(profiler.start()).To(Succeed())

	produceEndOfWorkAllocation()

	g.Expect(profiler.stop()).To(Succeed())

	memoryProfile := readProfile(g, memoryPath)
	g.Expect(profileContainsNonZeroSample(memoryProfile, "produceEndOfWorkAllocation")).To(BeTrue())
}

func readProfile(g Gomega, path string) *pprofprofile.Profile {
	data, err := os.ReadFile(path)
	g.Expect(err).NotTo(HaveOccurred())
	g.Expect(data).NotTo(BeEmpty())

	result, err := pprofprofile.Parse(bytes.NewReader(data))
	g.Expect(err).NotTo(HaveOccurred())
	return result
}

func profileContainsNonZeroSample(profile *pprofprofile.Profile, functionName string) bool {
	for _, sample := range profile.Sample {
		if !profileContainsFunctionInSample(sample, functionName) {
			continue
		}

		for _, value := range sample.Value {
			if value > 0 {
				return true
			}
		}
	}

	return false
}

func profileContainsFunctionInSample(sample *pprofprofile.Sample, functionName string) bool {
	for _, location := range sample.Location {
		for _, line := range location.Line {
			if line.Function != nil && strings.Contains(line.Function.Name, functionName) {
				return true
			}
		}
	}

	return false
}

//go:noinline
func allocateProfileBytes(size int) []byte {
	return make([]byte, size)
}

//go:noinline
func burnCPU(duration time.Duration) {
	deadline := time.Now().Add(duration)
	var accumulator uint64
	for time.Now().Before(deadline) {
		accumulator += 1
		accumulator ^= accumulator << 1
	}

	runtime.KeepAlive(accumulator)
}

//go:noinline
func produceEndOfWorkAllocation() {
	allocation := allocateProfileBytes(32 * 1024)
	runtime.KeepAlive(allocation)
}
