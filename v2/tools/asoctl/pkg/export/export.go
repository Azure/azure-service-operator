/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package export

import "github.com/spf13/cobra"

// NewExportCommand creates a new cobra Command when invoked from the command line
func NewExportCommand() (*cobra.Command, error) {

	cmd := &cobra.Command{
		Use:   "export",
		Short: "exports ARM resources as YAML resource definitions",
		Args:  cobra.ExactArgs(1),
	}

	exportResourceCommand, err := newExportResourceCommand()
	if err != nil {
		return nil, err
	}
	cmd.AddCommand(exportResourceCommand)

	return cmd, nil
}
