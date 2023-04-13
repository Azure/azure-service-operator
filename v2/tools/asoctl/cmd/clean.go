/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package cmd

import "github.com/spf13/cobra"

// newCleanCommand creates a new cobra command for working with CRDs
func newCleanCommand() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "clean",
		Short: "Clean Custom Resource Definitions (CRDs) prior to upgrade",
		Args:  cobra.NoArgs,
	}

	cleanCRDsCommand := newCleanCRDsCommand()
	cmd.AddCommand(cleanCRDsCommand)

	return cmd, nil
}
