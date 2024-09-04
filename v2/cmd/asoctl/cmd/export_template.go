/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package cmd

import (
	"os"
	"regexp"
	"strings"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/Azure/azure-service-operator/v2/cmd/asoctl/internal/template"
)

type templateOptions struct {
	source  string
	version string

	raw        bool
	crdPattern []string
}

var expectedVersionRegex = regexp.MustCompile(`^v\d\.\d?\d\.\d$`)

// newTemplateCommand creates a new cobra Command when invoked from the command line
func newTemplateCommand() *cobra.Command {
	var options templateOptions

	cmd := &cobra.Command{
		Use:   "template [--source <string> | --version <string>] [--crd-pattern <string> | --raw]",
		Short: "Template creates a YAML file from the specified ASO yaml template",
		Example: `asoctl export template --version v2.6.0 --crd-pattern "resources.azure.com/*" --crd-pattern "containerservice.azure.com/*"

With combined crd-pattern:
asoctl export template --version v2.6.0 --crd-pattern "resources.azure.com/*;containerservice.azure.com/*;keyvault.azure.com/*;managedidentity.azure.com/*;eventhub.azure.com/*"

With remote source:
asoctl export template --source https://github.com/Azure/azure-service-operator/releases/download/v2.6.0/azureserviceoperator_v2.6.0.yaml --crd-pattern "resources.azure.com/*" --crd-pattern "containerservice.azure.com/*"

With local source:
asoctl export template --source ~/Downloads/azureserviceoperator_v2.6.0.yaml --crd-pattern "resources.azure.com/*" --crd-pattern "containerservice.azure.com/*"

Raw:
asoctl export template --version v2.6.0  --raw

With kubectl:
asoctl export template --version v2.6.0 --crd-pattern "resources.azure.com/*;containerservice.azure.com/*;keyvault.azure.com/*;managedidentity.azure.com/*;eventhub.azure.com/*" | kubectl apply -f -`,
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()
			var uri string

			if options.source != "" {
				uri = options.source
			} else {
				if !expectedVersionRegex.MatchString(options.version) {
					return errors.New("specified version doesn't match expected format 'v#.#.#' (example: v2.6.0)")
				}

				uri = template.URIFromVersion(options.version)
			}

			result, err := template.Get(ctx, uri)
			if err != nil {
				return err
			}

			if !options.raw {
				// Render the CRDPattern w/ `;` separating them
				crdPattern := strings.Join(options.crdPattern, ";")
				result, err = template.Apply(result, crdPattern)
				if err != nil {
					return err
				}
			}

			_, err = os.Stdout.WriteString(result)
			if err != nil {
				return err
			}

			return nil
		},
	}

	cmd.Flags().StringVarP(
		&options.source,
		"source",
		"s",
		"",
		"File or URL path to the ASO YAML template. Use this if you've customized the base ASO YAML locally or are using a base YAML other than the one hosted at https://github.com/Azure/azure-service-operator/tags")

	cmd.Flags().StringVarP(
		&options.version,
		"version",
		"v",
		"",
		"Release version to use.")

	cmd.MarkFlagsOneRequired("source", "version")
	cmd.MarkFlagsMutuallyExclusive("source", "version")

	cmd.Flags().BoolVar(
		&options.raw,
		"raw",
		false,
		"Export the YAML without any variable replacements")

	cmd.Flags().StringSliceVarP(
		&options.crdPattern,
		"crd-pattern",
		"p",
		nil,
		"What new CRDs to install. Existing ASO CRDs in the cluster will always be upgraded even if crdPattern is empty. See https://azure.github.io/azure-service-operator/guide/crd-management/ for more details.")

	cmd.MarkFlagsOneRequired("crd-pattern", "raw")
	cmd.MarkFlagsMutuallyExclusive("crd-pattern", "raw")

	return cmd
}
