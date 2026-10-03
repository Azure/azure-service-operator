/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package main

import (
	"context"
	"errors"
	"os"

	"github.com/go-logr/logr"
	"github.com/go-logr/zerologr"
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"

	"github.com/Azure/azure-service-operator/v2/pkg/xcontext"
)

// Execute kicks off the command line
func Execute() {
	cmd, profiler, err := newRootCommand()
	if err != nil {
		log := CreateLogger()
		log.Error(err, "failed to create root command")
		return
	}

	ctx := xcontext.MakeInterruptibleContext(context.Background())
	if err = executeRootCommand(ctx, os.Args[1:], cmd, profiler); err != nil {
		log := CreateLogger()
		log.Error(err, "failed to execute root command")
		os.Exit(1)
	}
}

func newRootCommand() (*cobra.Command, *profiler, error) {
	profiler := newProfiler()
	rootCmd := &cobra.Command{
		Use:              "aso-gen",
		Short:            "aso-gen provides a cmdline interface for generating Azure Service Operator types from Azure deployment template schema",
		TraverseChildren: true,
		SilenceErrors:    true, // We show errors ourselves using our logger
		SilenceUsage:     true, // Let users ask for usage themselves
	}

	rootCmd.Flags().SortFlags = false

	rootCmd.PersistentFlags().BoolVar(&verbose, "verbose", false, "Enable verbose logging")
	rootCmd.PersistentFlags().BoolVar(&quiet, "quiet", false, "Suppress non-error logging")
	rootCmd.PersistentFlags().BoolVar(&trace, "trace", false, "Enable trace logging (very verbose)")
	rootCmd.PersistentFlags().StringVar(
		&profiler.memoryProfilePath,
		"memory-prof",
		"",
		"Write a cumulative memory allocation profile to this file",
	)
	rootCmd.PersistentFlags().StringVar(
		&profiler.cpuProfilePath,
		"cpu-prof",
		"",
		"Write a CPU profile to this file",
	)

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
		if err := cmd.ValidateRequiredFlags(); err != nil {
			return err
		}

		if err := cmd.ValidateFlagGroups(); err != nil {
			return err
		}

		// Configure logging; --trace overrides --verbose overrides --quiet
		if trace {
			zerologr.SetMaxV(2)
		} else if verbose {
			zerologr.SetMaxV(1)
		} else if quiet {
			// Can't use zerologr.SetMaxV(-1)
			zerolog.SetGlobalLevel(zerolog.ErrorLevel)
		} else {
			zerologr.SetMaxV(0)
		}

		return profiler.start()
	}

	return rootCmd, profiler, nil
}

type profileStopper interface {
	stop() error
}

func executeRootCommand(
	ctx context.Context,
	args []string,
	cmd *cobra.Command,
	profiler profileStopper,
) error {
	cmd.SetArgs(args)
	return executeCommand(ctx, cmd, profiler)
}

func executeCommand(
	ctx context.Context,
	cmd *cobra.Command,
	profiler profileStopper,
) (result error) {
	defer func() {
		result = errors.Join(result, profiler.stop())
	}()

	result = cmd.ExecuteContext(ctx)
	return result
}

var (
	quiet   bool
	trace   bool
	verbose bool
)

// CreateLogger creates a logger  for console output.
func CreateLogger() logr.Logger {
	// Configure console writer for ZeroLog
	output := zerolog.ConsoleWriter{
		Out:        os.Stderr,      // Write to StdErr
		TimeFormat: "15:04:05.999", // Display time to the millisecond
	}

	// Create zerolog logger
	zl := zerolog.New(output).
		With().Timestamp().
		Logger()

	// Use standard interface for logging
	zerologr.VerbosityFieldName = "" // Don't include verbosity in output

	log := zerologr.New(&zl)
	return log
}
