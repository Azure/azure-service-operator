/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package storage

import (
	"fmt"

	"github.com/pkg/errors"
	kerrors "k8s.io/apimachinery/pkg/util/errors"

	"github.com/Azure/azure-service-operator/hack/generator/pkg/astmodel"
)

// TypeConverter is used to create a storage variant of an API type
type TypeConverter struct {
	// visitor used to apply the modification
	visitor astmodel.TypeVisitor
	// types contains all the types for this group
	types astmodel.Types
	// propertyConverter is used to modify properties
	propertyConverter *PropertyConverter
}

// NewTypeConverter creates a new instance of the utility type
func NewTypeConverter(types astmodel.Types) *TypeConverter {
	result := &TypeConverter{
		types:             types,
		propertyConverter: NewPropertyConverter(types),
	}

	result.visitor = astmodel.TypeVisitorBuilder{
		VisitObjectType:    result.convertObjectType,
		VisitResourceType:  result.convertResourceType,
		VisitTypeName:      result.redirectTypeNamesToStoragePackage,
		VisitValidatedType: result.stripAllValidations,
		VisitFlaggedType:   result.stripAllFlags,
	}.Build()

	return result
}

// ConvertDefinition applies our type conversion to a specific type definition
func (t *TypeConverter) ConvertDefinition(def astmodel.TypeDefinition) (astmodel.TypeDefinition, error) {
	result, err := t.visitor.VisitDefinition(def, nil)
	if err != nil {
		// Don't need to wrap for context because all our callers do that with better precision
		return astmodel.TypeDefinition{}, err
	}

	description := t.descriptionForStorageVariant(def)
	result = result.WithDescription(description)

	return result, nil
}

/*
 * Functions used by the typeConverter TypeVisitor
 */

// convertResourceType creates a storage variation of a resource type
func (t *TypeConverter) convertResourceType(
	tv *astmodel.TypeVisitor,
	resource *astmodel.ResourceType,
	ctx interface{}) (astmodel.Type, error) {

	// storage resource types do not need defaulter/validator interfaces, they have no webhooks
	result := resource.WithoutInterface(astmodel.DefaulterInterfaceName).
		WithoutInterface(astmodel.ValidatorInterfaceName)

	return astmodel.IdentityVisitOfResourceType(tv, result, ctx)
}

// convertObjectType creates a storage variation of an object type
func (t *TypeConverter) convertObjectType(
	_ *astmodel.TypeVisitor, object *astmodel.ObjectType, _ interface{}) (astmodel.Type, error) {

	var errs []error
	properties := object.Properties()
	for i, prop := range properties {
		p, err := t.propertyConverter.ConvertProperty(prop)
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
func (t *TypeConverter) redirectTypeNamesToStoragePackage(
	_ *astmodel.TypeVisitor, name astmodel.TypeName, _ interface{}) (astmodel.Type, error) {
	if result, ok := t.tryConvertToStoragePackage(name); ok {
		return result, nil
	}

	return name, nil
}

// stripAllValidations removes all validations
func (t *TypeConverter) stripAllValidations(
	this *astmodel.TypeVisitor, v *astmodel.ValidatedType, ctx interface{}) (astmodel.Type, error) {
	// strip all type validations from storage types,
	// act as if they do not exist
	return this.Visit(v.ElementType(), ctx)
}

// stripAllFlags removes all flags
func (t *TypeConverter) stripAllFlags(
	tv *astmodel.TypeVisitor,
	flaggedType *astmodel.FlaggedType,
	ctx interface{}) (astmodel.Type, error) {
	if flaggedType.HasFlag(astmodel.ARMFlag) {
		// We don't want to do anything with ARM types
		return flaggedType, nil
	}

	return astmodel.IdentityVisitOfFlaggedType(tv, flaggedType, ctx)
}

func (_ *TypeConverter) tryConvertToStoragePackage(name astmodel.TypeName) (astmodel.TypeName, bool) {
	// Map the type name into our storage package
	localRef, ok := name.PackageReference.AsLocalPackage()
	if !ok {
		return astmodel.TypeName{}, false
	}

	storageRef := astmodel.MakeStoragePackageReference(localRef)
	visitedName := astmodel.MakeTypeName(storageRef, name.Name())
	return visitedName, true
}

// descriptionForStorageVariant creates a description for a storage variant, indicating which
// original type it is based upon
func (_ *TypeConverter) descriptionForStorageVariant(definition astmodel.TypeDefinition) []string {
	pkg := definition.Name().PackageReference.PackageName()

	result := []string{
		fmt.Sprintf("Storage version of %v.%v", pkg, definition.Name().Name()),
	}
	result = append(result, definition.Description()...)

	return result
}
