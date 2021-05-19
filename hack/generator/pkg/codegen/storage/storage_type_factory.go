/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package storage

import (
	"fmt"

	"github.com/Azure/k8s-infra/hack/generator/pkg/astmodel"
	"github.com/pkg/errors"
	kerrors "k8s.io/apimachinery/pkg/util/errors"
	"k8s.io/klog/v2"
)

// StorageTypeFactory is used to create storage types for a specific service
type StorageTypeFactory struct {
	service                    string                                                  // Name of the group we're handling (used mostly for logging)
	types                      astmodel.Types                                          // All the types for this group
	propertyConversions        []propertyConversion                                    // Conversion rules to use for properties when creating storage variants
	pendingStorageConversion   astmodel.TypeNameQueue                                  // Queue of types that need storage variants created for them
	pendingConversionInjection astmodel.TypeNameQueue                                  // Queue of types that need conversion functions injected
	pendingMarkAsHubVersion    astmodel.TypeNameQueue                                  // Queue of types that need to be flagged as the hub storage version
	idFactory                  astmodel.IdentifierFactory                              // Factory for creating identifiers
	storageConverter           astmodel.TypeVisitor                                    // a cached type visitor used to create storage variants
	propertyConverter          astmodel.TypeVisitor                                    // a cached type visitor used to simplify property types
	functionInjector           *FunctionInjector                                       // a utility used to inject functions into definitions
	resourceHubMarker          astmodel.TypeVisitor                                    // a cached type visitor used to mark resources as Storage Versions
	conversionMap              map[astmodel.PackageReference]astmodel.PackageReference // Map of conversion links for creating our conversion graph
}

// NewStorageTypeFactory creates a new instance of StorageTypeFactory ready for use
func NewStorageTypeFactory(service string, idFactory astmodel.IdentifierFactory) *StorageTypeFactory {
	result := &StorageTypeFactory{
		service:                    service,
		types:                      make(astmodel.Types),
		pendingStorageConversion:   astmodel.MakeTypeNameQueue(),
		pendingConversionInjection: astmodel.MakeTypeNameQueue(),
		pendingMarkAsHubVersion:    astmodel.MakeTypeNameQueue(),
		idFactory:                  idFactory,
		conversionMap:              make(map[astmodel.PackageReference]astmodel.PackageReference),
		functionInjector:           NewFunctionInjector(),
	}

	result.propertyConversions = []propertyConversion{
		result.preserveKubernetesResourceStorageProperties,
		result.defaultPropertyConversion,
	}

	result.storageConverter = astmodel.TypeVisitorBuilder{
		VisitObjectType:    result.convertObjectType,
		VisitResourceType:  result.convertResourceType,
		VisitTypeName:      result.redirectTypeNamesToStoragePackage,
		VisitValidatedType: result.stripAllValidations,
		VisitFlaggedType:   result.stripAllFlags,
	}.Build()

	result.propertyConverter = astmodel.TypeVisitorBuilder{
		VisitEnumType:      result.useBaseTypeForEnumerations,
		VisitValidatedType: result.stripAllValidations,
		VisitTypeName:      result.shortCircuitNamesOfSimpleTypes,
	}.Build()

	result.resourceHubMarker = astmodel.TypeVisitorBuilder{
		VisitResourceType: result.markResourceAsStorageVersion,
	}.Build()

	return result
}

// Add the supplied type definition to this factory
func (f *StorageTypeFactory) Add(def astmodel.TypeDefinition) {
	f.types.Add(def)

	// Add to our queue of types requiring storage variants
	f.pendingStorageConversion.Enqueue(def.Name())
}

// Types returns types contained by the factory, including all new storage variants and modified
// api types. If any errors occur during processing, they're returned here.
func (f *StorageTypeFactory) Types() (astmodel.Types, error) {
	err := f.process()
	if err != nil {
		return nil, err
	}

	return f.types, nil
}

func (f *StorageTypeFactory) process() error {
	err := f.pendingStorageConversion.Process(f.createStorageVariant)
	if err != nil {
		return err
	}

	err = f.pendingConversionInjection.Process(f.injectConversions)
	if err != nil {
		return err
	}

	err = f.pendingMarkAsHubVersion.Process(f.markAsHubVersion)
	if err != nil {
		return err
	}

	return nil
}

// createStorageVariant takes an existing object definition and creates a storage variant in a
// related package.
// def is the api definition on which to base the storage variant
// visitor is a type visitor that will do the creation
func (f *StorageTypeFactory) createStorageVariant(name astmodel.TypeName) error {

	// Only need to create storage variants of resources and objects
	underlyingType, err := f.types.FullyResolve(name)
	if err != nil {
		return errors.Wrapf(err,
			"expected to find underlying type for %q",
			name)
	}

	_, isObject := astmodel.AsObjectType(underlyingType)
	_, isResource := astmodel.AsResourceType(underlyingType)
	if !isObject && !isResource {
		// just skip it
		klog.V(4).Infof("Skipping %s as no storage variant needed", name)
		return nil
	}

	klog.V(3).Infof("Creating storage variant of %s", name)

	def, ok := f.types[name]
	if !ok {
		return errors.Errorf("failed to find definition for %q", name)
	}

	storageDef, err := f.storageConverter.VisitDefinition(def, nil)
	if err != nil {
		return errors.Wrapf(err, "creating storage variant for %q", name)
	}

	desc := f.descriptionForStorageVariant(def)
	storageDef = storageDef.WithDescription(desc)

	f.types.Add(storageDef)

	// Add API-Package -> Storage-Package link into the conversion map
	f.conversionMap[name.PackageReference] = storageDef.Name().PackageReference

	// Queue for injection of conversion functions
	f.pendingConversionInjection.Enqueue(name)

	//TODO: Queue storage type for injection of conversion too

	return nil
}

// injectConversions modifies the named type by injecting the required conversion methods using
// the conversionMap we've previously established
func (f *StorageTypeFactory) injectConversions(name astmodel.TypeName) error {
	klog.V(3).Infof("Injecting conversion functions into %s", name)

	// Find the definition to modify
	def, ok := f.types[name]
	if !ok {
		return errors.Errorf("failed to find definition for %q", name)
	}

	// Find the definition we want to convert to/from
	nextPackage, ok := f.conversionMap[name.PackageReference]
	if !ok {
		// No next package, so nothing to do
		// (this is expected if we have the hub storage package)
		// Flag the type as needing to be flagged as the storage version
		//TODO: Restore this - currently disabled until we get all the conversion functions injected
		//!! f.pendingMarkAsHubVersion.Enqueue(name)
		return nil
	}

	nextName := astmodel.MakeTypeName(nextPackage, name.Name())
	nextDef, ok := f.types[nextName]
	if !ok {
		// No next type so nothing to do
		// (this is expected if the type is discontinued)
		return nil
	}

	// Create conversion functions
	conversionContext := astmodel.NewStorageConversionContext(f.types)

	convertFrom, err := astmodel.NewStorageConversionFromFunction(def, nextDef, f.idFactory, conversionContext)
	if err != nil {
		return errors.Wrapf(err, "creating ConvertFrom() function for %q", name)
	}

	convertTo, err := astmodel.NewStorageConversionToFunction(def, nextDef, f.idFactory, conversionContext)
	if err != nil {
		return errors.Wrapf(err, "creating ConvertTo() function for %q", name)
	}

	def, err = f.functionInjector.Inject(def, convertFrom)
	if err != nil {
		return errors.Wrapf(err, "failed to inject ConvertFrom function into %q", name)
	}

	def, err = f.functionInjector.Inject(def, convertTo)
	if err != nil {
		return errors.Wrapf(err, "failed to inject ConvertFrom function into %q", name)
	}

	// Update our map
	f.types[name] = def

	return nil
}

func (f *StorageTypeFactory) markAsHubVersion(name astmodel.TypeName) error {
	// Find the definition to modify
	def, ok := f.types[name]
	if !ok {
		return errors.Errorf("failed to find definition for %q", name)
	}

	// Mark the resource as the hub storage version
	updated, err := f.resourceHubMarker.VisitDefinition(def, nil)
	if err != nil {
		return errors.Wrapf(err, "marking %q as hub storage version", name)
	}

	// Update our map
	f.types[name] = updated

	return nil
}

// makeStorageProperty applies a conversion to make a variant of the property for use when
// serializing to storage
func (f *StorageTypeFactory) makeStorageProperty(
	prop *astmodel.PropertyDefinition) (*astmodel.PropertyDefinition, error) {
	for _, conv := range f.propertyConversions {
		p, err := conv(prop)
		if err != nil {
			// Something went wrong, return the error
			return nil, err
		}
		if p != nil {
			// We have the conversion we need, return it promptly
			return p, nil
		}
	}

	return nil, fmt.Errorf("failed to find a conversion for property %v", prop.PropertyName())
}

// A property conversion accepts a property definition and optionally applies a conversion to make
// the property suitable for use on a storage type. Conversions return nil if they decline to
// convert, deferring the conversion to another.
type propertyConversion = func(property *astmodel.PropertyDefinition) (*astmodel.PropertyDefinition, error)

// descriptionForStorageVariant creates a description for a storage variant, indicating which
// original type it is based upon
func (f *StorageTypeFactory) descriptionForStorageVariant(definition astmodel.TypeDefinition) []string {
	pkg := definition.Name().PackageReference.PackageName()

	result := []string{
		fmt.Sprintf("Storage version of %v.%v", pkg, definition.Name().Name()),
	}
	result = append(result, definition.Description()...)

	return result
}

func (f *StorageTypeFactory) tryConvertToStorageNamespace(name astmodel.TypeName) (astmodel.TypeName, bool) {
	// Map the type name into our storage namespace
	localRef, ok := name.PackageReference.AsLocalPackage()
	if !ok {
		return astmodel.TypeName{}, false
	}

	storageRef := astmodel.MakeStoragePackageReference(localRef)
	visitedName := astmodel.MakeTypeName(storageRef, name.Name())
	return visitedName, true
}

/*
 * Functions used as propertyConversions
 */

// preserveKubernetesResourceStorageProperties preserves properties required by the
// KubernetesResource interface as they're always required exactly as declared
func (f *StorageTypeFactory) preserveKubernetesResourceStorageProperties(
	prop *astmodel.PropertyDefinition) (*astmodel.PropertyDefinition, error) {

	if astmodel.IsKubernetesResourceProperty(prop.PropertyName()) {
		// Keep these unchanged
		return prop, nil
	}

	// Not a kubernetes type, defer to another conversion
	return nil, nil
}

func (f *StorageTypeFactory) defaultPropertyConversion(
	prop *astmodel.PropertyDefinition) (*astmodel.PropertyDefinition, error) {
	propertyType, err := f.propertyConverter.Visit(prop.PropertyType(), nil)
	if err != nil {
		return nil, err
	}

	p := prop.WithType(propertyType).
		MakeOptional().
		WithDescription("")

	return p, nil
}

/*
 * Functions used by the storageConverter TypeVisitor
 */

// convertResourceType creates a storage variation of a resource type
func (f *StorageTypeFactory) convertResourceType(
	tv *astmodel.TypeVisitor,
	resource *astmodel.ResourceType,
	ctx interface{}) (astmodel.Type, error) {

	// storage resource types do not need defaulter interface, they have no webhooks
	rsrc := resource.WithoutInterface(astmodel.DefaulterInterfaceName)

	return astmodel.IdentityVisitOfResourceType(tv, rsrc, ctx)
}

// convertObjectType creates a storage variation of an object type
func (f *StorageTypeFactory) convertObjectType(
	_ *astmodel.TypeVisitor, object *astmodel.ObjectType, _ interface{}) (astmodel.Type, error) {

	var errs []error
	properties := object.Properties()
	for i, prop := range properties {
		p, err := f.makeStorageProperty(prop)
		if err != nil {
			errs = append(errs, errors.Wrapf(err, "property %s", prop.PropertyName()))
		} else {
			properties[i] = p
		}
	}

	if len(errs) > 0 {
		err := kerrors.NewAggregate(errs)
		return nil, err
	}

	objectType := astmodel.NewObjectType().WithProperties(properties...)
	return astmodel.StorageFlag.ApplyTo(objectType), nil
}

// redirectTypeNamesToStoragePackage modifies TypeNames to reference the current storage package
func (f *StorageTypeFactory) redirectTypeNamesToStoragePackage(
	_ *astmodel.TypeVisitor, name astmodel.TypeName, _ interface{}) (astmodel.Type, error) {
	if result, ok := f.tryConvertToStorageNamespace(name); ok {
		return result, nil
	}

	return name, nil
}

// stripAllValidations removes all validations
func (f *StorageTypeFactory) stripAllValidations(
	this *astmodel.TypeVisitor, v *astmodel.ValidatedType, ctx interface{}) (astmodel.Type, error) {
	// strip all type validations from storage types,
	// act as if they do not exist
	return this.Visit(v.ElementType(), ctx)
}

// stripAllFlags removes all flags
func (f *StorageTypeFactory) stripAllFlags(
	tv *astmodel.TypeVisitor,
	flaggedType *astmodel.FlaggedType,
	ctx interface{}) (astmodel.Type, error) {
	if flaggedType.HasFlag(astmodel.ARMFlag) {
		// We don't want to do anything with ARM types
		return flaggedType, nil
	}

	return astmodel.IdentityVisitOfFlaggedType(tv, flaggedType, ctx)
}

/*
 * Functions used by the propertyConverter TypeVisitor
 */

// useBaseTypeForEnumerations replaces an enumeration with its underlying base type
func (f *StorageTypeFactory) useBaseTypeForEnumerations(
	tv *astmodel.TypeVisitor, et *astmodel.EnumType, ctx interface{}) (astmodel.Type, error) {
	return tv.Visit(et.BaseType(), ctx)
}

// shortCircuitNamesOfSimpleTypes redirects TypeNames that reference resources or objects into our
// storage namespace, and replaces TypeNames that point to simple types (enumerations or
// primitives) with their underlying type.
func (f *StorageTypeFactory) shortCircuitNamesOfSimpleTypes(
	tv *astmodel.TypeVisitor, tn astmodel.TypeName, ctx interface{}) (astmodel.Type, error) {

	actualType, err := f.types.FullyResolve(tn)
	if err != nil {
		// Can't resolve to underlying type, give up
		return nil, err
	}

	_, isObject := astmodel.AsObjectType(actualType)
	_, isResource := astmodel.AsResourceType(actualType)

	if isObject || isResource {
		// We have an object or a resource, redirect to our storage namespace if we can
		if storageName, ok := f.tryConvertToStorageNamespace(tn); ok {
			return storageName, nil
		}

		// Otherwise just keep the name
		return tn, nil
	}

	// Replace the name with the underlying type
	return tv.Visit(actualType, ctx)
}

/*
 * Functions used by the resourceHubMarker TypeVisitor
 */

// injectFunctionIntoResource takes the function provided as a context and includes it on the
// provided resource type
func (f *StorageTypeFactory) markResourceAsStorageVersion(
	_ *astmodel.TypeVisitor, rt *astmodel.ResourceType, _ interface{}) (astmodel.Type, error) {
	return rt.MarkAsStorageVersion(), nil
}
