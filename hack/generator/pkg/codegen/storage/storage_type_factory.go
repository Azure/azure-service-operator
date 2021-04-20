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

// Each StorageTypeFactory is used to create storage types for a specific service
type StorageTypeFactory struct {
	service                  string                     // Name of the service we're handling (used mostly for logging)
	types                    astmodel.Types             // All the types for this service
	propertyConversions      []propertyConversion       // Conversion rules to use for properties when creating storage variants
	pendingStorageConversion astmodel.TypeNameQueue     // Queue of types that need storage variants created for them
	idFactory                astmodel.IdentifierFactory // Factory for creating identifiers
	storageTypesVisitor      *astmodel.TypeVisitor      // A cached type visitor used to create storage variants
}

/*
	apiTypes            astmodel.Types             // Modified types
	storageTypes        astmodel.Types             // Storage variants of apiTypes
*/

// NewStorageTypeFactory creates a new instance of StorageTypeFactory ready for use
func NewStorageTypeFactory(idFactory astmodel.IdentifierFactory) *StorageTypeFactory {
	result := &StorageTypeFactory{
		types:                    make(astmodel.Types),
		pendingStorageConversion: astmodel.MakeTypeNameQueue(),
		idFactory:                idFactory,
	}

	result.propertyConversions = []propertyConversion{
		result.preserveKubernetesResourceStorageProperties,
		result.convertPropertiesForStorage,
	}

	result.storageConverter = astmodel.TypeVisitorBuilder{
		VisitResourceType:  result.convertResourceType,
		VisitObjectType:    result.convertObjectType,
		VisitTypeName:      result.redirectTypeNamesToStoragePackage,
		VisitValidatedType: result.stripAllValidations,
		VisitFlaggedType:   result.stripAllFlags,
	}.Build()

	return result
}

func (f *StorageTypeFactory) Add(def astmodel.TypeDefinition) {
	f.types.Add(def)

	isArm := astmodel.ARMFlag.IsOn(def.Type())
	_, isEnum := astmodel.AsEnumType(def.Type())

	if !isArm && !isEnum {
		// Add to our queue of types requiring storage variants
		f.pendingStorageConversion.Enqueue(def.Name())
	}
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

	return nil
}

// createStorageVariant takes an existing object definition and creates a storage variant in a
// related package.
// def is the api definition on which to base the storage variant
// visitor is a type visitor that will do the creation
func (f *StorageTypeFactory) createStorageVariant(name astmodel.TypeName) error {

	klog.V(0).Infof("Creating storage variant of %s", name)

	def, ok := f.types[name]
	if !ok {
		return errors.Errorf("failed to find definition for %q", name)
	}

	vc := MakeStorageTypesVisitorContext()
	storageDef, err := f.storageTypesVisitor.VisitDefinition(def, vc)
	if err != nil {
		return errors.Wrapf(err, "creating storage variant for %q", name)
	}

	desc := f.descriptionForStorageVariant(def)
	storageDef = storageDef.WithDescription(desc)

	f.types.Add(storageDef)
	//TODO: Queue for injection of conversion functions

	return nil
}

/*
// createApiVariant modifies an existing object definition by adding the required conversion functions
func (f *StorageTypeFactory) createApiVariant(apiDef astmodel.TypeDefinition, storageDef astmodel.TypeDefinition) (astmodel.TypeDefinition, error) {
	objectType, isObjectType := astmodel.AsObjectType(apiDef.Type())
	if !isObjectType {
		return astmodel.TypeDefinition{}, errors.Errorf("Expected %q to be an object definition", apiDef.Name())
	}

	// Create conversion functions
	conversionContext := astmodel.NewStorageConversionContext(f.types)
	convertFrom, err := astmodel.NewStorageConversionFromFunction(apiDef, storageDef, f.idFactory, conversionContext)
	if err != nil {
		return astmodel.TypeDefinition{}, errors.Wrapf(err, "creating ConvertFrom() function for %q", apiDef.Name())
	}

	convertTo, err := astmodel.NewStorageConversionToFunction(apiDef, storageDef, f.idFactory, conversionContext)
	if err != nil {
		return astmodel.TypeDefinition{}, errors.Wrapf(err, "creating ConvertTo() function for %q", apiDef.Name())
	}

	objectType = objectType.WithFunction(convertFrom).WithFunction(convertTo)
	return apiDef.WithType(objectType), nil
}

*/

// newStorageTypesVisitor returns a TypeVisitor to do the creation of dedicated storage types
func (f *StorageTypeFactory) newStorageTypesVisitor() *astmodel.TypeVisitor {
	result := astmodel.MakeTypeVisitor()
	result.VisitValidatedType = f.visitValidatedType
	result.VisitTypeName = f.visitTypeName
	result.VisitObjectType = f.visitObjectType
	result.VisitResourceType = f.visitResourceType
	result.VisitFlaggedType = f.visitFlaggedType
	return &result
}

// A property conversion accepts a property definition and optionally applies a conversion to make
// the property suitable for use on a storage type. Conversions return nil if they decline to
// convert, deferring the conversion to another.
type propertyConversion = func(property *astmodel.PropertyDefinition, tv *astmodel.TypeVisitor, ctx StorageTypesVisitorContext) (*astmodel.PropertyDefinition, error)

func (f *StorageTypeFactory) visitValidatedType(this *astmodel.TypeVisitor,
	v *astmodel.ValidatedType, ctx interface{}) (astmodel.Type, error) {
	// strip all type validations from storage types,
	// act as if they do not exist
	return this.Visit(v.ElementType(), ctx)
}

func (f *StorageTypeFactory) visitTypeName(_ *astmodel.TypeVisitor, name astmodel.TypeName, ctx interface{}) (astmodel.Type, error) {
	visitorContext := ctx.(StorageTypesVisitorContext)

	// Resolve the type name to the actual referenced type
	actualType, err := f.types.FullyResolve(name)
	if err != nil {
		return nil, errors.Wrapf(err, "visiting type name %q", name)
	}

	// Check for property specific handling
	if visitorContext.property != nil {
		if et, ok := astmodel.AsEnumType(actualType); ok {
			// Property type refers to an enum, so we use the base type instead
			return et.BaseType(), nil
		}
	}

	// Map the type name into our storage namespace
	localRef, ok := name.PackageReference.AsLocalPackage()
	if !ok {
		return name, nil
	}

	storageRef := astmodel.MakeStoragePackageReference(localRef)
	visitedName := astmodel.MakeTypeName(storageRef, name.Name())
	return visitedName, nil
}

func (f *StorageTypeFactory) visitResourceType(
	tv *astmodel.TypeVisitor,
	resource *astmodel.ResourceType,
	ctx interface{}) (astmodel.Type, error) {

	// storage resource types do not need defaulter interface, they have no webhooks
	rsrc := resource.WithoutInterface(astmodel.DefaulterInterfaceName)

	return astmodel.IdentityVisitOfResourceType(tv, rsrc, ctx)
}

func (f *StorageTypeFactory) visitObjectType(
	tv *astmodel.TypeVisitor,
	object *astmodel.ObjectType,
	ctx interface{}) (astmodel.Type, error) {
	visitorContext := ctx.(StorageTypesVisitorContext)
	objectContext := visitorContext.forObject(object)

	var errs []error
	properties := object.Properties()
	for i, prop := range properties {
		p, err := f.makeStorageProperty(prop, tv, objectContext)
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

// makeStorageProperty applies a conversion to make a variant of the property for use when
// serializing to storage
func (f *StorageTypeFactory) makeStorageProperty(
	prop *astmodel.PropertyDefinition,
	tv *astmodel.TypeVisitor,
	objectContext StorageTypesVisitorContext) (*astmodel.PropertyDefinition, error) {
	for _, conv := range f.propertyConversions {
		p, err := conv(prop, tv, objectContext.forProperty(prop))
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

// preserveKubernetesResourceStorageProperties preserves properties required by the
// KubernetesResource interface as they're always required
func (f *StorageTypeFactory) preserveKubernetesResourceStorageProperties(
	prop *astmodel.PropertyDefinition,
	_ *astmodel.TypeVisitor,
	_ StorageTypesVisitorContext) (*astmodel.PropertyDefinition, error) {
	if astmodel.IsKubernetesResourceProperty(prop.PropertyName()) {
		// Keep these unchanged
		return prop, nil
	}

	// No opinion, defer to another conversion
	return nil, nil
}

func (f *StorageTypeFactory) convertPropertiesForStorage(
	prop *astmodel.PropertyDefinition,
	tv *astmodel.TypeVisitor,
	objectContext StorageTypesVisitorContext) (*astmodel.PropertyDefinition, error) {
	propertyType, err := tv.Visit(prop.PropertyType(), objectContext)
	if err != nil {
		return nil, err
	}

	p := prop.WithType(propertyType).
		MakeOptional().
		WithDescription("")

	return p, nil
}

func (f *StorageTypeFactory) visitFlaggedType(
	tv *astmodel.TypeVisitor,
	flaggedType *astmodel.FlaggedType,
	ctx interface{}) (astmodel.Type, error) {
	if flaggedType.HasFlag(astmodel.ArmFlag) {
		// We don't want to do anything with ARM types
		return flaggedType, nil
	}

	return astmodel.IdentityVisitOfFlaggedType(tv, flaggedType, ctx)
}

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
