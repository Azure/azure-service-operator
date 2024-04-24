/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package cmd

import (
	"context"

	"github.com/go-logr/zerologr"
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"

	"github.com/Azure/azure-service-operator/v2/internal/version"
	"github.com/Azure/azure-service-operator/v2/pkg/xcontext"
)

// Execute kicks off the command line
func Execute() {
	cmd, err := newRootCommand()
	if err != nil {
		log := CreateLogger()
		log.Error(err, "failed to build commands")
	}

	ctx := xcontext.MakeInterruptibleContext(context.Background())
	if err := cmd.ExecuteContext(ctx); err != nil {
		log := CreateLogger()
		log.Error(err, "failed to execute command")
		return
	}
}

func newRootCommand() (*cobra.Command, error) {
	rootCmd := &cobra.Command{
		Use:              "asoctl",
		Short:            "asoctl provides a cmdline interface for working with Azure Service Operator",
		TraverseChildren: true,
		SilenceErrors:    true, // We show errors ourselves using our logger
		SilenceUsage:     true, // Let users ask for usage themselves
	}

	rootCmd.PersistentFlags().BoolVar(&verbose, "verbose", false, "Enable verbose logging")
	rootCmd.PersistentFlags().BoolVar(&quiet, "quiet", false, "Silence most logging")

	rootCmd.Flags().SortFlags = false

	cmds := []func() (*cobra.Command, error){
		newCleanCommand,
		newImportCommand,
		newExportCommand,
		version.NewCommand,
	}

	for _, f := range cmds {
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
			if quiet {
				// Illegal combination, tell the user
				log := CreateLogger()
				log.Error(nil, "--quiet has no effect when --verbose is specified")
			}
		} else if quiet {
			// Can't use zerologr.SetMaxV(-1)
			zerolog.SetGlobalLevel(zerolog.ErrorLevel)
		} else {
			zerologr.SetMaxV(0)
		}
	}

	return rootCmd, nil
}
