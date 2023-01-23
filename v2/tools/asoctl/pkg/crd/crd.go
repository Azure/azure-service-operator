/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package crd

import "github.com/spf13/cobra"

// NewCommand creates a new cobra to clean deprecated crd versions
func NewCommand() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "crd",
		Short: "Custom Resource Definition (CRD) related actions",
		Args:  cobra.ExactArgs(1),
	}

	crdCleanCommand := newCRDCleanCommand()
	cmd.AddCommand(crdCleanCommand)

	return cmd, nil
}
