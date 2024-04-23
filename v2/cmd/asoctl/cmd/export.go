/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package cmd

import "github.com/spf13/cobra"

// newExportCommand creates a new cobra Command when invoked from the command line
func newExportCommand() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "export",
		Short: "Exports an ASO YAML file from a template",
		Args:  cobra.NoArgs,
	}

	cmd.AddCommand(newTemplateCommand())

	return cmd, nil
}
