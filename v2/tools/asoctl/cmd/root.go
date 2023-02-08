/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"k8s.io/klog/v2"

	"github.com/Azure/azure-service-operator/v2/internal/version"
	"github.com/Azure/azure-service-operator/v2/pkg/xcontext"
)

// Execute kicks off the command line
func Execute() {
	cmd, err := newRootCommand()
	if err != nil {
		klog.Fatalf("fatal error: commands failed to build! %s\n", err)
	}

	ctx := xcontext.MakeInterruptibleContext(context.Background())
	if err := cmd.ExecuteContext(ctx); err != nil {
		klog.Fatalln(err)
	}
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
