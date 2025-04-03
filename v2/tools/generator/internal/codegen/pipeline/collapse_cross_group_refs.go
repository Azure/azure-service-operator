/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"context"

	"github.com/rotisserie/eris"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
)

// CollapseCrossGroupReferencesStageID is the unique identifier for this pipeline stage
const CollapseCrossGroupReferencesStageID = "collapseCrossGroupReferences"

// CollapseCrossGroupReferences finds and removes references between API groups. This isn't particularly common
// but does occur in a few instances, for example from Microsoft.Compute -> Microsoft.Compute.Extensions.
func CollapseCrossGroupReferences(idFactory astmodel.IdentifierFactory) *Stage {
	return NewStage(
		CollapseCrossGroupReferencesStageID,
		"Find and remove cross group references",
		func(ctx context.Context, state *State) (*State, error) {
			result := make(astmodel.TypeDefinitionSet)
			for name, def := range state.Definitions().AllResources() {
				walker := newTypeWalker(idFactory, state.Definitions(), name)
				updatedTypes, err := walker.Walk(def)
				if err != nil {
					return nil, eris.Wrapf(err, "failed walking definitions")
				}

				for _, newDef := range updatedTypes {
					err := result.AddAllowDuplicates(newDef)
					if err != nil {
						return nil, err
					}
				}
			}

			return state.WithDefinitions(result), nil
		})
}

func newTypeWalker(
	idFactory astmodel.IdentifierFactory,
	definitions astmodel.TypeDefinitionSet,
	resourceName astmodel.InternalTypeName,
) *astmodel.TypeWalker[any] {
	visitor := astmodel.TypeVisitorBuilder[any]{}.Build()
	walker := astmodel.NewTypeWalker(definitions, visitor)

	walker.AfterVisit = func(original astmodel.TypeDefinition, updated astmodel.TypeDefinition, ctx interface{}) (astmodel.TypeDefinition, error) {
		if !resourceName.PackageReference().Equals(updated.Name().PackageReference()) {

			newName := astmodel.MakeInternalTypeName(resourceName.InternalPackageReference(), updated.Name().Name())
			if existingDef, ok := definitions[newName]; ok {
				if !astmodel.TypeEquals(existingDef.Type(), updated.Type()) {
					// There exists a type with this name already and it doesn't match this types shape. This is rare
					// but happens in a few instances. One example instance is Fleet + AKS which both have a
					// UserAssignedIdentity type that have slightly different shapes.
					// We disambiguate by including the name of the cross-resource pkg the type came from in the name of the type
					disambiguator := idFactory.CreateIdentifier(updated.Name().InternalPackageReference().Group(), astmodel.Exported)
					newName = astmodel.MakeInternalTypeName(resourceName.InternalPackageReference(), disambiguator+updated.Name().Name())
				}
			}
			updated = updated.WithName(newName)
		}
		return astmodel.IdentityAfterVisit(original, updated, ctx)
	}

	return walker
}
