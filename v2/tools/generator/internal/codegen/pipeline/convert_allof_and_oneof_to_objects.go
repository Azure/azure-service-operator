/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"context"
	"fmt"
	"strings"

	"github.com/pkg/errors"
	"golang.org/x/exp/maps"
	"k8s.io/klog/v2"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
)

// This is needed when we are processing a Resource that happens to be nested inside
// other types, and we are going to extract the type of the Resource to merge it
// with other types (e.g. when processing oneOf/allOf). In these cases we need to
// know if we are in a spec or status context, so we can pick out the correct "side"
// of the resource.
type resourceFieldSelector string

var (
	chooseSpec   resourceFieldSelector = "Spec"
	chooseStatus resourceFieldSelector = "Status"
)

// ConvertAllOfAndOneOfToObjects reduces the AllOfType and OneOfType to ObjectType
func ConvertAllOfAndOneOfToObjects(idFactory astmodel.IdentifierFactory) *Stage {
	return NewLegacyStage(
		"allof-anyof-objects",
		"Convert allOf and oneOf to object types",
		func(ctx context.Context, defs astmodel.TypeDefinitionSet) (astmodel.TypeDefinitionSet, error) {
			builder := astmodel.TypeVisitorBuilder{}

			// the context here is whether we are selecting spec or status fields
			builder.VisitAllOfType = func(this *astmodel.TypeVisitor, it *astmodel.AllOfType, ctx interface{}) (astmodel.Type, error) {
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

			builder.VisitOneOfType = func(this *astmodel.TypeVisitor, it *astmodel.OneOfType, ctx interface{}) (astmodel.Type, error) {

				// If the OneOf contains only one object type, we can just use that
				if it.Types().Len() == 1 {
					var only astmodel.Type
					it.Types().ForEach(func(t astmodel.Type, _ int) {
						only = t
					})

					if obj, ok := astmodel.AsObjectType(only); ok {
						return obj, nil
					}
				}

				// Otherwise synthesize
				synth := synthesizer{
					specOrStatus: ctx.(resourceFieldSelector),
					defs:         defs,
					idFactory:    idFactory,
				}

				// we want to preserve names of the inner types
				// even if they are converted to other (unnamed types)
				propNames, err := synth.getOneOfPropNames(it)
				if err != nil {
					return nil, err
				}

				// process children first so that allOfs are resolved
				result, err := astmodel.IdentityVisitOfOneOfType(this, it, ctx)
				if err != nil {
					return nil, err
				}

				if resultOneOf, ok := result.(*astmodel.OneOfType); ok {
					result = synth.oneOfObject(resultOneOf, propNames)
				}

				// we might end up with something that requires re-visiting
				return this.Visit(result, ctx)
			}

			builder.VisitResourceType = func(this *astmodel.TypeVisitor, it *astmodel.ResourceType, ctx interface{}) (astmodel.Type, error) {
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

			visitor := builder.Build()
			result := make(astmodel.TypeDefinitionSet)

			for _, def := range defs {
				resourceUpdater := chooseSpec
				// TODO: we need flags
				if def.Name().IsStatus() {
					resourceUpdater = chooseStatus
				}

				transformed, err := visitor.VisitDefinition(def, resourceUpdater)
				if err != nil {
					return nil, errors.Wrapf(err, "error processing type %s", def.Name())
				}

				result.Add(transformed)
			}

			return result, nil
		})
}

type synthesizer struct {
	specOrStatus resourceFieldSelector
	idFactory    astmodel.IdentifierFactory
	defs         astmodel.TypeDefinitionSet
}

type propertyNames struct {
	golang astmodel.PropertyName
	json   string

	// used to resolve conflicts:
	isGoodName bool
	depth      int
}

func (ns propertyNames) betterThan(other propertyNames) bool {
	if ns.isGoodName && !other.isGoodName {
		return true
	}

	if !ns.isGoodName && other.isGoodName {
		return false
	}

	// both are good or !good
	// return the name closer to the top
	// (i.e. “lower” in inheritance hierarchy)
	return ns.depth <= other.depth
}

func (s synthesizer) getOneOfPropNames(oneOf *astmodel.OneOfType) ([]propertyNames, error) {
	var result []propertyNames

	err := oneOf.Types().ForEachError(func(t astmodel.Type, ix int) error {
		name, err := s.getOneOfName(t, ix)
		if err == nil {
			result = append(result, name)
		}

		return err
	})

	return simplifyPropNames(result), err
}

// simplifyPropNames makes the OneOf names better by trimming any common suffix
func simplifyPropNames(names []propertyNames) []propertyNames {
	if len(names) == 1 {
		return names
	}

	trim := ""
	for nameIx, name := range names {
		if nameIx == 0 {
			trim = string(name.golang)
		} else {
			trim = commonUppercasedSuffix(trim, string(name.golang))
		}
	}

	if trim == "" {
		return names // nothing to do
	}

	result := make([]propertyNames, len(names))
	for ix := range names {
		it := names[ix]
		newName := strings.TrimSuffix(string(it.golang), trim)
		if newName == "" {
			return names // trimming common suffix would result in one name being empty, so trim nothing
		}

		it.golang = astmodel.PropertyName(newName)
		result[ix] = it
	}

	return result
}

func min(x, y int) int {
	if x < y {
		return x
	}

	return y
}

// commonUppercasedSuffix returns the longest common suffix that
// starts with an uppercase letter
func commonUppercasedSuffix(x, y string) string {
	ix := 1
	lastFound := -1
	for ix <= min(len(x), len(y)) {
		cx := x[len(x)-ix]
		cy := y[len(y)-ix]
		if cx != cy {
			break
		}
		if cx >= 'A' && cx <= 'Z' {
			lastFound = ix
		}
		ix++
	}

	if lastFound >= 0 {
		return x[len(x)-lastFound:]
	}

	return ""
}

func (s synthesizer) getOneOfName(t astmodel.Type, propIndex int) (propertyNames, error) {
	switch concreteType := t.(type) {
	case astmodel.TypeName:
		// JSON name is unimportant here because we will implement the JSON marshaller anyway,
		// but we still need it for controller-gen
		return propertyNames{
			golang:     s.idFactory.CreatePropertyName(concreteType.Name(), astmodel.Exported),
			json:       s.idFactory.CreateIdentifier(concreteType.Name(), astmodel.NotExported),
			isGoodName: true, // a typename name is good (everything else is not)
		}, nil
	case *astmodel.EnumType:
		// JSON name is unimportant here because we will implement the JSON marshaller anyway,
		// but we still need it for controller-gen
		name := fmt.Sprintf("enum%d", propIndex)
		return propertyNames{
			golang:     s.idFactory.CreatePropertyName(name, astmodel.Exported),
			json:       s.idFactory.CreateIdentifier(name, astmodel.NotExported),
			isGoodName: false, // TODO: This name sucks but what alternative do we have?
		}, nil
	case *astmodel.ObjectType:
		name := fmt.Sprintf("object%d", propIndex)
		return propertyNames{
			golang:     s.idFactory.CreatePropertyName(name, astmodel.Exported),
			json:       s.idFactory.CreateIdentifier(name, astmodel.NotExported),
			isGoodName: false, // TODO: This name sucks but what alternative do we have?
		}, nil
	case *astmodel.MapType:
		name := fmt.Sprintf("map%d", propIndex)
		return propertyNames{
			golang:     s.idFactory.CreatePropertyName(name, astmodel.Exported),
			json:       s.idFactory.CreateIdentifier(name, astmodel.NotExported),
			isGoodName: false, // TODO: This name sucks but what alternative do we have?
		}, nil

	case *astmodel.ValidatedType:
		// pass-through to inner type
		return s.getOneOfName(concreteType.ElementType(), propIndex)

	case *astmodel.PrimitiveType:
		var primitiveTypeName string
		if concreteType == astmodel.AnyType {
			primitiveTypeName = "anything"
		} else {
			primitiveTypeName = concreteType.Name()
		}

		name := fmt.Sprintf("%s%d", primitiveTypeName, propIndex)
		return propertyNames{
			golang:     s.idFactory.CreatePropertyName(name, astmodel.Exported),
			json:       s.idFactory.CreateIdentifier(name, astmodel.NotExported),
			isGoodName: false, // TODO: This name sucks but what alternative do we have?
		}, nil
	case *astmodel.ResourceType:
		name := fmt.Sprintf("resource%d", propIndex)
		return propertyNames{
			golang:     s.idFactory.CreatePropertyName(name, astmodel.Exported),
			json:       s.idFactory.CreateIdentifier(name, astmodel.NotExported),
			isGoodName: false, // TODO: This name sucks but what alternative do we have?
		}, nil

	case *astmodel.AllOfType:
		var result *propertyNames
		err := concreteType.Types().ForEachError(func(t astmodel.Type, ix int) error {
			inner, err := s.getOneOfName(t, ix)
			if err != nil {
				return err
			}

			if result == nil || inner.betterThan(*result) {
				result = &inner
			}

			return nil
		})
		if err != nil {
			return propertyNames{}, err
		}

		if result != nil {
			result.depth += 1
			return *result, nil
		}

		return propertyNames{}, errors.New("unable to produce name for AllOf")

	default:
		return propertyNames{}, errors.Errorf("unexpected oneOf member, type: %T", t)
	}
}

func (s synthesizer) oneOfObject(oneOf *astmodel.OneOfType, propNames []propertyNames) astmodel.Type {
	// If there's more than one option, synthesize a type.
	// Note that this is required because Kubernetes CRDs do not support OneOf the same way
	// OpenAPI does, see https://github.com/Azure/azure-service-operator/issues/1515
	var properties []*astmodel.PropertyDefinition

	commonProperties, err := s.findOneOfPropertyObject(oneOf)
	if err != nil {
		return astmodel.NewErroredType(oneOf, []string{err.Error()}, nil)
	}

	// propertyType is a local function to combine a oneof subtype with any common properties we've found earlier
	propertyType := func(t astmodel.Type) astmodel.Type {
		if commonProperties == nil {
			return t
		}

		result, err := s.intersectTypes(t, commonProperties)
		if err != nil {
			return astmodel.NewErroredType(t, []string{err.Error()}, nil)
		}

		return result
	}

	propertyDescription := "Mutually exclusive with all other properties"
	oneOf.Types().ForEach(func(t astmodel.Type, ix int) {
		// Don't create a property for the object type containing our common properties
		if t == commonProperties {
			return
		}

		names := propNames[ix]
		prop := astmodel.NewPropertyDefinition(names.golang, names.json, propertyType(t))
		prop = prop.MakeTypeOptional()
		prop = prop.WithDescription(propertyDescription)
		properties = append(properties, prop)
	})

	objectType := astmodel.NewObjectType().WithProperties(properties...)

	// We need this information later so save it as a flag
	result := astmodel.OneOfFlag.ApplyTo(objectType)

	return result
}

// findOneOfPropertyObject looks for a nested ObjectType that contains properties specific to this OneOf that are in
// common to all the subtypes. We need to push these properties down into each of the subtypes so that they can be
// serialized correctly.
// If we find exactly one nested ObjectType, it's returned.
// If we find multiple, that's an error.
// If we find none, that's ok.
func (s synthesizer) findOneOfPropertyObject(oneOf *astmodel.OneOfType) (*astmodel.ObjectType, error) {
	var result *astmodel.ObjectType
	err := oneOf.Types().ForEachError(func(t astmodel.Type, ix int) error {
		if obj, ok := astmodel.AsObjectType(t); ok {
			if result == nil {
				result = obj
				return nil
			}

			return errors.Errorf("found multiple nested ObjectTypes in OneOf")
		}

		// Not an object, keep looking
		return nil
	})

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s synthesizer) intersectTypes(left astmodel.Type, right astmodel.Type) (astmodel.Type, error) {
	return intersector.MergeWithContext(s, left, right)
}

var intersector *astmodel.TypeMerger

func init() {
	i := astmodel.NewTypeMerger(func(_ctx interface{}, left, right astmodel.Type) (astmodel.Type, error) {
		return nil, errors.Errorf("don't know how to intersect types: %s and %s", left, right)
	})

	i.Add(synthesizer.handleEqualTypes)
	i.AddUnordered(synthesizer.handleValidatedAndNonValidated)
	i.AddUnordered(synthesizer.handleAnyType)
	i.AddUnordered(synthesizer.handleAllOfType)
	i.AddUnordered(synthesizer.handleTypeName)
	i.AddUnordered(synthesizer.handleOneOf)
	i.AddUnordered(synthesizer.handleARMIDAndString)

	i.Add(synthesizer.handleOptionalOptional)
	i.AddUnordered(synthesizer.handleOptional)

	i.Add(synthesizer.handleResourceResource)
	i.AddUnordered(synthesizer.handleResourceType)

	i.Add(synthesizer.handleEnumEnum)
	i.AddUnordered(synthesizer.handleEnum)

	i.Add(synthesizer.handleObjectObject)
	i.Add(synthesizer.handleMapMap)
	i.Add(synthesizer.handleArrayArray)

	i.AddUnordered(synthesizer.handleMapObject)
	i.AddUnordered(synthesizer.handleErrored)

	intersector = i
}

func (s synthesizer) handleErrored(leftErrored *astmodel.ErroredType, right astmodel.Type) (astmodel.Type, error) {
	// can merge the contents of an ErroredType, if we preserve the errors
	if leftErrored.InnerType() == nil {
		return leftErrored.WithType(right), nil
	}

	combined, err := s.intersectTypes(leftErrored.InnerType(), right)
	if combined == nil && err == nil {
		return nil, nil // unable to combine
	}

	if err != nil {
		return nil, err
	}

	return leftErrored.WithType(combined), nil
}

func (s synthesizer) handleOptional(leftOptional *astmodel.OptionalType, right astmodel.Type) (astmodel.Type, error) {
	// is this wrong? it feels wrong, but needed for {optional{enum}, string}
	return s.intersectTypes(leftOptional.Element(), right)
}

func (s synthesizer) handleResourceResource(leftResource *astmodel.ResourceType, rightResource *astmodel.ResourceType) (astmodel.Type, error) {
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

func (s synthesizer) handleResourceType(leftResource *astmodel.ResourceType, right astmodel.Type) (astmodel.Type, error) {
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

func (s synthesizer) handleOptionalOptional(leftOptional *astmodel.OptionalType, rightOptional *astmodel.OptionalType) (astmodel.Type, error) {
	// if both optional merge their contents and put back in an optional
	result, err := s.intersectTypes(leftOptional.Element(), rightOptional.Element())
	if err != nil {
		return nil, err
	}

	return astmodel.NewOptionalType(result), nil
}

func (s synthesizer) handleMapMap(leftMap *astmodel.MapType, rightMap *astmodel.MapType) (astmodel.Type, error) {
	keyType, err := s.intersectTypes(leftMap.KeyType(), rightMap.KeyType())
	if err != nil {
		return nil, err
	}

	valueType, err := s.intersectTypes(leftMap.ValueType(), rightMap.ValueType())
	if err != nil {
		return nil, err
	}

	return leftMap.WithKeyType(keyType).WithValueType(valueType), nil
}

// intersection of array types is array of intersection of their element types
func (s synthesizer) handleArrayArray(leftArray *astmodel.ArrayType, rightArray *astmodel.ArrayType) (astmodel.Type, error) {
	intersected, err := s.intersectTypes(leftArray.Element(), rightArray.Element())
	if err != nil {
		return nil, err
	}

	return leftArray.WithElement(intersected), nil
}

func max(left, right int) int {
	if left > right {
		return left
	}

	return right
}

func (s synthesizer) handleObjectObject(leftObj *astmodel.ObjectType, rightObj *astmodel.ObjectType) (astmodel.Type, error) {
	leftProperties := leftObj.Properties()
	rightProperties := rightObj.Properties()
	mergedProps := make(map[astmodel.PropertyName]*astmodel.PropertyDefinition, max(leftProperties.Len(), rightProperties.Len()))

	leftProperties.ForEach(func(p *astmodel.PropertyDefinition) {
		mergedProps[p.PropertyName()] = p
	})

	rightProperties.ForEach(func(p *astmodel.PropertyDefinition) {
		existingProp, ok := mergedProps[p.PropertyName()]
		if !ok {
			// Property doesn't already exist, so just add it
			mergedProps[p.PropertyName()] = p
			return // continue
		}

		newType, err := s.intersectTypes(existingProp.PropertyType(), p.PropertyType())
		if err != nil {
			klog.Errorf("unable to combine properties: %s (%s)", p.PropertyName(), err)
			return // continue
		}

		// TODO: need to handle merging requiredness and tags and...
		newProp := existingProp.WithType(newType)
		if len(p.Description()) > len(newProp.Description()) {
			// When we merge properties, both of them may have comments, and we need to deterministically choose one
			// of them to include on the final property. Our simple heuristic is to choose the longer comment, as
			// that's likely to be more specific.
			// TODO: Is there a better heuristic? See https://github.com/Azure/azure-service-operator/issues/1768
			newProp = newProp.WithDescription(p.Description())
		}

		mergedProps[p.PropertyName()] = newProp
	})

	// flatten
	properties := maps.Values(mergedProps)

	// TODO: need to handle merging other bits of objects
	isResource := leftObj.IsResource() || rightObj.IsResource()
	return leftObj.WithProperties(properties...).WithIsResource(isResource), nil
}

func (s synthesizer) handleEnumEnum(leftEnum *astmodel.EnumType, rightEnum *astmodel.EnumType) (astmodel.Type, error) {
	if !astmodel.TypeEquals(leftEnum.BaseType(), rightEnum.BaseType()) {
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

	return astmodel.NewEnumType(leftEnum.BaseType(), inBoth...), nil
}

func (s synthesizer) handleEnum(leftEnum *astmodel.EnumType, right astmodel.Type) (astmodel.Type, error) {
	// we can restrict from a (maybe optional) base type to an enum type
	if astmodel.TypeEquals(leftEnum.BaseType(), right) ||
		astmodel.TypeEquals(astmodel.NewOptionalType(leftEnum.BaseType()), right) {
		return leftEnum, nil
	}

	opts := leftEnum.Options()
	strs := make([]string, 0, len(opts))
	for _, enumValue := range opts {
		strs = append(strs, enumValue.String())
	}

	return nil, errors.Errorf("don't know how to merge enum type (%s) with %s", strings.Join(strs, ", "), right)
}

func (s synthesizer) handleAllOfType(leftAllOf *astmodel.AllOfType, right astmodel.Type) (astmodel.Type, error) {
	result, err := s.allOfObject(leftAllOf)
	if err != nil {
		return nil, err
	}

	return s.intersectTypes(result, right)
}

// if combining a type with a oneOf that contains that type, the result is that type
func (s synthesizer) handleOneOf(leftOneOf *astmodel.OneOfType, right astmodel.Type) (astmodel.Type, error) {
	// if there is an equal case, use that
	if leftOneOf.Types().Contains(right, astmodel.EqualityOverrides{}) {
		return right, nil
	}

	// otherwise intersect with each type:
	newTypes := astmodel.MakeTypeSet()
	err := leftOneOf.Types().ForEachError(func(lType astmodel.Type, _ int) error {
		newType, err := s.intersectTypes(lType, right)
		if err != nil {
			return err
		}

		newTypes.Add(newType)
		return nil
	})
	if err != nil {
		return nil, errors.Wrapf(err, "unable to intersect oneOf with %s", astmodel.DebugDescription(right))
	}

	if only, ok := newTypes.Only(); ok {
		return only, nil
	}

	return leftOneOf.WithTypes(newTypes.AsSlice()), nil
}

func (s synthesizer) handleTypeName(leftName astmodel.TypeName, right astmodel.Type) (astmodel.Type, error) {
	found, ok := s.defs[leftName]
	if !ok {
		return nil, errors.Errorf("couldn't find type %s", leftName)
	}

	result, err := s.intersectTypes(found.Type(), right)
	if err != nil {
		return nil, errors.Wrapf(
			err,
			"couldn't intersect %s with %s", leftName, astmodel.DebugDescription(right))
	}

	// TODO: can we somehow process these pointed-to types first,
	// so that their innards are resolved already and we can retain
	// the names?

	if astmodel.TypeEquals(result, found.Type()) {
		// if we got back the same thing we are referencing, preserve the reference
		return leftName, nil
	}

	return result, nil
}

// any type always disappears when intersected with another type
func (synthesizer) handleAnyType(left *astmodel.PrimitiveType, right astmodel.Type) (astmodel.Type, error) {
	if astmodel.TypeEquals(left, astmodel.AnyType) {
		return right, nil
	}

	return nil, nil
}

// two identical types can become the same type
func (synthesizer) handleEqualTypes(left astmodel.Type, right astmodel.Type) (astmodel.Type, error) {
	if astmodel.TypeEquals(left, right) {
		return left, nil
	}

	return nil, nil
}

// a validated and non-validated version of the same type become the validated version
func (synthesizer) handleValidatedAndNonValidated(validated *astmodel.ValidatedType, right astmodel.Type) (astmodel.Type, error) {
	if astmodel.TypeEquals(validated.ElementType(), right) {
		return validated, nil
	}

	// validated(optional(T)) combined with a non-optional T becomes validated(T)
	if astmodel.TypeEquals(validated.ElementType(), astmodel.NewOptionalType(right)) {
		return validated.WithElement(right), nil
	}

	return nil, nil
}

// An ARM ID and a string become the ARM ID
func (synthesizer) handleARMIDAndString(left astmodel.Type, right astmodel.Type) (astmodel.Type, error) {
	if astmodel.TypeEquals(left, astmodel.ARMIDType) && astmodel.TypeEquals(right, astmodel.StringType) {
		return astmodel.ARMIDType, nil
	}

	return nil, nil
}

// a string map and object can be combined with the map type becoming additionalProperties
func (synthesizer) handleMapObject(leftMap *astmodel.MapType, rightObj *astmodel.ObjectType) (astmodel.Type, error) {
	if astmodel.TypeEquals(leftMap.KeyType(), astmodel.StringType) {
		if rightObj.Properties().IsEmpty() {
			// no properties, treat as map
			// TODO: there could be other things in the object to check?
			return leftMap, nil
		}

		additionalProps := astmodel.NewPropertyDefinition(
			astmodel.AdditionalPropertiesPropertyName,
			astmodel.AdditionalPropertiesJsonName,
			leftMap)

		return rightObj.WithProperties(additionalProps), nil
	}

	return nil, nil
}

// makes an ObjectType for an AllOf type
func (s synthesizer) allOfObject(allOf *astmodel.AllOfType) (astmodel.Type, error) {
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
