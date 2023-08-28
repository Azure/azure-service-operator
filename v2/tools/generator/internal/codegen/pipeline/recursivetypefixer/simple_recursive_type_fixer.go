/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package recursivetypefixer

import (
	"github.com/Azure/azure-service-operator/v2/internal/set"
	"github.com/go-logr/logr"
	"github.com/pkg/errors"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
)

// TODO: If we end up needing more complex cycle detection we can rework TypeWalker to do it, but for now this
// TODO: is simpler and seems to cover all the cases we need.
// SimpleRecursiveTypeFixer removes circular references from types that refer to themselves directly.
type SimpleRecursiveTypeFixer struct {
	visitor        *astmodel.TypeVisitor
	newDefinitions astmodel.TypeDefinitionSet
	log            logr.Logger
}

type simpleRecursiveTypeFixerContext struct {
	name         astmodel.TypeName
	mustUnroll   bool
	unrolledName astmodel.TypeName
}

func (c simpleRecursiveTypeFixerContext) WithUnrolledName(name astmodel.TypeName) simpleRecursiveTypeFixerContext {
	c.unrolledName = name
	c.mustUnroll = true
	return c
}

func NewSimpleRecursiveTypeFixer(log logr.Logger) *SimpleRecursiveTypeFixer {
	result := &SimpleRecursiveTypeFixer{
		newDefinitions: make(astmodel.TypeDefinitionSet),
		log:            log,
	}

	visitor := astmodel.TypeVisitorBuilder{
		VisitObjectType: astmodel.MakeIdentityVisitOfObjectType(result.unrollObjectTypeProperty),
		VisitTypeName:   result.unrollRecursiveReference,
	}.Build()
	result.visitor = &visitor

	return result
}

// Fix checks the supplied definition for a self reference and unrolls it.
func (s *SimpleRecursiveTypeFixer) Fix(def astmodel.TypeDefinition) (astmodel.TypeDefinition, error) {
	return s.visitor.VisitDefinition(def, simpleRecursiveTypeFixerContext{name: def.Name()})
}

// Types returns any new types created by this type fixer
func (s *SimpleRecursiveTypeFixer) Types() astmodel.TypeDefinitionSet {
	return s.newDefinitions
}

func (s *SimpleRecursiveTypeFixer) unrollObjectTypeProperty(ot *astmodel.ObjectType, prop *astmodel.PropertyDefinition, ctx interface{}) (interface{}, error) {
	typedCtx := ctx.(simpleRecursiveTypeFixerContext)
	name := typedCtx.name

	if !isPropertyMatchingTypeName(name, prop) {
		return ctx, nil
	}

	// We know we have at least one property we need to remove. There might be more than 1 direct reference, do a
	// quick scan for others. This is to deal with situations where there are multiple properties we need to remove.
	// Without this check, we end up generating N different types (each with 1 property removed) which isn't allowed
	// because types with the same name must match structurally when we go to add them to the type collection later.
	toRemove := set.Make(prop.PropertyName())
	for _, p := range ot.Properties().AsSlice() {
		if _, ok := toRemove[p.PropertyName()]; ok {
			continue
		}

		if isPropertyMatchingTypeName(name, p) {
			toRemove.Add(p.PropertyName())
		}
	}

	unrolledType := ot
	for propName := range toRemove {
		unrolledType = unrolledType.WithoutProperty(propName)
	}
	unrolledName := astmodel.MakeInternalTypeName(name.PackageReference(), name.Name()+"_Unrolled")
	unrolledDef := astmodel.MakeTypeDefinition(unrolledName, unrolledType)

	err := s.newDefinitions.AddAllowDuplicates(unrolledDef)
	if err != nil {
		return nil, errors.Wrapf(err, "error adding unrolled type %q", name)
	}

	return typedCtx.WithUnrolledName(unrolledName), nil
}

func (s *SimpleRecursiveTypeFixer) unrollRecursiveReference(this *astmodel.TypeVisitor, it astmodel.TypeName, ctx interface{}) (astmodel.Type, error) {
	typedCtx := ctx.(simpleRecursiveTypeFixerContext)

	if !typedCtx.mustUnroll {
		return astmodel.IdentityVisitOfTypeName(this, it, ctx)
	}

	s.log.V(2).Info(
		"unrolling recursive reference",
		"from", typedCtx.name,
		"to", typedCtx.unrolledName)

	return astmodel.IdentityVisitOfTypeName(this, typedCtx.unrolledName, typedCtx)
}

// asTypeName checks if the type is an astmodel.TypeName or can be unwrapped into an astmodel.TypeName.
// This differs from astmodel.AsTypeName() in that it unwraps Arrays and Maps too. Use with caution.
func asTypeName(t astmodel.Type) (astmodel.TypeName, bool) {
	typeName, ok := astmodel.AsTypeName(t)
	if ok {
		return typeName, true
	}

	arrayType, ok := astmodel.AsArrayType(t)
	if ok {
		return asTypeName(arrayType.Element())
	}

	mapType, ok := astmodel.AsMapType(t)
	if ok {
		return asTypeName(mapType.ValueType())
	}

	return nil, false
}

func isPropertyMatchingTypeName(name astmodel.TypeName, prop *astmodel.PropertyDefinition) bool {
	propTypeName, ok := asTypeName(prop.PropertyType())
	if !ok {
		return false
	}

	if !astmodel.TypeEquals(name, propTypeName) {
		return false
	}

	return true
}
