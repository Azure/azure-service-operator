/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"context"
	"fmt"
	kerrors "k8s.io/apimachinery/pkg/util/errors"
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
	return NewStage(
		"allof-anyof-objects",
		"Convert allOf and oneOf to object types",
		func(ctx context.Context, state *State) (*State, error) {

			baseSynthesizer := newSynthesizer(state.Definitions(), idFactory)

			newDefs := make(astmodel.TypeDefinitionSet)
			visitor := createVisitorForSynthesizer(baseSynthesizer)
			for _, def := range state.Definitions() {
				resourceUpdater := chooseSpec
				// TODO: we need flags
				if def.Name().IsStatus() {
					resourceUpdater = chooseStatus
				}

				transformed, err := visitor.VisitDefinition(def, resourceUpdater)
				if err != nil {
					return nil, errors.Wrapf(err, "error processing type %s", def.Name())
				}

				newDefs.Add(transformed)
			}

			finalDefs := newDefs.OverlayWith(baseSynthesizer.updatedDefs)
			return state.WithDefinitions(finalDefs), nil
		})
}

func createVisitorForSynthesizer(baseSynthesizer synthesizer) astmodel.TypeVisitor {
	builder := astmodel.TypeVisitorBuilder{}

	// the context here is whether we are selecting spec or status fields
	builder.VisitAllOfType = func(this *astmodel.TypeVisitor, it *astmodel.AllOfType, ctx interface{}) (astmodel.Type, error) {
		synth := baseSynthesizer.forField(ctx.(resourceFieldSelector))

		object, err := synth.allOfObject(it)
		if err != nil {
			return nil, errors.Wrapf(err, "creating object for allOf")
		}

		// we might end up with something that requires re-visiting
		// e.g. AllOf can turn into a OneOf that we then need to visit
		return this.Visit(object, ctx)
	}

	builder.VisitOneOfType = func(this *astmodel.TypeVisitor, it *astmodel.OneOfType, ctx interface{}) (astmodel.Type, error) {
		synth := baseSynthesizer.forField(ctx.(resourceFieldSelector))

		t, err := synth.oneOfToObject(it)
		if err != nil {
			return nil, errors.Wrapf(err, "creating object for oneOf")
		}

		// we might end up with something that requires re-visiting
		return this.Visit(t, ctx)
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

	builder.VisitErroredType = func(_ *astmodel.TypeVisitor, it *astmodel.ErroredType, ctx interface{}) (astmodel.Type, error) {
		// Nothing we can do to resolve errors, so just return the type as-is
		return it, nil
	}

	visitor := builder.Build()
	return visitor
}

type synthesizer struct {
	specOrStatus    resourceFieldSelector
	idFactory       astmodel.IdentifierFactory
	defs            astmodel.TypeDefinitionSet
	referenceCounts map[astmodel.TypeName]int
	activeNames     astmodel.TypeNameSet
	updatedDefs     astmodel.TypeDefinitionSet
}

// newSynthesizer returns a partially configured synthesizer which lacks a field selector
func newSynthesizer(
	defs astmodel.TypeDefinitionSet,
	idFactory astmodel.IdentifierFactory,
) synthesizer {
	return synthesizer{
		defs:            defs,
		idFactory:       idFactory,
		referenceCounts: countTypeReferences(defs),
		activeNames:     make(astmodel.TypeNameSet),
		updatedDefs:     make(astmodel.TypeDefinitionSet),
	}
}

// forField returns a synthesizer that will use the given field selector
func (s synthesizer) forField(specOrStatus resourceFieldSelector) synthesizer {
	result := s
	result.specOrStatus = specOrStatus
	return result
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
		newName = strings.Trim(newName, "_")

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
			isGoodName: true, // a typename name is good (little else is)
		}, nil

	case *astmodel.OneOfType:
		// If we have a name, use that
		if concreteType.Name() != "" {
			return propertyNames{
				golang:     s.idFactory.CreatePropertyName(concreteType.Name(), astmodel.Exported),
				json:       s.idFactory.CreateIdentifier(concreteType.Name(), astmodel.NotExported),
				isGoodName: true, // a oneOf name is good (little else is)
			}, nil
		}

		// If we have a discriminator value, use that as a name
		if concreteType.DiscriminatorValue() != "" {
			return propertyNames{
				golang:     s.idFactory.CreatePropertyName(concreteType.DiscriminatorValue(), astmodel.Exported),
				json:       s.idFactory.CreateIdentifier(concreteType.DiscriminatorValue(), astmodel.NotExported),
				isGoodName: true, // a discriminator value is good (little else is)
			}, nil
		}

		// Otherwise, if we only have one nested type, use that
		if only, ok := concreteType.Types().Single(); ok {
			return s.getOneOfName(only, propIndex)
		}

		// Otherwise return an error
		return propertyNames{}, errors.Errorf("expected nested oneOf member to have discriminator value, type: %T", t)

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

	case astmodel.MetaType:
		// Try unwrapping the meta type and basing the name on what's inside
		return s.getOneOfName(concreteType.Unwrap(), propIndex)

	case *astmodel.OptionalType:
		return s.getOneOfName(concreteType.Element(), propIndex)

	default:
		return propertyNames{}, errors.Errorf("unexpected oneOf member, type: %T", t)
	}
}

func (s synthesizer) oneOfToObject(
	oneOf *astmodel.OneOfType,
) (astmodel.Type, error) {

	if oneOf.DiscriminatorValue() != "" {
		// We have a leaf to assemble
		types := make([]astmodel.Type, 0, oneOf.Types().Len())
		oneOf.Types().ForEach(
			func(t astmodel.Type, _ int) {
				types = append(types, t)
			})
		allOf := astmodel.NewAllOfType(types...)
		result, err := s.allOfObject(allOf)

		return result, err
	}

	// Otherwise we have a root to assemble; we need to create a new object type to hold the oneOf
	// with properties for each of the leaves

	// Preserve names of the inner types
	propNames, err := s.getOneOfPropNames(oneOf)
	if err != nil {
		return nil, err
	}

	propertyDescription := "Mutually exclusive with all other properties"
	var properties []*astmodel.PropertyDefinition
	oneOf.Types().ForEach(func(t astmodel.Type, ix int) {
		names := propNames[ix]
		prop := astmodel.NewPropertyDefinition(names.golang, names.json, t)
		prop = prop.MakeTypeOptional()
		prop = prop.WithDescription(propertyDescription)
		properties = append(properties, prop)
	})

	objectType := astmodel.NewObjectType().WithProperties(properties...)

	// We need this information later so save it as a flag
	result := astmodel.OneOfFlag.ApplyTo(objectType)

	return result, nil
}

func (s synthesizer) pushCommonPropertiesToLeaves(oneOf *astmodel.OneOfType, commonProperties *astmodel.ObjectType) (*astmodel.OneOfType, error) {
	newTypes := make([]astmodel.Type, 0, oneOf.Types().Len())
	var errs []error
	oneOf.Types().ForEach(func(t astmodel.Type, _ int) {
		if _, ok := s.AsCommonProperties(t); ok {
			// This is a common properties object, so we don't need to do anything
			return
		}

		// This is a leaf object, so we need to add the common properties, then merge it with the leaf
		allOf := astmodel.NewAllOfType(t, commonProperties)
		t, err := s.allOfObject(allOf)
		if err != nil {
			errs = append(errs, err)
			return
		}

		newTypes = append(newTypes, t)
	})

	err := errors.Wrap(
		kerrors.NewAggregate(errs),
		"error pushing common properties to leaves of oneOf")
	if err != nil {
		return nil, err
	}

	return oneOf.WithTypes(newTypes), nil
}

// findCommonProperties looks for a nested ObjectType that contains common properties defined on this OneOf that need
// to be pushed to all the subtypes. We need to push these properties down into each of the subtypes so that they can be
// serialized correctly.
// If we find exactly one nested ObjectType, it's returned.
// If we find multiple, they're merged into a single ObjectType and returned.
// If we find none, that's ok.
func (s synthesizer) findCommonProperties(oneOf *astmodel.OneOfType) (*astmodel.ObjectType, error) {
	var result *astmodel.ObjectType

	err := oneOf.Types().ForEachError(func(t astmodel.Type, ix int) error {
		if obj, ok := s.AsCommonProperties(t); ok {
			if result == nil {
				result = obj
				return nil
			}

			// We already have a result, so we need to intersect this one with it
			intersection, err := s.intersectTypes(result, obj)
			if err != nil {
				return errors.Wrapf(err, "merging multiple nested ObjectTypes in OneOf")
			}

			newObj, ok := astmodel.AsObjectType(intersection)
			if !ok {
				return errors.New("expected merging two objects to return object")
			}

			result = newObj
		}

		// Not an object, keep looking
		return nil
	})

	if err != nil {
		return nil, err
	}

	return result, nil
}

// AsCommonProperties captures the test used to determine if a type is a common properties object, ensuring that
// we're consistent with our tests from multiple locations
func (s synthesizer) AsCommonProperties(t astmodel.Type) (*astmodel.ObjectType, bool) {
	return astmodel.AsObjectType(t)
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
	i.AddUnordered(synthesizer.handleFlaggedType)

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

	if only, ok := newTypes.Single(); ok {
		return only, nil
	}

	// Still have a OneOf, need to reprocess it
	oneOf := leftOneOf.WithTypes(newTypes.AsSlice())
	return s.oneOfToObject(oneOf)
}

func (s synthesizer) handleTypeName(leftName astmodel.TypeName, right astmodel.Type) (astmodel.Type, error) {
	found, ok := s.defs[leftName]
	if !ok {
		return nil, errors.Errorf("couldn't find type %s", leftName)
	}

	if leftName.Name() == "AKS_STATUS" {
		klog.Warningf("found AKS_STATUS")
	}

	if refs, ok := s.referenceCounts[leftName]; ok && refs == 1 && s.updatedDefs.Contains(leftName) {
		// s.updatedDefs only contains definitions for types that are referenced exactly once
		// If we're called a second time, it's because the type that contains this reference is referenced multiple
		// times. In that case, we're going to get the exact same answer this time, so we may as well short-circuit
		//
		// E.g. if we have:
		//
		// Foo references Bar
		// and Baz References Bar
		// and Bar references Zed
		//
		// Then we can end up handling Zed twice, once when we handle Foo and once when we handle Baz
		// But both times will have the exact same parameters (because they come from Bar) so we can just
		// short-circuit the second time
		return leftName, nil
	}

	if s.activeNames.Contains(leftName) {
		// Avoid recursion, just keep the existing TypeName
		return leftName, nil
	}

	// Track which TypeNames we're processing, so we can detect recursion
	s.activeNames.Add(leftName)
	defer func() { s.activeNames.Remove(leftName) }()

	result, err := s.intersectTypes(found.Type(), right)
	if err != nil {
		return nil, errors.Wrapf(
			err,
			"couldn't intersect %s with %s", leftName, astmodel.DebugDescription(right))
	}

	// TODO: can we somehow process these pointed-to types first,
	// so that their innards are resolved already and we can retain
	// the names?

	// If no change to the type, keep the typename
	if astmodel.TypeEquals(result, found.Type()) ||
		astmodel.TypeEquals(result, leftName) {
		// if we got back the same thing we already have, preserve the reference
		return leftName, nil
	}

	if refs, ok := s.referenceCounts[leftName]; ok && refs == 1 {
		// if this is the only reference to this type, we can redefine it
		s.updatedDefs.Add(found.WithType(result))
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
		return nil, errors.Wrapf(err, "combining AllOf with %d types", allOf.Types().Len())
	}

	return intersection, nil
}

func (s synthesizer) handleFlaggedType(left *astmodel.FlaggedType, right astmodel.Type) (astmodel.Type, error) {
	// Intersect the content and retain the flags
	internal, err := s.intersectTypes(left.Element(), right)
	if err != nil {
		return nil, err
	}

	return left.WithElement(internal), nil
}

func countTypeReferences(defs astmodel.TypeDefinitionSet) map[astmodel.TypeName]int {
	referenceCounts := make(map[astmodel.TypeName]int)

	visitor := astmodel.TypeVisitorBuilder{
		VisitTypeName: func(tn astmodel.TypeName) astmodel.Type {
			referenceCounts[tn]++
			return tn
		},
	}.Build()

	for _, def := range defs {
		// We visit the type, not the definition, so we don't count the definition of each type as a reference
		_, err := visitor.Visit(def.Type(), nil)
		if err != nil {
			// Never expected to error
			panic(err)
		}
	}

	return referenceCounts
}
