/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package codegen

import (
	"context"
	"fmt"

	"github.com/Azure/azure-service-operator/hack/generator/pkg/astmodel"
	"github.com/Azure/azure-service-operator/hack/generator/pkg/config"
	"k8s.io/klog/v2"
)

// applyExportFilters creates a PipelineStage to reduce our set of types for export
func applyExportFilters(configuration *config.Configuration) PipelineStage {
	return MakePipelineStage(
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

	filterer, err := configuration.BuildExportFilterer(definitions)
	if err != nil {
		return nil, err
	}

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

	return newDefinitions, nil
}
