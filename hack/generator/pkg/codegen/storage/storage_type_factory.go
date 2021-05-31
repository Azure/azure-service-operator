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
)

// StorageTypeFactory is used to create storage inputTypes for a specific api group
type StorageTypeFactory struct {
	group                  string                                                  // Name of the group we're handling (used mostly for logging)
	inputTypes             astmodel.Types                                          // All the types for this group
	outputTypes            astmodel.Types                                          // All the types created/modified by this factory
	idFactory              astmodel.IdentifierFactory                              // Factory for creating identifiers
	typeConverter          *TypeConverter                                          // a utility type type visitor used to create storage variants
	functionInjector       *FunctionInjector                                       // a utility used to inject functions into definitions
	implementationInjector *ImplementationInjector                                 // a utility used to inject interface implementations into definitions
	resourceHubMarker      *HubVersionMarker                                       // a utility used to mark resources as Storage Versions
	conversionMap          map[astmodel.PackageReference]astmodel.PackageReference // Map of conversion links for creating our conversion graph
	hubPackage             astmodel.PackageReference                               // Identifies the package that represents our hub
}

// NewStorageTypeFactory creates a new instance of StorageTypeFactory ready for use
func NewStorageTypeFactory(group string, idFactory astmodel.IdentifierFactory) *StorageTypeFactory {

	types := make(astmodel.Types)

	result := &StorageTypeFactory{
		group:                  group,
		inputTypes:             types,
		outputTypes:            make(astmodel.Types),
		idFactory:              idFactory,
		conversionMap:          make(map[astmodel.PackageReference]astmodel.PackageReference),
		functionInjector:       NewFunctionInjector(),
		implementationInjector: NewImplementationInjector(),
		resourceHubMarker:      NewHubVersionMarker(),
		typeConverter:          NewTypeConverter(types),
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

	err := f.createConversionGraph()
	if err != nil {
		return errors.Wrap(err, "creating conversion package graph")
	}

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

// createConversionMap creates our conversion graph of links between versions, leading towards our hub version
func (f *StorageTypeFactory) createConversionGraph() error {

	// Collect all distinct versions
	allApiVersions := astmodel.NewPackageReferenceSet()
	for _, t := range f.inputTypes {
		allApiVersions.AddReference(t.Name().PackageReference)
	}

	// And turn them into a sequence sorted by increasing version
	sortedApiVersions := allApiVersions.AsSlice()
	astmodel.SortPackageReferencesByPathAndVersion(sortedApiVersions)

	// For each API version create a matched storage package and to the graph
	sortedStorageVersions := make([]astmodel.StoragePackageReference, len(sortedApiVersions))
	for i, ref := range sortedApiVersions {
		localRef, ok := ref.AsLocalPackage()
		if !ok {
			return errors.Errorf("expected %q to be a local package reference", ref)
		}

		storageRef := astmodel.MakeStoragePackageReference(localRef)
		sortedStorageVersions[i] = storageRef

		// Add link into our graph
		f.conversionMap[ref] = storageRef
	}

	// For each Preview API version, the link goes from the associated storage version to the immediately prior storage version
	for i, ref := range sortedApiVersions {
		if i == 0 || !ref.IsPreview() {
			continue
		}

		f.conversionMap[sortedStorageVersions[i]] = sortedStorageVersions[i-1]
	}

	// For each GA (non-Preview) API version, the link goes from the associated storage version to the next GA release
	var gaRelease astmodel.PackageReference
	for i, ref := range sortedStorageVersions {
		if i == 0 {
			gaRelease = ref
			continue
		}

		if !ref.IsPreview() {
			f.conversionMap[gaRelease] = ref
			gaRelease = ref
		}
	}

	// Cache the hub version for later reference
	f.hubPackage = gaRelease

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
	knownTypes.AddTypes(f.inputTypes.Except(f.outputTypes))
	knownTypes.AddTypes(f.outputTypes)
	conversionFromContext := conversions.NewStorageConversionContext(knownTypes, conversions.ConvertFrom, f.idFactory)
	assignFromFn, err := conversions.NewPropertyAssignmentFromFunction(definition, nextDef, f.idFactory, conversionFromContext)
	if err != nil {
		return nil, errors.Wrapf(err, "creating PropertyAssignmentFrom() function for %q", name)
	}

	conversionToContext := conversions.NewStorageConversionContext(knownTypes, conversions.ConvertTo, f.idFactory)
	assignToFn, err := conversions.NewPropertyAssignmentToFunction(definition, nextDef, f.idFactory, conversionToContext)
	if err != nil {
		return nil, errors.Wrapf(err, "creating PropertyAssignmentTo() function for %q", name)
	}

	definitionWithAssignmentFunctions, err := f.functionInjector.Inject(definition, assignToFn, assignFromFn)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to inject ConvertFrom and ConvertTo functions into %q", name)
	}

	// Work out the name of our hub type
	// TODO: Make this work when types are renamed
	// TODO: Make this work when types are discontinued
	hubType := astmodel.MakeTypeName(f.hubPackage, name.Name())

	convertFromFn := conversions.NewConversionFromHubFunction(hubType, assignFromFn.OtherType(), assignFromFn.Name(), f.idFactory)
	convertToFn := conversions.NewConversionToHubFunction(hubType, assignToFn.OtherType(), assignToFn.Name(), f.idFactory)

	hubImplementation := astmodel.NewInterfaceImplementation(astmodel.ConvertibleInterface, convertFromFn, convertToFn)

	definitionWithHubInterface, err := f.implementationInjector.Inject(definitionWithAssignmentFunctions, hubImplementation)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to inject conversion.Convertible interface into %q", name)
	}

	return &definitionWithHubInterface, nil
}
