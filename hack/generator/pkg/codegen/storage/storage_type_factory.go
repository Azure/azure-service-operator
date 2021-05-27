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
	group                  string                                                  // Name of the group we're handling (used mostly for logging)
	referenceTypes         astmodel.Types                                          // All the types for this group
	outputTypes            astmodel.Types                                          // All the types created/modified by this factory
	specTypes              astmodel.TypeNameSet                                    // All the names of spec types
	statusTypes            astmodel.TypeNameSet                                    // All the names of status types
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
		referenceTypes:         types,
		outputTypes:            make(astmodel.Types),
		specTypes:              astmodel.NewTypeNameSet(),
		statusTypes:            astmodel.NewTypeNameSet(),
		idFactory:              idFactory,
		conversionMap:          make(map[astmodel.PackageReference]astmodel.PackageReference),
		functionInjector:       NewFunctionInjector(),
		implementationInjector: NewImplementationInjector(),
		resourceHubMarker:      NewHubVersionMarker(idFactory),
		typeConverter:          NewTypeConverter(types, idFactory),
	}

	return result
}

// Add the supplied type definition to this factory
func (f *StorageTypeFactory) Add(definitions ...astmodel.TypeDefinition) {
	for _, def := range definitions {
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
	err := f.createConversionGraph()
	if err != nil {
		return errors.Wrap(err, "creating conversion package graph")
	}

	// Inject OriginalVersion() methods into all our spec types
	modifiedSpecTypes, err := f.outputTypes.Process(f.injectOriginalVersionMethod)
	if err != nil {
		return err
	}
	f.outputTypes = f.outputTypes.OverlayWith(modifiedSpecTypes)

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

// createConversionMap creates our conversion graph of links between versions, leading towards our hub version
func (f *StorageTypeFactory) createConversionGraph() error {

	// Collect all distinct versions
	allApiVersions := astmodel.NewPackageReferenceSet()
	for _, t := range f.outputTypes {
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

	return f.injectConversionsBetween(definition, nextDef)
}

// inject conversion methods between the two specified definitions
// upstreamDef is the definition further away from our hub type in our directed conversion graph
// downstreamDef is the definition closer to our hub type in our directed conversion graph
func (f *StorageTypeFactory) injectConversionsBetween(
	upstreamDef astmodel.TypeDefinition, downstreamDef astmodel.TypeDefinition) (*astmodel.TypeDefinition, error) {

	// Create conversion functions
	knownTypes := f.referenceTypes.Copy()
	knownTypes = knownTypes.OverlayWith(f.outputTypes)
	conversionContext := conversions.NewPropertyConversionContext(knownTypes, f.idFactory)

	assignFromFn, err := conversions.NewPropertyAssignmentFromFunction(upstreamDef, downstreamDef, f.idFactory, conversionContext)
	upstreamName := upstreamDef.Name()
	if err != nil {
		return nil, errors.Wrapf(err, "creating AssignFrom() function for %q", upstreamName)
	}

	// Work out the name of our hub type
	// TODO: Make this work when types are renamed
	// TODO: Make this work when types are discontinued
	hub := astmodel.MakeTypeName(f.hubPackage, name.Name())
    
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
