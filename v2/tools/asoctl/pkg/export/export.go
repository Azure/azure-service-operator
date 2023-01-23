/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package export

import "github.com/spf13/cobra"

// NewCommand creates a new cobra Command when invoked from the command line
func NewCommand() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "export",
		Short: "exports ARM resources as YAML resource definitions",
		Args:  cobra.ExactArgs(1),
	}

	exportResourceCommand := newExportResourceCommand()
	cmd.AddCommand(exportResourceCommand)

	return cmd, nil
}
