/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package cmd

import (
	"context"
	"os"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-service-operator/v2/api"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"k8s.io/klog/v2"

	"github.com/Azure/azure-service-operator/v2/tools/asoctl/internal/importing"
)

func newImportAzureResourceCommand() *cobra.Command {
	var output *string

	cmd := &cobra.Command{
		Use:   "azure-resource <ARM/ID/of/resource>",
		Short: "imports an ARM resource as a CR",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			armID := args[0]
			ctx := cmd.Context()
			return importAzureResource(ctx, armID, output)
		},
	}

	output = cmd.Flags().StringP(
		"output",
		"o",
		"",
		"Write ARM resource CRD to a file")

	return cmd
}

// TODO: importing azure resource logic goes here
func importAzureResource(ctx context.Context, armID string, output *string) error {
	activeCloud := cloud.AzurePublic
	client, err := importing.CreateARMClient(activeCloud)
	if err != nil {
		return errors.Wrapf(err, "failed to create ARM client")
	}

	importer := importing.NewResourceImporter(api.CreateScheme())
	armImporter := importer.CreateARMImporter(
		client,
		activeCloud.Services[cloud.ResourceManager])

	result, err := armImporter.Import(ctx, armID)
	if err != nil {
		return errors.Wrapf(err, "failed to import resource %s:", armID)
	}

	if output == nil || *output == "" {
		err := result.SaveToWriter(os.Stdout)
		if err != nil {
			klog.Errorf("failed to write to stdout")
			return errors.Wrapf(err, "failed to write to stdout")
		}
	} else {
		err := result.SaveToFile(*output)
		if err != nil {
			klog.Errorf("failed to write to file %s", *output)
			return errors.Wrapf(err, "failed to write to file %s", *output)
		}
	}

	return nil
}
