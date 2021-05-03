/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package codegen

import (
	"context"

	"github.com/Azure/k8s-infra/hack/generator/pkg/astmodel"
	"github.com/Azure/k8s-infra/hack/generator/pkg/codegen/embeddedresources"
)

func removeEmbeddedResources() PipelineStage {
	return MakePipelineStage(
		"removeEmbeddedResources",
		"Removes properties that point to embedded resources. Only removes structural aspects of embedded resources, Id/ARMId references are retained.",
		func(ctx context.Context, types astmodel.Types) (astmodel.Types, error) {

			remover, err := embeddedresources.MakeEmbeddedResourceRemover(types)
			if err != nil {
				return nil, err
			}

			return remover.RemoveEmbeddedResources()
		})
}
