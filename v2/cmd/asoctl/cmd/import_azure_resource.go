/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package cmd

import (
	"context"
	"fmt"
	"io"
	"os"
	"sync"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/go-logr/logr"
	"github.com/rotisserie/eris"
	"github.com/spf13/cobra"
	"github.com/vbauerster/mpb/v8"

	"github.com/Azure/azure-service-operator/v2/api"
	"github.com/Azure/azure-service-operator/v2/cmd/asoctl/pkg/importreporter"
	"github.com/Azure/azure-service-operator/v2/cmd/asoctl/pkg/importresources"
	"github.com/Azure/azure-service-operator/v2/internal/genericarmclient"
	"github.com/Azure/azure-service-operator/v2/internal/version"
	asocloud "github.com/Azure/azure-service-operator/v2/pkg/common/cloud"
	"github.com/Azure/azure-service-operator/v2/pkg/common/config"
)

func newImportAzureResourceCommand() *cobra.Command {
	var options importAzureResourceOptions

	cmd := &cobra.Command{
		Use:   "azure-resource <ARM/ID/of/resource>",
		Short: "Import ARM resources as Custom Resources",
		Long: `Imports ARM resources as Custom Resources.

This command requires you to authenticate with Azure using an identity which has access to the resource(s) you would
like to import. The following authentication modes are supported:

Az-login token: az login and then use asoctl.
Managed Identity: Set the AZURE_CLIENT_ID environment variable and run on a machine with access to the managed identity.
Service Principal: Set the AZURE_SUBSCRIPTION_ID, AZURE_TENANT_ID, AZURE_CLIENT_ID, and AZURE_CLIENT_SECRET environment variables,

The following environment variables can be used to configure which cloud to use with asoctl:

AZURE_RESOURCE_MANAGER_ENDPOINT: The Azure Resource Manager endpoint. 
If not specified, the default is the Public cloud resource manager endpoint.
See https://docs.microsoft.com/cli/azure/manage-clouds-azure-cli#list-available-clouds for details
about how to find available resource manager endpoints for your cloud. Note that the resource manager
endpoint is referred to as "resourceManager" in the Azure CLI.

AZURE_RESOURCE_MANAGER_AUDIENCE: The Azure Resource Manager AAD audience.
If not specified, the default is the Public cloud resource manager audience https://management.core.windows.net/.
See https://docs.microsoft.com/cli/azure/manage-clouds-azure-cli#list-available-clouds for details
about how to find available resource manager audiences for your cloud. Note that the resource manager
audience is referred to as "activeDirectoryResourceId" in the Azure CLI.

AZURE_AUTHORITY_HOST: The URL of the AAD authority.
If not specified, the default
is the AAD URL for the public cloud: https://login.microsoftonline.com/. See
https://docs.microsoft.com/azure/active-directory/develop/authentication-national-cloud
`,
		Args: cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()
			return importAzureResource(ctx, args, &options)
		},
	}

	cmd.Flags().StringVarP(
		&options.outputPath,
		"output",
		"o",
		"",
		"Write ARM resource CRDs to a single file")

	cmd.Flags().StringVarP(
		&options.outputFolder,
		"output-folder",
		"f",
		"",
		"Write ARM resource CRDs to individual files in a folder")

	cmd.MarkFlagsMutuallyExclusive("output", "output-folder")

	cmd.Flags().StringVarP(
		&options.namespace,
		"namespace",
		"n",
		"",
		"Set the namespace of the the imported resources")

	cmd.Flags().StringSliceVarP(
		&options.labels,
		"label",
		"l",
		nil,
		"Add labels to the imported resources. Multiple comma-separated labels can be specified (--label example.com/mylabel=foo,example.com/mylabel2=bar) or the --label (-l) argument can be used multiple times (-l example.com/mylabel=foo -l example.com/mylabel2=bar)")

	cmd.Flags().StringSliceVarP(
		&options.annotations,
		"annotation",
		"a",
		nil,
		"Add annotations to the imported resources. Multiple comma-separated annotations can be specified (--annotation example.com/myannotation=foo,example.com/myannotation2=bar) or the --annotation (-a) argument can be used multiple times (-a example.com/myannotation=foo -a example.com/myannotation2=bar)")

	cmd.Flags().IntVarP(
		&options.workers,
		"workers",
		"w",
		4,
		"The number of parallel workers to use when importing resources")

	cmd.Flags().BoolVarP(
		&options.simpleLogging,
		"simple-logging",
		"s",
		false,
		"Use simple logging instead of progress bars")

	return cmd
}

// importAzureResource imports an ARM resource and writes the YAML to stdout or a file
func importAzureResource(
	ctx context.Context,
	armIDs []string,
	options *importAzureResourceOptions,
) error {
	// Check we're being asked to do something ... anything ...
	if len(armIDs) == 0 {
		return eris.New("no ARM IDs provided")
	}

	// Create an ARM client for requesting resources
	client, err := createARMClient(options)
	if err != nil {
		return eris.Wrapf(err, "failed to create ARM client")
	}

	done := make(chan struct{}) // signal for when we're done

	var log logr.Logger
	var progress importreporter.Interface
	var output io.Writer
	if options.simpleLogging {
		log = CreateLogger()
		progress = importreporter.NewLog("Import Azure Resources", log, done)
		output = os.Stderr
	} else {
		var bar *mpb.Progress

		// Caution: the progress bar can deadlock if no bar is ever created, so make sure the gap between
		// this and the call to importer.Import() is as small as possible.
		log, bar = CreateLoggerAndProgressBar()
		progress = importreporter.NewBar("Import Azure Resources", bar, done)

		// The progress bar can deadlock if no bar is ever created - which can happen if we need to return an error
		// between here and the call to importer.Import(). To avoid this, we create a dummy bar that will complete
		// asynchronously after a short delay.
		setup := progress.Create("Initalizing")
		setup.AddPending(1)
		go func() {
			time.Sleep(200 * time.Millisecond)
			setup.Completed(1)
		}()

		output = bar

		// Ensure the progress bar is closed when we're done, but skip if a panic has happened
		defer func() {
			if p := recover(); p != nil {
				fmt.Fprintf(os.Stderr, "panic: %s\n", p)
			} else {
				// No panic
				defer bar.Wait()
			}
		}()
	}

	importerOptions := importresources.ResourceImporterOptions{
		Workers: options.workers,
	}

	importer := importresources.New(api.CreateScheme(), client, log, progress, importerOptions)
	for _, armID := range armIDs {
		err = importer.AddARMID(armID)
		if err != nil {
			return eris.Wrapf(err, "failed to add %q to import list", armID)
		}
	}

	result, err := importer.Import(ctx, done)

	if ctx.Err() != nil {
		return ctx.Err()
	}

	if err != nil {
		if result.Count() == 0 {
			return eris.Wrap(err, "failed to import any resources")
		}

		log.Error(err, "Failed to import some resources.")
		log.Info("Will still save those resources that were imported successfully.")
	}

	if result.Count() == 0 {
		log.Info("No resources found, nothing to save.")
		return nil
	}

	err = configureImportedResources(options, result)
	if err != nil {
		return eris.Wrap(err, "failed to apply options to imported resources")
	}

	err = writeResources(result, options, log, output)
	if err != nil {
		return eris.Wrap(err, "failed to write resources")
	}

	return nil
}

// createARMClient creates our client for talking to ARM
func createARMClient(options *importAzureResourceOptions) (*genericarmclient.GenericClient, error) {
	creds, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		return nil, eris.Wrap(err, "unable to get default Azure credential")
	}

	clientOptions := &genericarmclient.GenericClientOptions{
		UserAgent: "asoctl/" + version.BuildVersion,
	}

	activeCloud := options.cloud()
	return genericarmclient.NewGenericClient(activeCloud, creds, clientOptions)
}

// configureImportedResources applies additional configuration to imported resources
func configureImportedResources(
	options *importAzureResourceOptions,
	result *importresources.Result,
) error {
	// Set the namespace on all resources
	if options.namespace != "" {
		result.SetNamespace(options.namespace)
	}

	// Apply labels
	if len(options.labels) > 0 {
		err := result.AddLabels(options.labels)
		if err != nil {
			return eris.Wrap(err, "failed to add labels")
		}
	}

	// Apply annotations
	if len(options.annotations) > 0 {
		err := result.AddAnnotations(options.annotations)
		if err != nil {
			return eris.Wrap(err, "failed to add annotations")
		}
	}

	return nil
}

func writeResources(
	result *importresources.Result,
	options *importAzureResourceOptions,
	log logr.Logger,
	out io.Writer,
) error {
	// Write all the resources to a single file
	if file, ok := options.writeToFile(); ok {
		log.Info(
			"Writing to a single file",
			"file", file)
		err := result.SaveToSingleFile(file)
		if err != nil {
			return eris.Wrapf(err, "failed to write to file %s", file)
		}

		return nil
	}

	// Write each resource to an individual file in a folder
	if folder, ok := options.writeToFolder(); ok {
		log.Info(
			"Writing to individual files in folder",
			"folder", folder)
		err := result.SaveToIndividualFilesInFolder(folder)
		if err != nil {
			return eris.Wrapf(err, "failed to write into folder %s", folder)
		}

		return nil
	}

	// Write all the resources to stdout
	err := result.SaveToWriter(out)
	if err != nil {
		return eris.Wrapf(err, "failed to write to stdout")
	}

	return nil
}

type importAzureResourceOptions struct {
	outputPath    string
	outputFolder  string
	namespace     string
	annotations   []string
	labels        []string
	workers       int
	simpleLogging bool

	readCloud sync.Once
	cloudCfg  asocloud.Configuration
}

func (option *importAzureResourceOptions) writeToFile() (string, bool) {
	if option.outputPath != "" {
		return option.outputPath, true
	}

	return "", false
}

func (option *importAzureResourceOptions) writeToFolder() (string, bool) {
	if option.outputFolder != "" {
		return option.outputFolder, true
	}

	return "", false
}

func (option *importAzureResourceOptions) cloud() cloud.Configuration {
	option.readCloud.Do(func() {
		option.cloudCfg = asocloud.Configuration{
			AzureAuthorityHost:      os.Getenv(config.AzureAuthorityHost),
			ResourceManagerEndpoint: os.Getenv(config.ResourceManagerEndpoint),
			ResourceManagerAudience: os.Getenv(config.ResourceManagerAudience),
		}
	})

	return option.cloudCfg.Cloud()
}
