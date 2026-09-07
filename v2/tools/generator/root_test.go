/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package main

import (
	"context"
	"errors"
	"os"
	"path/filepath"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/spf13/cobra"
)

//nolint:paralleltest // newRootCommand binds package-global logging flags.
func TestRootCommand_ExposesProfilingFlags(t *testing.T) {
	g := NewGomegaWithT(t)

	cmd, _, err := newRootCommand()
	g.Expect(err).NotTo(HaveOccurred())
	g.Expect(cmd.PersistentFlags().Lookup("cpu-prof")).NotTo(BeNil())
	g.Expect(cmd.PersistentFlags().Lookup("memory-prof")).NotTo(BeNil())
}

//nolint:paralleltest // newRootCommand binds package-global logging flags.
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

//nolint:paralleltest // newRootCommand binds package-global logging flags.
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

//nolint:paralleltest // newRootCommand binds package-global logging flags.
func TestExecuteRootCommand_InvalidFlagGroupsDoNotCreateProfiles(t *testing.T) {
	g := NewGomegaWithT(t)

	memoryPath := filepath.Join(t.TempDir(), "memory.pprof")
	cmd, profiler, err := newRootCommand()
	g.Expect(err).NotTo(HaveOccurred())

	executed := false
	cmd.AddCommand(&cobra.Command{
		Use: "successful-command",
		Run: func(*cobra.Command, []string) {
			executed = true
		},
	})

	err = executeRootCommand(
		context.Background(),
		[]string{"--verbose", "--quiet", "--memory-prof", memoryPath, "successful-command"},
		cmd,
		profiler,
	)
	g.Expect(err).To(MatchError(ContainSubstring("none of the others can be")))
	g.Expect(executed).To(BeFalse())

	_, statErr := os.Stat(memoryPath)
	g.Expect(os.IsNotExist(statErr)).To(BeTrue())
}

//nolint:paralleltest // newRootCommand binds package-global logging flags.
func TestExecuteRootCommand_AcceptsRootProfilingFlagsBeforeSubcommand(t *testing.T) {
	g := NewGomegaWithT(t)

	memoryPath := filepath.Join(t.TempDir(), "memory.pprof")
	cmd, profiler, err := newRootCommand()
	g.Expect(err).NotTo(HaveOccurred())

	executed := false
	cmd.AddCommand(&cobra.Command{
		Use: "successful-command",
		Run: func(*cobra.Command, []string) {
			executed = true
			_ = make([]byte, 1024*1024)
		},
	})

	err = executeRootCommand(
		context.Background(),
		[]string{"--memory-prof", memoryPath, "successful-command"},
		cmd,
		profiler,
	)
	g.Expect(err).NotTo(HaveOccurred())
	g.Expect(executed).To(BeTrue())
	g.Expect(readProfile(g, memoryPath)).NotTo(BeNil())
}

//nolint:paralleltest // newRootCommand binds package-global logging flags.
func TestExecuteRootCommand_InvalidArgumentsDoNotCreateProfiles(t *testing.T) {
	g := NewGomegaWithT(t)

	memoryPath := filepath.Join(t.TempDir(), "memory.pprof")
	cmd, profiler, err := newRootCommand()
	g.Expect(err).NotTo(HaveOccurred())

	err = executeRootCommand(
		context.Background(),
		[]string{"--memory-prof", memoryPath, "gen-types"},
		cmd,
		profiler,
	)
	g.Expect(err).To(MatchError(ContainSubstring("accepts 1 arg(s), received 0")))

	_, statErr := os.Stat(memoryPath)
	g.Expect(os.IsNotExist(statErr)).To(BeTrue())
}

//nolint:paralleltest // newRootCommand binds package-global logging flags.
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
		&fakeProfileStopper{err: profileErr},
	)

	g.Expect(errors.Is(err, commandErr)).To(BeTrue())
	g.Expect(errors.Is(err, profileErr)).To(BeTrue())
}

type fakeProfileStopper struct {
	err error
}

func (f *fakeProfileStopper) stop() error {
	return f.err
}
