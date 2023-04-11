/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package cmd

import (
	"context"

	"github.com/Azure/azure-service-operator/v2/internal/version"
	"github.com/Azure/azure-service-operator/v2/pkg/xcontext"
	"github.com/spf13/cobra"
)

// Execute kicks off the command line
func Execute() {
	progress := Progress()
	cmd, err := newRootCommand()
	if err != nil {
		log.Error(err, "failed to build commands")
	}

	ctx := xcontext.MakeInterruptibleContext(context.Background())
	if err := cmd.ExecuteContext(ctx); err != nil {
		log.Error(err, "failed to execute command")
		return
	}

	// Wait for the progress bar to finish updating before exiting
	progress.Wait()
}

func newRootCommand() (*cobra.Command, error) {
	rootCmd := &cobra.Command{
		Use:              "asoctl",
		Short:            "asoctl provides a cmdline interface for working with Azure Service Operator",
		TraverseChildren: true,
	}

	rootCmd.Flags().SortFlags = false

	cmdFuncs := []func() (*cobra.Command, error){
		newCRDCommand,
		newImportCommand,
		version.NewCommand,
	}

	for _, f := range cmdFuncs {
		cmd, err := f()
		if err != nil {
			return rootCmd, err
		}
		rootCmd.AddCommand(cmd)
	}

	return rootCmd, nil
}
