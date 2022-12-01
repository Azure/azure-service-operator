/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package embeddedresources

import (
	"fmt"

	"github.com/pkg/errors"
	kerrors "k8s.io/apimachinery/pkg/util/errors"
	"k8s.io/klog/v2"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/config"
)

type resourceRemovalVisitorContext struct {
	resource            astmodel.TypeName
	depth               int
	modifiedDefinitions astmodel.TypeDefinitionSet
}

func (e resourceRemovalVisitorContext) WithMoreDepth() resourceRemovalVisitorContext {
	e.depth += 1

	// Note that e.modifiedDefinitions is a pointer and so is shared between all instances
	// in order to allow tracking what definitions have been modified.
	return e
}

// EmbeddedResourceRemover uses a variety of heuristics to remove resources that are embedded inside other resources.
// There are a number of different kinds of embeddings:
//
// 1. A "Properties" embedding. When we process the Azure JSON schema/Swagger we manufacture a "Spec"
// type that doesn't exist in the JSON schema/Swagger. In the JSON schema the resource itself must comply with ARM
// resource requirements, meaning that all of the RP specific properties are stored in the "Properties"
// property which for the sake of example we will say has type "R1Properties".
// Other resources which have a property somewhere in their type hierarchy with that same "R1Properties"
// type are actually embedding the R1 resource entirely inside themselves. Since the R1 resource is its own
// resource it doesn't make sense to have it embedded inside another resource in Kubernetes. These embeddings
// should really just be cross resource references. This pipeline finds such embeddings and removes them. A concrete
// example of one such embedding is
// v20181001 Microsoft.Networking Connection.Spec.Properties.LocalNetworkGateway2.Properties.
// The LocalNetworkGateway2 property is of type "LocalNetworkGateway" which is itself a resource.
// The ideal shape of Connection.Spec.Properties.LocalNetworkGate2 would just be a reference to a
// LocalNetworkGateway resource.
//
// 2. A subresource embedding. For the same reasons above, embedded subresources don't make sense in Kubernetes.
// In the case of embedded subresources, the ideal shape would be a complete removal of the reference. We forbid
// parent resources directly referencing child resources as it complicates the Watches scenario for each resource
// reconciler. It's also not a common pattern in Kubernetes - usually you can identify children for a
// given parent via a label. An example of this type of embedding is
// v20180601 Microsoft.Networking RouteTable.Spec.Properties.Routes. The Routes property is of type RouteTableRoutes
// which is a child resource of RouteTable.
//
// Note that even though the above examples do not include Status definitions, the same rules apply to Status definitions, with
// the only difference being that for Status definitions the resource reference in Swagger (the source of the Status definitions)
// is to the Status type (as opposed to the "Properties" type for Spec).
type EmbeddedResourceRemover struct {
	definitions                  astmodel.TypeDefinitionSet
	resourceToSubresourceMap     map[astmodel.TypeName]astmodel.TypeNameSet
	resourcePropertiesTypes      astmodel.TypeNameSet
	resourceStatusTypes          astmodel.TypeNameSet
	typeSuffix                   string
	typeFlag                     astmodel.TypeFlag
	misbehavingEmbeddedResources astmodel.TypeNameSet
	renames                      map[astmodel.TypeName]embeddedResourceTypeName // A set of all the type renames made, indexed by the new name
}

// MakeEmbeddedResourceRemover creates an EmbeddedResourceRemover for the specified astmodel.TypeDefinitionSet collection.
func MakeEmbeddedResourceRemover(configuration *config.Configuration, definitions astmodel.TypeDefinitionSet) (EmbeddedResourceRemover, error) {
	resourceStatusTypeNames := findAllResourceStatusTypes(definitions)
	resourceToSubresourceMap, err := findSubResourcePropertiesTypeNames(definitions)
	if err != nil {
		return EmbeddedResourceRemover{}, errors.Wrap(err, "couldn't find subresource \"Properties\" type names")
	}

	resourcePropertiesTypes, err := findAllResourcePropertiesTypes(definitions)
	if err != nil {
		return EmbeddedResourceRemover{}, errors.Wrap(err, "couldn't find resource \"Properties\" type names")
	}

	misbehavingResources, err := findMisbehavingResources(configuration, definitions)
	if err != nil {
		return EmbeddedResourceRemover{}, errors.Wrap(err, "couldn't find all misbehaving embedded resources")
	}

	remover := EmbeddedResourceRemover{
		definitions:                  definitions,
		misbehavingEmbeddedResources: misbehavingResources,
		resourceToSubresourceMap:     resourceToSubresourceMap,
		resourcePropertiesTypes:      resourcePropertiesTypes,
		resourceStatusTypes:          resourceStatusTypeNames,
		typeSuffix:                   "SubResourceEmbedded",
		typeFlag:                     astmodel.TypeFlag("embeddedSubResource"), // TODO: Instead of flag we could just use a map here if we wanted
		renames:                      make(map[astmodel.TypeName]embeddedResourceTypeName),
	}

	return remover, nil
}

// RemoveEmbeddedResources removes any embedded resources according to the
func (e EmbeddedResourceRemover) RemoveEmbeddedResources() (astmodel.TypeDefinitionSet, error) {
	result := make(astmodel.TypeDefinitionSet)

	originalNames := make(map[astmodel.TypeName]embeddedResourceTypeName)

	visitor := e.makeEmbeddedResourceRemovalTypeVisitor()
	for _, def := range astmodel.FindResourceDefinitions(e.definitions) {

		// If this resource has any properties that are flagged as misbehaving embedded resources, we have to skip
		// validation
		if e.misbehavingEmbeddedResources.Contains(def.Name()) {
			tempWalker := astmodel.NewTypeWalker(e.definitions, astmodel.TypeVisitorBuilder{}.Build())
			updatedTypes, err := tempWalker.Walk(def)
			if err != nil {
				return nil, err
			}

			err = result.AddTypesAllowDuplicates(updatedTypes)
			if err != nil {
				return nil, err
			}
			continue
		}

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

	return RemoveEmptyObjects(result)
}

func (e EmbeddedResourceRemover) makeEmbeddedResourceRemovalTypeVisitor() astmodel.TypeVisitor {
	visitor := astmodel.TypeVisitorBuilder{
		VisitObjectType: func(this *astmodel.TypeVisitor, it *astmodel.ObjectType, ctx interface{}) (astmodel.Type, error) {
			typedCtx := ctx.(resourceRemovalVisitorContext)

			if typedCtx.depth <= 2 {
				// Avoid removing top level "Properties", which we never want to do. This is needed because
				// there are some resources (such as Microsoft.Web v20160801 Sites) where the resource
				// and some child resources reuse the same "Properties" type. This causes
				// the logic below to think that a resource is its own subresource. Without this
				// check the entire resource would be removed, leaving nothing.
				return astmodel.OrderedIdentityVisitOfObjectType(this, it, ctx)
			}

			// TODO: This is confusing...?
			// Before visiting, check if any properties are just referring to one of our sub-resources and remove them
			subResources := e.resourceToSubresourceMap[typedCtx.resource]
			it.Properties().ForEach(func(prop *astmodel.PropertyDefinition) {
				propTypeName, ok := astmodel.AsTypeName(prop.PropertyType())
				if !ok {
					return // continue
				}

				// TODO: This is currently no different than the below, but it likely will evolve to be different over time
				if subResources.Contains(propTypeName) {
					klog.V(5).Infof("Removing resource %q reference to subresource %q on property %q", typedCtx.resource, propTypeName, prop.PropertyName())
					it = removeResourceLikeProperties(it)
					return // continue
				}

				if e.resourcePropertiesTypes.Contains(propTypeName) {
					klog.V(5).Infof("Removing reference to resource %q on property %q", propTypeName, prop.PropertyName())
					it = removeResourceLikeProperties(it)
				}
			})

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

		klog.V(5).Infof("Updating %q to %q", original.Name(), updated.Name())

		return updated, nil
	}

	typeWalker.ShouldRemoveCycle = func(def astmodel.TypeDefinition, ctx interface{}) (bool, error) {
		// If we're about to walk a cycle that is to a known resource type, just skip it entirely
		if e.resourcePropertiesTypes.Contains(def.Name()) || e.resourceStatusTypes.Contains(def.Name()) {
			return true, nil
		}

		// TODO: Should this be replaced with a hardcoded list of resources (since most offending resources are in networking?)
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
			klog.V(5).Infof("Type %q is a resource lookalike", def.Name())
			return true, nil
		}

		return false, nil // Leave other cycles for now
	}

	typeWalker.MakeContext = func(it astmodel.TypeName, ctx interface{}) (interface{}, error) {
		if ctx == nil {
			return resourceRemovalVisitorContext{resource: def.Name(), depth: 0, modifiedDefinitions: make(astmodel.TypeDefinitionSet)}, nil
		}
		typedCtx := ctx.(resourceRemovalVisitorContext)
		return typedCtx.WithMoreDepth(), nil
	}

	return typeWalker
}

// findSubResourcePropertiesTypeNames finds the "Properties" type of each subresource and returns a map of
// parent resource to subresource "Properties" type names.
func findSubResourcePropertiesTypeNames(definitions astmodel.TypeDefinitionSet) (map[astmodel.TypeName]astmodel.TypeNameSet, error) {
	var errs []error
	result := make(map[astmodel.TypeName]astmodel.TypeNameSet)

	// Identify sub-resources and their "properties", associate them with parent resource
	// Look through parent resource for subresource properties
	for _, def := range astmodel.FindResourceDefinitions(definitions) {
		resource, ok := astmodel.AsResourceType(def.Type())
		if !ok {
			// Shouldn't be possible to get here
			panic(fmt.Sprintf("resource was somehow not a resource: %q", def.Name()))
		}

		if resource.Owner() == nil {
			continue
		}

		owner := *resource.Owner()
		specPropertiesTypeName, statusPropertiesTypeName, err := tryResolveSpecStatusTypes(definitions, resource)
		if err != nil {
			errs = append(errs, errors.Wrapf(err, "couldn't extract spec/status properties from %q", def.Name()))
			continue
		}

		if specPropertiesTypeName != nil {
			if result[owner] == nil {
				result[owner] = astmodel.NewTypeNameSet(*specPropertiesTypeName)
			} else {
				result[owner].Add(*specPropertiesTypeName)
			}
		}

		if statusPropertiesTypeName != nil {
			if result[owner] == nil {
				result[owner] = astmodel.NewTypeNameSet(*statusPropertiesTypeName)
			} else {
				result[owner].Add(*statusPropertiesTypeName)
			}
		}
	}

	err := kerrors.NewAggregate(errs)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// TODO: Move this to resourceType?
func tryResolveSpecStatusTypes(definitions astmodel.TypeDefinitionSet, resource *astmodel.ResourceType) (*astmodel.TypeName, *astmodel.TypeName, error) {
	specName, ok := astmodel.AsTypeName(resource.SpecType())
	if !ok {
		return nil, nil, errors.Errorf("resource spec was not a TypeName")
	}

	specPropertiesTypeName, err := extractPropertiesType(definitions, specName)
	if err != nil {
		return nil, nil, errors.Wrap(err, "couldn't extract spec properties")
	}

	var statusPropertiesTypeName *astmodel.TypeName
	statusName, ok := astmodel.AsTypeName(resource.StatusType())
	if ok {
		statusPropertiesTypeName, err = extractPropertiesType(definitions, statusName)
		if err != nil {
			return nil, nil, errors.Wrap(err, "couldn't extract status properties")
		}
	}

	return specPropertiesTypeName, statusPropertiesTypeName, nil
}

// findAllResourcePropertiesTypes finds the "Properties" type for each resource. The result is a astmodel.TypeNameSet containing
// each resources "Properties" type.
func findAllResourcePropertiesTypes(definitions astmodel.TypeDefinitionSet) (astmodel.TypeNameSet, error) {
	var errs []error
	result := astmodel.NewTypeNameSet()

	// Identify sub-resources and their "properties", associate them with parent resource
	// Look through parent resource for subresource properties
	for _, def := range astmodel.FindResourceDefinitions(definitions) {
		resource, ok := astmodel.AsResourceType(def.Type())
		if !ok {
			// Shouldn't be possible to get here
			panic(fmt.Sprintf("resource was somehow not a resource: %q", def.Name()))
		}

		specPropertiesTypeName, statusPropertiesTypeName, err := tryResolveSpecStatusTypes(definitions, resource)
		if err != nil {
			errs = append(errs, errors.Wrapf(err, "couldn't extract spec/status properties from %q", def.Name()))
			continue
		}

		if specPropertiesTypeName != nil {
			result.Add(*specPropertiesTypeName)
		}

		if statusPropertiesTypeName != nil {
			result.Add(*statusPropertiesTypeName)
		}
	}

	err := kerrors.NewAggregate(errs)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// findAllResourceStatusTypes finds the astmodel.TypeName's of each resources Status type. If the resource does not have a Status type then
// that TypeName is not included in the resulting astmodel.TypeNameSet (obviously).
func findAllResourceStatusTypes(definitions astmodel.TypeDefinitionSet) astmodel.TypeNameSet {
	resources := definitions.Where(func(def astmodel.TypeDefinition) bool {
		_, ok := astmodel.AsResourceType(def.Type())
		return ok
	})

	result := astmodel.NewTypeNameSet()

	// Identify sub-resources and their "properties", associate them with parent resource
	// Look through parent resource for subresource properties
	for _, def := range resources {
		resource, ok := astmodel.AsResourceType(def.Type())
		if !ok {
			// Shouldn't be possible to get here
			panic(fmt.Sprintf("resource was somehow not a resource: %q", def.Name()))
		}

		statusName, ok := astmodel.AsTypeName(resource.StatusType())
		if !ok {
			continue
		}

		result.Add(statusName)
	}

	return result
}

func extractPropertiesType(definitions astmodel.TypeDefinitionSet, typeName astmodel.TypeName) (*astmodel.TypeName, error) {
	resolved, err := definitions.GetDefinition(typeName)
	if err != nil {
		return nil, err
	}

	ot, ok := astmodel.AsObjectType(resolved.Type())
	if !ok {
		return nil, errors.Errorf("couldn't find object type %q", typeName)
	}

	propertiesProp, ok := ot.Property("Properties")
	if !ok {
		return nil, nil
	}

	propertiesTypeName, ok := astmodel.AsTypeName(propertiesProp.PropertyType())
	if !ok {
		return nil, nil
	}

	return &propertiesTypeName, nil
}

// requiredResourceProperties are properties that must be on a type for it to be considered a resource
func requiredResourceProperties() []string {
	return []string{
		"Name",
		"Properties",
	}
}

// optionalResourceProperties are properties which may or may not be on a resource. Technically all resources
// should have all of these properties, but because we drop the top-level AllOf that joins resource definitions with
// ResourceBase when parsing schemas sometimes they aren't defined.
func optionalResourceProperties() []string {
	return []string{
		"Type",
		"Etag",
		"Location",
		"Tags",
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

// removeResourceLikeProperties examines an astmodel.ObjectType and determines if it looks like an Azure resource.
// An object is "like" a resource if it has "name" and "properties" properties.
func removeResourceLikeProperties(o *astmodel.ObjectType) *astmodel.ObjectType {
	if !isObjectResourceLookalike(o) {
		// Doesn't match the shape we're looking for -- no change
		return o
	}

	result := o
	required := requiredResourceProperties()
	optional := optionalResourceProperties()

	for _, propName := range append(required, optional...) {
		result = result.WithoutProperty(astmodel.PropertyName(propName))
	}
	return result
}
