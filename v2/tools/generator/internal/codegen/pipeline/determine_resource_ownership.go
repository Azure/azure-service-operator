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

type latestNameKey struct {
	name  string
	group string
	// no version
}

func getLatestNameKey(tn astmodel.TypeName) latestNameKey {
	return latestNameKey{
		name:  tn.Name(),
		group: tn.PackageReference.(astmodel.LocalPackageReference).Group(),
	}
}

func isNewer(left, right astmodel.TypeDefinition) bool {
	leftVersion := left.Name().PackageReference.(astmodel.LocalPackageReference).ApiVersion()
	rightVersion := right.Name().PackageReference.(astmodel.LocalPackageReference).ApiVersion()

	// versions are ASCIIbetically ordered
	return leftVersion > rightVersion
}

func determineOwnership(definitions astmodel.TypeDefinitionSet, configuration *config.Configuration) (astmodel.TypeDefinitionSet, error) {
	updatedDefs := make(astmodel.TypeDefinitionSet)

	resources := astmodel.FindResourceDefinitions(definitions)
	latestResources := make(map[latestNameKey]astmodel.TypeDefinition)
	for name, resource := range resources {
		key := getLatestNameKey(name)
		if def, ok := latestResources[key]; ok {
			if isNewer(resource, def) {
				latestResources[key] = resource
			}
		} else {
			latestResources[key] = resource
		}
	}

	for _, def := range latestResources {
		resolved, err := definitions.ResolveResourceSpecAndStatus(def)
		if err != nil {
			return nil, errors.Wrapf(err, "unable to find resource %s spec and status", def.Name())
		}

		me := def.Type().(*astmodel.ResourceType)

		childResourceTypeNames := findChildren(me, resources)

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

func findChildren(rt *astmodel.ResourceType, others astmodel.TypeDefinitionSet) []astmodel.TypeName {

	// append "/" to the ARM URI so that if this is (e.g.):
	//     /resource/name
	// it doesn't match as a prefix of:
	//     /resource/namedThing
	// but it is a prefix of:
	//     /resource/name/subresource/subname
	myPrefix := rt.ARMURI() + "/"

	var result []astmodel.TypeName
	for otherName, otherDef := range others {
		if other, ok := astmodel.AsResourceType(otherDef.Type()); ok {
			if rt == other {
				continue // don’t self-own
			}

			otherURI := other.ARMURI()
			if strings.HasPrefix(otherURI, myPrefix) {
				// now, accept it only if it contains two '/' exactly:
				// so that the string is of the form:
				//     {prefix}/resourceType/resourceName
				// and not a grandchild resource:
				//     {prefix}/resourceType/resourceName/anotherResourceType/anotherResourceName
				withoutPrefix := otherURI[len(myPrefix)-1:]
				if strings.Count(withoutPrefix, "/") == 2 {
					result = append(result, otherName)
				}
			}
		}
	}

	return result
}

// this is the name we expect to see on "child resources" in the ARM JSON schema
const ChildResourceNameSuffix = "ChildResource"

func updateChildResourceDefinitionsWithOwner(
	definitions astmodel.TypeDefinitionSet,
	childResourceTypeNames []astmodel.TypeName,
	owningResourceName astmodel.TypeName,
	updatedDefs astmodel.TypeDefinitionSet,
) error {
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
		err := updatedDefs.AddAllowDuplicates(childResourceDef)
		if err != nil {
			// workaround: StorSimple has the same URIs on multiple "different" types
			// resolve in favour of the one that has a matching package
			if childResourceDef.Name().PackageReference.Equals(owningResourceName.PackageReference) {
				// override
				updatedDefs[childResourceDef.Name()] = childResourceDef
				continue // okay!
			} else {
				// double-check that existing one matches
				existingDef := updatedDefs[childResourceDef.Name()]
				rt := existingDef.Type().(*astmodel.ResourceType)
				if existingDef.Name().PackageReference.Equals(rt.Owner().PackageReference) {
					continue // okay!
				}
			}

			return errors.Wrapf(err, "conflicting child resource already defined for %s [%s]", typeName, childResource.ARMURI())
		}
	}

	return nil
}

// setDefaultOwner sets a default owner for all resources which don't have one. The default owner is ResourceGroup.
// Extension resources have no owner set, as they are a special case.
func setDefaultOwner(
	configuration *config.Configuration,
	definitions astmodel.TypeDefinitionSet,
	updatedDefs astmodel.TypeDefinitionSet,
) {
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
