/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package config

import (
	"fmt"

	"github.com/rotisserie/eris"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
)

// TransformResult is used to specify the result of a transformation
type TransformResult struct {
	Group    FieldMatcher `yaml:",omitempty"`
	Version  FieldMatcher `yaml:"version,omitempty"`
	Name     FieldMatcher `yaml:",omitempty"`
	Optional bool         `yaml:",omitempty"`
	Required bool         `yaml:",omitempty"`
	Map      *MapResult   `yaml:",omitempty"`
	Enum     *EnumResult  `yaml:",omitempty"`
}

type MapResult struct {
	Key   TransformResult `yaml:",omitempty"`
	Value TransformResult `yaml:",omitempty"`
}

type EnumResult struct {
	Base   string   `yaml:"base,omitempty"`
	Values []string `yaml:"values,omitempty"`
}

func (tr *TransformResult) produceTargetType(
	descriptor string,
	original astmodel.Type,
) (astmodel.Type, error) {
	if err := tr.validate(); err != nil {
		return nil, eris.Wrapf(err, "invalid transformation in %s", descriptor)
	}

	if tr == nil {
		// No transformation
		return original, nil
	}

	resultType := original

	if tr.Name.IsRestrictive() {
		t, err := tr.produceTargetNamedType(original)
		if err != nil {
			return nil, err
		}

		resultType = t
	}

	if tr.Map != nil {
		t, err := tr.produceTargetMapType(descriptor, original)
		if err != nil {
			return nil, err
		}

		resultType = t
	}

	if tr.Enum != nil {
		t, err := tr.produceTargetEnumType(descriptor)
		if err != nil {
			return nil, err
		}

		resultType = t
	}

	if resultType == nil {
		return nil, eris.Errorf("no target type found in %s", descriptor)
	}

	if tr.Optional {
		resultType = astmodel.NewOptionalType(resultType)
	}

	return resultType, nil
}

func (tr *TransformResult) produceTargetNamedType(original astmodel.Type) (astmodel.Type, error) {
	// Transform to name, ensure we have no other transformation
	if tr.Map != nil {
		return nil, eris.Errorf("cannot specify both Name transformation and Map transformation")
	}

	if tr.Enum != nil {
		return nil, eris.Errorf("cannot specify both Name transformation and Enum transformation")
	}

	// If we have *only* a name *and* that name represents a primitive type, we should use that
	// primitive type instead of a named type
	if !tr.Group.IsRestrictive() &&
		!tr.Version.IsRestrictive() {
		result, err := tr.asPrimitiveType(tr.Name.String())
		if err == nil {
			return result, nil
		}

		// If we got an error, it's because the name wasn't a primitive type, so we fall through to treating it
		// as a named type
	}

	tn, ok := astmodel.AsInternalTypeName(original)
	if !ok {
		return nil, eris.Errorf(
			"cannot apply type transformation; expected InternalTypeName, but have %s",
			astmodel.DebugDescription(original))
	}

	result := tn
	if tr.Group.IsRestrictive() || tr.Version.IsRestrictive() {
		ref := tr.produceTargetPackageReference(result.InternalPackageReference())
		result = result.WithPackageReference(ref)
	}

	return result.WithName(tr.Name.String()), nil
}

func (tr *TransformResult) produceTargetMapType(
	descriptor string,
	original astmodel.Type,
) (astmodel.Type, error) {
	// Transform to map, ensure we have no other transformation
	if tr.Name.IsRestrictive() {
		return nil, eris.Errorf("cannot specify both Name transformation and Map transformation")
	}

	if tr.Enum != nil {
		return nil, eris.Errorf("cannot specify both Map transformation and Enum transformation")
	}

	keyType, err := tr.Map.Key.produceTargetType(descriptor+"/map/key", original)
	if err != nil {
		return nil, err
	}

	valueType, err := tr.Map.Value.produceTargetType(descriptor+"/map/value", original)
	if err != nil {
		return nil, err
	}

	return astmodel.NewMapType(keyType, valueType), nil
}

func (tr *TransformResult) produceTargetEnumType(
	_ string,
) (astmodel.Type, error) {
	// Transform to enum, ensure we have no other transformation
	if tr.Name.IsRestrictive() {
		return nil, eris.Errorf("cannot specify both Name transformation and Enum transformation")
	}

	if tr.Map != nil {
		return nil, eris.Errorf("cannot specify both Map transformation and Enum transformation")
	}

	if tr.Enum.Base == "" {
		return nil, eris.Errorf("enum transformation requires a base type")
	}

	baseType, err := tr.asPrimitiveType(tr.Enum.Base)
	if err != nil {
		return nil, err
	}

	values := make([]astmodel.EnumValue, 0, len(tr.Enum.Values))
	caser := cases.Title(language.English, cases.NoLower)
	for _, value := range tr.Enum.Values {
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

	return astmodel.NewEnumType(baseType, values...), nil
}

func (tr *TransformResult) produceTargetPackageReference(ref astmodel.InternalPackageReference) astmodel.InternalPackageReference {
	switch t := ref.(type) {
	case astmodel.LocalPackageReference:
		if tr.Group.IsRestrictive() || tr.Version.IsRestrictive() {
			prefix := t.LocalPathPrefix()
			group := t.Group()
			versionPrefix := t.GeneratorVersion()
			version := t.Version()

			if tr.Group.IsRestrictive() {
				group = tr.Group.String()
			}

			if tr.Version.IsRestrictive() {
				version = tr.Version.String()
			}

			return astmodel.MakeLocalPackageReference(prefix, group, versionPrefix, version)
		}

		return ref

	case astmodel.SubPackageReference:
		base := tr.produceTargetPackageReference(t.Base())
		return astmodel.MakeSubPackageReference(t.PackageName(), base)
	}

	return ref
}

func (tr *TransformResult) asPrimitiveType(name string) (*astmodel.PrimitiveType, error) {
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
		return nil, eris.Errorf("unknown primitive type transformation target: %s", name)
	}
}

func (tr *TransformResult) validate() error {
	if !tr.Name.IsRestrictive() &&
		tr.Map == nil &&
		tr.Enum == nil &&
		!tr.Optional &&
		!tr.Required {
		return eris.Errorf("no result transformation specified")
	}

	return nil
}
