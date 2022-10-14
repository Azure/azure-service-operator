/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"context"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/pkg/errors"
)

const AssembleOneOfTypesID = "assembleOneOfTypes"

func AssembleOneOfTypes() *Stage {
	stage := NewStage(
		AssembleOneOfTypesID,
		"Assemble OneOf types from OpenAPI Fragments",
		func(ctx context.Context, state *State) (*State, error) {
			assembler := newOneOfAssembler(state.Definitions())

			newDefs := assembler.assembleOneOfs()
			return state.WithDefinitions(
					state.Definitions().OverlayWith(newDefs)),
				nil
		})

	return stage
}

type oneOfAssembler struct {
	// A set of all the TypeDefinitions we're assembling from
	defs astmodel.TypeDefinitionSet
	// A map of all the leaf OneOfs, indexed by the root they reference
	leavesByRoot map[astmodel.TypeName][]astmodel.TypeName
	// The set of definitions we've modified
	updatedDefs astmodel.TypeDefinitionSet
}

// newOneOfAssembler creates a new assembler to build our oneOf types
func newOneOfAssembler(defs astmodel.TypeDefinitionSet) *oneOfAssembler {
	result := &oneOfAssembler{
		defs:        defs,
		updatedDefs: make(astmodel.TypeDefinitionSet, 100),
	}

	result.indexLeaves()

	return result
}

// assembleOneOfs is called to build up the oneOf types
func (o *oneOfAssembler) assembleOneOfs() astmodel.TypeDefinitionSet {
	for rootName, leaves := range o.leavesByRoot {
		err := o.restructureRoot(rootName, leaves)
		if err != nil {
			o.saveError(rootName, errors.Wrapf(err, "error restructuring root oneOf %v", rootName))
			continue
		}

		err = o.restructureLeaves(rootName, leaves)
		if err != nil {
			o.saveError(rootName, errors.Wrapf(err, "error restructuring leaf oneOfs for %v", rootName))
		}
	}

	return o.updatedDefs
}

// restructureRoot embeds names of the leaves into the root and updates our underlying TypeDefinitionSet.
func (o *oneOfAssembler) restructureRoot(rootName astmodel.TypeName, leaves []astmodel.TypeName) error {

	root, err := o.lookupOneOf(rootName)
	if err != nil {
		return errors.Wrapf(err, "root type name %s didn't map to a known OneOf", rootName)
	}

	// Add all the leaves into the base type so that it knows about them
	for _, leaf := range leaves {
		root = root.WithAdditionalType(leaf)
	}

	// Replace the old root with the new one
	o.saveOneOf(rootName, root)

	return nil
}

// restructureLeaves removes the name of the root from each of the leaves and updates our underlying TypeDefinitionSet.
// We don't need the name reference as they're now embedded in the root.
// Plus, failing to remove it results in a circular reference.
func (o *oneOfAssembler) restructureLeaves(root astmodel.TypeName, leaves []astmodel.TypeName) error {
	for _, leafName := range leaves {
		leaf, err := o.lookupOneOf(leafName)
		if err != nil {
			return errors.Wrapf(err, "leaf type name %s didn't map to a known OneOf", leafName)
		}

		leaf = leaf.WithoutType(root)
		o.saveOneOf(leafName, leaf)
	}

	return nil
}

// indexLeaves is called once to build up our index of all the oneOf leaf instances
func (o *oneOfAssembler) indexLeaves() {
	oneOfs := o.defs.Where(func(def astmodel.TypeDefinition) bool {
		if oneOf, ok := astmodel.AsOneOfType(def.Type()); ok {
			isLeaf := oneOf.DiscriminatorValue() != ""
			return isLeaf
		}

		return false
	})

	o.leavesByRoot = make(map[astmodel.TypeName][]astmodel.TypeName, len(oneOfs)/4)
	for _, def := range oneOfs {
		o.addLeafToIndex(def)
	}
}

// addLeafToIndex adds a single leaf to our index, based on the root types it references.
// We expect there to be one, but nothing restricts that from happening
func (o *oneOfAssembler) addLeafToIndex(def astmodel.TypeDefinition) {
	oneOf, _ := astmodel.AsOneOfType(def.Type())
	oneOf.Types().ForEach(func(t astmodel.Type, _ int) {
		tn, ok := astmodel.AsTypeName(t)
		if !ok {
			// Not a typename, can't point to a root
			return
		}

		root, ok := o.defs[tn]
		if !ok {
			// Not a known type, doesn't point to a root
			return
		}

		rootOneOf, ok := astmodel.AsOneOfType(root.Type())
		if !ok {
			// TypeName doesn't identify a oneOf, cannot be the root
			return
		}

		if rootOneOf.DiscriminatorValue() == "" {
			// Not a root oneOf, don't index it
			return
		}

		// Capturing names requires an extra indirection, but ensures we don't end up with multiple (different!) copies
		// of the same type definition lying around.
		o.leavesByRoot[tn] = append(o.leavesByRoot[tn], def.Name())
	})
}

// lookupOneOf returns a OneOf type with the given name, if it exists
// We look first in the updatedDefs (as we may have already modified the type), then in the original defs
func (o *oneOfAssembler) lookupOneOf(name astmodel.TypeName) (*astmodel.OneOfType, error) {
	def, ok := o.updatedDefs[name]
	if !ok {
		// try looking in the original defs instead
		def, ok = o.defs[name]
	}

	if !ok {
		// Still not found
		return nil, errors.Errorf("didn't find type %s", name)
	}

	if tn, ok := astmodel.AsTypeName(def.Type()); ok {
		one, err := o.lookupOneOf(tn)
		if err != nil {
			return nil, errors.Wrapf(err, "looking up alias for %s", name)
		}

		return one, nil
	}

	oneOf, ok := astmodel.AsOneOfType(def.Type())
	if !ok {
		// Not a OneOf
		return nil, errors.Errorf(
			"type %s is not a OneOf, but actually a %s",
			name,
			astmodel.DebugDescription(def.Type()))
	}

	return oneOf, nil
}

func (o *oneOfAssembler) saveOneOf(name astmodel.TypeName, oneOf *astmodel.OneOfType) {
	// We want to explicitly overwrite any pre-existing version as we've made another change
	o.updatedDefs[name] = astmodel.MakeTypeDefinition(name, oneOf)
}

func (o *oneOfAssembler) saveError(name astmodel.TypeName, err error) {
	def, _ := o.defs[name]
	erroredType := astmodel.NewErroredType(def.Type(), []string{err.Error()}, nil)
	o.updatedDefs[name] = astmodel.MakeTypeDefinition(name, erroredType)
}
