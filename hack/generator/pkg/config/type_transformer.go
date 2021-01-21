/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package config

import (
	"fmt"
	"regexp"

	"github.com/pkg/errors"

	"github.com/Azure/k8s-infra/hack/generator/pkg/astmodel"
)

// A transformation target
type TransformTarget struct {
	Group    string `yaml:",omitempty"`
	Version  string `yaml:"version,omitempty"`
	Name     string `yaml:",omitempty"`
	Optional bool   `yaml:",omitempty"`
	Map      *MapType
}

type MapType struct {
	Key   TransformTarget `yaml:",omitempty"`
	Value TransformTarget `yaml:",omitempty"`
}

// A TypeTransformer is used to remap types
type TypeTransformer struct {
	TypeMatcher `yaml:",inline"`

	// Property is a wildcard matching specific properties on the types selected by this filter
	Property      string `yaml:",omitempty"`
	propertyRegex *regexp.Regexp

	// IfType only performs the transform if the original type matches (only usable with Property at the moment)
	IfType *TransformTarget `yaml:"ifType,omitempty"`
	ifType astmodel.Type    // cache the astmodel type

	// Target is the type to turn the type into
	Target     *TransformTarget `yaml:",omitempty"`
	targetType astmodel.Type    // cache the astmodel type

	// Remove indicates that the property should be removed from the
	// type. This is only usable with Property. Target and Remove are
	// mutually exclusive.
	Remove bool `yaml:",omitempty"`

	// makeLocalPackageReferenceFunc is a function creating a local package reference
	makeLocalPackageReferenceFunc func(group string, version string) astmodel.LocalPackageReference
}

// TODO: passing this here is a bit icky
func produceTargetType(target TransformTarget, descriptor string, makeLocalPackageReferenceFunc func(group string, version string) astmodel.LocalPackageReference) (astmodel.Type, error) {
	if target.Name != "" && target.Map != nil {
		return nil, errors.Errorf("multiple target types defined")
	}

	var result astmodel.Type

	if target.Name != "" {
		if target.Group != "" && target.Version != "" {
			result = astmodel.MakeTypeName(
				makeLocalPackageReferenceFunc(target.Group, target.Version),
				target.Name)
		} else {
			var err error
			result, err = primitiveTypeTarget(target.Name)
			if err != nil {
				return nil, err
			}
		}
	}

	if target.Map != nil {
		keyType, err := produceTargetType(target.Map.Key, descriptor+"/map/key", makeLocalPackageReferenceFunc)
		if err != nil {
			return nil, err
		}

		valueType, err := produceTargetType(target.Map.Value, descriptor+"/map/value", makeLocalPackageReferenceFunc)
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

	if transformer.Remove {
		if transformer.Property == "" {
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

	transformer.propertyRegex = createGlobbingRegex(transformer.Property)
	if transformer.IfType != nil {
		if transformer.Property == "" {
			return errors.Errorf("ifType is only usable with property matches (for now)")
		}

		ifType, err := produceTargetType(*transformer.IfType, "ifType", transformer.makeLocalPackageReferenceFunc)
		if err != nil {
			return err
		}

		transformer.ifType = ifType
	}

	if transformer.Target != nil {
		targetType, err := produceTargetType(*transformer.Target, "target", transformer.makeLocalPackageReferenceFunc)
		if err != nil {
			return errors.Wrapf(
				err,
				"type transformer for group: %s, version: %s, name: %s",
				transformer.Group,
				transformer.Version,
				transformer.Name)
		}

		transformer.targetType = targetType
	}
	return nil
}

func primitiveTypeTarget(name string) (astmodel.Type, error) {
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

func (transformer *TypeTransformer) propertyNameMatches(propName astmodel.PropertyName) bool {
	return transformer.matches(transformer.Property, &transformer.propertyRegex, string(propName))
}

// TransformTypeName transforms the type with the specified name into the TypeTransformer target type if
// the provided type name matches the pattern(s) specified in the TypeTransformer
func (transformer *TypeTransformer) TransformTypeName(typeName astmodel.TypeName) astmodel.Type {
	if localRef, ok := typeName.PackageReference.AsLocalPackage(); ok {
		group := localRef.Group()
		version := localRef.Version()
		name := typeName.Name()

		if transformer.groupMatches(group) &&
			transformer.versionMatches(version) &&
			transformer.nameMatches(name) {
			return transformer.targetType
		}
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

// TransformProperty transforms the property on the given object type
func (transformer *TypeTransformer) TransformProperty(name astmodel.TypeName, objectType *astmodel.ObjectType) *PropertyTransformResult {
	if !transformer.AppliesToType(name) {
		return nil
	}

	found := false
	var propName astmodel.PropertyName
	var newProps []*astmodel.PropertyDefinition

	for _, prop := range objectType.Properties() {
		if transformer.propertyNameMatches(prop.PropertyName()) &&
			(transformer.ifType == nil || transformer.ifType.Equals(prop.PropertyType())) {

			found = true
			propName = prop.PropertyName()

			if transformer.targetType != nil {
				newProps = append(newProps, prop.WithType(transformer.targetType))
			}
			// Otherwise this is a removal - we don't copy the prop across.
		} else {
			newProps = append(newProps, prop)
		}
	}

	if !found {
		return nil
	}

	result := PropertyTransformResult{
		TypeName: name,
		NewType:  objectType.WithoutProperties().WithProperties(newProps...),
		Property: propName,
		Because:  transformer.Because,
	}

	if transformer.Remove {
		result.Removed = true
	} else {
		result.NewPropertyType = transformer.targetType
	}
	return &result
}
