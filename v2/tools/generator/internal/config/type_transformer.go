/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package config

import (
	"fmt"

	"github.com/go-logr/logr"
	"github.com/rotisserie/eris"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
)

// A TypeTransformer is used to remap types
type TypeTransformer struct {
	TypeMatcher `yaml:",inline"`

	// Property is a wildcard matching specific properties on the types selected by this filter
	Property FieldMatcher `yaml:",omitempty"`

	// IfType only performs the transform if the original type matches
	IfType *TransformSelector `yaml:"ifType,omitempty"`

	// Target is the type to turn the type into
	Target *TransformResult `yaml:",omitempty"`

	// Remove indicates that the property should be removed from the
	// type. This is only usable with Property. Target and Remove are
	// mutually exclusive.
	Remove bool `yaml:",omitempty"`

	// RenameTo allows for renaming of type definitions.
	RenameTo string `yaml:"renameTo,omitempty"`

	// matchedProperties is a map of types that were matched to the property found on the type.
	// This is used to ensure that the type transformer matched at least one property
	matchedProperties map[astmodel.TypeName]string
}

// AppliesToDefinition returns true if this transform should be applied to the passed TypeDefinition.
func (transformer *TypeTransformer) AppliesToDefinition(def astmodel.TypeDefinition) bool {
	// If our GVK doesn't match, we don't apply
	if !transformer.AppliesToType(def.Name()) {
		return false
	}

	// If we're not matching a specific property, check IfType for a match.
	if !transformer.Property.IsRestrictive() {
		if transformer.IfType != nil {
			return transformer.IfType.AppliesToType(def.Type())
		}

		return true
	}

	// We're matching on a specific property, can only do this if this is definition is an object
	_, isObject := astmodel.AsObjectType(def.Type())
	return isObject
}

func (transformer *TypeTransformer) TransformDefinition(def astmodel.TypeDefinition) (astmodel.TypeDefinition, error) {
	if err := transformer.validate(); err != nil {
		return astmodel.TypeDefinition{}, eris.Wrapf(err, "validating transformer for %s", def.Name())
	}

	// Apply our rename if we have one configured
	if transformer.RenameTo != "" {
		def = def.WithName(
			def.Name().WithName(transformer.RenameTo))
	}

	// Transform the whole type if we're not looking for a specific property
	if transformer.Target != nil && !transformer.Property.IsRestrictive() {
		resultType, err := transformer.Target.produceTargetType("target", def.Type())
		if err != nil {
			return astmodel.TypeDefinition{}, eris.Wrapf(
				err,
				"type transformer for group: %s, version: %s, name: %s",
				transformer.Group.String(),
				transformer.Version.String(),
				transformer.Name.String())
		}

		def = def.WithType(resultType)
		return def, nil
	}

	// Transform any matching properties
	if transformer.Property.IsRestrictive() {
		visitor := astmodel.TypeVisitorBuilder[any]{
			VisitObjectType: func(
				it *astmodel.ObjectType,
			) (astmodel.Type, error) {
				return transformer.transformProperties(it)
			},
		}.Build()

		d, err := visitor.VisitDefinition(def, nil)
		if err != nil {
			return astmodel.TypeDefinition{}, eris.Wrapf(
				err,
				"visiting definition %s",
				def.Name())
		}

		def = d
	}

	return def, nil
}

// TransformTypeName transforms the type with the specified name into the TypeTransformer target type if
// the provided type name matches the pattern(s) specified in the TypeTransformer
func (transformer *TypeTransformer) TransformTypeName(typeName astmodel.InternalTypeName) (astmodel.Type, error) {
	if err := transformer.validate(); err != nil {
		return nil, err
	}

	if transformer.AppliesToType(typeName) {
		result, err := transformer.Target.produceTargetType("target", typeName)
		if err != nil {
			return nil, eris.Wrapf(
				err,
				"type transformer for group: %s, version: %s, name: %s",
				transformer.Group.String(),
				transformer.Version.String(),
				transformer.Name.String())
		}

		return result, nil
	}

	// Didn't match so return nil
	return nil, nil
}

func (transformer *TypeTransformer) validate() error {
	// Using Remove precludes other transforms
	if transformer.Remove {
		if !transformer.Property.IsRestrictive() {
			return eris.Errorf("remove is only usable with property transforms")
		}
		if transformer.Target != nil {
			return eris.Errorf("remove and target can't both be set")
		}
	}

	// If we're not removing a property, we need either a target or a rename
	if !transformer.Remove {
		if transformer.Target == nil && transformer.RenameTo == "" {
			return eris.Errorf("transformer must either rename or modify")
		}
	}

	// Property specific transforms should either modify or remove
	if transformer.Property.IsRestrictive() && !transformer.Remove && transformer.Target == nil {
		return eris.Errorf("property transforms must either remove or modify")
	}

	return nil
}

// PropertyTransformResult is the result of applying a property type transform
type PropertyTransformResult struct {
	TypeName        astmodel.TypeName
	NewType         *astmodel.ObjectType
	Property        astmodel.PropertyName
	NewPropertyType astmodel.Type
	Removed         bool
	Because         string
}

// String generates a printable representation of this result.
func (r PropertyTransformResult) String() string {
	if r.Removed {
		return fmt.Sprintf("%s.%s removed because %s",
			r.TypeName,
			r.Property,
			r.Because,
		)
	}
	return fmt.Sprintf("%s.%s -> %s because %s",
		r.TypeName,
		r.Property,
		r.NewPropertyType.String(),
		r.Because,
	)
}

// LogTo creates a log message for the transformation
func (r PropertyTransformResult) LogTo(log logr.Logger) {
	if r.Removed {
		log.V(3).Info(
			"Removing property",
			"type", r.TypeName,
			"property", r.Property,
			"because", r.Because)
		return
	}

	log.V(2).Info(
		"Transforming property",
		"type", r.TypeName,
		"property", r.Property,
		"newType", r.NewPropertyType.String(),
		"because", r.Because)
}

func (transformer *TypeTransformer) RequiredPropertiesWereMatched() error {
	// If this transformer applies to entire types (instead of just properties on types), we just defer to
	// transformer.MatchedRequiredTypes
	if !transformer.Property.IsRestrictive() {
		return transformer.RequiredTypesWereMatched()
	}

	if !transformer.MustMatch() {
		return nil
	}

	if err := transformer.RequiredTypesWereMatched(); err != nil {
		return err
	}

	if err := transformer.Property.WasMatched(); err != nil {
		return eris.Wrapf(
			err,
			"%s matched types but all properties were excluded by name or type",
			transformer.String())
	}

	return nil
}

// TransformProperty transforms the property on the given object type
func (transformer *TypeTransformer) TransformProperty(
	name astmodel.InternalTypeName,
	objectType *astmodel.ObjectType,
) (*PropertyTransformResult, error) {
	if !transformer.AppliesToType(name) {
		return nil, nil
	}

	found := false
	var propName astmodel.PropertyName
	var newProps []*astmodel.PropertyDefinition
	var newPropertyType astmodel.Type

	for _, prop := range objectType.Properties().AsSlice() {
		if transformer.Property.Matches(string(prop.PropertyName())).Matched &&
			(transformer.IfType == nil || transformer.IfType.AppliesToType(prop.PropertyType())) {

			found = true
			propName = prop.PropertyName()

			if transformer.Target != nil {
				propertyType, err := transformer.Target.produceTargetType(string(propName), prop.PropertyType())
				if err != nil {
					return nil, eris.Wrapf(err, "transforming property %s", propName)
				}

				newProps = append(newProps, prop.WithType(propertyType))
				newPropertyType = propertyType
			}
			// Otherwise, this is a removal - we don't copy the prop across.
		} else {
			newProps = append(newProps, prop)
		}
	}

	if !found {
		return nil, nil
	}

	// Ensure we've JIT created the map
	if transformer.matchedProperties == nil {
		transformer.matchedProperties = make(map[astmodel.TypeName]string)
	}

	transformer.matchedProperties[name] = propName.String()

	result := PropertyTransformResult{
		TypeName: name,
		NewType:  objectType.WithoutProperties().WithProperties(newProps...),
		Property: propName,
		Because:  transformer.Because,
	}

	if transformer.Remove {
		result.Removed = true
	} else {
		result.NewPropertyType = newPropertyType
	}

	return &result, nil
}

// TransformProperty transforms the property on the given object type
func (transformer *TypeTransformer) transformProperties(
	objectType *astmodel.ObjectType,
) (*astmodel.ObjectType, error) {
	for _, prop := range objectType.Properties().AsSlice() {
		propertyName := string(prop.PropertyName())
		if !transformer.Property.Matches(propertyName).Matched {
			// Skip properties where the name doesn't match
			continue
		}

		if transformer.IfType != nil && !transformer.IfType.AppliesToType(prop.PropertyType()) {
			// Skip properties with the wrong type
			continue
		}

		if transformer.Remove {
			objectType = objectType.WithoutProperty(prop.PropertyName())
		} else if transformer.Target != nil {
			propertyType, err := transformer.Target.produceTargetType(propertyName, prop.PropertyType())
			if err != nil {
				return nil, eris.Wrapf(err, "transforming property %s", propertyName)
			}

			// Can't specify both required and optional
			if transformer.Target.Optional && transformer.Target.Required {
				return nil, eris.Errorf("transformer for %s must not specify both required and optional", propertyName)
			}

			if transformer.Target.Optional {
				prop = prop.MakeOptional()
			} else if transformer.Target.Required {
				prop = prop.MakeRequired()
			}

			objectType = objectType.WithProperty(prop.WithType(propertyType))
		}
	}

	return objectType, nil
}
