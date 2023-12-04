/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"context"

	"github.com/pkg/errors"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/armconversion"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/config"
)

// PruneResourcesWithLifecycleOwnedByParentStageID is the unique identifier for this pipeline stage
const PruneResourcesWithLifecycleOwnedByParentStageID = "pruneResourcesWithLifecycleOwnedByParentStage"

// PruneResourcesWithLifecycleOwnedByParent prunes networking embedded types
func PruneResourcesWithLifecycleOwnedByParent(configuration *config.Configuration) *Stage {
	stage := NewStage(
		PruneResourcesWithLifecycleOwnedByParentStageID,
		"Prune embedded resources whose lifecycle is owned by the parent.",
		func(ctx context.Context, state *State) (*State, error) {
			result := make(astmodel.TypeDefinitionSet)

			// A previous stage may have used these flags, but we want to make sure we're using them too so reset
			// the consumed bit
			err := configuration.ObjectModelConfiguration.ResourceLifecycleOwnedByParent.MarkUnconsumed()
			if err != nil {
				return nil, err
			}

			pruner := newMisbehavingEmbeddedTypeVisitor(configuration)

			// TODO: This is a hack placed here to protect future releases from include VNET but not
			// TODO: the corresponding Subnet. Each networking APIVersion that supports VNET must also
			// TODO: support subnet or the code in virtual_network_extensions.go will not function properly.
			// TODO: It's likely that failure to function would be caught by tests, but blocking it here
			// TODO: as an extra precaution.
			for _, def := range state.Definitions() {
				if def.Name().Name() == "VirtualNetwork" {
					subnetName := def.Name().WithName("VirtualNetworksSubnet")
					if !state.Definitions().Contains(subnetName) {
						return nil, errors.Errorf("Couldn't find subnet type matching %s. VirtualNetwork and VirtualNetworksSubnet must always be exported together", def.Name())
					}
				}
			}

			for _, def := range state.Definitions() {
				var updatedDef astmodel.TypeDefinition
				updatedDef, err = pruner.visitor.VisitDefinition(def, def.Name())
				if err != nil {
					return nil, errors.Wrapf(err, "failed to visit definition %s", def.Name())
				}
				result.Add(updatedDef)
			}

			result, err = flagPrunedEmptyProperties(result, pruner.emptyPrunedProperties)
			if err != nil {
				return nil, err
			}

			err = configuration.ObjectModelConfiguration.ResourceLifecycleOwnedByParent.VerifyConsumed()
			if err != nil {
				return nil, err
			}

			return state.WithDefinitions(result), nil
		})

	stage.RequiresPrerequisiteStages(CreateARMTypesStageID)
	return stage
}

func flagPrunedEmptyProperties(
	defs astmodel.TypeDefinitionSet,
	emptyPrunedProps astmodel.InternalTypeNameSet,
) (astmodel.TypeDefinitionSet, error) {
	emptyObjectVisitor := astmodel.TypeVisitorBuilder[astmodel.InternalTypeNameSet]{
		VisitObjectType: tagEmptyObjectARMProperty,
	}.Build()

	emptyPrunedPropertiesArm := astmodel.NewInternalTypeNameSet()
	for emptyPrunedProp := range emptyPrunedProps {
		// we need to add the noConversion tag on ARM type for the empty pruned property to relax the validation for convertToARM function.
		armDef, err := GetARMTypeDefinition(defs, emptyPrunedProp)
		if err != nil {
			return nil, err
		}
		emptyPrunedPropertiesArm.Add(armDef.Name())
	}

	result, err := emptyObjectVisitor.VisitDefinitions(defs, emptyPrunedPropertiesArm)
	if err != nil {
		return nil, err

	}

	return result, nil
}

type misbehavingEmbeddedTypePruner struct {
	configuration         *config.Configuration
	emptyPrunedProperties astmodel.InternalTypeNameSet
	visitor               astmodel.TypeVisitor[astmodel.InternalTypeName]
}

func newMisbehavingEmbeddedTypeVisitor(configuration *config.Configuration) *misbehavingEmbeddedTypePruner {
	pruner := &misbehavingEmbeddedTypePruner{
		configuration:         configuration,
		emptyPrunedProperties: astmodel.NewInternalTypeNameSet(),
	}

	visitor := astmodel.TypeVisitorBuilder[astmodel.InternalTypeName]{
		VisitObjectType: pruner.pruneMisbehavingEmbeddedResourceProperties,
	}.Build()

	pruner.visitor = visitor
	return pruner
}

// tagEmptyObjectARMProperty finds the empty properties in an Object and adds the ConversionTag:NoARMConversionValue property tag.
func tagEmptyObjectARMProperty(
	this *astmodel.TypeVisitor[astmodel.InternalTypeNameSet],
	it *astmodel.ObjectType,
	ctx astmodel.InternalTypeNameSet,
) (astmodel.Type, error) {
	prop, ok := it.Properties().Find(func(prop *astmodel.PropertyDefinition) bool {
		typeName, ok := astmodel.ExtractTypeName(prop.PropertyType())
		if !ok {
			return false
		}

		return ctx.Contains(typeName)
	})

	if ok {
		prop = prop.WithTag(armconversion.ConversionTag, armconversion.NoARMConversionValue)
		it = it.WithProperty(prop)
	}

	return astmodel.IdentityVisitOfObjectType(this, it, ctx)
}

func (m *misbehavingEmbeddedTypePruner) pruneMisbehavingEmbeddedResourceProperties(
	this *astmodel.TypeVisitor[astmodel.InternalTypeName],
	it *astmodel.ObjectType,
	ctx astmodel.InternalTypeName,
) (astmodel.Type, error) {
	for _, prop := range it.Properties().Copy() {
		if _, ok := m.configuration.ObjectModelConfiguration.ResourceLifecycleOwnedByParent.Lookup(ctx, prop.PropertyName()); !ok {
			continue
		}

		it = it.WithoutProperty(prop.PropertyName())
		if it.Properties().Len() == 0 {
			m.emptyPrunedProperties.Add(ctx)
		}
	}

	return astmodel.IdentityVisitOfObjectType(this, it, ctx)
}
