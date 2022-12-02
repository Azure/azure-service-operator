/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"context"
	"fmt"

	"github.com/pkg/errors"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
)

// AddCrossplaneOwnerProperties adds the 3-tuple of (xName, xNameRef, xNameSelector) for each owning resource
func AddCrossplaneOwnerProperties(idFactory astmodel.IdentifierFactory) *Stage {
	return NewLegacyStage(
		"addCrossplaneOwnerProperties",
		"Add the 3-tuple of (xName, xNameRef, xNameSelector) for each owning resource",
		func(ctx context.Context, definitions astmodel.TypeDefinitionSet) (astmodel.TypeDefinitionSet, error) {
			referenceTypeName := astmodel.MakeTypeName(
				CrossplaneRuntimeV1Alpha1Package,
				idFactory.CreateIdentifier("Reference", astmodel.Exported))
			selectorTypeName := astmodel.MakeTypeName(
				CrossplaneRuntimeV1Alpha1Package,
				idFactory.CreateIdentifier("Selector", astmodel.Exported))

			result := make(astmodel.TypeDefinitionSet)
			for _, typeDef := range definitions {
				// TODO: Do we need to rewrite this to deal with wrapping?
				if resource, ok := typeDef.Type().(*astmodel.ResourceType); ok {

					owners, err := lookupOwners(definitions, typeDef)
					if err != nil {
						return nil, errors.Wrapf(err, "failed to look up owners for %s", typeDef.Name())
					}

					// The right-most owner is this type, so remove it
					// Note that lookupOwners will error rather than return an empty list so this is safe
					owners = owners[0 : len(owners)-1]

					// This type has no owners so no modification needed
					if len(owners) == 0 {
						continue
					}

					specDef, err := definitions.ResolveResourceSpecDefinition(resource)
					if err != nil {
						return nil, errors.Wrapf(err, "getting resource spec definition")
					}

					for _, owner := range owners {
						nameSubset := fmt.Sprintf("%sName", owner.Name())
						name := idFactory.CreatePropertyName(nameSubset, astmodel.Exported)
						nameRef := name + "Ref"
						nameSelector := name + "Selector"

						updatedDef, err := specDef.ApplyObjectTransformation(func(o *astmodel.ObjectType) (astmodel.Type, error) {
							nameProperty := astmodel.NewPropertyDefinition(
								name,
								idFactory.CreateStringIdentifier(string(name), astmodel.NotExported),
								astmodel.StringType)
							nameRefProperty := astmodel.NewPropertyDefinition(
								nameRef,
								idFactory.CreateStringIdentifier(string(nameRef), astmodel.NotExported),
								referenceTypeName).MakeTypeOptional()
							nameSelectorProperty := astmodel.NewPropertyDefinition(
								nameSelector,
								idFactory.CreateStringIdentifier(string(nameSelector), astmodel.NotExported),
								selectorTypeName).MakeTypeOptional()

							result := o.WithProperty(nameProperty).WithProperty(nameRefProperty).WithProperty(nameSelectorProperty)
							return result, nil
						})
						if err != nil {
							return nil, errors.Wrapf(err, "adding ownership properties to spec")
						}
						specDef = updatedDef
					}
					result.Add(typeDef)
					result.Add(specDef)
				}
			}

			// Second pass that adds anything that we haven't already added
			for _, typeDef := range definitions {
				if !result.Contains(typeDef.Name()) {
					result.Add(typeDef)
				}
			}

			return result, nil
		})
}

func lookupOwners(defs astmodel.TypeDefinitionSet, resourceDef astmodel.TypeDefinition) ([]astmodel.TypeName, error) {
	resourceType, ok := resourceDef.Type().(*astmodel.ResourceType)
	if !ok {
		return nil, errors.Errorf("type %s is not a resource", resourceDef.Name())
	}

	if resourceType.Owner() == nil {
		return []astmodel.TypeName{resourceDef.Name()}, nil
	}

	if resourceType.Owner().Name() == "ResourceGroup" {
		return []astmodel.TypeName{*resourceType.Owner(), resourceDef.Name()}, nil
	}

	owner := *resourceType.Owner()
	ownerDef, ok := defs[owner]
	if !ok {
		return nil, errors.Errorf("couldn't find definition for owner %s", owner)
	}

	result, err := lookupOwners(defs, ownerDef)
	if err != nil {
		return nil, err
	}

	return append(result, resourceDef.Name()), nil
}
