/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package codegen

import (
	"context"
	"fmt"

	"github.com/Azure/azure-service-operator/hack/generator/pkg/astmodel"
	kerrors "k8s.io/apimachinery/pkg/util/errors"
)

// createStorageTypes returns a pipeline stage that creates dedicated storage types for each resource and nested object.
// Storage versions are created for *all* API versions to allow users of older versions of the operator to easily
// upgrade. This is of course a bit odd for the first release, but defining the approach from day one is useful.
func createStorageTypes() PipelineStage {
	return MakePipelineStage(
		"createStorage",
		"Create storage versions of CRD types",
		func(ctx context.Context, types astmodel.Types) (astmodel.Types, error) {

			storageTypes := make(astmodel.Types)
			visitor := makeStorageTypesVisitor(types)
			vc := makeStorageTypesVisitorContext()
			var errs []error
			for _, d := range types {
				d := d

				if astmodel.ARMFlag.IsOn(d.Type()) {
					// Skip ARM definitions, we don't need to create storage variants of those
					continue
				}

				if _, ok := types.ResolveEnumDefinition(&d); ok {
					// Skip Enum definitions as we use the base type for storage
					continue
				}

				def, err := visitor.VisitDefinition(d, vc)
				if err != nil {
					errs = append(errs, err)
					continue
				}

				finalDef := def.WithDescription(descriptionForStorageVariant(d))
				storageTypes[finalDef.Name()] = finalDef
			}

			if len(errs) > 0 {
				err := kerrors.NewAggregate(errs)
				return nil, err
			}

			types.AddTypes(storageTypes)

			return types, nil
		})
}

// makeStorageTypesVisitor returns a TypeVisitor to do the creation of dedicated storage types
func makeStorageTypesVisitor(types astmodel.Types) astmodel.TypeVisitor {
	factory := &StorageTypeFactory{
		types: types,
	}

	visitor := astmodel.TypeVisitorBuilder{
		VisitValidatedType: factory.stripValidations,
		VisitTypeName:      factory.remapNames,
		VisitObjectType:    factory.convertObjects,
		VisitResourceType:  factory.convertResources,
		VisitFlaggedType:   factory.skipArmTypes,
	}.Build()

	factory.visitor = visitor
	factory.propertyConversions = []propertyConversion{
		factory.preserveKubernetesResourceStorageProperties,
		factory.convertPropertiesForStorage,
	}

	return visitor
}

type StorageTypeFactory struct {
	types               astmodel.Types
	propertyConversions []propertyConversion
	visitor             astmodel.TypeVisitor
}

// A property conversion accepts a property definition and optionally applies a conversion to make
// the property suitable for use on a storage type. Conversions return nil if they decline to
// convert, deferring the conversion to another.
type propertyConversion = func(property *astmodel.PropertyDefinition, ctx StorageTypesVisitorContext) (*astmodel.PropertyDefinition, error)

func (factory *StorageTypeFactory) stripValidations(this *astmodel.TypeVisitor, v *astmodel.ValidatedType, ctx interface{}) (astmodel.Type, error) {
	// strip all type validations from storage types,
	// act as if they do not exist
	return this.Visit(v.ElementType(), ctx)
}

func (factory *StorageTypeFactory) remapNames(_ *astmodel.TypeVisitor, name astmodel.TypeName, ctx interface{}) (astmodel.Type, error) {
	visitorContext := ctx.(StorageTypesVisitorContext)

	// Resolve the type name to the actual referenced type
	actualDefinition, actualDefinitionFound := factory.types[name]

	// Check for property specific handling
	if visitorContext.property != nil && actualDefinitionFound {
		if et, ok := actualDefinition.Type().(*astmodel.EnumType); ok {
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

func (factory *StorageTypeFactory) convertResources(
	this *astmodel.TypeVisitor,
	resource *astmodel.ResourceType,
	ctx interface{}) (astmodel.Type, error) {

	// storage resource types do not need defaulter/validator interface, they have no webhooks
	return resource.WithoutInterface(astmodel.DefaulterInterfaceName).WithoutInterface(astmodel.ValidatorInterfaceName), nil
}

func (factory *StorageTypeFactory) convertObjects(
	_ *astmodel.TypeVisitor,
	object *astmodel.ObjectType,
	ctx interface{}) (astmodel.Type, error) {
	visitorContext := ctx.(StorageTypesVisitorContext)
	objectContext := visitorContext.forObject(object)

	var errs []error
	properties := object.Properties()
	for i, prop := range properties {
		p, err := factory.makeStorageProperty(prop, objectContext)
		if err != nil {
			errs = append(errs, err)
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
func (factory *StorageTypeFactory) makeStorageProperty(
	prop *astmodel.PropertyDefinition,
	objectContext StorageTypesVisitorContext) (*astmodel.PropertyDefinition, error) {
	for _, conv := range factory.propertyConversions {
		p, err := conv(prop, objectContext.forProperty(prop))
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

// preserveKubernetesResourceStorageProperties preserves properties required by the KubernetesResource interface as they're always required
func (factory *StorageTypeFactory) preserveKubernetesResourceStorageProperties(
	prop *astmodel.PropertyDefinition,
	_ StorageTypesVisitorContext) (*astmodel.PropertyDefinition, error) {
	if astmodel.IsKubernetesResourceProperty(prop.PropertyName()) {
		// Keep these unchanged
		return prop, nil
	}

	// No opinion, defer to another conversion
	return nil, nil
}

func (factory *StorageTypeFactory) convertPropertiesForStorage(
	prop *astmodel.PropertyDefinition,
	objectContext StorageTypesVisitorContext) (*astmodel.PropertyDefinition, error) {
	propertyType, err := factory.visitor.Visit(prop.PropertyType(), objectContext)
	if err != nil {
		return nil, err
	}

	p := prop.WithType(propertyType).
		MakeOptional().
		WithDescription("")

	return p, nil
}

func (factory *StorageTypeFactory) skipArmTypes(
	tv *astmodel.TypeVisitor,
	flaggedType *astmodel.FlaggedType,
	ctx interface{}) (astmodel.Type, error) {
	if flaggedType.HasFlag(astmodel.ARMFlag) {
		// We don't want to do anything with ARM types
		return flaggedType, nil
	}

	return astmodel.IdentityVisitOfFlaggedType(tv, flaggedType, ctx)
}

func descriptionForStorageVariant(definition astmodel.TypeDefinition) []string {
	pkg := definition.Name().PackageReference.PackageName()

	result := []string{
		fmt.Sprintf("Storage version of %v.%v", pkg, definition.Name().Name()),
	}
	result = append(result, definition.Description()...)

	return result
}

type StorageTypesVisitorContext struct {
	object   *astmodel.ObjectType
	property *astmodel.PropertyDefinition
}

func makeStorageTypesVisitorContext() StorageTypesVisitorContext {
	return StorageTypesVisitorContext{}
}

func (context StorageTypesVisitorContext) forObject(object *astmodel.ObjectType) StorageTypesVisitorContext {
	context.object = object
	context.property = nil
	return context
}

func (context StorageTypesVisitorContext) forProperty(property *astmodel.PropertyDefinition) StorageTypesVisitorContext {
	context.property = property
	return context
}
