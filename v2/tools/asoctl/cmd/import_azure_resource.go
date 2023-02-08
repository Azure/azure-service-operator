/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package importing

import (
	"github.com/spf13/cobra"
	"k8s.io/klog/v2"
)

func newImportAzureResourceCommand() *cobra.Command {
	var output *string

	cmd := &cobra.Command{
		Use:   "azure-resource <ARM/ID/of/resource>",
		Short: "imports an ARM resource as a CR",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			armID := args[0]
			return importAzureResource(armID, output)
		}),
	}

	output = cmd.Flags().StringP(
		"output",
		"o",
		"",
		"Write ARM resource CRD to a file")

	return cmd
}

// TODO: importing azure resource logic goes here
func importAzureResource(armID string, output *string) error {
	klog.Info("importing azure resource")
	importer := azureresource.NewImporter()
	err := importer.Import(armID)
	if err != nil {
		klog.Errorf("failed to import resource %s", armID)
		return errors.Wrapf(err, "failed to import resource %s:", armID)
	}

	return nil
}
