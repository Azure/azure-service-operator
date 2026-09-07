# ASO Generator Profiling Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Add opt-in CPU and cumulative allocation profiles to every ASO generator subcommand through `--cpu-prof <file>` and `--memory-prof <file>`.

**Architecture:** A focused profiler lifecycle owns profile files and calls `runtime/pprof`. Cobra persistent flags configure that lifecycle, a persistent pre-run hook starts it immediately before command work, and the command executor always finalises it after `ExecuteContext` returns, including error paths.

**Tech Stack:** Go 1.26, Cobra, `runtime/pprof`, `errors.Join`, Gomega, `github.com/google/pprof/profile`

## Global Constraints

- Both options must work alone or concurrently.
- `--memory-prof` must write the cumulative `allocs` profile, not a live-heap snapshot.
- Profiles must cover the complete execution of the selected command.
- Profile finalisation must run after both successful and failed commands.
- Generated profile files must be readable by standard pprof tooling.
- Do not edit generated files.

---

### Task 1: Implement the profile lifecycle

**Files:**
- Create: `v2/tools/generator/profiling.go`
- Create: `v2/tools/generator/profiling_test.go`
- Modify: `v2/tools/generator/go.mod`
- Modify: `v2/tools/generator/go.sum`

**Interfaces:**
- Consumes: Go's `os.Create()`, `runtime/pprof.StartCPUProfile()`, `runtime/pprof.StopCPUProfile()`, and `runtime/pprof.Lookup("allocs")`.
- Produces: `newProfiler() *profiler`, `(*profiler).start() error`, and `(*profiler).stop() error`. The `profiler` exposes package-private `cpuProfilePath` and `memoryProfilePath` fields for Cobra flag binding in Task 2.

- [ ] **Step 1: Add the pprof parser used to validate binary profiles**

Run:

```bash
cd v2/tools/generator
go get github.com/google/pprof/profile@v0.0.0-20250403155104-27863c87afa6
```

Expected: `go.mod` gains a direct `github.com/google/pprof` requirement and `go.sum` gains the module checksums.

- [ ] **Step 2: Write failing lifecycle tests**

Create `v2/tools/generator/profiling_test.go`:

```go
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

	pprofprofile "github.com/google/pprof/profile"
	. "github.com/onsi/gomega"
)

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

func TestProfiler_NoPathsIsANoOp(t *testing.T) {
	g := NewGomegaWithT(t)

	profiler := newProfiler()

	g.Expect(profiler.start()).To(Succeed())
	g.Expect(profiler.stop()).To(Succeed())
}

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
```

Do not mark CPU-profiling tests parallel because Go permits only one process-wide CPU profile at a time.

- [ ] **Step 3: Run the focused tests to verify they fail**

Run:

```bash
cd v2/tools/generator
go test . -run '^TestProfiler_' -count=1
```

Expected: FAIL because `newProfiler` is undefined.

- [ ] **Step 4: Implement the profiler lifecycle**

Create `v2/tools/generator/profiling.go`:

```go
/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package main

import (
	"errors"
	"fmt"
	"os"
	"runtime/pprof"
)

type profiler struct {
	cpuProfilePath    string
	memoryProfilePath string
	cpuProfile        *os.File
	memoryProfile     *os.File
}

func newProfiler() *profiler {
	return &profiler{}
}

func (p *profiler) start() error {
	var err error

	if p.memoryProfilePath != "" {
		p.memoryProfile, err = os.Create(p.memoryProfilePath)
		if err != nil {
			return fmt.Errorf("creating memory profile %q: %w", p.memoryProfilePath, err)
		}
	}

	if p.cpuProfilePath != "" {
		p.cpuProfile, err = os.Create(p.cpuProfilePath)
		if err != nil {
			return errors.Join(
				fmt.Errorf("creating CPU profile %q: %w", p.cpuProfilePath, err),
				p.closeMemoryProfile())
		}

		if err = pprof.StartCPUProfile(p.cpuProfile); err != nil {
			return errors.Join(
				fmt.Errorf("starting CPU profile %q: %w", p.cpuProfilePath, err),
				p.closeCPUProfile(),
				p.closeMemoryProfile())
		}
	}

	return nil
}

func (p *profiler) stop() error {
	var result error

	if p.cpuProfile != nil {
		pprof.StopCPUProfile()
		result = errors.Join(result, p.closeCPUProfile())
	}

	if p.memoryProfile != nil {
		allocations := pprof.Lookup("allocs")
		if allocations == nil {
			result = errors.Join(result, errors.New("finding cumulative allocation profile"))
		} else if err := allocations.WriteTo(p.memoryProfile, 0); err != nil {
			result = errors.Join(result, fmt.Errorf("writing memory profile %q: %w", p.memoryProfilePath, err))
		}

		result = errors.Join(result, p.closeMemoryProfile())
	}

	return result
}

func (p *profiler) closeCPUProfile() error {
	if p.cpuProfile == nil {
		return nil
	}

	err := p.cpuProfile.Close()
	p.cpuProfile = nil
	if err != nil {
		return fmt.Errorf("closing CPU profile %q: %w", p.cpuProfilePath, err)
	}

	return nil
}

func (p *profiler) closeMemoryProfile() error {
	if p.memoryProfile == nil {
		return nil
	}

	err := p.memoryProfile.Close()
	p.memoryProfile = nil
	if err != nil {
		return fmt.Errorf("closing memory profile %q: %w", p.memoryProfilePath, err)
	}

	return nil
}
```

- [ ] **Step 5: Run the focused tests**

Run:

```bash
cd v2/tools/generator
go test . -run '^TestProfiler_' -count=1
```

Expected: PASS, with both output files parsed successfully.

- [ ] **Step 6: Commit the lifecycle**

```bash
git add v2/tools/generator/profiling.go v2/tools/generator/profiling_test.go v2/tools/generator/go.mod v2/tools/generator/go.sum
git commit -m "feat: add generator profiling lifecycle" -m "Co-authored-by: Copilot <223556219+Copilot@users.noreply.github.com>"
```

---

### Task 2: Wire profiling through the root command

**Files:**
- Modify: `v2/tools/generator/root.go`
- Create: `v2/tools/generator/root_test.go`

**Interfaces:**
- Consumes: `newProfiler() *profiler`, `(*profiler).start() error`, and `(*profiler).stop() error` from Task 1.
- Produces: `newRootCommand() (*cobra.Command, *profiler, error)`, a narrow `profileStopper` interface containing `stop() error`, and `executeCommand(context.Context, *cobra.Command, profileStopper) error`.

- [ ] **Step 1: Write failing root command tests**

Create `v2/tools/generator/root_test.go`:

```go
/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package main

import (
	"context"
	"errors"
	"path/filepath"
	"testing"

	"github.com/spf13/cobra"
	. "github.com/onsi/gomega"
)

func TestRootCommand_ExposesProfilingFlags(t *testing.T) {
	g := NewGomegaWithT(t)

	cmd, _, err := newRootCommand()
	g.Expect(err).NotTo(HaveOccurred())
	g.Expect(cmd.PersistentFlags().Lookup("cpu-prof")).NotTo(BeNil())
	g.Expect(cmd.PersistentFlags().Lookup("memory-prof")).NotTo(BeNil())
}

func TestRootCommand_FinalizesProfilesAfterCommandFailure(t *testing.T) {
	g := NewGomegaWithT(t)

	memoryPath := filepath.Join(t.TempDir(), "memory.pprof")
	cmd, profiler, err := newRootCommand()
	g.Expect(err).NotTo(HaveOccurred())

	commandErr := errors.New("command failed")
	cmd.AddCommand(&cobra.Command{
		Use: "failing-command",
		RunE: func(*cobra.Command, []string) error {
			_ = make([]byte, 1024*1024)
			return commandErr
		},
	})
	cmd.SetArgs([]string{"failing-command", "--memory-prof", memoryPath})

	err = executeCommand(context.Background(), cmd, profiler)
	g.Expect(err).To(MatchError(ContainSubstring(commandErr.Error())))
	g.Expect(readProfile(g, memoryPath)).NotTo(BeNil())
}

func TestRootCommand_InvalidProfilePathPreventsExecution(t *testing.T) {
	g := NewGomegaWithT(t)

	cmd, profiler, err := newRootCommand()
	g.Expect(err).NotTo(HaveOccurred())

	executed := false
	cmd.AddCommand(&cobra.Command{
		Use: "successful-command",
		Run: func(*cobra.Command, []string) {
			executed = true
		},
	})
	memoryPath := filepath.Join(t.TempDir(), "missing", "memory.pprof")
	cmd.SetArgs([]string{"successful-command", "--memory-prof", memoryPath})

	err = executeCommand(context.Background(), cmd, profiler)
	g.Expect(err).To(MatchError(ContainSubstring("creating memory profile")))
	g.Expect(executed).To(BeFalse())
}

func TestExecuteCommand_PreservesCommandAndProfileErrors(t *testing.T) {
	g := NewGomegaWithT(t)

	commandErr := errors.New("command failed")
	profileErr := errors.New("profile failed")
	cmd := &cobra.Command{
		RunE: func(*cobra.Command, []string) error {
			return commandErr
		},
	}

	err := executeCommand(
		context.Background(),
		cmd,
		&fakeProfileStopper{err: profileErr})

	g.Expect(errors.Is(err, commandErr)).To(BeTrue())
	g.Expect(errors.Is(err, profileErr)).To(BeTrue())
}

type fakeProfileStopper struct {
	err error
}

func (f *fakeProfileStopper) stop() error {
	return f.err
}
```

- [ ] **Step 2: Run the focused tests to verify they fail**

Run:

```bash
cd v2/tools/generator
go test . -run '^TestRootCommand_' -count=1
```

Expected: FAIL because `newRootCommand` has the old return signature and `executeCommand` is undefined.

- [ ] **Step 3: Bind persistent flags and start profiling in the root pre-run**

Update `newRootCommand()` in `v2/tools/generator/root.go` to return the profiler alongside the command:

```go
func newRootCommand() (*cobra.Command, *profiler, error) {
	profiler := newProfiler()
	rootCmd := &cobra.Command{
		Use:              "aso-gen",
		Short:            "aso-gen provides a cmdline interface for generating Azure Service Operator types from Azure deployment template schema",
		TraverseChildren: true,
		SilenceErrors:    true,
		SilenceUsage:     true,
	}

	rootCmd.Flags().SortFlags = false

	rootCmd.PersistentFlags().BoolVar(&verbose, "verbose", false, "Enable verbose logging")
	rootCmd.PersistentFlags().BoolVar(&quiet, "quiet", false, "Suppress non-error logging")
	rootCmd.PersistentFlags().BoolVar(&trace, "trace", false, "Enable trace logging (very verbose)")
	rootCmd.PersistentFlags().StringVar(
		&profiler.memoryProfilePath,
		"memory-prof",
		"",
		"Write a cumulative memory allocation profile to this file")
	rootCmd.PersistentFlags().StringVar(
		&profiler.cpuProfilePath,
		"cpu-prof",
		"",
		"Write a CPU profile to this file")

	rootCmd.MarkFlagsMutuallyExclusive("verbose", "quiet", "trace")

	cmdFuncs := []func() (*cobra.Command, error){
		NewGenTypesCommand,
		NewGenKustomizeCommand,
	}

	for _, f := range cmdFuncs {
		cmd, err := f()
		if err != nil {
			return rootCmd, profiler, err
		}
		rootCmd.AddCommand(cmd)
	}

	rootCmd.PersistentPreRunE = func(cmd *cobra.Command, args []string) error {
		if trace {
			zerologr.SetMaxV(2)
		} else if verbose {
			zerologr.SetMaxV(1)
		} else if quiet {
			zerolog.SetGlobalLevel(zerolog.ErrorLevel)
		} else {
			zerologr.SetMaxV(0)
		}

		return profiler.start()
	}

	return rootCmd, profiler, nil
}
```

Keep the existing logging behaviour unchanged apart from changing `PersistentPreRun` to `PersistentPreRunE`.

- [ ] **Step 4: Always stop profiling after command execution**

Add `errors` to the imports in `v2/tools/generator/root.go`, define a narrow finalisation interface, then split command execution into a testable helper:

```go
func Execute() {
	cmd, profiler, err := newRootCommand()
	if err != nil {
		log := CreateLogger()
		log.Error(err, "failed to create root command")
		return
	}

	ctx := xcontext.MakeInterruptibleContext(context.Background())
	if err = executeCommand(ctx, cmd, profiler); err != nil {
		log := CreateLogger()
		log.Error(err, "failed to execute root command")
		os.Exit(1)
	}
}

type profileStopper interface {
	stop() error
}

func executeCommand(
	ctx context.Context,
	cmd *cobra.Command,
	profiler profileStopper,
) error {
	executeErr := cmd.ExecuteContext(ctx)
	profileErr := profiler.stop()
	return errors.Join(executeErr, profileErr)
}
```

This must call `stop()` even when Cobra returns an argument, startup, or command error. Calling `stop()` after a failed `start()` is safe because Task 1 clears every file it closes.

- [ ] **Step 5: Run the focused root and profiler tests**

Run:

```bash
cd v2/tools/generator
go test . -run '^(TestProfiler_|TestRootCommand_)' -count=1
```

Expected: PASS.

- [ ] **Step 6: Format and run all generator checks**

Run from the repository root:

```bash
./hack/tools/task generator:format-code
./hack/tools/task generator:quick-checks
```

Expected: both commands exit successfully; generator unit tests, basic checks, and lint pass.

- [ ] **Step 7: Manually verify both profiles in one generator invocation**

Run from the repository root:

```bash
mkdir -p /tmp/aso-gen-profile-check
cd v2
./bin/aso-gen gen-types azure-arm.yaml \
  --cpu-prof /tmp/aso-gen-profile-check/cpu.pprof \
  --memory-prof /tmp/aso-gen-profile-check/memory.pprof
go tool pprof -top /tmp/aso-gen-profile-check/cpu.pprof
go tool pprof -top -alloc_space /tmp/aso-gen-profile-check/memory.pprof
rm -rf /tmp/aso-gen-profile-check
```

Expected: generation succeeds and both `go tool pprof` commands print profile summaries without parse errors.

- [ ] **Step 8: Commit the CLI integration**

```bash
git add v2/tools/generator/root.go v2/tools/generator/root_test.go
git commit -m "feat: add generator profiling options" -m "Co-authored-by: Copilot <223556219+Copilot@users.noreply.github.com>"
```
