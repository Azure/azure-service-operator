/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package embeddedresources

import (
	"fmt"

	"github.com/go-logr/logr"
	"github.com/pkg/errors"
	kerrors "k8s.io/apimachinery/pkg/util/errors"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/config"
)

type resourceRemovalVisitorContext struct {
	resource            astmodel.TypeName
	name                astmodel.TypeName
	depth               int
	modifiedDefinitions astmodel.TypeDefinitionSet
}

func (e resourceRemovalVisitorContext) WithMoreDepth() resourceRemovalVisitorContext {
	e.depth += 1

	// Note that e.modifiedDefinitions is a pointer and so is shared between all instances
	// in order to allow tracking what definitions have been modified.
	return e
}

func (e resourceRemovalVisitorContext) WithName(name astmodel.TypeName) resourceRemovalVisitorContext {
	e.name = name
	return e
}

// EmbeddedResourceRemover uses the "x-ms-azure-resource" extension to detect resources that are embedded
// inside other resources.
// There are two different kinds of embeddings:
//  1. Peer resource embedding: This embedding allows users to create the peer resource inline rather than
//     use an ARM ID reference. An example of this embedding can be seen at
//     https://github.com/Azure/azure-rest-api-specs/blob/main/specification/network/resource-manager/Microsoft.Network/stable/2020-11-01/publicIpAddress.json#L453
//     where PublicIPAddress has a natGateway property which refers to "./natGateway.json#/definitions/NatGateway".
//     This results in the ability to define a NatGateway inline while creating a PublicIPAddress. This is possibly
//     useful in ARM templates but for our purposes would be better served with just an ID string representing the
//     NatGateway.
//  2. A subresource embedding. For the same reasons above, embedded sub-resources don't make sense in Kubernetes.
//     In the case of embedded sub-resources, the ideal shape would be a complete removal of the reference. We forbid
//     parent resources directly referencing child resources as it complicates the Watches scenario for each resource
//     reconciler. It's also not a common pattern in Kubernetes - usually you can identify children for a
//     given parent via a label. An example of this type of embedding is
//     https://github.com/Azure/azure-rest-api-specs/blob/main/specification/network/resource-manager/Microsoft.Network/stable/2020-11-01/routeTable.json#L668
//     The Routes property is of type RouteTableRoutes  which is a child resource of RouteTable.
type EmbeddedResourceRemover struct {
	definitions              astmodel.TypeDefinitionSet
	resourceToSubresourceMap map[resourceKey]astmodel.TypeNameSet
	typeSuffix               string
	typeFlag                 astmodel.TypeFlag

	// resourcesWhichOwnSubresourceLifecycle is a collection of resources which own 1 or more subresource lifecycles.
	// Examples of this include VirtualNetwork (owns Subnet) and RouteTable (owns Route). Ownership in this case means that
	// the property subnets/routes on the parent will delete subresources if that collection doesn't include all existing
	// subnets/routes.
	// This is a map of resource name to details about the owned resource
	resourcesWhichOwnSubresourceLifecycle map[astmodel.TypeName][]misbehavingResourceDetails

	// resourcesEmbeddedInParent is a collection of subresources whose IsResource() is conditional based on the context it's used in.
	// In some places, it is a resource and should be pruned. In other places, it must not be pruned. This usually boils down
	// to pseudo-resources in networking that while they look like a resource can only be created as properties on another
	// resource and so must NOT be pruned from that context as otherwise they cannot be created anywhere.
	// This map is from subresource type name to resource type name.
	resourcesEmbeddedInParent map[astmodel.TypeName]astmodel.TypeName
	renames                   map[astmodel.TypeName]embeddedResourceTypeName // A set of all the type renames made, indexed by the new name
}

// MakeEmbeddedResourceRemover creates an EmbeddedResourceRemover for the specified astmodel.TypeDefinitionSet collection.
func MakeEmbeddedResourceRemover(configuration *config.Configuration, definitions astmodel.TypeDefinitionSet) (EmbeddedResourceRemover, error) {
	resourceToSubresourceMap := findResourceSubResources(definitions)

	resourcesEmbeddedInParent, err := findResourcesEmbeddedInParent(configuration, definitions)
	if err != nil {
		return EmbeddedResourceRemover{}, errors.Wrap(err, "couldn't find all resources embedded in parent")
	}

	misbehavingResources, err := findMisbehavingResources(configuration, definitions)
	if err != nil {
		return EmbeddedResourceRemover{}, errors.Wrap(err, "couldn't find all misbehaving embedded resources")
	}

	remover := EmbeddedResourceRemover{
		definitions:                           definitions,
		resourcesWhichOwnSubresourceLifecycle: misbehavingResources,
		resourcesEmbeddedInParent:             resourcesEmbeddedInParent,
		resourceToSubresourceMap:              resourceToSubresourceMap,
		typeSuffix:                            "SubResourceEmbedded",
		typeFlag:                              astmodel.TypeFlag("embeddedSubResource"),
		renames:                               make(map[astmodel.TypeName]embeddedResourceTypeName),
	}

	return remover, nil
}

// RemoveEmbeddedResources removes any embedded resources according to the
func (e EmbeddedResourceRemover) RemoveEmbeddedResources(
	log logr.Logger,
) (astmodel.TypeDefinitionSet, error) {
	result := make(astmodel.TypeDefinitionSet)

	originalNames := make(map[astmodel.TypeName]embeddedResourceTypeName)

	visitor := e.makeEmbeddedResourceRemovalTypeVisitor()
	for _, def := range astmodel.FindResourceDefinitions(e.definitions) {
		typeWalker := e.newResourceRemovalTypeWalker(visitor, def)

		updatedTypes, err := typeWalker.Walk(def)
		if err != nil {
			return nil, err
		}

		for _, newDef := range updatedTypes {
			err := result.AddAllowDuplicates(newDef)
			if err != nil {
				return nil, err
			}
		}

		// Aggregate all renames
		for nw, og := range e.renames {
			originalNames[nw] = og
		}
	}

	result, err := simplifyTypeNames(result, e.typeFlag, originalNames)
	if err != nil {
		return nil, err
	}

	return RemoveEmptyObjects(result, log)
}

func (e EmbeddedResourceRemover) makeEmbeddedResourceRemovalTypeVisitor() astmodel.TypeVisitor {
	visitor := astmodel.TypeVisitorBuilder{
		VisitObjectType: func(this *astmodel.TypeVisitor, it *astmodel.ObjectType, ctx interface{}) (astmodel.Type, error) {
			typedCtx := ctx.(resourceRemovalVisitorContext)

			parent := e.resourcesEmbeddedInParent[typedCtx.name]
			isResourceEmbeddedInParent := astmodel.TypeEquals(parent, typedCtx.resource)
			if isResourceEmbeddedInParent {
				// Remove the ID field if there is one, and it's not a status type.
				// The expectation is that these resources must be created through
				// their parent, so you won't ever use the ID field
				if !typedCtx.name.IsStatus() {
					it = it.WithoutSpecificProperties("Id")
				}
				return astmodel.OrderedIdentityVisitOfObjectType(this, it, ctx)
			}

			// If this resource has any properties that are flagged as misbehaving embedded resources, we have to skip
			// pruning that property
			if detailsCollection, ok := e.resourcesWhichOwnSubresourceLifecycle[typedCtx.resource]; ok {
				for _, details := range detailsCollection {
					if details.propertyType == typedCtx.name {
						return astmodel.OrderedIdentityVisitOfObjectType(this, it, ctx)
					}
				}
			}

			if typedCtx.depth <= 1 || !it.IsResource() {
				// Avoid removing top level Spec
				return astmodel.OrderedIdentityVisitOfObjectType(this, it, ctx)
			}

			if subResources, ok := e.resourceToSubresourceMap[getResourceKey(typedCtx.resource)]; ok {
				isSubresource := it.IsResource() && it.Resources() != nil && subResources.ContainsAny(it.Resources())
				if subResources.Contains(typedCtx.name) || isSubresource {
					it = astmodel.EmptyObjectType // Remove this object
					return it, nil
				}
			}

			var keep []*astmodel.PropertyDefinition
			it.Properties().ForEach(func(def *astmodel.PropertyDefinition) {
				if def.HasName("Id") {
					keep = append(keep, def)
				}
			})

			it = it.WithoutProperties().WithProperties(keep...)
			return astmodel.OrderedIdentityVisitOfObjectType(this, it, ctx)
		},
	}.Build()

	return visitor
}

func (e EmbeddedResourceRemover) newResourceRemovalTypeWalker(visitor astmodel.TypeVisitor, def astmodel.TypeDefinition) *astmodel.TypeWalker {
	typeWalker := astmodel.NewTypeWalker(e.definitions, visitor)
	typeWalker.AfterVisit = func(original astmodel.TypeDefinition, updated astmodel.TypeDefinition, ctx interface{}) (astmodel.TypeDefinition, error) {
		typedCtx := ctx.(resourceRemovalVisitorContext)

		if !astmodel.TypeEquals(original.Name(), updated.Name()) {
			panic(fmt.Sprintf("Unexpected name mismatch during type walk: %q -> %q", original.Name(), updated.Name()))
		}

		if astmodel.TypeEquals(original.Type(), updated.Type()) {
			return updated, nil
		}

		flaggedType := e.typeFlag.ApplyTo(updated.Type())

		// Generate a unique TypeName for this usage.
		// A particular type may be used in multiple contexts in the same resource, or in multiple contexts in different resources. Since the pruning we are
		// doing is context specific, a single type may end up with multiple shapes after pruning. In order to cater for this possibility we generate a
		// unique name below and then collapse unneeded uniqueness away with simplifyTypeNames.
		var newName astmodel.TypeName
		var embeddedName embeddedResourceTypeName
		exists := false
		for count := 0; ; count++ {
			embeddedName = embeddedResourceTypeName{
				original: original.Name(),
				context:  typedCtx.resource.Name(),
				suffix:   e.typeSuffix,
				count:    count,
			}
			newName = embeddedName.ToTypeName()
			existing, ok := typedCtx.modifiedDefinitions[newName]
			if !ok {
				break
			}
			if astmodel.TypeEquals(existing.Type(), flaggedType) {
				exists = true
				// Shape matches what we have already, can proceed
				break
			}
		}

		e.renames[newName] = embeddedName

		updated = updated.WithName(newName)
		updated = updated.WithType(flaggedType)
		if !exists {
			typedCtx.modifiedDefinitions.Add(updated)
		}

		return updated, nil
	}

	typeWalker.ShouldRemoveCycle = func(def astmodel.TypeDefinition, ctx interface{}) (bool, error) {
		ot, ok := astmodel.AsObjectType(def.Type())
		if !ok {
			return false, nil
		}
		if ot.IsResource() {
			return true, nil
		}

		// This is here because some microsoft.networking resources are resources (in the sense that they have ARM IDs)
		// but can only be created as children of another resource. The resources in question don't have
		// their own PUT and so are not actually classified as a top level resource by the JSON schema. We don't want to
		// remove ALL cycles in the type graph currently as we can't know for sure that the cycles are structurally meaningless.
		// This is an attempt at a middle-ground heuristic that lets us find cycles between things that are resource-like.
		// For example see the cycle between NetworkInterfaceIPConfiguration in microsoft.network 20180601:
		// NetworkInterfaceIPConfiguration_Status -> NetworkInterfaceIPConfigurationPropertiesFormat_Status ->
		// ApplicationGatewayBackendAddressPool_Status -> ApplicationGatewayBackendAddressPoolPropertiesFormat_Status -> NetworkInterfaceIPConfiguration_Status
		// Sometimes these resource-like things are promoted to real resources in future APIs as in the case of Subnet in the 2017-06-01
		// API version.
		if isTypeResourceLookalike(def.Type()) {
			return true, nil
		}

		return false, nil // Leave other cycles for now
	}

	typeWalker.MakeContext = func(it astmodel.TypeName, ctx interface{}) (interface{}, error) {
		if ctx == nil {
			return resourceRemovalVisitorContext{resource: def.Name(), depth: 0, modifiedDefinitions: make(astmodel.TypeDefinitionSet)}, nil
		}
		typedCtx := ctx.(resourceRemovalVisitorContext)
		return typedCtx.WithMoreDepth().WithName(it), nil
	}

	return typeWalker
}

func findResourceSubResources(definitions astmodel.TypeDefinitionSet) map[resourceKey]astmodel.TypeNameSet {
	result := make(map[resourceKey]astmodel.TypeNameSet)

	resources := astmodel.FindResourceDefinitions(definitions)
	for _, def := range resources {
		resource, ok := astmodel.AsResourceType(def.Type())
		if !ok {
			// Shouldn't be possible to get here
			panic(fmt.Sprintf("resource was somehow not a resource: %q", def.Name()))
		}

		if resource.Owner() == nil {
			continue
		}

		owner := *resource.Owner()
		ownerKey := getResourceKey(owner)
		if result[ownerKey] == nil {
			result[ownerKey] = astmodel.NewTypeNameSet()
		}
		result[ownerKey].Add(def.Name())
	}

	return result
}

// requiredResourceProperties are properties that must be on a type for it to be considered a resource
func requiredResourceProperties() []string {
	return []string{
		"Name",
		"Properties",
	}
}

func isTypeResourceLookalike(t astmodel.Type) bool {
	o, ok := astmodel.AsObjectType(t)
	if !ok {
		return false
	}

	return isObjectResourceLookalike(o)
}

func isObjectResourceLookalike(o *astmodel.ObjectType) bool {
	hasRequiredProperties := true
	for _, propName := range requiredResourceProperties() {
		_, hasProp := o.Property(astmodel.PropertyName(propName))
		hasRequiredProperties = hasRequiredProperties && hasProp
	}

	return hasRequiredProperties
}

func findResourcesEmbeddedInParent(configuration *config.Configuration, defs astmodel.TypeDefinitionSet) (map[astmodel.TypeName]astmodel.TypeName, error) {
	result := make(map[astmodel.TypeName]astmodel.TypeName)

	var errs []error
	for name, def := range defs {
		objectType, ok := def.Type().(*astmodel.ObjectType)
		if !ok {
			continue
		}

		parentResource, err := configuration.ObjectModelConfiguration.LookupResourceEmbeddedInParent(name)
		if err != nil {
			if config.IsNotConfiguredError(err) {
				// $isResource is not configured, skip this object
				continue
			}

			// If something else went wrong, keep details
			errs = append(errs, err)
			continue
		}

		// Perform some validation that this annotation makes sense before we accept it
		if !objectType.IsResource() {
			errs = append(errs, errors.Errorf("%s is not labelled as a resource, so cannot be a resource embedded in a parent", name))
			continue
		}
		parentTypeName := name.WithName(parentResource)
		if !defs.Contains(parentTypeName) {
			errs = append(errs, errors.Errorf("cannot find %s parent %s", name, parentTypeName))
			continue
		}

		result[name] = parentTypeName
	}

	var err error
	err = kerrors.NewAggregate(errs)
	if err != nil {
		return nil, err
	}

	// Ensure that all the $isResource properties were used
	err = configuration.ObjectModelConfiguration.VerifyResourceEmbeddedInParentConsumed()
	if err != nil {
		return nil, err
	}

	return result, nil
}

type resourceKey struct {
	name  string
	group string
}

func getResourceKey(name astmodel.TypeName) resourceKey {
	group, _ := name.PackageReference.GroupVersion()

	return resourceKey{
		group: group,
		name:  name.Name(),
	}
}
