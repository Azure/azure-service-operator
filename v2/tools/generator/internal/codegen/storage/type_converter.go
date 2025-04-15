/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package storage

import (
	"fmt"

	"github.com/rotisserie/eris"
	kerrors "k8s.io/apimachinery/pkg/util/errors"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
)

// TypeConverter is used to create a storage variant of an API type
type TypeConverter struct {
	// visitor used to apply the modification
	visitor astmodel.TypeVisitor[any]
	// definitions contains all the definitions for this group
	definitions astmodel.TypeDefinitionSet
	// propertyConverter is used to modify properties
	propertyConverter *PropertyConverter
}

// NewTypeConverter creates a new converter for the creation of storage variants
func NewTypeConverter(definitions astmodel.TypeDefinitionSet) *TypeConverter {
	result := &TypeConverter{
		definitions:       definitions,
		propertyConverter: NewPropertyConverter(definitions),
	}

	result.visitor = astmodel.TypeVisitorBuilder[any]{
		VisitObjectType:       result.convertObjectType,
		VisitResourceType:     result.convertResourceType,
		VisitInternalTypeName: result.redirectTypeNamesToStoragePackage,
		VisitFlaggedType:      result.stripAllFlags,
		// No need to VisitValidatedType and strip validations here as it's already taken care of by the property conversion logic inside of convertObjectType
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
	result = result.WithDescription(description...)

	return result, nil
}

/*
 * Functions used by the typeConverter TypeVisitor
 */

// convertResourceType creates a storage variation of a resource type
func (t *TypeConverter) convertResourceType(
	tv *astmodel.TypeVisitor[any],
	resource *astmodel.ResourceType,
	ctx any,
) (astmodel.Type, error) {
	// storage resource definitions do not need defaulter/validator interfaces, they have no webhooks
	result := resource.WithoutInterface(astmodel.DefaulterInterfaceName).
		WithoutInterface(astmodel.ValidatorInterfaceName)

	return astmodel.IdentityVisitOfResourceType(tv, result, ctx)
}

// convertObjectType creates a storage variation of an object type
func (t *TypeConverter) convertObjectType(
	_ *astmodel.TypeVisitor[any], object *astmodel.ObjectType, _ any,
) (astmodel.Type, error) {
	var errs []error
	properties := object.Properties().Copy()
	for name, prop := range properties {
		p, err := t.propertyConverter.ConvertProperty(prop)
		if err != nil {
			errs = append(errs, eris.Wrapf(err, "property %s", name))
		} else {
			properties[name] = p
		}
	}

	bagName, err := t.selectPropertyBagName(object)
	if err != nil {
		errs = append(errs, err)
	}

	// We use the JSON identifier $propertyBag because it can't possibly conflict with any identifier generated from
	// an ARM schema (none of those use the prefix `$`)
	bagProperty := astmodel.NewPropertyDefinition(bagName, "$propertyBag", astmodel.PropertyBagType).
		WithTag("json", "omitempty")
	properties.Add(bagProperty)

	if len(errs) > 0 {
		err := kerrors.NewAggregate(errs)
		return nil, err
	}

	objectType := astmodel.NewObjectType().WithProperties(properties.AsSlice()...)

	return astmodel.StorageFlag.ApplyTo(objectType), nil
}

// redirectTypeNamesToStoragePackage modifies TypeNames to reference the current storage package
func (t *TypeConverter) redirectTypeNamesToStoragePackage(name astmodel.InternalTypeName) (astmodel.Type, error) {
	if result, ok := t.tryConvertToStoragePackage(name); ok {
		return result, nil
	}

	// Failed to redirect into a storage package, return an error
	return nil, eris.Errorf("unable to redirect %s into a storage package", name)
}

// stripAllFlags removes all flags
func (t *TypeConverter) stripAllFlags(
	tv *astmodel.TypeVisitor[any],
	flaggedType *astmodel.FlaggedType,
	ctx any,
) (astmodel.Type, error) {
	if flaggedType.HasFlag(astmodel.ARMFlag) {
		// We don't want to do anything with ARM definitions
		return flaggedType, nil
	}

	return astmodel.IdentityVisitOfFlaggedType(tv, flaggedType, ctx)
}

// tryConvertToStoragePackage converts the supplied TypeName to reference the parallel type in a storage package if it
// is a local reference; if not, it returns false.
func (t *TypeConverter) tryConvertToStoragePackage(name astmodel.InternalTypeName) (astmodel.InternalTypeName, bool) {
	local, ok := name.PackageReference().(astmodel.LocalPackageReference)
	if !ok {
		return astmodel.InternalTypeName{}, false
	}

	storage := astmodel.MakeStoragePackageReference(local)
	return name.WithPackageReference(storage), true
}

// descriptionForStorageVariant creates a description for a storage variant, indicating which
// original type it is based upon
func (*TypeConverter) descriptionForStorageVariant(definition astmodel.TypeDefinition) []string {
	pkg := definition.Name().PackageReference().PackageName()

	result := []string{
		fmt.Sprintf("Storage version of %s.%s", pkg, definition.Name().Name()),
	}
	result = append(result, definition.Description()...)

	return result
}

// selectPropertyBagName chooses a name for the property bag that doesn't clash with any existing property name
// We have to allow for the possibility of a clash given that individual product teams choose their own property names
// when they are building their ARM RPs.
// We'll have a problem if a team introduces a clash by changing an EXISTING API version, but that shouldn't happen
func (t *TypeConverter) selectPropertyBagName(object *astmodel.ObjectType) (astmodel.PropertyName, error) {
	candidateNames := []astmodel.PropertyName{
		"PropertyBag",
		"PropertyStash",
		"ASOPropertyBag",
		"ASOPropertyStash",
	}

	for _, name := range candidateNames {
		if _, exists := object.Property(name); exists {
			continue
		}

		return name, nil
	}

	return "", eris.Errorf("failed to find non-clashing name for PropertyBag (tried %q)", candidateNames)
}
