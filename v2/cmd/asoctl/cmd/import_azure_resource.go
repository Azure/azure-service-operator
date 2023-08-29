/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package cmd

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/Azure/azure-service-operator/v2/api"
	"github.com/Azure/azure-service-operator/v2/internal/genericarmclient"
	"github.com/Azure/azure-service-operator/v2/internal/version"

	"github.com/Azure/azure-service-operator/v2/cmd/asoctl/internal/importing"
)

func newImportAzureResourceCommand() *cobra.Command {
	var options importAzureResourceOptions

	cmd := &cobra.Command{
		Use:   "azure-resource <ARM/ID/of/resource>",
		Short: "Import ARM resources as Custom Resources",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()
			return importAzureResource(ctx, args, options)
		},
	}

	options.outputPath = cmd.Flags().StringP(
		"output",
		"o",
		"",
		"Write ARM resource CRDs to a single file")

	options.outputFolder = cmd.Flags().StringP(
		"output-folder",
		"f",
		"",
		"Write ARM resource CRDs to individual files in a folder")

	cmd.MarkFlagsMutuallyExclusive("output", "output-folder")

	return cmd
}

// importAzureResource imports an ARM resource and writes the YAML to stdout or a file
func importAzureResource(ctx context.Context, armIDs []string, options importAzureResourceOptions) error {

	log, progress := CreateLoggerAndProgressBar()

	//TODO: Support other Azure clouds
	activeCloud := cloud.AzurePublic
	creds, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		return errors.Wrap(err, "unable to get default Azure credential")
	}

	clientOptions := &genericarmclient.GenericClientOptions{
		UserAgent: "asoctl/" + version.BuildVersion,
	}

	client, err := genericarmclient.NewGenericClient(activeCloud, creds, clientOptions)
	if err != nil {
		return errors.Wrapf(err, "failed to create ARM client")
	}

	importer := importing.NewResourceImporter(api.CreateScheme(), client, log, progress)
	for _, armID := range armIDs {
		err = importer.AddARMID(armID)
		if err != nil {
			return errors.Wrapf(err, "failed to add %q to import list", armID)
		}
	}

	result, err := importer.Import(ctx)

	if ctx.Err() != nil {
		return ctx.Err()
	}

	if err != nil {
		if result.Count() == 0 {
			return errors.Wrap(err, "failed to import any resources")
		}

		log.Error(err, "Failed to import some resources.")
		log.Info("Will still save those resources that were imported successfully.")
	}

	if result.Count() == 0 {
		log.Info("No resources found, nothing to save.")
		return nil
	}

	if file, ok := options.writeToFile(); ok {
		log.Info(
			"Writing to a single file",
			"file", file)
		err := result.SaveToSingleFile(file)
		if err != nil {
			return errors.Wrapf(err, "failed to write to file %s", file)
		}
	} else if folder, ok := options.writeToFolder(); ok {
		log.Info(
			"Writing to individual files in folder",
			"folder", folder)
		err := result.SaveToIndividualFilesInFolder(folder)
		if err != nil {
			return errors.Wrapf(err, "failed to write into folder %s", folder)
		}
	} else {
		err := result.SaveToWriter(progress)
		if err != nil {
			return errors.Wrapf(err, "failed to write to stdout")
		}
	}

	// No error, wait for progress bar to finish & flush
	progress.Wait()

	return nil
}

type importAzureResourceOptions struct {
	outputPath   *string
	outputFolder *string
}

func (option *importAzureResourceOptions) writeToFile() (string, bool) {
	if option.outputPath != nil && *option.outputPath != "" {
		return *option.outputPath, true
	}

	return "", false
}

func (option *importAzureResourceOptions) writeToFolder() (string, bool) {
	if option.outputFolder != nil && *option.outputFolder != "" {
		return *option.outputFolder, true
	}

	return "", false
}
