/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"context"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/codegen/embeddedresources"
)

// RemoveEmbeddedResourcesStageID is the unique identifier for this pipeline stage
const RemoveEmbeddedResourcesStageID = "removeEmbeddedResources"

func RemoveEmbeddedResources() Stage {
	return MakeLegacyStage(
		RemoveEmbeddedResourcesStageID,
		"Remove properties that point to embedded resources. Only removes structural aspects of embedded resources, Id/ARMId references are retained.",
		func(ctx context.Context, types astmodel.Types) (astmodel.Types, error) {
			remover, err := embeddedresources.MakeEmbeddedResourceRemover(types)
			if err != nil {
				return nil, err
			}

			return remover.RemoveEmbeddedResources()
		})
}
