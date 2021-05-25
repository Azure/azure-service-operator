/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package storage

import (
	"github.com/Azure/azure-service-operator/hack/generator/pkg/astmodel"
	"github.com/pkg/errors"
	"k8s.io/klog/v2"
)

// StorageTypeFactory is used to create storage inputTypes for a specific service
type StorageTypeFactory struct {
	service           string                                                  // Name of the group we're handling (used mostly for logging)
	inputTypes        astmodel.Types                                          // All the types for this group
	outputTypes       astmodel.Types                                          // All the types created/modified by this factory
	idFactory         astmodel.IdentifierFactory                              // Factory for creating identifiers
	typeConverter     *TypeConverter                                          // a utility type type visitor used to create storage variants
	functionInjector  *FunctionInjector                                       // a utility used to inject functions into definitions
	resourceHubMarker *HubVersionMarker                                       // a utility used to mark resources as Storage Versions
	conversionMap     map[astmodel.PackageReference]astmodel.PackageReference // Map of conversion links for creating our conversion graph
}

// NewStorageTypeFactory creates a new instance of StorageTypeFactory ready for use
func NewStorageTypeFactory(service string, idFactory astmodel.IdentifierFactory) *StorageTypeFactory {

	types := make(astmodel.Types)

	result := &StorageTypeFactory{
		service:                    service,
		inputTypes:                 types,
		idFactory:                  idFactory,
		conversionMap:              make(map[astmodel.PackageReference]astmodel.PackageReference),
		functionInjector:           NewFunctionInjector(),
		resourceHubMarker:          NewHubVersionMarker(),
		typeConverter:              NewTypeConverter(types),
	}

	return result
}

// Add the supplied type definition to this factory
func (f *StorageTypeFactory) Add(def astmodel.TypeDefinition) {
	f.inputTypes.Add(def)
}

// Types returns inputTypes contained by the factory, including all new storage variants and modified
// api inputTypes. If any errors occur during processing, they're returned here.
func (f *StorageTypeFactory) Types() (astmodel.Types, error) {
	if len(f.outputTypes) == 0 {
		err := f.process()
		if err != nil {
			return nil, err
		}
	}

	return f.outputTypes, nil
}

func (f *StorageTypeFactory) process() error {

	f.outputTypes = make(astmodel.Types)

	// Create storage variants and stash them in outputTypes so injectConversions can find them
	storageVariants, err := f.inputTypes.Process(f.createStorageVariant)
	if err != nil {
		return err
	}

	f.outputTypes.AddTypes(storageVariants)

	// Inject conversions into our original types and stash them in outputTypes
	inputTypesWithConversions, err := f.inputTypes.Process(f.injectConversions)
	if err != nil {
		return err
	}

	f.outputTypes.AddTypes(inputTypesWithConversions)

	// Inject conversions into our storage variants and replace the definitions in outputTypes
	storageVariantsWithConversions, err := storageVariants.Process(f.injectConversions)
	if err != nil {
		return err
	}

	for _, d := range storageVariantsWithConversions {
		f.outputTypes[d.Name()] = d
	}


	// Add anything missing into outputTypes
	f.outputTypes.AddTypes(f.inputTypes.Except(f.outputTypes))

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
	conversionContext := astmodel.NewStorageConversionContext(knownTypes, f.idFactory)

	convertFromFn, err := astmodel.NewStorageConversionFromFunction(definition, nextDef, f.idFactory, conversionContext)
	if err != nil {
		return nil, errors.Wrapf(err, "creating ConvertFrom() function for %q", name)
	}

	convertToFn, err := astmodel.NewStorageConversionToFunction(definition, nextDef, f.idFactory, conversionContext)
	if err != nil {
		return nil, errors.Wrapf(err, "creating ConvertTo() function for %q", name)
	}

	updatedDefinition, err := f.functionInjector.Inject(definition, convertFromFn)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to inject ConvertFrom function into %q", name)
	}

	updatedDefinition, err = f.functionInjector.Inject(updatedDefinition, convertToFn)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to inject ConvertFrom function into %q", name)
	}

	return &updatedDefinition, nil
}
