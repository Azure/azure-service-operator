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

	. "github.com/onsi/gomega"

	"github.com/spf13/cobra"
)

func TestRootCommand_ExposesProfilingFlags(t *testing.T) {
	t.Parallel()

	g := NewGomegaWithT(t)

	cmd, _, err := newRootCommand()
	g.Expect(err).NotTo(HaveOccurred())
	g.Expect(cmd.PersistentFlags().Lookup("cpu-prof")).NotTo(BeNil())
	g.Expect(cmd.PersistentFlags().Lookup("memory-prof")).NotTo(BeNil())
}

func TestRootCommand_FinalizesProfilesAfterCommandFailure(t *testing.T) {
	t.Parallel()

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
	t.Parallel()

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
	t.Parallel()

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
