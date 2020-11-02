/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package cmd

import (
	"github.com/spf13/cobra"
	"k8s.io/klog/v2"
)

// Execute kicks off the command line
func Execute() {
	cmd, err := newRootCommand()
	if err != nil {
		klog.Fatalf("fatal error: commands failed to build! %v\n", err)
	}

	if err := cmd.Execute(); err != nil {
		klog.Fatalln(err)
	}
}

func newRootCommand() (*cobra.Command, error) {
	rootCmd := &cobra.Command{
		Use:              "k8sinfra",
		Short:            "k8sinfra provides a cmdline interface for generating k8s-infra types from Azure deployment template schema",
		TraverseChildren: true,
	}

	rootCmd.Flags().SortFlags = false

	cmdFuncs := []func() (*cobra.Command, error){
		NewGenTypesCommand,
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
