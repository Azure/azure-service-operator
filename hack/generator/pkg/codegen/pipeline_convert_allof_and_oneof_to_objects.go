/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package codegen

import (
	"context"
	"fmt"
	"strings"

	"github.com/Azure/k8s-infra/hack/generator/pkg/astmodel"
	"github.com/pkg/errors"
	"k8s.io/klog/v2"
)

// This is needed when we are processing a Resource that happens to be nested inside
// other types and we are going to extract the type of the Resource to merge it
// with other types (e.g. when processing oneOf/allOf). In these cases we need to
// know if we are in a spec or status context so we can pick out the correct "side"
// of the resource.
type resourceFieldSelector string

var (
	chooseSpec   resourceFieldSelector = "Spec"
	chooseStatus resourceFieldSelector = "Status"
)

// convertAllOfAndOneOfToObjects reduces the AllOfType and OneOfType to ObjectType
func convertAllOfAndOneOfToObjects(idFactory astmodel.IdentifierFactory) PipelineStage {
	return MakePipelineStage(
		"allof-anyof-objects",
		"Convert allOf and oneOf to object types",
		func(ctx context.Context, defs astmodel.Types) (astmodel.Types, error) {
			visitor := astmodel.MakeTypeVisitor()

			// the context here is whether we are selecting spec or status fields
			visitor.VisitAllOfType = func(this *astmodel.TypeVisitor, it astmodel.AllOfType, ctx interface{}) (astmodel.Type, error) {
				synth := synthesizer{
					specOrStatus: ctx.(resourceFieldSelector),
					defs:         defs,
					idFactory:    idFactory,
				}

				object, err := synth.allOfObject(it)
				if err != nil {
					return nil, err
				}

				// we might end up with something that requires re-visiting
				// e.g. AllOf can turn into a OneOf that we then need to visit
				return this.Visit(object, ctx)
			}

			visitor.VisitOneOfType = func(this *astmodel.TypeVisitor, it astmodel.OneOfType, ctx interface{}) (astmodel.Type, error) {
				// process children first so that allOfs are resolved
				result, err := astmodel.IdentityVisitOfOneOfType(this, it, ctx)
				if err != nil {
					return nil, err
				}

				if resultOneOf, ok := result.(astmodel.OneOfType); ok {
					synth := synthesizer{
						specOrStatus: ctx.(resourceFieldSelector),
						defs:         defs,
						idFactory:    idFactory,
					}

					result, err = synth.oneOfObject(resultOneOf)
					if err != nil {
						return nil, errors.Wrapf(err, "unable to synthesize oneOf")
					}
				}

				// we might end up with something that requires re-visiting
				return this.Visit(result, ctx)
			}

			visitor.VisitResourceType = func(this *astmodel.TypeVisitor, it *astmodel.ResourceType, ctx interface{}) (astmodel.Type, error) {
				spec, err := this.Visit(it.SpecType(), chooseSpec)
				if err != nil {
					return nil, errors.Wrapf(err, "unable to visit resource spec type")
				}

				status, err := this.Visit(it.StatusType(), chooseStatus)
				if err != nil {
					return nil, errors.Wrapf(err, "unable to visit resource status type")
				}

				return it.WithSpec(spec).WithStatus(status), nil
			}

			result := make(astmodel.Types)

			for _, def := range defs {
				resourceUpdater := chooseSpec
				// TODO: we need flags
				if strings.HasSuffix(def.Name().Name(), "_Status") {
					resourceUpdater = chooseStatus
				}

				transformed, err := visitor.VisitDefinition(def, resourceUpdater)
				if err != nil {
					return nil, errors.Wrapf(err, "error processing type %v", def.Name())
				}

				result.Add(*transformed)
			}

			return result, nil
		})
}

type synthesizer struct {
	specOrStatus resourceFieldSelector
	idFactory    astmodel.IdentifierFactory
	defs         astmodel.Types
}

func (s synthesizer) oneOfObject(oneOf astmodel.OneOfType) (astmodel.Type, error) {
	// If there's more than one option, synthesize a type.
	// Note that this is required because Kubernetes CRDs do not support OneOf the same way
	// OpenAPI does, see https://github.com/Azure/k8s-infra/issues/71
	var properties []*astmodel.PropertyDefinition

	propertyDescription := "Mutually exclusive with all other properties"
	err := oneOf.Types().ForEachError(func(t astmodel.Type, i int) error {
		prop, err := s.convertToOneOfProperty(i, t)
		if err != nil {
			return err
		}

		prop = prop.MakeOptional()
		prop = prop.WithDescription(propertyDescription)

		properties = append(properties, prop)
		return nil
	})

	if err != nil {
		return nil, err
	}

	objectType := astmodel.NewObjectType().WithProperties(properties...)
	objectType = objectType.WithFunction(astmodel.JSONMarshalFunctionName, astmodel.NewOneOfJSONMarshalFunction(objectType, s.idFactory))

	return objectType, nil
}

func (s synthesizer) convertToOneOfProperty(propIndex int, from astmodel.Type) (*astmodel.PropertyDefinition, error) {
	switch concreteType := from.(type) {
	case astmodel.TypeName:
		propertyName := s.idFactory.CreatePropertyName(concreteType.Name(), astmodel.Exported)

		// JSON name is unimportant here because we will implement the JSON marshaller anyway,
		// but we still need it for controller-gen
		jsonName := s.idFactory.CreateIdentifier(concreteType.Name(), astmodel.NotExported)
		return astmodel.NewPropertyDefinition(propertyName, jsonName, concreteType), nil
	case *astmodel.EnumType:
		// TODO: This name sucks but what alternative do we have?
		name := fmt.Sprintf("enum%v", propIndex)
		propertyName := s.idFactory.CreatePropertyName(name, astmodel.Exported)

		// JSON name is unimportant here because we will implement the JSON marshaller anyway,
		// but we still need it for controller-gen
		jsonName := s.idFactory.CreateIdentifier(name, astmodel.NotExported)
		return astmodel.NewPropertyDefinition(propertyName, jsonName, concreteType), nil
	case *astmodel.ObjectType:
		// TODO: This name sucks but what alternative do we have?
		name := fmt.Sprintf("object%v", propIndex)
		propertyName := s.idFactory.CreatePropertyName(name, astmodel.Exported)

		// JSON name is unimportant here because we will implement the JSON marshaller anyway,
		// but we still need it for controller-gen
		jsonName := s.idFactory.CreateIdentifier(name, astmodel.NotExported)
		return astmodel.NewPropertyDefinition(propertyName, jsonName, concreteType), nil
	case *astmodel.PrimitiveType:
		var primitiveTypeName string
		if concreteType == astmodel.AnyType {
			primitiveTypeName = "anything"
		} else {
			primitiveTypeName = concreteType.Name()
		}

		// TODO: This name sucks but what alternative do we have?
		name := fmt.Sprintf("%v%v", primitiveTypeName, propIndex)
		propertyName := s.idFactory.CreatePropertyName(name, astmodel.Exported)

		// JSON name is unimportant here because we will implement the JSON marshaller anyway,
		// but we still need it for controller-gen
		jsonName := s.idFactory.CreateIdentifier(name, astmodel.NotExported)
		return astmodel.NewPropertyDefinition(propertyName, jsonName, concreteType).MakeOptional(), nil

	case *astmodel.ResourceType:
		name := fmt.Sprintf("resource%v", propIndex)
		propertyName := s.idFactory.CreatePropertyName(name, astmodel.Exported)
		jsonName := s.idFactory.CreateIdentifier(name, astmodel.NotExported)
		return astmodel.NewPropertyDefinition(propertyName, jsonName, concreteType).MakeOptional(), nil

	default:
		return nil, errors.Errorf("unexpected oneOf member, type: %T", from)
	}
}

func (s synthesizer) intersectTypes(left astmodel.Type, right astmodel.Type) (astmodel.Type, error) {
	for _, handler := range intersectHandlers {
		result, err := handler(s, left, right)
		if err != nil {
			return nil, errors.Wrapf(err, "error intersecting types")
		}

		if result != nil {
			return result, nil
		}
	}

	return nil, errors.Errorf("don't know how to intersect types: %s and %s", left, right)
}

// intersectHandler knows how to do intersection for one case only. It is biased to examine
// the LHS type (but some are symmetric and handle both sides at once). The handler can return
// nil,nil which indicates it doesn't apply to this combination of types.
type intersectHandler = func(synthesizer, astmodel.Type, astmodel.Type) (astmodel.Type, error)

// handles the same case for the right hand side
func flip(f intersectHandler) intersectHandler {
	return func(s synthesizer, left astmodel.Type, right astmodel.Type) (astmodel.Type, error) {
		return f(s, right, left)
	}
}

var intersectHandlers []intersectHandler

func init() {
	intersectHandlers = []intersectHandler{
		synthesizer.handleEqualTypes, // equality is symmetric
		synthesizer.handleAnyType, flip(synthesizer.handleAnyType),
		synthesizer.handleAllOfType, flip(synthesizer.handleAllOfType),
		synthesizer.handleTypeName, flip(synthesizer.handleTypeName),
		synthesizer.handleOneOf, flip(synthesizer.handleOneOf),
		synthesizer.handleOptionalOptional,                           // symmetric
		synthesizer.handleOptional, flip(synthesizer.handleOptional), // needs to be before Enum
		synthesizer.handleResourceResource, // symmetric
		synthesizer.handleResourceType, flip(synthesizer.handleResourceType),
		synthesizer.handleEnumEnum, // symmetric
		synthesizer.handleEnum, flip(synthesizer.handleEnum),
		synthesizer.handleObjectObject, // symmetric
		synthesizer.handleMapMap,       // symmetric
		synthesizer.handleMapObject, flip(synthesizer.handleMapObject),
	}
}

func (s synthesizer) handleOptional(left astmodel.Type, right astmodel.Type) (astmodel.Type, error) {
	leftOptional, ok := left.(*astmodel.OptionalType)
	if !ok {
		return nil, nil
	}

	// is this wrong? it feels wrong, but needed for {optional{enum}, string}
	return s.intersectTypes(leftOptional.Element(), right)
}

func (s synthesizer) handleResourceResource(left astmodel.Type, right astmodel.Type) (astmodel.Type, error) {
	leftResource, ok := left.(*astmodel.ResourceType)
	if !ok {
		return nil, nil
	}

	rightResource, ok := right.(*astmodel.ResourceType)
	if !ok {
		return nil, nil
	}

	// merge two resources: merge spec/status
	spec, err := s.intersectTypes(leftResource.SpecType(), rightResource.SpecType())
	if err != nil {
		return nil, err
	}

	// handle combinations of nil statuses
	var status astmodel.Type
	if leftResource.StatusType() != nil && rightResource.StatusType() != nil {
		status, err = s.intersectTypes(leftResource.StatusType(), rightResource.StatusType())
		if err != nil {
			return nil, err
		}
	} else if leftResource.StatusType() != nil {
		status = leftResource.StatusType()
	} else {
		status = rightResource.StatusType()
	}

	return leftResource.WithSpec(spec).WithStatus(status), nil
}

func (s synthesizer) handleResourceType(left astmodel.Type, right astmodel.Type) (astmodel.Type, error) {
	leftResource, ok := left.(*astmodel.ResourceType)
	if !ok {
		return nil, nil
	}

	if s.specOrStatus == chooseStatus {
		if leftResource.StatusType() != nil {
			newT, err := s.intersectTypes(leftResource.StatusType(), right)
			if err != nil {
				return nil, err
			}

			return leftResource.WithStatus(newT), nil
		} else {
			return leftResource.WithStatus(right), nil
		}
	} else if s.specOrStatus == chooseSpec {
		newT, err := s.intersectTypes(leftResource.SpecType(), right)
		if err != nil {
			return nil, err
		}

		return leftResource.WithSpec(newT), nil
	} else {
		panic("invalid specOrStatus")
	}
}

func (s synthesizer) handleOptionalOptional(left astmodel.Type, right astmodel.Type) (astmodel.Type, error) {
	// if both optional merge their contents and put back in an optional
	leftOptional, ok := left.(*astmodel.OptionalType)
	if !ok {
		return nil, nil
	}

	rightOptional, ok := right.(*astmodel.OptionalType)
	if !ok {
		return nil, nil
	}

	result, err := s.intersectTypes(leftOptional.Element(), rightOptional.Element())
	if err != nil {
		return nil, err
	}

	return astmodel.NewOptionalType(result), nil
}

func (s synthesizer) handleMapMap(left astmodel.Type, right astmodel.Type) (astmodel.Type, error) {
	leftMap, ok := left.(*astmodel.MapType)
	if !ok {
		return nil, nil
	}

	rightMap, ok := right.(*astmodel.MapType)
	if !ok {
		return nil, nil
	}

	keyType, err := s.intersectTypes(leftMap.KeyType(), rightMap.KeyType())
	if err != nil {
		return nil, err
	}

	valueType, err := s.intersectTypes(leftMap.ValueType(), rightMap.ValueType())
	if err != nil {
		return nil, err
	}

	return astmodel.NewMapType(keyType, valueType), nil
}

func (s synthesizer) handleObjectObject(left astmodel.Type, right astmodel.Type) (astmodel.Type, error) {
	leftObj, ok := left.(*astmodel.ObjectType)
	if !ok {
		return nil, nil
	}

	rightObj, ok := right.(*astmodel.ObjectType)
	if !ok {
		return nil, nil
	}

	mergedProps := make(map[astmodel.PropertyName]*astmodel.PropertyDefinition)

	for _, p := range leftObj.Properties() {
		mergedProps[p.PropertyName()] = p
	}

	for _, p := range rightObj.Properties() {
		if existingProp, ok := mergedProps[p.PropertyName()]; ok {
			newType, err := s.intersectTypes(existingProp.PropertyType(), p.PropertyType())
			if err != nil {
				klog.Errorf("unable to combine properties: %s (%v)", p.PropertyName(), err)
				continue
				//return nil, err
			}

			// TODO: need to handle merging requiredness and tags and...
			mergedProps[p.PropertyName()] = existingProp.WithType(newType)
		} else {
			mergedProps[p.PropertyName()] = p
		}
	}

	// flatten
	var properties []*astmodel.PropertyDefinition
	for _, p := range mergedProps {
		properties = append(properties, p)
	}

	// TODO: need to handle merging other bits of objects
	return leftObj.WithProperties(properties...), nil
}

func (s synthesizer) handleEnumEnum(left astmodel.Type, right astmodel.Type) (astmodel.Type, error) {
	leftEnum, ok := left.(*astmodel.EnumType)
	if !ok {
		return nil, nil
	}

	// if both are enums then the elements must be common
	rightEnum, ok := right.(*astmodel.EnumType)
	if !ok {
		return nil, nil
	}

	if !leftEnum.BaseType().Equals(rightEnum.BaseType()) {
		return nil, errors.Errorf("cannot merge enums with differing base types")
	}

	var inBoth []astmodel.EnumValue

	for _, option := range leftEnum.Options() {
		for _, otherOption := range rightEnum.Options() {
			if option == otherOption {
				inBoth = append(inBoth, option)
				break
			}
		}
	}

	return astmodel.NewEnumType(leftEnum.BaseType(), inBoth), nil
}

func (s synthesizer) handleEnum(left astmodel.Type, right astmodel.Type) (astmodel.Type, error) {
	leftEnum, ok := left.(*astmodel.EnumType)
	if !ok {
		return nil, nil
	}

	// we can restrict from a (maybe optional) base type to an enum type
	if leftEnum.BaseType().Equals(right) ||
		astmodel.NewOptionalType(leftEnum.BaseType()).Equals(right) {
		return leftEnum, nil
	}

	var strs []string
	for _, enumValue := range leftEnum.Options() {
		strs = append(strs, enumValue.String())
	}

	return nil, errors.Errorf("don't know how to merge enum type (%s) with %s", strings.Join(strs, ", "), right)
}

func (s synthesizer) handleAllOfType(left astmodel.Type, right astmodel.Type) (astmodel.Type, error) {
	leftAllOf, ok := left.(astmodel.AllOfType)
	if !ok {
		return nil, nil
	}

	result, err := s.allOfObject(leftAllOf)
	if err != nil {
		return nil, err
	}

	return s.intersectTypes(result, right)
}

// if combining a type with a oneOf that contains that type, the result is that type
func (s synthesizer) handleOneOf(left astmodel.Type, right astmodel.Type) (astmodel.Type, error) {
	leftOneOf, ok := left.(astmodel.OneOfType)
	if !ok {
		return nil, nil
	}

	// if there is an equal case, use that:
	{
		var result astmodel.Type
		leftOneOf.Types().ForEach(func(lType astmodel.Type, _ int) {
			if lType.Equals(right) {
				result = lType
			}
		})

		if result != nil {
			return result, nil
		}
	}

	// otherwise intersect with each type:
	var newTypes []astmodel.Type
	err := leftOneOf.Types().ForEachError(func(lType astmodel.Type, _ int) error {
		newType, err := s.intersectTypes(lType, right)
		if err != nil {
			return err
		}

		newTypes = append(newTypes, newType)
		return nil
	})

	if err != nil {
		return nil, err
	}

	return astmodel.MakeOneOfType(newTypes...), nil
}

func (s synthesizer) handleTypeName(left astmodel.Type, right astmodel.Type) (astmodel.Type, error) {
	leftName, ok := left.(astmodel.TypeName)
	if !ok {
		return nil, nil
	}

	if found, ok := s.defs[leftName]; !ok {
		return nil, errors.Errorf("couldn't find type %s", leftName)
	} else {
		result, err := s.intersectTypes(found.Type(), right)
		if err != nil {
			return nil, err
		}

		// TODO: can we somehow process these pointed-to types first,
		// so that their innards are resolved already and we can retain
		// the names?

		if result.Equals(found.Type()) {
			// if we got back the same thing we are referencing, preserve the reference
			return leftName, nil
		}

		return result, nil
	}
}

// any type always disappears when intersected with another type
func (synthesizer) handleAnyType(left astmodel.Type, right astmodel.Type) (astmodel.Type, error) {
	if left.Equals(astmodel.AnyType) {
		return right, nil
	}

	return nil, nil
}

// two identical types can become the same type
func (synthesizer) handleEqualTypes(left astmodel.Type, right astmodel.Type) (astmodel.Type, error) {
	if left.Equals(right) {
		return left, nil
	}

	return nil, nil
}

// a string map and object can be combined with the map type becoming additionalProperties
func (synthesizer) handleMapObject(left astmodel.Type, right astmodel.Type) (astmodel.Type, error) {
	if leftMap, ok := left.(*astmodel.MapType); ok {
		if leftMap.KeyType().Equals(astmodel.StringType) {
			if rightObj, ok := right.(*astmodel.ObjectType); ok {
				if len(rightObj.Properties()) == 0 {
					// no properties, treat as map
					// TODO: there could be other things in the object to check?
					return leftMap, nil
				}

				additionalProps := astmodel.NewPropertyDefinition("additionalProperties", "additionalProperties", leftMap)
				return rightObj.WithProperties(additionalProps), nil
			}
		}
	}

	return nil, nil
}

// makes an ObjectType for an AllOf type
func (s synthesizer) allOfObject(allOf astmodel.AllOfType) (astmodel.Type, error) {

	var intersection astmodel.Type = astmodel.AnyType
	err := allOf.Types().ForEachError(func(t astmodel.Type, _ int) error {
		var err error
		intersection, err = s.intersectTypes(intersection, t)
		return err
	})

	if err != nil {
		return nil, err
	}

	return intersection, nil
}
