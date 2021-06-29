/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package storage

import (
	"github.com/pkg/errors"
	"k8s.io/klog/v2"

	"github.com/Azure/azure-service-operator/hack/generator/pkg/astmodel"
	"github.com/Azure/azure-service-operator/hack/generator/pkg/conversions"
	"github.com/Azure/azure-service-operator/hack/generator/pkg/functions"
)

// StorageTypeFactory is used to create storage types for a specific api group
type StorageTypeFactory struct {
	group             string                     // Name of the group we're handling (used mostly for logging)
	types             astmodel.Types             // All the types for this group
	specTypes         astmodel.TypeNameSet       // All the names of spec types
	statusTypes       astmodel.TypeNameSet       // All the names of status types
	idFactory         astmodel.IdentifierFactory // Factory for creating identifiers
	typeConverter     *TypeConverter             // a utility type type visitor used to create storage variants
	functionInjector  *FunctionInjector          // a utility used to inject functions into definitions
	resourceHubMarker *HubVersionMarker          // a utility used to mark resources as Storage Versions
	conversionGraph   *GroupConversionGraph      // Map of conversion links for creating our conversion graph
	processed         bool                       // Flag to track whether we've done all our processing or not
}

// NewStorageTypeFactory creates a new instance of StorageTypeFactory ready for use
func NewStorageTypeFactory(group string, idFactory astmodel.IdentifierFactory) *StorageTypeFactory {

	types := make(astmodel.Types)

	result := &StorageTypeFactory{
		group:             group,
		types:             types,
		specTypes:         astmodel.NewTypeNameSet(),
		statusTypes:       astmodel.NewTypeNameSet(),
		idFactory:         idFactory,
		conversionGraph:   NewGroupConversionGraph(group),
		functionInjector:  NewFunctionInjector(),
		resourceHubMarker: NewHubVersionMarker(),
		typeConverter:     NewTypeConverter(types, idFactory),
	}

	return result
}

// Add the supplied type definition to this factory
func (f *StorageTypeFactory) Add(definitions ...astmodel.TypeDefinition) {
	for _, def := range definitions {
		f.types.Add(def)

		if rt, ok := astmodel.AsResourceType(def.Type()); ok {
			// We have a resource type
			if tn, ok := astmodel.AsTypeName(rt.SpecType()); ok {
				// Keep track of all our spec types
				f.specTypes.Add(tn)
			}
			if tn, ok := astmodel.AsTypeName(rt.StatusType()); ok {
				// Keep track of all our status types
				f.statusTypes.Add(tn)
			}
		}
	}
}

// Types returns types contained by the factory, including all new storage variants and modified
// api types. If any errors occur during processing, they're returned here.
func (f *StorageTypeFactory) Types() (astmodel.Types, error) {

	if !f.processed {
		err := f.process()
		f.processed = true
		if err != nil {
			return nil, err
		}
	}

	return f.types, nil
}

// process carries out our transformations
// Each step reads from outputTypes and puts the results back in there
func (f *StorageTypeFactory) process() error {

	// Inject OriginalVersion() methods into all our spec types and update f.types with the new definitions
	modifiedSpecTypes, err := f.types.Process(f.injectOriginalVersionMethod)
	if err != nil {
		return err
	}
	f.types = f.types.OverlayWith(modifiedSpecTypes)

	// Create Storage Variants (injectConversions will look for them) and add them to f.types
	storageVariants, err := f.types.Process(f.createStorageVariant)
	if err != nil {
		return err
	}
	f.types.AddTypes(storageVariants)

	// Inject conversion functions where required and update f.types with the new definitions
	typesWithConversions, err := f.types.Process(f.injectConversions)
	if err != nil {
		return err
	}
	f.types = f.types.OverlayWith(typesWithConversions)

	return nil
}

// createStorageVariant takes an existing object definition and creates a storage variant in a
// related package.
// def is the api definition on which to base the storage variant
// visitor is a type visitor that will do the creation
func (f *StorageTypeFactory) createStorageVariant(definition astmodel.TypeDefinition) (*astmodel.TypeDefinition, error) {
	name := definition.Name()

	_, isObject := astmodel.AsObjectType(definition.Type())
	_, isResource := astmodel.AsResourceType(definition.Type())
	if !isObject && !isResource {
		// just skip it
		klog.V(4).Infof("Skipping %s as no storage variant needed", name)
		return nil, nil
	}

	klog.V(3).Infof("Creating storage variant of %s", name)

	var storageDef astmodel.TypeDefinition
	var err error
	if isResource {
		storageDef, err = f.typeConverter.ConvertResourceDefinition(definition)
		if err != nil {
			return nil, errors.Wrapf(err, "creating storage variant for resource %q", name)
		}
	} else if f.specTypes.Contains(definition.Name()) {
		storageDef, err = f.typeConverter.ConvertSpecDefinition(definition)
		if err != nil {
			return nil, errors.Wrapf(err, "creating storage variant for spec %q", name)
		}
	} else if f.statusTypes.Contains(definition.Name()) {
		storageDef, err = f.typeConverter.ConvertStatusDefinition(definition)
		if err != nil {
			return nil, errors.Wrapf(err, "creating storage variant for status %q", name)
		}
	} else if isObject {
		storageDef, err = f.typeConverter.ConvertObjectDefinition(definition)
		if err != nil {
			return nil, errors.Wrapf(err, "creating storage variant for object %q", name)
		}
	}

	// Add API-Package -> Storage-Package link into the conversion map
	f.conversionGraph.AddLink(name.PackageReference, storageDef.Name().PackageReference)

	return &storageDef, nil
}

// injectConversions modifies the named type by injecting the required conversion methods using
// the conversionMap we've previously established
func (f *StorageTypeFactory) injectConversions(definition astmodel.TypeDefinition) (*astmodel.TypeDefinition, error) {
	name := definition.Name()
	_, isObject := astmodel.AsObjectType(definition.Type())
	_, isResource := astmodel.AsResourceType(definition.Type())
	if !isObject && !isResource {
		// just skip it
		klog.V(4).Infof("Skipping %s as no conversion functions needed", name)
		return nil, nil
	}

	klog.V(3).Infof("Injecting conversion functions into %s", name)

	// Find the definition we want to convert to/from
	nextPackage, ok := f.conversionGraph.LookupTransition(name.PackageReference)
	if !ok {
		// No next package, so nothing to do
		// (this is expected if we have the hub storage package)
		// Flag the type as needing to be flagged as the storage version
		//TODO: Restore this - currently disabled until we get all the conversion functions injected
		//hubDefintion, err := f.resourceHubMarker.markResourceAsStorageVersion(definition)
		//if err != nil {
		//	return nil, errors.Wrapf(err, "marking %q as hub version", name)
		//}
		//return &hubDefintion, nil
		return nil, nil
	}

	nextName := astmodel.MakeTypeName(nextPackage, name.Name())
	nextDef, ok := f.types[nextName]
	if !ok {
		// No next type so nothing to do
		// (this is expected if the type is discontinued or we're looking at the hub type)
		return nil, nil
	}

	return f.injectConversionsBetween(definition, nextDef)
}

// inject conversion methods between the two specified definitions
// upstreamDef is the definition further away from our hub type in our directed conversion graph
// downstreamDef is the definition closer to our hub type in our directed conversion graph
func (f *StorageTypeFactory) injectConversionsBetween(
	upstreamDef astmodel.TypeDefinition, downstreamDef astmodel.TypeDefinition) (*astmodel.TypeDefinition, error) {

	// Create conversion functions
	conversionContext := conversions.NewPropertyConversionContext(f.types, f.idFactory)

	assignFromFn, err := conversions.NewPropertyAssignmentFromFunction(upstreamDef, downstreamDef, f.idFactory, conversionContext)
	upstreamName := upstreamDef.Name()
	if err != nil {
		return nil, errors.Wrapf(err, "creating AssignFrom() function for %q", upstreamName)
	}

	assignToFn, err := conversions.NewPropertyAssignmentToFunction(upstreamDef, downstreamDef, f.idFactory, conversionContext)
	if err != nil {
		return nil, errors.Wrapf(err, "creating AssignTo() function for %q", upstreamName)
	}

	updatedDefinition, err := f.functionInjector.Inject(upstreamDef, assignFromFn)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to inject %s function into %q", assignFromFn.Name(), upstreamName)
	}

	updatedDefinition, err = f.functionInjector.Inject(updatedDefinition, assignToFn)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to inject %s function into %q", assignToFn.Name(), upstreamName)
	}

	return &updatedDefinition, nil
}

// injectOriginalVersionMethod modifies spec types by injecting an OriginalVersion() function
func (f *StorageTypeFactory) injectOriginalVersionMethod(definition astmodel.TypeDefinition) (*astmodel.TypeDefinition, error) {
	if !f.specTypes.Contains(definition.Name()) {
		// No error, no transform
		return nil, nil
	}

	fn := functions.NewOriginalVersionFunction(f.idFactory)
	result, err := f.functionInjector.Inject(definition, fn)
	if err != nil {
		return nil, errors.Wrapf(err, "injecting OriginalVersion() into %s", definition.Name())
	}

	return &result, nil
}
