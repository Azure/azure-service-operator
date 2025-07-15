/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"context"
	"regexp"
	"strings"

	"github.com/rotisserie/eris"
	kerrors "k8s.io/apimachinery/pkg/util/errors"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/config"
)

const resourcesPropertyName = astmodel.PropertyName("Resources")

const DetermineResourceOwnershipStageID = "determineResourceOwnership"

func DetermineResourceOwnership(
	configuration *config.Configuration,
) *Stage {
	return NewStage(
		DetermineResourceOwnershipStageID,
		"Determine ARM resource relationships",
		func(ctx context.Context, state *State) (*State, error) {
			determiner := newOwnershipStage(configuration, state.Definitions())
			defs, err := determiner.assignOwners()
			if err != nil {
				return nil, eris.Wrapf(err, "failed to determine resource ownership")
			}

			return state.WithOverlaidDefinitions(defs), nil
		})
}

type ownershipStage struct {
	configuration        *config.Configuration
	definitions          astmodel.TypeDefinitionSet
	resourcesByParentURI map[string][]astmodel.InternalTypeName
}

func newOwnershipStage(
	configuration *config.Configuration,
	definitions astmodel.TypeDefinitionSet,
) *ownershipStage {
	result := &ownershipStage{
		configuration:        configuration,
		definitions:          definitions,
		resourcesByParentURI: make(map[string][]astmodel.InternalTypeName),
	}

	result.indexByParent()
	return result
}

func (o *ownershipStage) indexByParent() {
	// Index all resources by canonical URL of their parent
	for _, def := range o.definitions.AllResources() {
		rt, _ := astmodel.AsResourceType(def.Type())
		canonical := o.canonicalizeURI(rt.ARMURI())
		parent := o.uriOfParentResource(canonical)
		if parent != "" {
			o.resourcesByParentURI[parent] = append(o.resourcesByParentURI[parent], def.Name())
		}
	}
}

func (o *ownershipStage) assignOwners() (astmodel.TypeDefinitionSet, error) {
	updatedDefs := make(astmodel.TypeDefinitionSet)

	// Loop through and associate children with parents, if found
	var errs []error
	for _, def := range o.definitions.AllResources() {
		resolved, err := o.definitions.ResolveResourceSpecAndStatus(def)
		if err != nil {
			return nil, eris.Wrapf(err, "unable to find resource %s spec and status", def.Name())
		}

		childResourceTypeNames := o.findChildren(def)

		err = o.updateChildResourceDefinitionsWithOwner(childResourceTypeNames, def.Name(), updatedDefs)
		if err != nil {
			errs = append(
				errs,
				eris.Wrapf(err, "failed to update ownership for resource %s", def.Name()))
			continue
		}

		// Remove the resources property from the owning resource spec
		// TODO: Can delete this once we drop JSON schema golden files
		if _, ok := resolved.SpecType.Property(resourcesPropertyName); ok {
			// Remove the property from the Spec while preserving the structure of the original type
			remover := astmodel.NewPropertyRemover()
			newDef, err := remover.Remove(resolved.SpecDef, resourcesPropertyName)
			if err != nil {
				errs = append(
					errs,
					eris.Wrapf(err, "failed to remove resources property from resource %s", def.Name()))
				continue
			}

			updatedDefs[resolved.SpecDef.Name()] = newDef
		}
	}

	if len(errs) > 0 {
		return nil, eris.Wrapf(
			kerrors.NewAggregate(errs),
			"failed to update ownership for some resources")
	}

	o.setDefaultOwner(updatedDefs)

	return updatedDefs, nil
}

var urlParamRegex = regexp.MustCompile(`\{.*?}`)

func (o *ownershipStage) findChildren(
	def astmodel.TypeDefinition,
) []astmodel.InternalTypeName {
	rt, ok := astmodel.AsResourceType(def.Type())
	if !ok {
		return nil
	}

	resourceURI := o.canonicalizeURI(rt.ARMURI())
	return o.resourcesByParentURI[resourceURI]
}

func (*ownershipStage) canonicalizeURI(uri string) string {
	// Replace all {.*}'s with {}, in case different URIs use different names for the same
	// parameter
	uri = urlParamRegex.ReplaceAllString(uri, "{}")
	uri = strings.TrimSuffix(uri, "/")
	return strings.ToLower(uri)
}

// uriOfParentResource removes the last two segments of the URI, which are the resource type and resource name
func (*ownershipStage) uriOfParentResource(uri string) string {
	parts := strings.Split(uri, "/")
	if len(parts) < 2 {
		return ""
	}

	return strings.Join(parts[:len(parts)-2], "/")
}

func (o *ownershipStage) updateChildResourceDefinitionsWithOwner(
	childResourceTypeNames []astmodel.InternalTypeName,
	owningResourceName astmodel.InternalTypeName,
	updatedDefs astmodel.TypeDefinitionSet,
) error {
	applyOwnerToChild := func(
		owner astmodel.InternalTypeName,
		def astmodel.TypeDefinition,
		rsrc *astmodel.ResourceType,
	) {
		// Update the child resource to have the requested owner
		rsrc = rsrc.WithOwner(owner)
		def = def.WithType(rsrc)

		// We update the definition set with overwrite semantics, as we may have already added this child
		// with a worse choice for owner
		updatedDefs[def.Name()] = def
	}

	for _, typeName := range childResourceTypeNames {
		// Use the singular form of the name
		typeName = typeName.Singular()

		// Confirm the child type really exists by finding its definition.
		// Look in updatedDefs first, as we may have already modified it
		childResourceDef, ok := updatedDefs[typeName]
		if !ok {
			// If not found, look in the original definitions
			childResourceDef, ok = o.definitions[typeName]
			if !ok {
				return eris.Errorf("couldn't find child resource type %s", typeName)
			}
		}

		childResource, ok := childResourceDef.Type().(*astmodel.ResourceType)
		if !ok {
			return eris.Errorf("child resource %s not of type *astmodel.ResourceType, instead %T", typeName, childResourceDef.Type())
		}

		// Ownership transcends APIVersion, but in order for things like $exportAs to work, it's best if
		// ownership for each resource points to the owner in the same package. This ensures that standard tools
		// like renamingVisitor work.
		//
		// Thus, we prefer to set owners within the same package as the child if we can.
		// If we must have an owner in a different package, we need to disambiguate by choosing the one in the
		// newest package so that our order of processing produces a consistent result.
		//
		// Two examples for why we need to do this:
		//  (i) StorSimple has the ame URIs on "different" types (different potential parents for
		//     the same child), so we disambiguate in favour of the one with the same package.
		// (ii) Insights don't publish all resources in all versions, so we have parents from
		//      different versions than their children (e.g. Component vs PricingPlan)

		// If the child doesn't have an owner, set it to the candidate we have
		existingOwner := childResource.Owner()
		if existingOwner.IsEmpty() {
			applyOwnerToChild(owningResourceName, childResourceDef, childResource)
			continue
		}

		// if the child already has an owner in the same package as the child,
		// keep that owner unchanged
		if existingOwner.PackageReference().Equals(typeName.PackageReference()) {
			continue
		}

		// If the candidate new owner is in the same package as the child, use that
		if owningResourceName.InternalPackageReference().Equals(typeName.InternalPackageReference()) {
			applyOwnerToChild(owningResourceName, childResourceDef, childResource)
			continue
		}

		// Given two candidate owners, neither of which is in the same package as the child,
		// we need to choose the one with the most recent package reference.
		if astmodel.ComparePathAndVersion(
			existingOwner.InternalPackageReference().ImportPath(),
			owningResourceName.InternalPackageReference().ImportPath()) < 0 {
			// New owner is more recent, so use it
			applyOwnerToChild(owningResourceName, childResourceDef, childResource)
			continue
		}
	}

	return nil
}

// setDefaultOwner sets a default owner for all resources which don't have one. The default owner is ResourceGroup.
// Extension resources have no owner set, as they are a special case.
func (o *ownershipStage) setDefaultOwner(
	updatedDefs astmodel.TypeDefinitionSet,
) {
	// Go over all the resource types and flag any that don't have an owner as having resource group as their owner
	for _, def := range o.definitions {
		// Check if we've already modified this type - we need to use the already modified value
		if updatedDef, ok := updatedDefs[def.Name()]; ok {
			def = updatedDef
		}

		resourceType, ok := def.Type().(*astmodel.ResourceType)
		if !ok {
			continue
		}

		if resourceType.Owner().IsEmpty() && resourceType.Scope() == astmodel.ResourceScopeResourceGroup {
			ownerTypeName := astmodel.MakeInternalTypeName(
				// Note that the version doesn't really matter here -- it's removed later. We just need to refer to the logical
				// resource group really
				o.configuration.MakeLocalPackageReference("resources", "v20191001"),
				"ResourceGroup")
			updatedType := resourceType.WithOwner(ownerTypeName) // TODO: Note that right now... this type doesn't actually exist...
			// This can overwrite because a resource with no owner may have had child resources,
			// and earlier on in this process we removed the resources property from the parent resource,
			// so it may already be in updatedDefs. In this case, that's okay so we allow it to overwrite.
			updatedDefs[def.Name()] = def.WithType(updatedType)
		}
	}
}
