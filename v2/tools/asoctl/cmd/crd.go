/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package cmd

import "github.com/spf13/cobra"

// newCRDCommand creates a new cobra command for working with CRDs
func newCRDCommand() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "crd",
		Short: "Custom Resource Definition (CRD) related actions",
		Args:  cobra.NoArgs,
	}

	crdCleanCommand := newCRDCleanCommand()
	cmd.AddCommand(crdCleanCommand)

	return cmd, nil
}
