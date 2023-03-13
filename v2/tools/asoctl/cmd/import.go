/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package main

import "github.com/spf13/cobra"

// newImportCommand creates a new cobra command for importing other resources as ASO custom resources
func newImportCommand() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "import",
		Short: "Import related actions",
		Args:  cobra.ExactArgs(1),
	}

	importAzureResourceCommand := newImportAzureResourceCommand()
	cmd.AddCommand(importAzureResourceCommand)

	return cmd, nil
}
