/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"context"
	"strings"

	"github.com/pkg/errors"
	kerrors "k8s.io/apimachinery/pkg/util/errors"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
)

// NameTypesForCRD - for CRDs all inner enums and objects and validated types must be named, so we do it here
func NameTypesForCRD(idFactory astmodel.IdentifierFactory) *Stage {
	return NewLegacyStage(
		"nameTypes",
		"Name inner types for CRD",
		func(ctx context.Context, definitions astmodel.TypeDefinitionSet) (astmodel.TypeDefinitionSet, error) {
			result := make(astmodel.TypeDefinitionSet)

			// this is a little bit of a hack, better way to do it?
			getDescription := func(typeName astmodel.TypeName) []string {
				if typeDef, ok := definitions[typeName]; ok {
					return typeDef.Description()
				}

				return []string{}
			}

			for typeName, typeDef := range definitions {

				newDefs, err := nameInnerTypes(typeDef, getDescription)
				if err != nil {
					return nil, errors.Wrapf(err, "failed to name inner definitions")
				}

				err = result.AddAllAllowDuplicates(newDefs)
				if err != nil {
					return nil, errors.Wrapf(err, "failed to add new definitions")
				}

				if _, ok := result[typeName]; !ok {
					// if we didn't regenerate the “input” type in nameInnerTypes then it won’t
					// have been added to the output; do it here
					result.Add(typeDef)
				}
			}

			return result, nil
		})
}

func nameInnerTypes(
	def astmodel.TypeDefinition,
	getDescription func(typeName astmodel.TypeName) []string,
) ([]astmodel.TypeDefinition, error) {
	var resultTypes []astmodel.TypeDefinition

	builder := astmodel.TypeVisitorBuilder{}
	builder.VisitEnumType = func(this *astmodel.TypeVisitor, it *astmodel.EnumType, ctx interface{}) (astmodel.Type, error) {
		hint := ctx.(nameHint)

		enumName := hint.AsTypeName(def.Name().PackageReference)
		namedEnum := astmodel.MakeTypeDefinition(enumName, it)
		namedEnum = namedEnum.WithDescription(getDescription(enumName))

		resultTypes = append(resultTypes, namedEnum)

		return namedEnum.Name(), nil
	}

	builder.VisitValidatedType = func(this *astmodel.TypeVisitor, v *astmodel.ValidatedType, ctx interface{}) (astmodel.Type, error) {
		// a validated type anywhere except directly under a property
		// must be named so that we can put the validations on it
		hint := ctx.(nameHint)
		newElementType, err := this.Visit(v.ElementType(), hint.WithBasePart("Validated"))
		if err != nil {
			return nil, err
		}

		name := hint.AsTypeName(def.Name().PackageReference)
		namedType := astmodel.MakeTypeDefinition(name, v.WithType(newElementType))
		resultTypes = append(resultTypes, namedType)
		return namedType.Name(), nil
	}

	builder.VisitFlaggedType = func(this *astmodel.TypeVisitor, it *astmodel.FlaggedType, ctx interface{}) (astmodel.Type, error) {
		// Because we're returning type names here, we need to look up the name returned by visit and wrap that with the correct flags
		hint := ctx.(nameHint)

		name, err := this.Visit(it.Element(), hint)
		if err != nil {
			return nil, err
		}

		// The above visit of ObjectType will have mutated resultTypes to include a mapping of the type name
		// to the object type. Because we need to preserve flag types, we must find the ObjectType and re-wrap
		// it in the flags it had before. Note that we cannot just bypass the ObjectType visit as it may make mutations
		// to the Object (to name the types of its properties) which we also need to preserve.
		// There are no words for how much I want LINQ right here
		var found astmodel.TypeDefinition
		for i, item := range resultTypes {
			if astmodel.TypeEquals(item.Name(), name) {
				found = item
				resultTypes[i] = resultTypes[len(resultTypes)-1]
				resultTypes = resultTypes[:len(resultTypes)-1]
				break
			}
		}

		resultTypes = append(resultTypes, found.WithType(it.WithElement(found.Type())))
		return name, nil
	}

	builder.VisitObjectType = func(this *astmodel.TypeVisitor, it *astmodel.ObjectType, ctx interface{}) (astmodel.Type, error) {
		hint := ctx.(nameHint)

		var errs []error
		var props []*astmodel.PropertyDefinition
		// first map the inner types:
		it.Properties().ForEach(func(prop *astmodel.PropertyDefinition) {
			propType := prop.PropertyType()
			propHint := hint.WithBasePart(string(prop.PropertyName()))
			if validated, ok := propType.(*astmodel.ValidatedType); ok {
				// handle validated types in properties specially,
				// they don't need to be named, so skip directly to element type
				newElementType, err := this.Visit(validated.ElementType(), propHint)
				if err != nil {
					errs = append(errs, err)
				} else {
					props = append(props, prop.WithType(validated.WithType(newElementType)))
				}
			} else {
				newPropType, err := this.Visit(propType, propHint)
				if err != nil {
					errs = append(errs, err)
				} else {
					props = append(props, prop.WithType(newPropType))
				}
			}
		})

		if len(errs) > 0 {
			return nil, kerrors.NewAggregate(errs)
		}

		objectName := hint.AsTypeName(def.Name().PackageReference)

		namedObjectType := astmodel.MakeTypeDefinition(objectName, it.WithProperties(props...))
		namedObjectType = namedObjectType.WithDescription(getDescription(objectName))

		resultTypes = append(resultTypes, namedObjectType)

		return namedObjectType.Name(), nil
	}

	builder.VisitResourceType = func(this *astmodel.TypeVisitor, it *astmodel.ResourceType, ctx interface{}) (astmodel.Type, error) {
		hint := ctx.(nameHint)

		spec, err := this.Visit(it.SpecType(), hint.WithSuffix(astmodel.SpecSuffix))
		if err != nil {
			return nil, errors.Wrapf(err, "failed to name spec type %s", it.SpecType())
		}

		var status astmodel.Type
		if it.StatusType() != nil {
			status, err = this.Visit(it.StatusType(), hint.WithSuffix(astmodel.StatusSuffix))
			if err != nil {
				return nil, errors.Wrapf(err, "failed to name status type %s", it.StatusType())
			}
		}

		resourceName := hint.AsTypeName(def.Name().PackageReference)

		it = it.WithSpec(spec).WithStatus(status)
		resource := astmodel.MakeTypeDefinition(resourceName, it).WithDescription(getDescription(resourceName))
		resultTypes = append(resultTypes, resource)

		return resource.Name(), nil
	}

	visitor := builder.Build()

	_, err := visitor.Visit(def.Type(), newNameHint(def.Name()))
	if err != nil {
		return nil, errors.Wrapf(err, "failed to name inner types of %s", def.Name())
	}

	return resultTypes, nil
}

type nameHint struct {
	baseName string
	suffix   string
}

func newNameHint(name astmodel.TypeName) nameHint {

	suffixesToFloat := []string{
		astmodel.SpecSuffix,
		astmodel.StatusSuffix,
		astmodel.ArmSuffix,
	}

	baseName := name.Name()
	var suffixes []string
	done := false
	for !done {
		done = true
		for _, s := range suffixesToFloat {
			if strings.HasSuffix(baseName, s) {
				baseName = strings.TrimSuffix(baseName, s)
				suffixes = append(
					[]string{strings.TrimPrefix(s, "_")},
					suffixes...)
				done = false
				break
			}
		}
	}

	result := nameHint{
		baseName: baseName,
		suffix:   strings.Join(suffixes, "_"),
	}

	return result
}

func (n nameHint) WithBasePart(part string) nameHint {
	return nameHint{
		baseName: n.baseName + "_" + part,
		suffix:   n.suffix,
	}
}

func (n nameHint) WithSuffix(suffix string) nameHint {
	return nameHint{
		baseName: n.baseName,
		suffix:   strings.TrimPrefix(suffix, "_"),
	}
}

func (n nameHint) String() string {
	if n.suffix != "" {
		return n.baseName + "_" + n.suffix
	}

	return n.baseName
}

func (n nameHint) AsTypeName(ref astmodel.PackageReference) astmodel.TypeName {
	if n.suffix != "" {
		return astmodel.MakeTypeName(ref, n.baseName+"_"+n.suffix)
	}
	return astmodel.MakeTypeName(ref, n.baseName)
}
