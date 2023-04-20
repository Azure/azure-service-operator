/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package main

import (
	"context"
	"os"

	"github.com/Azure/azure-service-operator/v2/internal/version"
	"github.com/Azure/azure-service-operator/v2/pkg/xcontext"
	"github.com/go-logr/logr"
	"github.com/go-logr/zerologr"
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
)

// Execute kicks off the command line
func Execute() {
	cmd, err := newRootCommand()
	if err != nil {
		log := CreateLogger()
		log.Error(err, "failed to create root command")
		return
	}

	ctx := xcontext.MakeInterruptibleContext(context.Background())
	if err := cmd.ExecuteContext(ctx); err != nil {
		log := CreateLogger()
		log.Error(err, "failed to execute root command")
	}
}

func newRootCommand() (*cobra.Command, error) {
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

	cmdFuncs := []func() (*cobra.Command, error){
		NewGenTypesCommand,
		NewGenKustomizeCommand,
		version.NewCommand,
	}

	for _, f := range cmdFuncs {
		cmd, err := f()
		if err != nil {
			return rootCmd, err
		}
		rootCmd.AddCommand(cmd)
	}

	rootCmd.PersistentPreRun = func(cmd *cobra.Command, args []string) {
		// Configure logging; --verbose overrides --quiet
		if verbose {
			zerologr.SetMaxV(1)
		} else if quiet {
			// Can't use zerologr.SetMaxV(-1)
			zerolog.SetGlobalLevel(zerolog.ErrorLevel)
		} else {
			zerologr.SetMaxV(0)
		}
	}

	return rootCmd, nil
}

var (
	verbose bool
	quiet   bool
)

// CreateLogger creates a logger  for console output.
// Use this when your command wants to show only log messages
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
