/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package cmd

import (
	"context"
	"os"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"k8s.io/klog/v2"

	"github.com/Azure/azure-service-operator/v2/internal/genericarmclient"
	"github.com/Azure/azure-service-operator/v2/internal/version"

	"github.com/Azure/azure-service-operator/v2/api"

	"github.com/Azure/azure-service-operator/v2/tools/asoctl/internal/importing"
)

func newImportAzureResourceCommand() *cobra.Command {
	var outputPath *string

	cmd := &cobra.Command{
		Use:   "azure-resource <ARM/ID/of/resource>",
		Short: "imports an ARM resource as a CR",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			armID := args[0]
			ctx := cmd.Context()
			return importAzureResource(ctx, armID, outputPath)
		},
	}

	outputPath = cmd.Flags().StringP(
		"output",
		"o",
		"",
		"Write ARM resource CRD to a file")

	return cmd
}

// importAzureResource imports an ARM resource and writes the YAML to stdout or a file
func importAzureResource(ctx context.Context, armID string, outputPath *string) error {
	//TODO: Support other clouds

	activeCloud := cloud.AzurePublic
	creds, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		return errors.Wrap(err, "unable to get default azure credential")
	}

	options := &genericarmclient.GenericClientOptions{
		UserAgent: "asoctl/" + version.BuildVersion,
	}

	client, err := genericarmclient.NewGenericClient(activeCloud, creds, options)
	if err != nil {
		return errors.Wrapf(err, "failed to create ARM client")
	}

	importer := importing.NewResourceImporter(api.CreateScheme(), client)
	importer.AddARMID(armID)

	result, err := importer.Import(ctx)
	if err != nil {
		return errors.Wrapf(err, "failed to import resource %s:", armID)
	}

	if outputPath == nil || *outputPath == "" {
		err := result.SaveToWriter(os.Stdout)
		if err != nil {
			klog.Errorf("failed to write to stdout")
			return errors.Wrapf(err, "failed to write to stdout")
		}
	} else {
		err := result.SaveToFile(*outputPath)
		if err != nil {
			return errors.Wrapf(err, "failed to write to file %s", *outputPath)
		}
	}

	return nil
}
