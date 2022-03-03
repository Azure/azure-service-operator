/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"context"
	"strings"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/config"

	"github.com/pkg/errors"
)

const resourcesPropertyName = astmodel.PropertyName("Resources")

func DetermineResourceOwnership(configuration *config.Configuration) *Stage {
	return NewLegacyStage(
		"determineResourceOwnership",
		"Determine ARM resource relationships",
		func(ctx context.Context, definitions astmodel.TypeDefinitionSet) (astmodel.TypeDefinitionSet, error) {
			return determineOwnership(definitions, configuration)
		})
}

func determineOwnership(definitions astmodel.TypeDefinitionSet, configuration *config.Configuration) (astmodel.TypeDefinitionSet, error) {
	updatedDefs := make(astmodel.TypeDefinitionSet)

	resources := astmodel.FindResourceDefinitions(definitions)
	for _, def := range resources {
		resolved, err := definitions.ResolveResourceSpecAndStatus(def)
		if err != nil {
			return nil, errors.Wrapf(err, "unable to find resource %s spec and status", def.Name())
		}

		childResourcePropertyTypeDef, err := extractChildResourcePropertyTypeDef(
			definitions,
			def.Name(),
			resolved.SpecDef.Name(),
			resolved.SpecType)
		if err != nil {
			return nil, err
		}
		if childResourcePropertyTypeDef == nil {
			continue // This just means skip
		}

		childResourceTypeNames, err := extractChildResourceTypeNames(*childResourcePropertyTypeDef)
		if err != nil {
			return nil, err
		}

		err = updateChildResourceDefinitionsWithOwner(definitions, childResourceTypeNames, def.Name(), updatedDefs)
		if err != nil {
			return nil, err
		}

		// Remove the resources property from the owning resource spec
		newDef := resolved.SpecDef.WithType(resolved.SpecType.WithoutProperty(resourcesPropertyName))
		updatedDefs[resolved.SpecDef.Name()] = newDef
	}

	setDefaultOwner(configuration, definitions, updatedDefs)

	return definitions.OverlayWith(updatedDefs), nil
}

func extractChildResourcePropertyTypeDef(
	definitions astmodel.TypeDefinitionSet,
	resourceName astmodel.TypeName,
	resourceSpecName astmodel.TypeName,
	specType *astmodel.ObjectType) (*astmodel.TypeDefinition, error) {

	// We're looking for a magical "Resources" property - if we don't find
	// one just move on
	resourcesProp, ok := specType.Property(resourcesPropertyName)
	if !ok {
		return nil, nil
	}

	// The resources property should be an array
	resourcesPropArray, ok := resourcesProp.PropertyType().(*astmodel.ArrayType)
	if !ok {
		return nil, errors.Errorf(
			"Resource %s has spec %s with Resources property whose type is %T not array",
			resourceName,
			resourceSpecName,
			resourcesProp.PropertyType())
	}

	// We're really interested in the type of this array
	resourcesPropertyTypeName, ok := resourcesPropArray.Element().(astmodel.TypeName)
	if !ok {
		return nil, errors.Errorf(
			"Resource %s has spec %s with Resources property of type array but whose inner type is not TypeName, instead being %T",
			resourceName,
			resourceSpecName,
			resourcesPropArray.Element())
	}

	resourcesDef, ok := definitions[resourcesPropertyTypeName]
	if !ok {
		return nil, errors.Errorf("couldn't find definition Resources property type %s", resourcesPropertyTypeName)
	}

	return &resourcesDef, nil
}

func resolveResourcesTypeNames(
	resourcesPropertyName astmodel.TypeName,
	resourcesPropertyType *astmodel.ObjectType) ([]astmodel.TypeName, error) {
	var results []astmodel.TypeName

	// Each property type is a subresource type
	for _, prop := range resourcesPropertyType.Properties() {
		optionalType, ok := prop.PropertyType().(*astmodel.OptionalType)
		if !ok {
			return nil, errors.Errorf(
				"OneOf type %s property %s not of type *astmodel.OptionalType",
				resourcesPropertyName.Name(),
				prop.PropertyName())
		}

		propTypeName, ok := optionalType.Element().(astmodel.TypeName)
		if !ok {
			return nil, errors.Errorf(
				"OneOf type %s optional property %s not of type astmodel.TypeName",
				resourcesPropertyName.Name(),
				prop.PropertyName())
		}
		results = append(results, propTypeName)
	}

	return results, nil
}

func extractChildResourceTypeNames(resourcesPropertyTypeDef astmodel.TypeDefinition) ([]astmodel.TypeName, error) {
	// This type should be ResourceType, or ObjectType if modelling a OneOf/AllOf
	_, isResource := resourcesPropertyTypeDef.Type().(*astmodel.ResourceType)

	resourcesPropertyTypeAsObject, ok := astmodel.AsObjectType(resourcesPropertyTypeDef.Type())
	if !isResource && !ok {
		return nil, errors.Errorf(
			"Resources property type %s was not of type *astmodel.ResourceType and didn't wrap *astmodel.ObjectType, instead %T",
			resourcesPropertyTypeDef.Name(),
			resourcesPropertyTypeDef.Type())
	}

	// Determine if this is a OneOf/AllOf
	if ok && astmodel.OneOfFlag.IsOn(resourcesPropertyTypeDef.Type()) {
		return resolveResourcesTypeNames(resourcesPropertyTypeDef.Name(), resourcesPropertyTypeAsObject)
	} else {
		return []astmodel.TypeName{resourcesPropertyTypeDef.Name()}, nil
	}
}

// this is the name we expect to see on "child resources" in the ARM JSON schema
const ChildResourceNameSuffix = "ChildResource"

func updateChildResourceDefinitionsWithOwner(
	definitions astmodel.TypeDefinitionSet,
	childResourceTypeNames []astmodel.TypeName,
	owningResourceName astmodel.TypeName,
	updatedDefs astmodel.TypeDefinitionSet) error {

	for _, typeName := range childResourceTypeNames {
		// If the typename ends in ChildResource, remove that
		if strings.HasSuffix(typeName.Name(), ChildResourceNameSuffix) {
			typeName = astmodel.MakeTypeName(typeName.PackageReference, strings.TrimSuffix(typeName.Name(), ChildResourceNameSuffix))
		}

		// If type typename is ExtensionsChild, remove Child -- this is a special case due to
		// compute...
		if typeName.Name() == "ExtensionsChild" {
			typeName = astmodel.MakeTypeName(typeName.PackageReference, strings.TrimSuffix(typeName.Name(), "Child"))
		}

		// Confirm the type really exists
		childResourceDef, ok := definitions[typeName]
		if !ok {
			return errors.Errorf("couldn't find child resource type %s", typeName)
		}

		// Update the definition of the child resource type to point to its owner
		childResource, ok := childResourceDef.Type().(*astmodel.ResourceType)
		if !ok {
			return errors.Errorf("child resource %s not of type *astmodel.ResourceType, instead %T", typeName, childResourceDef.Type())
		}

		childResourceDef = childResourceDef.WithType(childResource.WithOwner(&owningResourceName))
		if updatedDef, ok := updatedDefs[typeName]; ok {
			// already exists, make sure it is the same
			if !astmodel.TypeEquals(updatedDef.Type(), childResourceDef.Type()) {
				return errors.Errorf("conflicting child resource already defined for %s", typeName)
			}
		} else {
			updatedDefs.Add(childResourceDef)
		}
	}

	return nil
}

// setDefaultOwner sets a default owner for all resources which don't have one. The default owner is ResourceGroup.
// Extension resources have no owner set, as they are a special case.
func setDefaultOwner(
	configuration *config.Configuration,
	definitions astmodel.TypeDefinitionSet,
	updatedDefs astmodel.TypeDefinitionSet) {

	// Go over all of the resource types and flag any that don't have an owner as having resource group as their owner
	for _, def := range definitions {
		// Check if we've already modified this type - we need to use the already modified value
		if updatedDef, ok := updatedDefs[def.Name()]; ok {
			def = updatedDef
		}

		resourceType, ok := def.Type().(*astmodel.ResourceType)
		if !ok {
			continue
		}

		if resourceType.Owner() == nil && resourceType.Kind() == astmodel.ResourceKindNormal {
			ownerTypeName := astmodel.MakeTypeName(
				// Note that the version doesn't really matter here -- it's removed later. We just need to refer to the logical
				// resource group really
				configuration.MakeLocalPackageReference("resources", "v20191001"),
				"ResourceGroup")
			updatedType := resourceType.WithOwner(&ownerTypeName) // TODO: Note that right now... this type doesn't actually exist...
			// This can overwrite because a resource with no owner may have had child resources,
			// and earlier on in this process we removed the resources property from the parent resource,
			// so it may already be in updatedDefs. In this case, that's okay so we allow it to overwrite.
			updatedDefs[def.Name()] = def.WithType(updatedType)
		}
	}
}
