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
	PackagePath string `yaml:",omitempty"`
	Name        string `yaml:",omitempty"`
	Map         *MapType
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

	Target TransformTarget

	// This is used purely for "caching" the actual astmodel type after Initialize()
	targetType astmodel.Type
}

func produceTargetType(target TransformTarget, descriptor string) (astmodel.Type, error) {
	if target.Name != "" && target.Map != nil {
		return nil, errors.Errorf("multiple target types defined")
	}

	if target.Name != "" {
		if target.PackagePath == "" {
			return primitiveTypeTarget(target.Name)
		}
		return astmodel.MakeTypeName(
			astmodel.MakePackageReference(target.PackagePath),
			target.Name), nil
	}

	if target.Map != nil {
		keyType, err := produceTargetType(target.Map.Key, descriptor+"/map/key")
		if err != nil {
			return nil, err
		}

		valueType, err := produceTargetType(target.Map.Value, descriptor+"/map/value")
		if err != nil {
			return nil, err
		}

		return astmodel.NewMapType(keyType, valueType), nil
	}

	return nil, errors.Errorf("no target type found in %s", descriptor)
}

func (transformer *TypeTransformer) Initialize() error {
	err := transformer.TypeMatcher.Initialize()
	if err != nil {
		return err
	}

	transformer.propertyRegex = createGlobbingRegex(transformer.Property)

	targetType, err := produceTargetType(transformer.Target, "target")
	if err != nil {
		return errors.Wrapf(
			err,
			"type transformer for group: %s, version: %s, name: %s",
			transformer.Group,
			transformer.Version,
			transformer.Name)
	}

	transformer.targetType = targetType
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
	name := typeName.Name()

	if typeName.PackageReference.IsLocalPackage() {
		group, version, err := typeName.PackageReference.GroupAndPackage()
		if err != nil {
			// This shouldn't ever happen because IsLocalPackage is true -- checking just to be safe
			panic(fmt.Sprintf("%s was flagged as a local package but has no group and package", typeName.PackageReference))
		}

		if transformer.groupMatches(group) && transformer.versionMatches(version) && transformer.nameMatches(name) {
			return transformer.targetType
		}
	} else {
		// TODO: Support external types better rather than doing everything in terms of GVK?
		if transformer.nameMatches(name) {
			return transformer.targetType
		}
	}

	// Didn't match so return nil
	return nil
}
