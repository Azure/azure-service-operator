/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package main

import (
	"github.com/spf13/cobra"
	"k8s.io/klog/v2"
)

// NewExportCommand creates a new cobra Command when invoked from the command line
func NewExportCommand() (*cobra.Command, error) {

	cmd := &cobra.Command{
		Use:   "export",
		Short: "exports ARM templates",
		Args:  cobra.ExactArgs(1),
	}

	exportResourceCommand, err := newExportResourceCommand()
	if err != nil {
		return nil, err
	}
	cmd.AddCommand(exportResourceCommand)

	return cmd, nil
}

func newExportResourceCommand() (*cobra.Command, error) {
	var output *string

	cmd := &cobra.Command{
		Use:   "resource <ARM/ID/of/resource>",
		Short: "exports ARM resource templates",
		Args:  cobra.ExactArgs(1),
		RunE:  exportResourceFunction(output),
	}

	output = cmd.Flags().StringP(
		"output",
		"o",
		"",
		"Write ARM resource template to a file")

	return cmd, nil
}

// TODO: export resource logic goes here
func exportResourceFunction(output *string) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		armID := args[0]
		klog.Infof("armID : %s", armID)

		if output != nil && *output != "" {
			klog.Infof("output : %s", *output)
		}

		return nil
	}
}
