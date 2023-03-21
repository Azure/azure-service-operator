/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/spf13/cobra"
	"k8s.io/klog/v2"
)

// Execute kicks off the command line
func Execute() {
	cmd, err := newRootCommand()
	if err != nil {
		klog.Fatalf("fatal error: commands failed to build! %s\n", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	// Wait for a signal to quit:
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, os.Kill)

	go func() {
		<-signalChan
		cancel()
	}()

	if err := cmd.ExecuteContext(ctx); err != nil {
		klog.Fatalln(err)
	}
}

func newRootCommand() (*cobra.Command, error) {
	rootCmd := &cobra.Command{
		Use:              "aso-gen",
		Short:            "aso-gen provides a cmdline interface for generating Azure Service Operator types from Azure deployment template schema",
		TraverseChildren: true,
	}

	rootCmd.Flags().SortFlags = false

	cmdFuncs := []func() (*cobra.Command, error){
		NewGenTypesCommand,
		NewGenKustomizeCommand,
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
