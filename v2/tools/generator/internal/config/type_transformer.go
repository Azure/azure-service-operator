/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package config

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
)

// A TransformTarget represents the target of a transformation
type TransformTarget struct {
	Group        FieldMatcher `yaml:",omitempty"`
	Version      FieldMatcher `yaml:"version,omitempty"`
	Name         FieldMatcher `yaml:",omitempty"`
	Optional     bool         `yaml:",omitempty"`
	Map          *MapType
	actualType   astmodel.Type
	appliesCache map[astmodel.Type]bool // cache for the results of AppliesToType()
}

type MapType struct {
	Key   TransformTarget `yaml:",omitempty"`
	Value TransformTarget `yaml:",omitempty"`
}

// A TypeTransformer is used to remap types
type TypeTransformer struct {
	TypeMatcher `yaml:",inline"`

	// Property is a wildcard matching specific properties on the types selected by this filter
	Property FieldMatcher `yaml:",omitempty"`

	// IfType only performs the transform if the original type matches (only usable with Property at the moment)
	IfType *TransformTarget `yaml:"ifType,omitempty"`

	// Target is the type to turn the type into
	Target *TransformTarget `yaml:",omitempty"`

	// Remove indicates that the property should be removed from the
	// type. This is only usable with Property. Target and Remove are
	// mutually exclusive.
	Remove bool `yaml:",omitempty"`

	// makeLocalPackageReferenceFunc is a function creating a local package reference
	makeLocalPackageReferenceFunc func(group string, version string) astmodel.LocalPackageReference

	// matchedProperties is a map of types that were matched to the property found on the type.
	// This is used to ensure that the type transformer matched at least one property
	matchedProperties map[astmodel.TypeName]string
}

func (target *TransformTarget) AppliesToType(t astmodel.Type) bool {
	if target.appliesCache == nil {
		target.appliesCache = make(map[astmodel.Type]bool)
	}

	if result, ok := target.appliesCache[t]; ok {
		return result
	}

	result := target.appliesToType(t)
	target.appliesCache[t] = result
	return result
}

func (target *TransformTarget) appliesToType(t astmodel.Type) bool {
	if target == nil {
		return true
	}

	inspect := t
	if target.Optional {
		// Need optional
		opt, ok := astmodel.AsOptionalType(inspect)
		if !ok {
			// but don't have optional
			return false
		}

		inspect = opt.Element()
	}

	if target.Name.IsRestrictive() {
		if target.Group.IsRestrictive() || target.Version.IsRestrictive() {
			// Expecting TypeName
			if tn, ok := astmodel.AsTypeName(inspect); ok {
				return target.appliesToTypeName(tn)
			}

			return false
		}

		// Expecting primitive type
		if pt, ok := astmodel.AsPrimitiveType(inspect); ok {
			return target.appliesToPrimitiveType(pt)
		}

		return false
	}

	if target.Map != nil {
		// Expecting map type
		if mp, ok := astmodel.AsMapType(inspect); ok {
			return target.appliesToMapType(mp)
		}

		return false
	}

	return true
}

func (target *TransformTarget) appliesToTypeName(tn astmodel.TypeName) bool {
	if !target.Name.Matches(tn.Name()) {
		// No match on name
		return false
	}

	g, v := tn.PackageReference.GroupVersion()

	if target.Group.IsRestrictive() {
		if !target.Group.Matches(g) {
			// No match on group
			return false
		}
	}

	if target.Version.IsRestrictive() {

		// Need to handle both full (v1beta20200101) and API (2020-01-01) formats
		switch ref := tn.PackageReference.(type) {
		case astmodel.LocalPackageReference:
			if !ref.HasApiVersion(target.Version.String()) && !target.Version.Matches(v) {
				return false
			}
		case astmodel.StoragePackageReference:
			if !ref.Local().HasApiVersion(target.Version.String()) && target.Version.Matches(v) {
				return false
			}
		default:
			return false
		}
	}

	return true
}

func (target *TransformTarget) appliesToPrimitiveType(pt *astmodel.PrimitiveType) bool {
	if target.Name.Matches(pt.Name()) {
		return true
	}

	// Special case, allowing config to use `any` as a synonym for `interface{}`
	if strings.EqualFold(target.Name.String(), "any") &&
		strings.EqualFold(pt.Name(), astmodel.AnyType.Name()) {
		return true
	}

	return false
}

func (target *TransformTarget) appliesToMapType(mp *astmodel.MapType) bool {
	return target.Map.Key.AppliesToType(mp.KeyType()) &&
		target.Map.Value.AppliesToType(mp.ValueType())
}

func (target *TransformTarget) assignActualType(
	descriptor string,
	makeLocalPackageReferenceFunc func(group string, version string) astmodel.LocalPackageReference) error {
	t, err := target.produceTargetType(descriptor, makeLocalPackageReferenceFunc)
	if err != nil {
		return err
	}

	target.actualType = t
	return nil
}

func (target *TransformTarget) produceTargetType(
	descriptor string,
	makeLocalPackageReferenceFunc func(group string, version string) astmodel.LocalPackageReference) (astmodel.Type, error) {
	if target.Name.IsRestrictive() && target.Map != nil {
		return nil, errors.Errorf("multiple target types defined")
	}

	var result astmodel.Type

	if target.Name.IsRestrictive() {
		if target.Group.IsRestrictive() || target.Version.IsRestrictive() {
			result = astmodel.MakeTypeName(
				makeLocalPackageReferenceFunc(target.Group.String(), target.Version.String()),
				target.Name.String())
		} else {
			var err error
			result, err = target.asPrimitiveType(target.Name.String())
			if err != nil {
				return nil, err
			}
		}
	}

	if target.Map != nil {
		keyType, err := target.Map.Key.produceTargetType(descriptor+"/map/key", makeLocalPackageReferenceFunc)
		if err != nil {
			return nil, err
		}

		valueType, err := target.Map.Value.produceTargetType(descriptor+"/map/value", makeLocalPackageReferenceFunc)
		if err != nil {
			return nil, err
		}

		result = astmodel.NewMapType(keyType, valueType)
	}

	if result == nil {
		return nil, errors.Errorf("no target type found in %s", descriptor)
	}

	if target.Optional {
		result = astmodel.NewOptionalType(result)
	}

	return result, nil
}

func (transformer *TypeTransformer) Initialize(makeLocalPackageReferenceFunc func(group string, version string) astmodel.LocalPackageReference) error {
	transformer.makeLocalPackageReferenceFunc = makeLocalPackageReferenceFunc
	err := transformer.TypeMatcher.Initialize()
	if err != nil {
		return err
	}
	transformer.matchedProperties = make(map[astmodel.TypeName]string)

	if transformer.Remove {
		if !transformer.Property.IsRestrictive() {
			return errors.Errorf("remove is only usable with property matches")
		}
		if transformer.Target != nil {
			return errors.Errorf("remove and target can't both be set")
		}
	} else {
		if transformer.Target == nil {
			return errors.Errorf("no target type and remove is not set")
		}
	}

	if transformer.IfType != nil {
		if !transformer.Property.IsRestrictive() {
			return errors.Errorf("ifType is only usable with property matches (for now)")
		}

		err := transformer.IfType.assignActualType("ifType", transformer.makeLocalPackageReferenceFunc)
		if err != nil {
			return err
		}
	}

	if transformer.Target != nil {
		err := transformer.Target.assignActualType("target", transformer.makeLocalPackageReferenceFunc)
		if err != nil {
			return errors.Wrapf(
				err,
				"type transformer for group: %s, version: %s, name: %s",
				transformer.Group.String(),
				transformer.Version.String(),
				transformer.Name.String())
		}
	}

	return nil
}

func (target *TransformTarget) asPrimitiveType(name string) (astmodel.Type, error) {
	switch name {
	case "bool":
		return astmodel.BoolType, nil
	case "float":
		return astmodel.FloatType, nil
	case "int":
		return astmodel.IntType, nil
	case "string":
		return astmodel.StringType, nil
	case "any":
		return astmodel.AnyType, nil
	default:
		return nil, errors.Errorf("unknown primitive type transformation target: %s", name)
	}
}

// TransformTypeName transforms the type with the specified name into the TypeTransformer target type if
// the provided type name matches the pattern(s) specified in the TypeTransformer
func (transformer *TypeTransformer) TransformTypeName(typeName astmodel.TypeName) astmodel.Type {
	if transformer.AppliesToType(typeName) {
		return transformer.Target.actualType
	}

	// Didn't match so return nil
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

func (transformer *TypeTransformer) RequiredPropertiesWereMatched() error {
	// If this transformer applies to entire types (instead of just properties on types), we just defer to
	// transformer.MatchedRequiredTypes
	if !transformer.Property.IsRestrictive() {
		return transformer.RequiredTypesWereMatched()
	}

	if !*transformer.MatchRequired {
		return nil
	}

	if err := transformer.RequiredTypesWereMatched(); err != nil {
		return err
	}

	if err := transformer.Property.WasMatched(); err != nil {
		return errors.Wrap(
			err,
			"matched types but all properties were excluded by name or type")
	}

	return nil
}

// TransformProperty transforms the property on the given object type
func (transformer *TypeTransformer) TransformProperty(name astmodel.TypeName, objectType *astmodel.ObjectType) *PropertyTransformResult {
	if !transformer.AppliesToType(name) {
		return nil
	}

	found := false
	var propName astmodel.PropertyName
	var newProps []*astmodel.PropertyDefinition

	for _, prop := range objectType.Properties().AsSlice() {
		if transformer.Property.Matches(string(prop.PropertyName())) &&
			(transformer.IfType == nil || transformer.IfType.AppliesToType(prop.PropertyType())) {

			found = true
			propName = prop.PropertyName()

			if transformer.Target != nil && transformer.Target.actualType != nil {
				newProps = append(newProps, prop.WithType(transformer.Target.actualType))
			}
			// Otherwise, this is a removal - we don't copy the prop across.
		} else {
			newProps = append(newProps, prop)
		}
	}

	if !found {
		return nil
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
		result.NewPropertyType = transformer.Target.actualType
	}
	return &result
}
