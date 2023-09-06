/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"fmt"

	"github.com/pkg/errors"
)

// TODO: This is conceptually kinda close to ReferenceGraph except more powerful. We may be able to refactor ReferenceGraph to use this

// typeWalkerRemoveType is a special TypeName that informs the type walker to remove the property containing this TypeName
// entirely.
var typeWalkerRemoveType = MakeInternalTypeName(LocalPackageReference{}, "TypeWalkerRemoveProperty")

// TypeWalker performs a depth first traversal across the types provided, applying the visitor to each TypeDefinition.
// MakeContext is called before each visit, and AfterVisit is called after each visit. ShouldRemoveCycle is called
// if a cycle is detected.
type TypeWalker[C any] struct {
	allDefinitions TypeDefinitionSet
	visitor        TypeVisitor[C]

	// MakeContext is called before any type is visited - including before the root type is visited. It is given the current context and returns
	// a new context for use in the upcoming Visit. When visiting the root type, the ctx parameter is always nil.
	MakeContext func(it InternalTypeName, ctx C) (C, error)

	// AfterVisit is called after the type walker has applied the visitor to a TypeDefinition.
	AfterVisit func(original TypeDefinition, updated TypeDefinition, ctx C) (TypeDefinition, error)

	// ShouldRemoveCycle is called if a cycle is detected. If true is returned the cycle will be pruned, otherwise it will be preserved as-is.
	ShouldRemoveCycle func(def TypeDefinition, ctx C) (bool, error)

	state                         typeWalkerState
	originalVisitInternalTypeName func(this *TypeVisitor[C], it InternalTypeName, ctx C) (Type, error)
	originalVisitObjectType       func(this *TypeVisitor[C], it *ObjectType, ctx C) (Type, error)
}

type typeWalkerState struct {
	result     TypeDefinitionSet
	processing TypeNameSet
}

// NewTypeWalker returns a TypeWalker.
// The provided visitor visitTypeName function must return a TypeName and visitObjectType must return an ObjectType or calls to
// Walk will panic.
func NewTypeWalker[C any](allDefs TypeDefinitionSet, visitor TypeVisitor[C]) *TypeWalker[C] {
	typeWalker := TypeWalker[C]{
		allDefinitions:                allDefs,
		originalVisitInternalTypeName: visitor.visitInternalTypeName,
		originalVisitObjectType:       visitor.visitObjectType,
		AfterVisit:                    IdentityAfterVisit[C],
		MakeContext:                   IdentityMakeContext[C],
		ShouldRemoveCycle:             IdentityShouldRemoveCycle[C],
	}

	// visitor is a copy - modifications won't impact passed visitor
	visitor.visitInternalTypeName = typeWalker.visitInternalTypeName
	visitor.visitObjectType = typeWalker.visitObjectType

	typeWalker.visitor = visitor

	return &typeWalker
}

func (t *TypeWalker[C]) visitInternalTypeName(this *TypeVisitor[C], it InternalTypeName, ctx C) (Type, error) {
	updatedCtx, err := t.MakeContext(it, ctx)
	if err != nil {
		return nil, errors.Wrapf(err, "MakeContext failed for name %q", it)
	}

	if t.originalVisitInternalTypeName != nil {
		visitedTypeName, terr := t.originalVisitInternalTypeName(this, it, updatedCtx)
		if terr != nil {
			return nil, errors.Wrapf(terr, "visitTypeName failed for name %q", it)
		}

		itn, ok := visitedTypeName.(InternalTypeName)
		if !ok {
			panic(fmt.Sprintf("TypeWalker visitor visitTypeName must return an InternalTypeName, instead returned %T", visitedTypeName))
		}

		it = itn
	}

	if IsExternalPackageReference(it.PackageReference()) {
		// Non-local type names are fine, we can exit early
		return it, nil
	}

	def, err := t.allDefinitions.GetDefinition(it)
	if err != nil {
		return nil, err
	}

	// Prevent loops by bypassing this type if it's currently being processed.
	if _, ok := t.state.processing[def.Name()]; ok {
		remove, cycleErr := t.ShouldRemoveCycle(def, updatedCtx)
		if cycleErr != nil {
			return nil, errors.Wrapf(cycleErr, "ShouldRemoveCycle failed for def %q", def.Name())
		}
		if remove {
			return typeWalkerRemoveType, nil
		} else {
			// Preserve
			return def.Name(), nil
		}
	}
	t.state.processing[def.Name()] = struct{}{}

	updatedType, err := this.Visit(def.Type(), updatedCtx)
	if err != nil {
		return nil, errors.Wrapf(err, "error visiting type %q", def.Name())
	}
	updatedDef, err := t.AfterVisit(def, def.WithType(updatedType), updatedCtx)
	if err != nil {
		return nil, errors.Wrapf(err, "AfterVisit failed for def %q", def.Name())
	}

	delete(t.state.processing, def.Name())

	err = t.state.result.AddAllowDuplicates(updatedDef)
	if err != nil {
		return nil, err
	}

	return updatedDef.Name(), nil
}

func (t *TypeWalker[C]) visitObjectType(this *TypeVisitor[C], it *ObjectType, ctx C) (Type, error) {
	result, err := t.originalVisitObjectType(this, it, ctx)
	if err != nil {
		return nil, err
	}

	ot, ok := result.(*ObjectType)
	if !ok {
		panic(fmt.Sprintf("TypeWalker visitor visitObjectType must return a *ObjectType, instead returned %T", result))
	}

	var toRemove []PropertyName
	ot.Properties().ForEach(func(prop *PropertyDefinition) {
		shouldRemove := shouldRemove(prop.PropertyType())
		if shouldRemove {
			toRemove = append(toRemove, prop.PropertyName())
		}
	})

	ot = ot.WithoutSpecificProperties(toRemove...)

	return ot, nil
}

// Walk returns a type definition collection constructed by applying the Visitor to each type in the graph of types
// reachable from the provided TypeDefinition 'def'. TypeDefinitionSet are visited in a depth-first order. Cycles are
// not followed (so each type in a cycle will be visited only once).
func (t *TypeWalker[C]) Walk(def TypeDefinition) (TypeDefinitionSet, error) {
	t.state = typeWalkerState{
		result:     make(TypeDefinitionSet),
		processing: make(TypeNameSet),
	}

	// Visit our own name to start the walk.
	// Initial ctx is nil -- MakeContext will get called and fabricate a context if needed
	var ctx C
	_, err := t.visitor.Visit(def.Name(), ctx)
	if err != nil {
		return nil, err
	}

	return t.state.result, nil
}

// IdentityAfterVisit is the default AfterVisit function for TypeWalker. It returns the TypeDefinition from Visit unmodified.
func IdentityAfterVisit[C any](_ TypeDefinition, updated TypeDefinition, _ C) (TypeDefinition, error) {
	return updated, nil
}

// IdentityMakeContext returns the context unmodified
func IdentityMakeContext[C any](_ InternalTypeName, ctx C) (C, error) {
	return ctx, nil
}

// IdentityShouldRemoveCycle is the default cycle removal behavior. It preserves all cycles unmodified.
func IdentityShouldRemoveCycle[C any](_ TypeDefinition, _ C) (bool, error) {
	return false, nil
}

func shouldRemove(t Type) bool {
	switch cast := t.(type) {
	case TypeName:
		return TypeEquals(cast, typeWalkerRemoveType)
	case *PrimitiveType:
		return false
	case MetaType:
		return shouldRemove(cast.Unwrap())
	case *ArrayType:
		return shouldRemove(cast.Element())
	case *MapType:
		return shouldRemove(cast.KeyType()) || shouldRemove(cast.ValueType())
	}

	panic(fmt.Sprintf("Unknown Type: %T", t))
}
