/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"context"
	"fmt"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/pkg/errors"
	kerrors "k8s.io/apimachinery/pkg/util/errors"
	"k8s.io/klog/v2"
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
				return nil, errors.Wrapf(err, "couldn't assemble OneOf types")
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
	result := &oneOfAssembler{
		defs:      defs,
		idFactory: idFactory,
		updates:   make(map[astmodel.TypeName]*astmodel.OneOfType),
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

	root, ok := oa.findParentFor(name)
	if !ok {
		// No parent, so we've found a root
		return oa.removeCommonPropertiesFromRoot(name)
	}

	// We have a parent, so we're a leaf
	for {
		// Remove any direct reference to the parent from the leaf
		err := oa.removeParentReferenceFromLeaf(root, name)
		if err != nil {
			return errors.Wrapf(err, "couldn't remove parent reference to %s from leaf %s", root, name)
		}

		// Copy any common properties from the parent to the child
		err = oa.embedCommonPropertiesInLeaf(root, name)
		if err != nil {
			return errors.Wrapf(err, "couldn't embed common properties from %s in leaf %s", root, name)
		}

		// Now we need to check if the parent is a leaf itself
		parent, ok := oa.findParentFor(root)
		if !ok {
			break
		}

		root = parent
	}

	// We've found the actual root, add a reference to the leaf
	err := oa.addLeafReferenceToRoot(root, name)
	if err != nil {
		return errors.Wrapf(err, "couldn't add leaf reference to %s from root %s", name, root)
	}

	// Ensure the discriminator property exists on the leaf
	err = oa.addDiscriminatorProperty(name, root)
	if err != nil {
		return errors.Wrapf(err, "couldn't ensure discriminator property exists on %s", name)
	}

	return nil
}

// findParentFor returns the parent of the provided oneOf, if any.
// leaf is the name of the oneOf to find the parent for.
// If the provided oneOf has a parent, the parent's name is returned along with true.
// Otherwise, we return an empty name and false.
func (oa *oneOfAssembler) findParentFor(
	leaf astmodel.TypeName,
) (astmodel.TypeName, bool) {
	// Look up the definition
	def, ok := oa.defs[leaf]
	if !ok {
		// No definition, this may point to an external type
		return astmodel.EmptyTypeName, false
	}

	oneOf, ok := astmodel.AsOneOfType(def.Type())
	if !ok {
		// Not a OneOF, can't have a parent
		return astmodel.EmptyTypeName, false
	}

	result := astmodel.EmptyTypeName
	found := false

	// Look for a reference to a parent
	// We iterate through all the referenced types - finding more than one typename that points to a OneOf violates
	// our understanding of the way Azure uses OpenAPI specs, so we'll error out if we find more than one.
	oneOf.Types().ForEach(func(t astmodel.Type, _ int) {
		if tn, ok := astmodel.AsTypeName(t); ok {
			if _, ok := oa.asOneOf(tn); ok {
				if found {
					// We've already found a parent, this is an error
					panic(fmt.Sprintf("Found multiple parents for %s", leaf))
				}

				result = tn
				found = true
			}
		}
	})

	return result, found
}

// removeCommonPropertiesFromRoot removes any common properties from the root of a oneOf hierarchy.
// root is the name of the root of the oneOf hierarchy.
func (oa *oneOfAssembler) removeCommonPropertiesFromRoot(root astmodel.TypeName) error {
	err := oa.updateOneOf(
		root,
		func(oneOf *astmodel.OneOfType) (*astmodel.OneOfType, error) {
			result := oneOf
			oneOf.Types().ForEach(
				func(t astmodel.Type, _ int) {
					if obj, ok := oa.AsCommonProperties(t); ok {
						// Found an object representing common properties
						result = result.WithoutType(obj)
					}
				})

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

// embedCommonPropertiesInLeaf embeds any common properties found on the parent into the leaf
func (oa *oneOfAssembler) embedCommonPropertiesInLeaf(parent astmodel.TypeName, leaf astmodel.TypeName) error {
	parentOneOf, ok := oa.asOneOf(parent)
	if !ok {
		// Not a parent, nothing to do
		return nil
	}

	err := oa.updateOneOf(
		leaf,
		func(oneOf *astmodel.OneOfType) (*astmodel.OneOfType, error) {
			result := oneOf
			parentOneOf.Types().ForEach(
				func(t astmodel.Type, _ int) {
					if obj, ok := oa.AsCommonProperties(t); ok {
						// Found an object representing common properties
						result = result.WithType(obj)
					}
				})

			return result, nil
		})
	return errors.Wrapf(err, "embedding common properties in leaf %s", leaf)
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
	propertyJson := oa.idFactory.CreateIdentifier(discriminatorProperty, astmodel.NotExported)

	err := oa.updateOneOf(
		name,
		func(oneOf *astmodel.OneOfType) (*astmodel.OneOfType, error) {
			// Create a name for the discriminator value
			discriminatorValue := oneOf.DiscriminatorValue()
			valueName := oa.idFactory.CreateIdentifier(discriminatorValue, astmodel.Exported)

			// Create the discriminator property as a single valued enum
			property := astmodel.NewPropertyDefinition(
				propertyName,
				propertyJson,
				astmodel.NewEnumType(
					astmodel.StringType,
					astmodel.MakeEnumValue(valueName, fmt.Sprintf("%q", discriminatorValue))))

			if discriminatorValue == "" {
				klog.Warning("Weirdness")
			}

			obj := astmodel.NewObjectType().WithProperty(property)
			return oneOf.WithType(obj), nil
		})

	return err
}

// asOneOf returns true if the provided name identifies is a OneOf, false otherwise.
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
		return errors.Wrapf(err, "couldn't transform oneOf %s", name)
	}

	oa.updates[name] = updated
	return nil
}

/*










 */

func (*oneOfAssembler) findAllOneOfs(defs astmodel.TypeDefinitionSet) map[astmodel.TypeName]*astmodel.OneOfType {
	result := make(map[astmodel.TypeName]*astmodel.OneOfType)

	for _, def := range defs {
		if oneOf, ok := astmodel.AsOneOfType(def.Type()); ok {
			result[def.Name()] = oneOf
		}
	}

	return result
}

// findCommonProperties returns a (possibly empty) slice of Objects that contain properties that need to be hoisted
// to the provided oneOf. This list is made up of properties from both directly referenced roots, and indirectly
// referenced roots.
func (oa *oneOfAssembler) findCommonProperties(oneOf *astmodel.OneOfType) ([]*astmodel.ObjectType, error) {
	result := make([]*astmodel.ObjectType, 0, oneOf.Types().Len())

	// We need to look at all the types referenced by this oneOf
	oneOf.Types().ForEach(func(t astmodel.Type, _ int) {

	})

	return result, nil
}

// AsCommonProperties captures the test used to determine if a type is a common properties object, ensuring that
// we're consistent with our tests from multiple locations
func (oa *oneOfAssembler) AsCommonProperties(t astmodel.Type) (*astmodel.ObjectType, bool) {
	return astmodel.AsObjectType(t)
}
