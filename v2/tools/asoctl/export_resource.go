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

func newExportResourceCommand() (*cobra.Command, error) { //nolint:unparam // TODO: Remove this comment when the tool is actually functional
	var output *string

	cmd := &cobra.Command{
		Use:   "resource <ARM/ID/of/resource>",
		Short: "exports an ARM resource CRD",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error { // TODO: Should consider xcobra.RunWithCtx here
			armID := args[0]

			return exportResource(armID, output)
		},
	}

	output = cmd.Flags().StringP(
		"output",
		"o",
		"",
		"Write ARM resource CRD to a file")

	return cmd, nil
}

// TODO: export resource logic goes here
func exportResource(armID string, output *string) error {
	klog.Infof("armID : %s", armID)

	if output != nil && *output != "" {
		klog.Infof("output : %s", *output)
	}

	return nil
}
