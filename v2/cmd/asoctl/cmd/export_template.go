/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package cmd

import (
	"fmt"
	"os"
	"reflect"
	"regexp"
	"strings"

	. "github.com/Azure/azure-service-operator/v2/internal/logging"

	"github.com/go-logr/logr"
	"github.com/rotisserie/eris"
	"github.com/spf13/cobra"
	"k8s.io/apimachinery/pkg/runtime/schema"

	"github.com/Azure/azure-service-operator/v2/api"
	"github.com/Azure/azure-service-operator/v2/cmd/asoctl/internal/template"
	"github.com/Azure/azure-service-operator/v2/internal/util/match"
)

type templateOptions struct {
	source  string
	version string

	raw        bool
	crdPattern []string
}

var expectedVersionRegex = regexp.MustCompile(`^v\d\.\d+\.\d+$`)

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
					return eris.New("specified version doesn't match expected format 'v#.#.#' (example: v2.6.0)")
				}

				uri = template.URIFromVersion(options.version)
			}

			log := CreateLogger()
			log.V(Status).Info("Export ASO Template", "version", options.version)
			log.V(Status).Info("Downloading template", "uri", uri)

			result, err := template.Get(ctx, uri)
			if err != nil {
				return err
			}

			if !options.raw {
				// Check to see if the CRD patterns are valid
				validateCRDPatternsForTemplate(options.crdPattern, log)

				// Render the CRDPattern w/ `;` separating them and inject into the template
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

func validateCRDPatternsForTemplate(
	crdPatterns []string,
	log logr.Logger,
) {
	scheme := api.CreateScheme()
	knownCRDs := scheme.AllKnownTypes()
	warnings := 0
	for _, p := range crdPatterns {
		// We split patterns by `;` so we can check them individually and provide more fine-grained feedback
		for _, pattern := range strings.Split(p, ";") {
			count := countCRDsMatchingPattern(knownCRDs, pattern)
			if count == 0 {
				log.V(Status).Info("No CRDs matched pattern", "pattern", pattern)
				warnings++
			} else {
				log.V(Info).Info("Pattern matched CRDs", "pattern", pattern, "count", count)
			}
		}
	}

	if warnings > 0 {
		log.V(Status).Info("Possibly ineffective CRD patterns detected")
		log.V(Status).Info("Please check the CRDs supported by the installed version of ASO")
	}
}

func countCRDsMatchingPattern(
	knownCRDs map[schema.GroupVersionKind]reflect.Type,
	pattern string,
) int {
	result := 0
	matcher := match.NewStringMatcher(pattern)
	for gvk := range knownCRDs {
		key := fmt.Sprintf("%s/%s", gvk.Group, gvk.Kind)
		if matcher.Matches(key).Matched {
			result++
		}
	}

	return result
}
