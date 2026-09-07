/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package main

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"
	"time"

	. "github.com/onsi/gomega"

	pprofprofile "github.com/google/pprof/profile"
)

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
		allocations = append(allocations, make([]byte, 1024))
		if len(allocations) > 1_000 {
			allocations = allocations[:0]
		}
	}

	g.Expect(profiler.stop()).To(Succeed())
	g.Expect(readProfile(g, cpuPath)).NotTo(BeNil())

	memoryProfile := readProfile(g, memoryPath)
	g.Expect(memoryProfile.SampleType).To(ContainElements(
		And(HaveField("Type", "alloc_objects"), HaveField("Unit", "count")),
		And(HaveField("Type", "alloc_space"), HaveField("Unit", "bytes")),
	))
}

//nolint:paralleltest // Keep profiler lifecycle tests serialized for consistency.
func TestProfiler_NoPathsIsANoOp(t *testing.T) {
	g := NewGomegaWithT(t)

	profiler := newProfiler()

	g.Expect(profiler.start()).To(Succeed())
	g.Expect(profiler.stop()).To(Succeed())
}

//nolint:paralleltest // Keep profiler lifecycle tests serialized for consistency.
func TestProfiler_InvalidPathFailsStart(t *testing.T) {
	g := NewGomegaWithT(t)

	profiler := newProfiler()
	profiler.memoryProfilePath = filepath.Join(t.TempDir(), "missing", "memory.pprof")

	g.Expect(profiler.start()).To(MatchError(ContainSubstring("creating memory profile")))
	g.Expect(profiler.stop()).To(Succeed())
}

func readProfile(g Gomega, path string) *pprofprofile.Profile {
	data, err := os.ReadFile(path)
	g.Expect(err).NotTo(HaveOccurred())
	g.Expect(data).NotTo(BeEmpty())

	result, err := pprofprofile.Parse(bytes.NewReader(data))
	g.Expect(err).NotTo(HaveOccurred())
	return result
}
