/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package codegen

import (
	"context"
	"fmt"

	"k8s.io/klog/v2"

	"github.com/Azure/azure-service-operator/hack/generator/pkg/astmodel"
	"github.com/Azure/azure-service-operator/hack/generator/pkg/codegen/pipeline"
	"github.com/Azure/azure-service-operator/hack/generator/pkg/config"
)

// applyExportFilters creates a PipelineStage to reduce our set of types for export
func applyExportFilters(configuration *config.Configuration) pipeline.PipelineStage {
	return pipeline.MakePipelineStage(
		"filterTypes",
		"Filter generated types",
		func(ctx context.Context, types astmodel.Types) (astmodel.Types, error) {
			return filterTypes(configuration, types)
		})
}

// filterTypes applies the configuration include/exclude filters to the generated definitions
func filterTypes(
	configuration *config.Configuration,
	definitions astmodel.Types) (astmodel.Types, error) {

	newDefinitions := make(astmodel.Types)

	filterer := configuration.BuildExportFilterer(definitions)

	for _, def := range definitions {
		defName := def.Name()
		shouldExport, reason := filterer(defName)

		switch shouldExport {
		case config.Skip:
			klog.V(3).Infof("Skipping %s because %s", defName, reason)

		case config.Export:
			if reason == "" {
				klog.V(3).Infof("Exporting %s", defName)
			} else {
				klog.V(2).Infof("Exporting %s because %s", defName, reason)
			}

			newDefinitions[def.Name()] = def
		default:
			panic(fmt.Sprintf("unhandled shouldExport case %q", shouldExport))
		}
	}

	// Ensure that the export filters had no issues
	err := configuration.GetExportFiltersError()
	if err != nil {
		return nil, err
	}

	return newDefinitions, nil
}
