/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package cmd

import (
	"context"

	"github.com/Azure/azure-service-operator/v2/internal/version"
	"github.com/Azure/azure-service-operator/v2/pkg/xcontext"
	"github.com/go-logr/zerologr"
	"github.com/spf13/cobra"
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
	}

	rootCmd.PersistentFlags().BoolVar(&verbose, "verbose", false, "Enable verbose logging")

	rootCmd.Flags().SortFlags = false

	cmds := []func() (*cobra.Command, error){
		newCleanCommand,
		newImportCommand,
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
		if verbose {
			zerologr.SetMaxV(1)
		}
	}

	return rootCmd, nil
}
