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
	group             string                                                  // Name of the group we're handling (used mostly for logging)
	referenceTypes    astmodel.Types                                          // All the types for this group
	outputTypes       astmodel.Types                                          // All the types created/modified by this factory
	idFactory         astmodel.IdentifierFactory                              // Factory for creating identifiers
	typeConverter     *TypeConverter                                          // a utility type type visitor used to create storage variants
	functionInjector  *FunctionInjector                                       // a utility used to inject functions into definitions
	resourceHubMarker *HubVersionMarker                                       // a utility used to mark resources as Storage Versions
	conversionMap     map[astmodel.PackageReference]astmodel.PackageReference // Map of conversion links for creating our conversion graph
}

// NewStorageTypeFactory creates a new instance of StorageTypeFactory ready for use
func NewStorageTypeFactory(group string, idFactory astmodel.IdentifierFactory) *StorageTypeFactory {

	types := make(astmodel.Types)

	result := &StorageTypeFactory{
		group:             group,
		referenceTypes:    types,
		outputTypes:       make(astmodel.Types),
		idFactory:         idFactory,
		conversionMap:     make(map[astmodel.PackageReference]astmodel.PackageReference),
		functionInjector:  NewFunctionInjector(),
		resourceHubMarker: NewHubVersionMarker(),
		typeConverter:     NewTypeConverter(types),
	}

	return result
}

// Add the supplied type definition to this factory
func (f *StorageTypeFactory) Add(def astmodel.TypeDefinition) {
	f.referenceTypes.Add(def)

	if rt, ok := astmodel.AsResourceType(def.Type()); ok {
		// We have a resource type
		if tn, ok := astmodel.AsTypeName(rt.SpecType()); ok {
			// Keep track of all our spec types
			f.specTypes = f.specTypes.Add(tn)
		}
		if tn, ok := astmodel.AsTypeName(rt.StatusType()); ok {
			// Keep track of all our status types
			f.statusTypes = f.statusTypes.Add(tn)
		}
	}
}

// Types returns referenceTypes contained by the factory, including all new storage variants and modified
// api referenceTypes. If any errors occur during processing, they're returned here.
func (f *StorageTypeFactory) Types() (astmodel.Types, error) {

	if len(f.outputTypes) == 0 {
		err := f.process()
		if err != nil {
			return nil, err
		}
	}

	return f.outputTypes, nil
}

// process carries out our transformations
// Each step reads from outputTypes and puts the results back in there
func (f *StorageTypeFactory) process() error {

	f.outputTypes = f.referenceTypes.Copy()

	// Inject OriginalGVK() methods into our resource types
	modifiedResourceTypes, err := f.referenceTypes.Process(f.injectOriginalGVK)
	if err != nil {
		return err
	}
	f.outputTypes = f.outputTypes.OverlayWith(modifiedResourceTypes)

	// Create Storage Variants (injectConversions will look for them)
	storageVariants, err := f.outputTypes.Process(f.createStorageVariant)
	if err != nil {
		return err
	}
	f.outputTypes.AddTypes(storageVariants)

	// Inject conversion functions where required and stash them in output types
	typesWithConversions, err := f.outputTypes.Process(f.injectConversions)
	if err != nil {
		return err
	}
	f.outputTypes = f.outputTypes.OverlayWith(typesWithConversions)

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

	storageDef, err := f.typeConverter.ConvertDefinition(definition)
	if err != nil {
		return nil, errors.Wrapf(err, "creating storage variant for %q", name)
	}

	// Add API-Package -> Storage-Package link into the conversion map
	f.conversionMap[name.PackageReference] = storageDef.Name().PackageReference

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
	nextPackage, ok := f.conversionMap[name.PackageReference]
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
	nextDef, ok := f.outputTypes[nextName]
	if !ok {
		// No next type so nothing to do
		// (this is expected if the type is discontinued or we're looking at the hub type)
		return nil, nil
	}

	// Create conversion functions
	knownTypes := make(astmodel.Types)
	knownTypes.AddTypes(f.inputTypes)
	knownTypes.AddTypes(f.outputTypes)
	conversionContext := conversions.NewPropertyConversionContext(knownTypes, f.idFactory)

	convertFromFn, err := conversions.NewPropertyAssignmentFromFunction(definition, nextDef, f.idFactory, conversionContext)
	if err != nil {
		return nil, errors.Wrapf(err, "creating ConvertFrom() function for %q", name)
	}

	convertToFn, err := conversions.NewPropertyAssignmentToFunction(definition, nextDef, f.idFactory, conversionContext)
	if err != nil {
		return nil, errors.Wrapf(err, "creating ConvertTo() function for %q", name)
	}

	updatedDefinition, err := f.functionInjector.Inject(definition, convertFromFn)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to inject %s function into %q", convertFromFn.Name(), name)
	}

	updatedDefinition, err = f.functionInjector.Inject(updatedDefinition, convertToFn)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to inject %s function into %q", convertToFn.Name(), name)
	}

	return &updatedDefinition, nil
}

// injectOriginalGVK modifies spec types by injecting an OriginalGVK() function
func (f *StorageTypeFactory) injectOriginalGVK(definition astmodel.TypeDefinition) (*astmodel.TypeDefinition, error) {
	if _, isResource := astmodel.AsResourceType(definition.Type()); isResource {
		// We have a Resource, inject the OriginalGVK() function
		fn := functions.NewOriginalGVKFunction(f.idFactory)

		result, err := f.functionInjector.Inject(definition, fn)
		if err != nil {
			return nil, errors.Wrapf(err, "injecting OriginalGVK() into %s", definition.Name())
		}

		return &result, nil
	}

	return nil, nil
}
