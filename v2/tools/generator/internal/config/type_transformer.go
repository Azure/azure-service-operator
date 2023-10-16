/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package config

import (
	"fmt"
	"strings"

	"github.com/go-logr/logr"
	"github.com/pkg/errors"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
)

// TransformTarget is used either to select a type for transformation, or to specify the result of that transformation
type TransformTarget struct {
	Group        FieldMatcher           `yaml:",omitempty"`
	Version      FieldMatcher           `yaml:"version,omitempty"`
	Name         FieldMatcher           `yaml:",omitempty"`
	Optional     bool                   `yaml:",omitempty"`
	Map          *MapType               `yaml:",omitempty"`
	Enum         *EnumType              `yaml:",omitempty"`
	appliesCache map[astmodel.Type]bool // cache for the results of AppliesToType()
}

type MapType struct {
	Key   TransformTarget `yaml:",omitempty"`
	Value TransformTarget `yaml:",omitempty"`
}

type EnumType struct {
	Base   string   `yaml:"base,omitempty"`
	Values []string `yaml:"values,omitempty"`
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
			if tn, ok := astmodel.AsInternalTypeName(inspect); ok {
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

func (target *TransformTarget) appliesToTypeName(tn astmodel.InternalTypeName) bool {
	if !target.Name.Matches(tn.Name()).Matched {
		// No match on name
		return false
	}

	grp, ver := tn.InternalPackageReference().GroupVersion()

	if target.Group.IsRestrictive() {
		if !target.Group.Matches(grp).Matched {
			// No match on group
			return false
		}
	}

	if target.Version.IsRestrictive() {

		// Need to handle both full (v1beta20200101) and API (2020-01-01) formats
		switch ref := tn.PackageReference().(type) {
		case astmodel.LocalPackageReference:
			if !ref.HasApiVersion(target.Version.String()) && !target.Version.Matches(ver).Matched {
				return false
			}
		case astmodel.StoragePackageReference:
			if !ref.Local().HasApiVersion(target.Version.String()) && target.Version.Matches(ver).Matched {
				return false
			}
		default:
			return false
		}
	}

	return true
}

func (target *TransformTarget) appliesToPrimitiveType(pt *astmodel.PrimitiveType) bool {
	if target.Name.Matches(pt.Name()).Matched {
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

func (target *TransformTarget) produceTargetType(
	descriptor string,
	original astmodel.Type,
) (astmodel.Type, error) {

	var result astmodel.Type

	if target.Name.IsRestrictive() {
		t, err := target.produceTargetNamedType(original)
		if err != nil {
			return nil, err
		}

		result = t
	}

	if target.Map != nil {
		t, err := target.produceTargetMapType(descriptor, original)
		if err != nil {
			return nil, err
		}

		result = t
	}

	if target.Enum != nil {
		t, err := target.produceTargetEnumType(descriptor)
		if err != nil {
			return nil, err
		}

		result = t
	}

	if result == nil {
		return nil, errors.Errorf("no target type found in %s", descriptor)
	}

	if target.Optional {
		result = astmodel.NewOptionalType(result)
	}

	return result, nil
}

func (target *TransformTarget) produceTargetNamedType(original astmodel.Type) (astmodel.Type, error) {
	// Transform to name, ensure we have no other transformation
	if target.Map != nil {
		return nil, errors.Errorf("cannot specify both Name transformation and Map transformation")
	}

	if target.Enum != nil {
		return nil, errors.Errorf("cannot specify both Name transformation and Enum transformation")
	}

	// If we have *only* a name and that name represents a primitive type, we should use that
	// primitive type instead of a named type
	if !target.Group.IsRestrictive() &&
		!target.Version.IsRestrictive() {
		result, err := target.asPrimitiveType(target.Name.String())
		if err == nil {
			return result, nil
		}

		// If we got an error, it's because the name wasn't a primitive type, so we fall through to treating it
		// as a named type
	}

	tn, ok := astmodel.AsInternalTypeName(original)
	if !ok {
		return nil, errors.Errorf(
			"cannot apply type transformation; expected InternalTypeName, but have %s",
			astmodel.DebugDescription(original))
	}

	result := tn
	if target.Group.IsRestrictive() || target.Version.IsRestrictive() {
		ref := target.produceTargetPackageReference(result.InternalPackageReference())
		result = result.WithPackageReference(ref)
	}

	return result.WithName(target.Name.String()), nil
}

func (target *TransformTarget) produceTargetMapType(
	descriptor string,
	original astmodel.Type,
) (astmodel.Type, error) {
	// Transform to map, ensure we have no other transformation
	if target.Name.IsRestrictive() {
		return nil, errors.Errorf("cannot specify both Name transformation and Map transformation")
	}

	if target.Enum != nil {
		return nil, errors.Errorf("cannot specify both Map transformation and Enum transformation")
	}

	keyType, err := target.Map.Key.produceTargetType(descriptor+"/map/key", original)
	if err != nil {
		return nil, err
	}

	valueType, err := target.Map.Value.produceTargetType(descriptor+"/map/value", original)
	if err != nil {
		return nil, err
	}

	result := astmodel.NewMapType(keyType, valueType)
	return result, nil
}

func (target *TransformTarget) produceTargetEnumType(
	_ string,
) (astmodel.Type, error) {
	// Transform to enum, ensure we have no other transformation
	if target.Name.IsRestrictive() {
		return nil, errors.Errorf("cannot specify both Name transformation and Enum transformation")
	}

	if target.Map != nil {
		return nil, errors.Errorf("cannot specify both Map transformation and Enum transformation")
	}

	if target.Enum.Base == "" {
		return nil, errors.Errorf("enum transformation requires a base type")
	}

	baseType, err := target.asPrimitiveType(target.Enum.Base)
	if err != nil {
		return nil, err
	}

	values := make([]astmodel.EnumValue, 0, len(target.Enum.Values))
	caser := cases.Title(language.English, cases.NoLower)
	for _, value := range target.Enum.Values {
		id := caser.String(value)
		v := value

		if baseType == astmodel.StringType {
			// String enums need to be quoted
			v = fmt.Sprintf("%q", value)
		}

		values = append(values, astmodel.EnumValue{
			Identifier: id,
			Value:      v,
		})
	}

	result := astmodel.NewEnumType(baseType, values...)
	return result, nil
}

func (target *TransformTarget) produceTargetPackageReference(ref astmodel.InternalPackageReference) astmodel.InternalPackageReference {
	switch t := ref.(type) {
	case astmodel.LocalPackageReference:
		if target.Group.IsRestrictive() || target.Version.IsRestrictive() {
			prefix := t.LocalPathPrefix()
			group := t.Group()
			versionPrefix := t.GeneratorVersion()
			version := t.Version()

			if target.Group.IsRestrictive() {
				group = target.Group.String()
			}

			if target.Version.IsRestrictive() {
				version = target.Version.String()
			}

			return astmodel.MakeLocalPackageReference(prefix, group, versionPrefix, version)
		}

		return ref

	case astmodel.StoragePackageReference:
		base := target.produceTargetPackageReference(t.Base())
		return astmodel.MakeStoragePackageReference(base)

	case astmodel.SubPackageReference:
		base := target.produceTargetPackageReference(t.Base())
		return astmodel.MakeSubPackageReference(t.PackageName(), base)
	}

	return ref
}

func (transformer *TypeTransformer) Initialize() error {
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

		_, err := transformer.IfType.produceTargetType("ifType", astmodel.InternalTypeName{})
		if err != nil {
			return err
		}
	}

	if transformer.Target != nil {
		_, err := transformer.Target.produceTargetType("target", astmodel.InternalTypeName{})
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

func (target *TransformTarget) asPrimitiveType(name string) (*astmodel.PrimitiveType, error) {
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
func (transformer *TypeTransformer) TransformTypeName(typeName astmodel.InternalTypeName) astmodel.Type {
	if transformer.AppliesToType(typeName) {
		result, err := transformer.Target.produceTargetType("target", typeName)
		if err != nil {
			// Temporary
			panic(err)
		}

		return result
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

// LogTo creates a log message for the transformation
func (r PropertyTransformResult) LogTo(log logr.Logger) {
	if r.Removed {
		log.V(2).Info(
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

	if !*transformer.MatchRequired {
		return nil
	}

	if err := transformer.RequiredTypesWereMatched(); err != nil {
		return err
	}

	if err := transformer.Property.WasMatched(); err != nil {
		return errors.Wrapf(
			err,
			"%s matched types but all properties were excluded by name or type",
			transformer.TypeMatcher.String())
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
					return nil, errors.Wrapf(err, "transforming property %s", propName)
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
