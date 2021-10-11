/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"context"

	"github.com/pkg/errors"

	"github.com/Azure/azure-service-operator/v2/internal/generator/astmodel"
)

// CollapseCrossGroupReferencesStageID is the unique identifier for this pipeline stage
const CollapseCrossGroupReferencesStageID = "collapseCrossGroupReferences"

// CollapseCrossGroupReferences finds and removes references between API groups. This isn't particularly common
// but does occur in a few instances, for example from Microsoft.Compute -> Microsoft.Compute.Extensions.
func CollapseCrossGroupReferences() Stage {
	return MakeLegacyStage(
		CollapseCrossGroupReferencesStageID,
		"Find and remove cross group references",
		func(ctx context.Context, types astmodel.Types) (astmodel.Types, error) {
			resources := astmodel.CollectResourceDefinitions(types)
			result := make(astmodel.Types)

			for resourceName := range resources {
				walker := newTypeWalker(types, resourceName)
				updatedTypes, err := walker.Walk(types[resourceName])
				if err != nil {
					return nil, errors.Wrapf(err, "failed walking types")
				}

				for _, newDef := range updatedTypes {
					err := result.AddAllowDuplicates(newDef)
					if err != nil {
						return nil, err
					}
				}
			}

			return result, nil
		})
}

func newTypeWalker(types astmodel.Types, resourceName astmodel.TypeName) *astmodel.TypeWalker {
	visitor := astmodel.TypeVisitorBuilder{}.Build()
	walker := astmodel.NewTypeWalker(types, visitor)
	walker.AfterVisit = func(original astmodel.TypeDefinition, updated astmodel.TypeDefinition, ctx interface{}) (astmodel.TypeDefinition, error) {
		if !resourceName.PackageReference.Equals(updated.Name().PackageReference) {
			// Note: If we ever find this generating colliding names, we might need to introduce a unique suffix.
			// For now though it doesn't seem to, so preserving the shorter names as they're clearer.
			updated = updated.WithName(astmodel.MakeTypeName(resourceName.PackageReference, updated.Name().Name()))
		}
		return astmodel.IdentityAfterVisit(original, updated, ctx)
	}

	return walker
}
