/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package config

import (
	"fmt"
	"github.com/Azure/k8s-infra/hack/generator/pkg/astmodel"
)

type TransformTarget struct {
	PackagePath string `yaml:",omitempty"`
	Name        string `yaml:",omitempty"`
}

// A TypeTransformer is used to remap types
type TypeTransformer struct {
	TypeMatcher `yaml:",inline"`

	Target TransformTarget

	// This is used purely for "caching" the actual astmodel type after Initialize()
	targetType astmodel.Type
}

func (transformer *TypeTransformer) Initialize() error {
	err := transformer.TypeMatcher.Initialize()
	if err != nil {
		return err
	}

	if transformer.Target.Name == "" {
		return fmt.Errorf(
			"type transformer for group: %s, version: %s, name: %s is missing type name to transform to",
			transformer.Group,
			transformer.Version,
			transformer.Name)
	}

	// Type has no package -- must be a primitive type
	if transformer.Target.PackagePath == "" {
		return transformer.initializePrimitiveTypeTarget()
	}

	transformer.targetType = astmodel.NewTypeName(
		*astmodel.NewPackageReference(transformer.Target.PackagePath),
		transformer.Target.Name)
	return nil
}

func (transformer *TypeTransformer) initializePrimitiveTypeTarget() error {
	switch transformer.Target.Name {
	case "bool":
		transformer.targetType = astmodel.BoolType
	case "float":
		transformer.targetType = astmodel.FloatType
	case "int":
		transformer.targetType = astmodel.IntType
	case "string":
		transformer.targetType = astmodel.StringType
	default:
		return fmt.Errorf(
			"type transformer for group: %s, version: %s, name: %s has unknown"+
				"primtive type transformation target: %s",
			transformer.Group,
			transformer.Version,
			transformer.Name,
			transformer.Target.Name)

	}

	return nil
}

func (transformer *TypeTransformer) TransformTypeName(typeName *astmodel.TypeName) astmodel.Type {
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
