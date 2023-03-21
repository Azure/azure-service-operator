/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"context"
	"regexp"
	"strings"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/config"

	"github.com/pkg/errors"
)

const resourcesPropertyName = astmodel.PropertyName("Resources")

const DetermineResourceOwnershipStageId = "determineResourceOwnership"

func DetermineResourceOwnership(
	configuration *config.Configuration,
	idFactory astmodel.IdentifierFactory,
) *Stage {
	return NewLegacyStage(
		DetermineResourceOwnershipStageId,
		"Determine ARM resource relationships",
		func(ctx context.Context, definitions astmodel.TypeDefinitionSet) (astmodel.TypeDefinitionSet, error) {
			return determineOwnership(definitions, configuration, idFactory)
		})
}

func determineOwnership(
	definitions astmodel.TypeDefinitionSet,
	configuration *config.Configuration,
	idFactory astmodel.IdentifierFactory,
) (astmodel.TypeDefinitionSet, error) {
	updatedDefs := make(astmodel.TypeDefinitionSet)

	resources := astmodel.FindResourceDefinitions(definitions)
	for _, def := range resources {
		resolved, err := definitions.ResolveResourceSpecAndStatus(def)
		if err != nil {
			return nil, errors.Wrapf(err, "unable to find resource %s spec and status", def.Name())
		}

		rt := def.Type().(*astmodel.ResourceType)
		childResourceTypeNames := findChildren(rt, def.Name(), resources)

		err = updateChildResourceDefinitionsWithOwner(definitions, childResourceTypeNames, def.Name(), updatedDefs, idFactory)
		if err != nil {
			return nil, err
		}

		// Remove the resources property from the owning resource spec
		// TODO: Can delete this once we drop JSON schema golden files
		newDef := resolved.SpecDef.WithType(resolved.SpecType.WithoutProperty(resourcesPropertyName))
		updatedDefs[resolved.SpecDef.Name()] = newDef
	}

	setDefaultOwner(configuration, definitions, updatedDefs)

	return definitions.OverlayWith(updatedDefs), nil
}

var urlParamRegex = regexp.MustCompile("\\{.*?}")

func findChildren(rt *astmodel.ResourceType, resourceName astmodel.TypeName, others astmodel.TypeDefinitionSet) []astmodel.TypeName {
	// append "/" to the ARM URI so that if this is (e.g.):
	//     /resource/name
	// it doesn't match as a prefix of:
	//     /resource/namedThing
	// but it is a prefix of:
	//     /resource/name/subresource/subname
	myPrefix := rt.ARMURI() + "/"
	myPrefix = canonicalizeURI(myPrefix)

	var result []astmodel.TypeName
	for otherName, otherDef := range others {
		other, ok := astmodel.AsResourceType(otherDef.Type())
		if !ok {
			continue
		}
		if rt == other {
			continue // don’t self-own
		}

		// TODO: If it ever arises that we have a resource whose owner doesn't exist in the same API
		// TODO: version this might be an issue.
		// Ownership transcends APIVersion, but in order for things like $exportAs to work, it's best if
		// ownership for each resource points to the owner in the same package. This ensures that standard tools
		// like renamingVisitor work.
		if !otherDef.Name().PackageReference.Equals(resourceName.PackageReference) {
			continue // Don't own if in a different package
		}

		otherURI := canonicalizeURI(other.ARMURI())

		// Compare case-insensitive in case specs have different URL casings in their Swagger
		if strings.HasPrefix(strings.ToLower(otherURI), strings.ToLower(myPrefix)) {
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

	return result
}

func canonicalizeURI(uri string) string {
	// Replace all {.*}'s with {}, in case different URIs use different names for the same
	// parameter
	uri = urlParamRegex.ReplaceAllString(uri, "{}")
	return uri
}

func updateChildResourceDefinitionsWithOwner(
	definitions astmodel.TypeDefinitionSet,
	childResourceTypeNames []astmodel.TypeName,
	owningResourceName astmodel.TypeName,
	updatedDefs astmodel.TypeDefinitionSet,
	idFactory astmodel.IdentifierFactory,
) error {

	for _, typeName := range childResourceTypeNames {
		// Use the singular form of the name
		typeName = typeName.Singular()

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
	// Go over all the resource types and flag any that don't have an owner as having resource group as their owner
	for _, def := range definitions {
		// Check if we've already modified this type - we need to use the already modified value
		if updatedDef, ok := updatedDefs[def.Name()]; ok {
			def = updatedDef
		}

		resourceType, ok := def.Type().(*astmodel.ResourceType)
		if !ok {
			continue
		}

		if resourceType.Owner() == nil && resourceType.Scope() == astmodel.ResourceScopeResourceGroup {
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
