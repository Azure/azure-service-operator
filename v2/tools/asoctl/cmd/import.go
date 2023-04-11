/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package cmd

import "github.com/spf13/cobra"

// newImportCommand creates a new cobra Command when invoked from the command line
func newImportCommand() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "import",
		Short: "imports ARM resources as YAML resource definitions",
		Args:  cobra.ExactArgs(1),
	}

	importAzureResourceCommand := newImportAzureResourceCommand()
	cmd.AddCommand(importAzureResourceCommand)

	return cmd, nil
}
