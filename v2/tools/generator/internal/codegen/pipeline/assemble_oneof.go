/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	kerrors "k8s.io/apimachinery/pkg/util/errors"
	"k8s.io/klog/v2"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
)

const AssembleOneOfTypesID = "assembleOneOfTypes"

func AssembleOneOfTypes(idFactory astmodel.IdentifierFactory) *Stage {
	stage := NewStage(
		AssembleOneOfTypesID,
		"Assemble OneOf types from OpenAPI Fragments",
		func(ctx context.Context, state *State) (*State, error) {
			assembler := newOneOfAssembler(state.Definitions(), idFactory)
			newDefs, err := assembler.assembleOneOfs()
			if err != nil {
				return nil, errors.Wrapf(err, "assembling OneOf types")
			}
			return state.WithDefinitions(
					state.Definitions().OverlayWith(newDefs)),
				nil
		})

	return stage
}

// oneOfAssembler is used to build up our oneOf types.
type oneOfAssembler struct {
	// A set of all the TypeDefinitions we're assembling from
	defs astmodel.TypeDefinitionSet

	// our usual factory for creating identifiers
	idFactory astmodel.IdentifierFactory

	// A set of all the changes we're going to make to TypeDefinitions
	// We batch these up for processing at the end.
	updates map[astmodel.TypeName]*astmodel.OneOfType
}

// newOneOfAssembler creates a new assembler to build our oneOf types
func newOneOfAssembler(
	defs astmodel.TypeDefinitionSet,
	idFactory astmodel.IdentifierFactory,
) *oneOfAssembler {
	// As a reasonable initial size for 'updates', we assume 25% of our definitions are oneOfs
	// Actual benchmarking: of 88k original definitions, we have 18k oneOfs (approximately 20%)
	updateSize := len(defs) / 4

	result := &oneOfAssembler{
		defs:      defs,
		idFactory: idFactory,
		updates:   make(map[astmodel.TypeName]*astmodel.OneOfType, updateSize),
	}

	return result
}

// assembleOneOfs is called to build up the oneOf types.
// We return a TypeDefinitionSet containing the new types.
func (oa *oneOfAssembler) assembleOneOfs() (astmodel.TypeDefinitionSet, error) {
	// Find all OneOfs and work out the modifications we need to make to each one
	var errs []error
	for tn := range oa.defs {
		if _, ok := oa.asOneOf(tn); ok {
			err := oa.assemble(tn)
			if err != nil {
				errs = append(errs, err)
			}
		}
	}

	// If we encountered any errors, bail out
	err := kerrors.NewAggregate(errs)
	if err != nil {
		return nil, err
	}

	// Assemble our results by applying our accumulated modifications
	result := make(astmodel.TypeDefinitionSet)
	for tn, oneOf := range oa.updates {
		def := oa.defs.MustGetDefinition(tn)
		result.Add(def.WithType(oneOf))
	}

	return result, nil
}

// assemble is called to build up a single oneOf type.
func (oa *oneOfAssembler) assemble(name astmodel.TypeName) error {
	for _, tn := range oa.findParentsOf(name) {
		err := oa.assemblePair(tn, name)
		if err != nil {
			return err
		}
	}

	return nil
}

// assemblePair updates a pair of nodes that have a relationship and recursively invokes itself to traverse the entire
// tree of referenced imports.
// ancestor is the parent or super-type. It may be a less specialised OneOf, a root OneOf, an Object or an AllOf
// descendent is the child or subtype. It will be an intermediate or leaf OneOf.
func (oa *oneOfAssembler) assemblePair(ancestor astmodel.TypeName, descendant astmodel.TypeName) error {
	// Remove any direct reference to the parent from the leaf
	err := oa.removeParentReferenceFromLeaf(ancestor, descendant)
	if err != nil {
		return errors.Wrapf(err, "removing parent reference %s from leaf %s", ancestor, descendant)
	}

	if oa.hasDiscriminatorValue(descendant) {
		// Copy any common properties from the parent to the child
		err = oa.embedCommonPropertiesInLeaf(ancestor, descendant)
		if err != nil {
			return errors.Wrapf(err, "embedding common properties from %s in leaf %s", ancestor, descendant)
		}
	}

	if oa.isDiscriminatorRoot(ancestor) {
		// Strip common properties from the root (we replace these with other properties in a later stage)
		err := oa.removeCommonPropertiesFromRoot(ancestor)
		if err != nil {
			return errors.Wrapf(err, "removing common properties from OneOf root %s", ancestor)
		}

		// We've found the actual root, add a reference to the leaf
		err = oa.addLeafReferenceToRoot(ancestor, descendant)
		if err != nil {
			return errors.Wrapf(err, "adding leaf reference to %s from root %s", descendant, ancestor)
		}

		// Ensure the discriminator property exists on the leaf
		err = oa.addDiscriminatorProperty(descendant, ancestor)
		if err != nil {
			return errors.Wrapf(err, "ensuring discriminator property exists on %s", descendant)
		}
	}

	// Recursively walk any parent's of the current ancestor
	for _, tn := range oa.findParentsOf(ancestor) {
		err := oa.assemblePair(tn, descendant)
		if err != nil {
			return err
		}
	}

	return nil
}

// findParentsFor returns the parents of the provided typename, if any.
// name is the name of the type to find parents for.
func (oa *oneOfAssembler) findParentsOf(
	name astmodel.TypeName,
) []astmodel.TypeName {
	// Look up the definition
	def, ok := oa.defs[name]
	if !ok {
		// No definition, this may point to an external type
		return nil
	}

	if tn, ok := astmodel.AsTypeName(def.Type()); ok {
		// Expand aliases
		return oa.findParentsOf(tn)
	}

	if oneOf, ok := astmodel.AsOneOfType(def.Type()); ok {
		// Found a oneOf, look for parents here
		types := oneOf.Types()
		result := make([]astmodel.TypeName, 0, types.Len())
		types.ForEach(func(t astmodel.Type, ix int) {
			if tn, ok := astmodel.AsTypeName(t); ok {
				result = append(result, tn)
			}
		})

		return result
	}

	if allOf, ok := astmodel.AsAllOfType(def.Type()); ok {
		// Found an allOf, look for parents here
		types := allOf.Types()
		result := make([]astmodel.TypeName, 0, types.Len())
		types.ForEach(func(t astmodel.Type, ix int) {
			if tn, ok := astmodel.AsTypeName(t); ok {
				result = append(result, tn)
			}
		})

		return result
	}

	// No parents
	return nil
}

// removeCommonPropertiesFromRoot removes any common properties from the root of a oneOf hierarchy.
// root is the name of the root of the oneOf hierarchy.
func (oa *oneOfAssembler) removeCommonPropertiesFromRoot(root astmodel.TypeName) error {
	err := oa.updateOneOf(
		root,
		func(oneOf *astmodel.OneOfType) (*astmodel.OneOfType, error) {
			result := oneOf.WithoutAnyPropertyObjects()
			return result, nil
		})

	return errors.Wrapf(err, "removing common properties from root %s", root)
}

// removeParentReferenceFromLeaf removes any reference to the root from the leaf.
// root is the name of the root of the oneOf hierarchy.
// leaf is the name of the leaf from which to remove the reference.
func (oa *oneOfAssembler) removeParentReferenceFromLeaf(parent astmodel.TypeName, leaf astmodel.TypeName) error {
	err := oa.updateOneOf(
		leaf,
		func(oneOf *astmodel.OneOfType) (*astmodel.OneOfType, error) {
			return oneOf.WithoutType(parent), nil
		})
	return errors.Wrapf(err, "removing parent reference from leaf %s", leaf)
}

// embedCommonPropertiesInLeaf embeds any common properties found on the parent into the leaf.
// Parents can be OneOf, AllOf or object types.
// A OneOf parent may be either the root of the OneOf tree, or an intermediate OneOf option that is also specialised
// An AllOf parent is always an intermediate node, which isn't a valid option in itself, but is used to embed common properties
//
// One case where this happens is in MachineLearningServices, v2022-50-1, where DataVersionBase references AssetBase
// to pull in common properties.
func (oa *oneOfAssembler) embedCommonPropertiesInLeaf(parent astmodel.TypeName, leaf astmodel.TypeName) error {
	commonProperties := oa.findCommonProperties(parent)
	err := oa.updateOneOf(
		leaf,
		func(oneOf *astmodel.OneOfType) (*astmodel.OneOfType, error) {
			result := oneOf
			for _, obj := range commonProperties {
				result = result.WithAdditionalPropertyObject(obj)
			}

			return result, nil
		})

	return errors.Wrapf(err, "embedding common properties from OneOf in leaf %s", leaf)
}

// findCommonProperties finds all the object properties referenced by name to embed in a leaf OneOf.
// Recursively walking the reference tree is handled elsewhere.
func (oa *oneOfAssembler) findCommonProperties(name astmodel.TypeName) []*astmodel.ObjectType {
	if parentOneOf, ok := oa.asOneOf(name); ok {
		// Parent is a OneOf
		return parentOneOf.PropertyObjects()
	}

	if parentAllOf, ok := oa.asAllOf(name); ok {
		// Parent is an AllOf
		result := make([]*astmodel.ObjectType, 0, parentAllOf.Types().Len())
		parentAllOf.Types().ForEach(func(t astmodel.Type, ix int) {
			if obj, ok := astmodel.AsObjectType(t); ok {
				result = append(result, obj)
			}
		})

		return result
	}

	if obj, ok := oa.asObject(name); ok {
		return []*astmodel.ObjectType{obj}
	}

	return nil
}

// addLeafReferenceToRoot adds a reference to the leaf to the root.
// root is the name of the root of the oneOf hierarchy.
// leaf is the name of the leaf to add a reference to.
func (oa *oneOfAssembler) addLeafReferenceToRoot(root astmodel.TypeName, leaf astmodel.TypeName) error {
	err := oa.updateOneOf(
		root,
		func(oneOf *astmodel.OneOfType) (*astmodel.OneOfType, error) {
			return oneOf.WithType(leaf), nil
		})

	return errors.Wrapf(err, "adding leaf reference %s to root %s", leaf, root)
}

func (oa *oneOfAssembler) addDiscriminatorProperty(name astmodel.TypeName, rootName astmodel.TypeName) error {
	// Find the name of the discriminator property from the root
	root, ok := oa.asOneOf(rootName)
	if !ok {
		return errors.Errorf("couldn't find root %s", rootName)
	}

	discriminatorProperty := root.DiscriminatorProperty()
	propertyName := oa.idFactory.CreatePropertyName(discriminatorProperty, astmodel.Exported)
	propertyJson := oa.idFactory.CreateStringIdentifier(discriminatorProperty, astmodel.NotExported)

	err := oa.updateOneOf(
		name,
		func(oneOf *astmodel.OneOfType) (*astmodel.OneOfType, error) {
			// Create a name for the discriminator value
			discriminatorValue := oneOf.DiscriminatorValue()
			valueName := oa.idFactory.CreateIdentifier(discriminatorValue, astmodel.Exported)

			// Create the discriminator property as a single valued enum
			enumType := astmodel.NewEnumType(
				astmodel.StringType,
				astmodel.MakeEnumValue(valueName, fmt.Sprintf("%q", discriminatorValue)))

			property := astmodel.NewPropertyDefinition(
				propertyName,
				propertyJson,
				astmodel.NewOptionalType(enumType))

			if discriminatorValue == "" {
				klog.Warning("Weirdness")
			}

			obj := astmodel.NewObjectType().WithProperty(property)
			return oneOf.WithAdditionalPropertyObject(obj), nil
		})

	return err
}

// asOneOf returns a OneOf and true if the provided name identifies is a OneOf, nil and false otherwise.
// name is the name of the type to check.
func (oa *oneOfAssembler) asOneOf(name astmodel.TypeName) (*astmodel.OneOfType, bool) {
	def, ok := oa.defs[name]
	if !ok {
		// Name doesn't identify anything!
		// (This can happen if the name references our standard library, for example)
		return nil, false
	}

	return astmodel.AsOneOfType(def.Type())
}

// asAllOf returns an AllOf and true if the provided name identifies an AllOf, nil and false otherwise.
// name is the name of the type to check.
func (oa *oneOfAssembler) asAllOf(name astmodel.TypeName) (*astmodel.AllOfType, bool) {
	def, ok := oa.defs[name]
	if !ok {
		// Name doesn't identify anything!
		// (This can happen if the name references our standard library, for example)
		return nil, false
	}

	return astmodel.AsAllOfType(def.Type())
}

// asObject returns an object and true if the provided name identifies an object, false otherwise
// name is the name of the type to check.
func (oa *oneOfAssembler) asObject(name astmodel.TypeName) (*astmodel.ObjectType, bool) {
	def, ok := oa.defs[name]
	if !ok {
		// Name doesn't identify anything!
		// (This can happen if the name references our standard library, for example)
		return nil, false
	}

	return astmodel.AsObjectType(def.Type())
}

// hasDiscriminatorValue returns true if the oneOf has a discriminator value
func (oa *oneOfAssembler) hasDiscriminatorValue(name astmodel.TypeName) bool {
	oneOf, ok := oa.asOneOf(name)
	if !ok {
		return false
	}

	return oneOf.HasDiscriminatorValue()
}

// isDiscriminatorRoot returns true if the oneOf is the root of a set of OneOf options
func (oa *oneOfAssembler) isDiscriminatorRoot(name astmodel.TypeName) bool {
	oneOf, ok := oa.asOneOf(name)
	if !ok {
		return false
	}

	return oneOf.HasDiscriminatorProperty()
}

// updateOneOf is used to make a change to a oneOf type.
// name is the TypeName of the oneOf to update.
// transform is a function that takes the existing oneOf and returns a new oneOf.
// All updates funnel through here to ensure we're consistent with storing our changes in oa.updates without
// overwriting prior changes.
func (oa *oneOfAssembler) updateOneOf(
	name astmodel.TypeName,
	transform func(ofType *astmodel.OneOfType) (*astmodel.OneOfType, error),
) error {
	// Default to pulling an instance with existing changes applied
	oneOf, ok := oa.updates[name]
	if !ok {
		def := oa.defs.MustGetDefinition(name)

		oneOf, ok = astmodel.AsOneOfType(def.Type())
		if !ok {
			return errors.Errorf("found definition for %s, but it wasn't a OneOf", name)
		}
	}

	updated, err := transform(oneOf)
	if err != nil {
		return errors.Wrapf(err, "transforming oneOf %s", name)
	}

	oa.updates[name] = updated
	return nil
}
